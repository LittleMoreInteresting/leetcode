package main

import (
	"fmt"
	"reflect"
	"time"
)

type Token struct{}

func newWorker(id int, ch chan Token, nextCh chan Token) {
	for {
		token := <-ch         // 取得令牌
		fmt.Println((id + 1)) // id从1开始
		time.Sleep(time.Second)
		nextCh <- token
	}
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()

	<-or(
		sig(10*time.Second),
		sig(20*time.Second),
		sig(30*time.Second),
		sig(40*time.Second),
		sig(50*time.Second),
		sig(01*time.Minute),
	)
	<-or2(
		sig(10*time.Second),
		sig(20*time.Second),
		sig(30*time.Second),
		sig(40*time.Second),
		sig(50*time.Second),
		sig(01*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	// 特殊情况，只有零个或者1个chan
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2: // 2个也是一种特殊情况
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default: //超过两个，二分法递归处理
			m := len(channels) / 2
			select {
			case <-or(channels[:m]...):
			case <-or(channels[m:]...):
			}
		}
	}()

	return orDone
}

func or2(channels ...<-chan interface{}) <-chan interface{} {
	//特殊情况，只有0个或者1个
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		// 利用反射构建SelectCase
		var cases []reflect.SelectCase
		for _, c := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		// 随机选择一个可用的case
		reflect.Select(cases)
	}()

	return orDone
}

func fanInReflect(chans ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		// 构造SelectCase slice
		var cases []reflect.SelectCase
		for _, c := range chans {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		// 循环，从cases中选择一个可用的
		for len(cases) > 0 {
			i, v, ok := reflect.Select(cases)
			if !ok { // 此channel已经close
				cases = append(cases[:i], cases[i+1:]...)
				continue
			}
			out <- v.Interface()
		}
	}()
	return out
}

func fanOut(ch <-chan interface{}, out []chan interface{}, async bool) {
	go func() {
		defer func() { //退出时关闭所有的输出chan
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()

		for v := range ch { // 从输入chan中读取数据
			v := v
			for i := 0; i < len(out); i++ {
				i := i
				if async { //异步
					go func() {
						out[i] <- v // 放入到输出chan中,异步方式
					}()
				} else {
					out[i] <- v // 放入到输出chan中，同步方式
				}
			}
		}
	}()
}

func mapChan(in <-chan interface{}, fn func(interface{}) interface{}) <-chan interface{} {
	out := make(chan interface{}) //创建一个输出chan
	if in == nil {                // 异常检查
		close(out)
		return out
	}

	go func() { // 启动一个goroutine,实现map的主要逻辑
		defer close(out)
		for v := range in { // 从输入chan读取数据，执行业务操作，也就是map操作
			out <- fn(v)
		}
	}()

	return out
}

func reduce(in <-chan interface{}, fn func(r, v interface{}) interface{}) interface{} {
	if in == nil { // 异常检查
		return nil
	}

	out := <-in         // 先读取第一个元素
	for v := range in { // 实现reduce的主要逻辑
		out = fn(out, v)
	}

	return out
}

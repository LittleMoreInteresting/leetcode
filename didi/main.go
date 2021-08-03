package main

import (
	"fmt"
	"time"
)

func main() {
	mutil := MutilCurl(4)
	fmt.Printf("%v\n", mutil)
}

func MutilCurl(n int) []error {
	ch := make(chan error, n)
	for i := 0; i < n; i++ {
		go func() {
			err := curl()
			ch <- err
		}()
	}
	var errors []error
	for i := 0; i < n; i++ {
		select {
		case err := <-ch:
			if err != nil {
				errors = append(errors, err)
			}
		}
	}
	return errors
}

func curl() error {
	fmt.Println("curl ………………")
	time.Sleep(time.Second)
	return new(MyError)
}

type MyError struct {
}

func (e MyError) Error() string {
	return "ERROR"
}

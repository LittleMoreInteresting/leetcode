package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(evalRPN([]string{"2","1","+","3","*"}))
	fmt.Println(evalRPN([]string{"4","13","5","/","+"}))
	fmt.Println(evalRPN([]string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}))
}

func evalRPN(tokens []string) int {
	st := []int{}
	for _, v := range tokens {
		if strings.Contains("+-*/",v) {
			right,left := st[1],st[0]
			st = Pop(st)
			st = Pop(st)
			res :=0
			switch v {
			case "+":
				res = right + left 
			case "-":
				res = right - left 
			case "*":
				res = right * left 
			case "/":
				res = right / left 
			}
			st = Push(st,res)
		}else{
			vNum,_ := strconv.Atoi(v)
			st = Push(st,vNum)
		}
	}
	return st[0]
}

func Pop(nums []int) []int {
	return nums[1:]
}

func Push(nums []int, num int) []int {
	temp := []int{num}
	temp = append(temp, nums...)
	return temp
}
package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(restoreIpAddresses("0000"))
	fmt.Println(restoreIpAddresses("25525511135"))
}

var res []string
var ip []int

func restoreIpAddresses(s string) []string {
	res = []string{}
	ip = make([]int, 4)
	IpAddr(s, 0, 0)
	return res
}

func IpAddr(s string, start, point int) {
	if point == 4 {
		if start == len(s) {
			ipAddr := ""
			for i := 0; i < 4; i++ {
				ipAddr += strconv.Itoa(ip[i])
				if i != 3 {
					ipAddr += "."
				}
			}
			res = append(res, ipAddr)
		}
		return
	}
	if start == len(s) {
		return
	}
	if s[start] == '0' {
		ip[point] = 0
		IpAddr(s, start+1, point+1)
		return
	}

	ipAddr := 0
	for i := start; i < len(s); i++ {
		ipAddr = ipAddr*10 + int(s[i]-'0')
		if ipAddr > 0 && ipAddr <= 255 {
			ip[point] = ipAddr
			IpAddr(s, i+1, point+1)
		} else {
			break
		}
	}
}

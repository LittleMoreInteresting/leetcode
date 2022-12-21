package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var back, forward []string
var current string

const (
	VISIT   = "VISIT"
	QUIT    = "QUIT"
	FORWARD = "FORWARD"
	BACK    = "BACK"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	back, forward, current = []string{}, []string{}, "***###.acm.org/"
	for scanner.Scan() {
		str := scanner.Text()
		if strings.HasPrefix(str, VISIT) {
			split := strings.Split(str, " ")
			if len(split) == 2 {
				forward = []string{}
				fmt.Println("curl :", split[1])
				if len(current) > 0 {
					back = append(back, current)
				}
				current = split[1]
			}
			continue
		}
		switch str {
		case QUIT:
			return
		case BACK:
			blen := len(back)
			if blen > 0 {
				url := back[blen-1]
				fmt.Println(url)
				back = back[:blen-1]
				forward = append(forward, current)
				current = url
			} else {
				fmt.Println("Ignored")
			}
		case FORWARD:
			flen := len(forward)
			if flen > 0 {
				url := forward[flen-1]
				fmt.Println(url)
				forward = forward[:flen-1]
				back = append(back, current)
				current = url
			} else {
				fmt.Println("Ignored")
			}
		}
	}
}

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	res:=strings.Split(s," ")
	a:=make(map[string]int)
	for _,j:= range res {
		a[j] = a[j] + 1
	}
	return a

}

func main() {
	wc.Test(WordCount)
}

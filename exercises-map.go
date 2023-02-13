package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	split := strings.Split(s," ")
	var retorno = make(map[string]int)
	for _, value := range split {
		if(retorno[value] == 1) {
			retorno[value] += 1
		} else {
			retorno[value] = 1
		}
	}
	return retorno
}

func main() {
	wc.Test(WordCount)
}

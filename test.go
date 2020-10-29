package main

import "fmt"

func main() {
	conf := 0
	bg := 0
	text := 9
	msg := "xxxxx"
	a := fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, conf, bg, text, msg, 0x1B)
	fmt.Println(a)
}


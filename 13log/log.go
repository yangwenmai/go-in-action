// Package Log from http://www.flysnow.org/2017/05/06/go-in-action-go-log.html
package main

import (
	"log"
)

func init() {
	log.SetPrefix("[Log]")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	log.Println("我的博客地址:", "http://maiyang.me/")
	log.Printf("我的微信公众号是: %s \n", "tech_tea")
	log.Panicln("panice")
	log.Fatalln("fatal")
}

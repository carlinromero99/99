package main

import "log"

func GetName() string {
	return "carlos"
}

func main() {
	log.Printf("hello %s!", GetName())
}

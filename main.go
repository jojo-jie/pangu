package main

import (
	"log"
	"pangu/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err:%v", err)
	}
}

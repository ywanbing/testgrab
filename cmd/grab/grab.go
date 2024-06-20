package main

import (
	"log"

	"testGrab/internal"
	"testGrab/internal/exit"
)

func main() {
	internal.Run()

	log.Println("按下 Ctrl+C 退出程序")
	exit.WaitSignal()
}

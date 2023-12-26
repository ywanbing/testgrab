package internal

import (
	"log"
	"time"

	"testGrab/internal/config"
	"testGrab/internal/exit"
	"testGrab/internal/session"
)

func Run() {
	app := session.NewSession()
	if !app.Login(config.GetName(), config.GetPwd()) {
		return
	}

	if !app.Home() {
		return
	}

	for _, course := range app.CourseList {
		for i := 0; i < config.GetLoopNum(); i++ {
			log.Println("start "+course.Name+" ----> ", i)
			if !app.Course(course.BatchId, course.Id) {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}

	app.WriteFile()

	log.Println("请查看docs目录下的文件")
	log.Println("按下 Ctrl+C 退出程序")
	exit.WaitSignal()
}

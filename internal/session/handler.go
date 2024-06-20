package session

import (
	"fmt"
	"log"

	"testGrab/internal/common"
	"testGrab/internal/constant"
)

func (s *Session) Login(name string, pwd string) bool {
	err := s.Post(constant.LOGIN_URL, map[string]string{
		"login_name":     name,
		"login_password": pwd,
		"mac_address":    "",
		"ip_address":     "",
	})
	if err != nil {
		log.Println(err)
		return false
	}

	if s.killed {
		log.Println("Login failed")
		return false
	}

	return true
}

func (s *Session) Home() bool {
	err := s.Visit(common.GetHomeUrl())
	if err != nil {
		log.Println("Visit home err:" + err.Error())
		return false
	}

	if s.killed {
		log.Println("visit home err")
		return false
	}
	return true
}

func (s *Session) Course(batchId string, courseId string) bool {
	err := s.Visit(fmt.Sprintf(constant.Course_URL, batchId, courseId))
	if err != nil {
		log.Println("visit course err:" + err.Error())
		return false
	}

	if s.killed {
		log.Println("visit course err")
		return false
	}
	return true
}

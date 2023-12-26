package session

import (
	"log"
	"os"

	"github.com/gocolly/colly"
	"testGrab/internal/config"
	"testGrab/internal/constant"
)

type Session struct {
	*colly.Collector

	// 课程列表
	CourseList []*Course

	// 课程内容
	CourseContent map[string]*CourseContent
}

func NewSession() *Session {
	collector := colly.NewCollector(
		colly.AllowedDomains("www.zk028.cn"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edg/120.0.0.0"))

	collector.CheckHead = false
	collector.AllowURLRevisit = true
	s := &Session{
		Collector:     collector,
		CourseList:    make([]*Course, 0, 6),
		CourseContent: make(map[string]*CourseContent, 6),
	}

	s.RegisterRequestHandler()
	s.RegisterResponseHandler()
	s.RegisterOnHTML()

	return s
}

func (s *Session) GetCourseNameById(id string) string {
	for _, course := range s.CourseList {
		if course.Id == id {
			return course.Name
		}
	}
	return ""
}

// WriteFile 写入文件
func (s *Session) WriteFile() {
	switch config.GenFileType() {
	case string(constant.FileType_MD):
		for _, content := range s.CourseContent {
			ct, _ := content.GenMarkDownFile()
			err := os.WriteFile("./docs/"+content.Name+".md", ct, 0666)
			if err != nil {
				log.Println("生成文件:", "./docs/"+content.Name+".md", "出现了问题：", err.Error())
				continue
			}
			log.Println("生成文件:", "./docs/"+content.Name+".md")
		}
	case string(constant.FileType_JSON):
		// TODO 还没有搞
	}
}

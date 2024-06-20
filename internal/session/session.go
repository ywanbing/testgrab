package session

import (
	"log"
	"os"
	"os/exec"

	"github.com/gocolly/colly"
	"testGrab/internal/common"
	"testGrab/internal/config"
	"testGrab/internal/constant"
)

type Session struct {
	*colly.Collector

	// 课程列表
	CourseList []*Course

	// 课程内容
	CourseContent map[string]*CourseContent

	// 出错强制退出
	killed bool
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
	checkDir("./docs")

	switch config.GenFileType() {
	case string(constant.FileType_MD):
		for _, content := range s.CourseContent {
			ct, _ := content.GenMarkDownFile()
			filePath := "./docs/" + content.Name
			if config.IsSimple() {
				filePath += "_simple"
			}

			filePath += ".md"
			err := os.WriteFile(filePath, ct, 0666)
			if err != nil {
				log.Println("生成文件:", filePath, "出现了问题：", err.Error())
				continue
			}
			log.Println("生成文件:", filePath)
		}
	case string(constant.FileType_JSON):
		// TODO 还没有搞
	case string(constant.FileType_HTML):
		for _, content := range s.CourseContent {
			ct, _ := content.GenMarkDownFile()
			filePath := "./docs/" + content.Name
			if config.IsSimple() {
				filePath += "_simple"
			}

			filePath += ".html"
			html := common.GetHtml(content.Name, ct)
			err := os.WriteFile(filePath, html, 0666)
			if err != nil {
				log.Println("生成文件:", filePath, "出现了问题：", err.Error())
				continue
			}
			log.Println("生成文件:", filePath)
		}
	case string(constant.FileType_PDF):
		for _, content := range s.CourseContent {
			ct, _ := content.GenMarkDownFile()
			filePath := "./docs/" + content.Name
			if config.IsSimple() {
				filePath += "_simple"
			}

			filePath += ".pdf"

			// 先生成html
			html := common.GetHtml(content.Name, ct)
			temp, err := os.CreateTemp(os.TempDir(), "*.html")
			if err != nil {
				log.Println("生成html临时文件:", temp, "出现了问题：", err.Error())
				continue
			}
			defer os.Remove(temp.Name())
			_, _ = temp.Write(html)
			_ = temp.Close()

			err = exec.Command("./tools/wkhtmltopdf.exe", temp.Name(), filePath).Run()
			if err != nil {
				log.Println("生成文件:", filePath, "出现了问题：", err.Error())
				continue
			}
			log.Println("生成文件:", filePath)
		}
	}
}

func checkDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0666)
		if err != nil {
			log.Println("创建目录:", path, "出现了问题：", err.Error())
		}
		log.Println("检查到没有存放文档的目录，创建目录:", path)
	}
}

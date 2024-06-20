package session

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
	"testGrab/internal/common"
	"testGrab/internal/config"
	"testGrab/internal/constant"
)

func (s *Session) RegisterRequestHandler() {
	s.OnRequest(func(r *colly.Request) {
		url := r.URL.String()
		switch url {
		case common.GetHomeUrl():
			r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
			r.Headers.Set("Accept-Encoding", "gzip, deflate, br")
			r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
			r.Headers.Set("Sec-Ch-Ua", `"Not_A Brand";v="8", "Chromium";v="120", "Microsoft Edge";v="120"`)
			r.Headers.Set("Sec-Ch-Ua-Mobile", "?0")
			r.Headers.Set("Sec-Ch-Ua-Platform", `"Windows"`)
			r.Headers.Set("Sec-Fetch-Dest", "document")
			r.Headers.Set("Sec-Fetch-Mode", "navigate")
			r.Headers.Set("Sec-Fetch-Site", "none")
			r.Headers.Set("Sec-Fetch-User", "?1")
			r.Headers.Set("Upgrade-Insecure-Requests", "1")
		}
		log.Println("Visit site :", url)
	})
}

// RegisterResponseHandler 注册响应处理器
func (s *Session) RegisterResponseHandler() {
	s.OnResponse(func(r *colly.Response) {
		reqUrl := r.Request.URL.String()
		switch reqUrl {
		case constant.LOGIN_URL:
			log.Println(string(r.Body))
			if strings.Contains(string(r.Body), "登录成功") {
				log.Println("登录成功")
			} else {
				log.Fatal("登录失败")
				return
			}
			// 设置cookie
			cookies := r.Headers.Values("Set-Cookie")
			_ = s.SetCookies("www.zk028.cn", common.Parse(cookies))
		}
	})
}

func (s *Session) RegisterOnHTML() {
	// 记录课程参数
	s.OnHTML(".page_body a", func(element *colly.HTMLElement) {
		url := element.Request.URL.String()
		if url == common.GetHomeUrl() {

			// javascript:enter_drill(5,'02141')
			attr := element.Attr("href")
			// 5,'02141'
			arg := attr[23 : len(attr)-2]

			// 5 02141
			args := strings.Split(strings.Replace(arg, "'", "", -1), ",")

			course := Course{
				Name:    element.Text,
				Index:   args[0],
				BatchId: config.GetBatchId(),
				Id:      args[1],
			}
			s.CourseList = append(s.CourseList, &course)
		}
	})

	// 记录课程内容
	s.OnHTML(".paper_body>div>div", func(element *colly.HTMLElement) {
		// 只处理对应的网址
		if element.Request.URL.Path != constant.Course_Base_URL {
			return
		}

		// 获取课程名称
		courseName := s.GetCourseNameById(element.Request.URL.Query().Get("course_code"))
		if _, ok := s.CourseContent[courseName]; !ok {
			s.CourseContent[courseName] = &CourseContent{
				Name:          courseName,
				selects:       make([]*XZT, 0, 10),
				selectHash:    make(map[string]struct{}, 10),
				pdts:          make([]*PDT, 0, 10),
				pdtHash:       make(map[string]struct{}, 10),
				tiankongs:     make([]*TKT, 0, 10),
				tiankongsHash: make(map[string]struct{}, 10),
				mcjs:          make([]*MCJS, 0, 10),
				mcjsHash:      make(map[string]struct{}, 10),
				jds:           make([]*JDT, 0, 10),
				jdsHash:       make(map[string]struct{}, 10),
				jss:           make([]*JST, 0, 10),
				jssHash:       make(map[string]struct{}, 10),
				zhs:           make([]*ZHT, 0, 10),
				zhsHash:       make(map[string]struct{}, 10),
				lst:           make([]*LST, 0, 10),
				lstHash:       make(map[string]struct{}, 10),
				qtt:           make([]*QTT, 0, 10),
				qttHash:       make(map[string]struct{}, 10),
			}
		}

		// 获取课程对象
		courseContent := s.CourseContent[courseName]

		s.parseContent(element, courseContent)
	})
}

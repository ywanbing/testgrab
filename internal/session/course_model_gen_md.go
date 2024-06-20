package session

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"testGrab/internal/config"
)

// GenMarkDownFile 生成文件（采用MD格式生成）
func (c *CourseContent) GenMarkDownFile() ([]byte, error) {
	var buff bytes.Buffer
	buff.WriteString("# " + c.Name + "\n\n")

	if len(c.selects) > 0 {
		log.Println(c.Name, "获取到 [", len(c.selects), "] 个选择题")
		buff.WriteString("## 选择题\n\n")
		for i, xzt := range c.selects {
			buff.WriteString(fmt.Sprintf("%d. %s  \n", i+1, xzt.TIMU))
			if config.IsSimple() {
				anies := xzt.OPTION[getAnswerIndex(xzt.ANSWER)]
				if strings.Contains(anies, "http") {
					buff.WriteString(fmt.Sprintf("\t%s  \n", anies))
					continue
				}
				buff.WriteString(fmt.Sprintf("答案：`%s`  \n\n", anies))
				continue
			}
			for _, option := range xzt.OPTION {
				if strings.Contains(option, "http") {
					buff.WriteString(fmt.Sprintf("\t%s  \n", option))
					continue
				}
				buff.WriteString(fmt.Sprintf("\t`%s`  \n", option))
			}
			buff.WriteString(fmt.Sprintf("%s\n\n  ", xzt.ANSWER))
		}
	}

	if len(c.pdts) > 0 {
		log.Println(c.Name, "获取到 [", len(c.pdts), "] 个判断题")
		buff.WriteString("## 判断题 \n\n")
		for i, pdt := range c.pdts {
			buff.WriteString(fmt.Sprintf("%d. %s  \n", i+1, pdt.TIMU))
			if config.IsSimple() {
				buff.WriteString(fmt.Sprintf("答案：%s  \n\n", pdt.OPTION[getAnswerIndex(pdt.ANSWER)]))
				continue
			}
			for _, option := range pdt.OPTION {
				buff.WriteString(fmt.Sprintf("\t%s  \n", option))
			}
			buff.WriteString(fmt.Sprintf("%s  \n\n", pdt.ANSWER))
		}
	}

	if len(c.tiankongs) > 0 {
		log.Println(c.Name, "获取到 [", len(c.tiankongs), "] 个填空题")
		buff.WriteString("## 填空题\n\n")
		for i, tkt := range c.tiankongs {
			buff.WriteString(fmt.Sprintf("%d. %s  \n", i+1, tkt.TIMU))
			buff.WriteString(fmt.Sprintf("%s  \n\n", tkt.ANSWER))
		}
	}

	if len(c.mcjs) > 0 {
		log.Println(c.Name, "获取到 [", len(c.mcjs), "] 个名称解释题")
		buff.WriteString("## 名称解释题\n\n")
		for i, mcjs := range c.mcjs {
			buff.WriteString(fmt.Sprintf("%d. %s  \n", i+1, mcjs.TIMU))
			buff.WriteString(fmt.Sprintf("%s  \n\n", mcjs.ANSWER))
		}
	}

	if len(c.jds) > 0 {
		log.Println(c.Name, "获取到 [", len(c.jds), "] 个简答题")
		buff.WriteString("## 简答题\n\n")
		for i, jdt := range c.jds {
			buff.WriteString(fmt.Sprintf("%d. %s  \n", i+1, jdt.TIMU))
			buff.WriteString(fmt.Sprintf("%s  \n\n", jdt.ANSWER))
		}
	}

	if len(c.jss) > 0 {
		log.Println(c.Name, "获取到 [", len(c.jss), "] 个计算题")
		buff.WriteString("## 计算题\n\n")
		for i, jst := range c.jss {
			buff.WriteString(fmt.Sprintf("%d. %s  \n", i+1, jst.TIMU))
			buff.WriteString(fmt.Sprintf("%s  \n\n", jst.ANSWER))
		}
	}

	if len(c.zhs) > 0 {
		log.Println(c.Name, "获取到 [", len(c.zhs), "] 个综合题")
		buff.WriteString("## 综合题\n\n")
		for i, zht := range c.zhs {
			buff.WriteString(fmt.Sprintf("%d. %s  \n", i+1, zht.TIMU))
			buff.WriteString(fmt.Sprintf("%s  \n\n", zht.ANSWER))
		}
	}

	if len(c.lst) > 0 {
		log.Println(c.Name, "获取到 [", len(c.lst), "] 个论述题")
		buff.WriteString("## 论述题\n\n")
		for i, lst := range c.lst {
			buff.WriteString(fmt.Sprintf("%d. %s  \n", i+1, lst.TIMU))
			buff.WriteString(fmt.Sprintf("%s  \n\n", lst.ANSWER))
		}
	}

	if len(c.qtt) > 0 {
		log.Println(c.Name, "获取到 [", len(c.qtt), "] 个其他题")
		buff.WriteString("## 其他题\n\n")
		for i, qtt := range c.qtt {
			buff.WriteString(fmt.Sprintf("%d. %s  \n", i+1, qtt.TIMU))
			buff.WriteString(fmt.Sprintf("%s  \n\n", qtt.ANSWER))
		}
	}

	return buff.Bytes(), nil
}

func getAnswerIndex(answer string) int {
	switch {
	case strings.Contains(answer, "A"):
		return 0
	case strings.Contains(answer, "B"):
		return 1
	case strings.Contains(answer, "C"):
		return 2
	case strings.Contains(answer, "D"):
		return 3
	}
	return 0
}

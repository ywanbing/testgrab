package session

import (
	"strings"

	"github.com/gocolly/colly"
	"testGrab/internal/config"
	"testGrab/internal/constant"
)

func (s *Session) parseContent(element *colly.HTMLElement, courseContent *CourseContent) {
	var timuType constant.TimuType
	element.ForEach(".paper_part_title h2", func(_ int, el *colly.HTMLElement) {
		// XZT
		switch {
		case strings.Contains(el.Text, "选择题") || strings.Contains(el.Text, "单选题"):
			timuType = constant.CHOICE
		case strings.Contains(el.Text, "判断题"):
			timuType = constant.Bool
		case strings.Contains(el.Text, "填空题"):
			timuType = constant.FILL
		case strings.Contains(el.Text, "名词解释题"):
			timuType = constant.NOMENON
		case strings.Contains(el.Text, "简答题"):
			timuType = constant.SIMPLE
		case strings.Contains(el.Text, "计算题"):
			timuType = constant.CALCULATE
		case strings.Contains(el.Text, "综合题"):
			timuType = constant.COMPLEX
		case strings.Contains(el.Text, "论述题"):
			timuType = constant.DISCUSSION
		default:
			timuType = constant.UNKNOWN
		}
	})

	switch timuType {
	case constant.CHOICE:
		getXZT(element, courseContent)
	case constant.Bool:
		getPDT(element, courseContent)
	case constant.FILL:
		getTKT(element, courseContent)
	case constant.SIMPLE:
		getJDT(element, courseContent)
	case constant.NOMENON:
		getMCJST(element, courseContent)
	case constant.CALCULATE:
		getJST(element, courseContent)
	case constant.COMPLEX:
		getZHT(element, courseContent)
	case constant.DISCUSSION:
		getLST(element, courseContent)
	case constant.UNKNOWN:
		getQTT(element, courseContent)
	}
}

// 获取选择题
func getXZT(element *colly.HTMLElement, courseContent *CourseContent) {
	element.ForEach(".item_div", func(_ int, el *colly.HTMLElement) {
		var xzt = &XZT{}
		el.ForEach(".item_title", func(_ int, el *colly.HTMLElement) {
			var imgAddr string
			el.ForEach("p", func(_ int, el *colly.HTMLElement) {
				imgAddr = getImgAddr(el)
			})

			// 没有图片直接获取题目
			if imgAddr == "" {
				xzt.TIMU = strings.SplitN(el.Text, "、", 2)[1]
			} else {
				xzt.TIMU = imgAddr
			}
		})
		el.ForEach(".item_li_option_ul_li", func(_ int, el *colly.HTMLElement) {
			xzt.OPTION = append(xzt.OPTION, el.Text+getImgAddr(el))
		})
		el.ForEach(".item_answer", func(_ int, el *colly.HTMLElement) {
			xzt.ANSWER = el.Text
		})

		hash := xzt.Hash()
		if _, ok := courseContent.selectHash[hash]; ok {
			return
		}
		courseContent.selects = append(courseContent.selects, xzt)
		courseContent.selectHash[hash] = struct{}{}
	})
}

// 获取判断题
func getPDT(element *colly.HTMLElement, courseContent *CourseContent) {
	element.ForEach(".item_div", func(_ int, el *colly.HTMLElement) {
		var pdt = &PDT{}
		el.ForEach(".item_title", func(_ int, el *colly.HTMLElement) {
			var imgAddr string
			el.ForEach("p", func(_ int, el *colly.HTMLElement) {
				imgAddr = getImgAddr(el)
			})

			// 没有图片直接获取题目
			if imgAddr == "" {
				pdt.TIMU = strings.SplitN(el.Text, "、", 2)[1]
			} else {
				pdt.TIMU = imgAddr
			}
		})
		el.ForEach(".item_li_option_ul_li", func(_ int, el *colly.HTMLElement) {
			pdt.OPTION = append(pdt.OPTION, el.Text+getImgAddr(el))
		})
		el.ForEach(".item_answer", func(_ int, el *colly.HTMLElement) {
			pdt.ANSWER = el.Text
		})

		hash := pdt.Hash()
		if _, ok := courseContent.pdtHash[hash]; ok {
			return
		}
		courseContent.pdts = append(courseContent.pdts, pdt)
		courseContent.pdtHash[hash] = struct{}{}
	})
}

// 获取填空题
func getTKT(element *colly.HTMLElement, courseContent *CourseContent) {
	element.ForEach(".item_div", func(_ int, el *colly.HTMLElement) {
		var tkt = &TKT{}
		el.ForEach(".item_title", func(_ int, el *colly.HTMLElement) {
			var imgAddr string
			el.ForEach("p", func(_ int, el *colly.HTMLElement) {
				imgAddr = getImgAddr(el)
			})
			// 没有图片直接获取题目
			if imgAddr == "" {
				tkt.TIMU = strings.SplitN(el.Text, "、", 2)[1]
			} else {
				tkt.TIMU = imgAddr
			}
		})
		el.ForEach(".item_answer", func(_ int, el *colly.HTMLElement) {
			tkt.ANSWER = el.Text + getImgAddr(el)
		})

		if tkt.ANSWER == "" {
			el.ForEach(".div_answer", func(_ int, el *colly.HTMLElement) {
				tkt.ANSWER = el.Text + getImgAddr(el)
			})
		}

		hash := tkt.Hash()
		if _, ok := courseContent.tiankongsHash[hash]; ok {
			return
		}
		courseContent.tiankongs = append(courseContent.tiankongs, tkt)
		courseContent.tiankongsHash[hash] = struct{}{}
	})
}

// 获取简答题
func getJDT(element *colly.HTMLElement, courseContent *CourseContent) {
	element.ForEach(".item_div", func(_ int, el *colly.HTMLElement) {
		var jdt = &JDT{}
		el.ForEach(".item_title", func(_ int, el *colly.HTMLElement) {
			var imgAddr string
			el.ForEach("p", func(_ int, el *colly.HTMLElement) {
				imgAddr = getImgAddr(el)
			})
			// 没有图片直接获取题目
			if imgAddr == "" {
				jdt.TIMU = strings.SplitN(el.Text, "、", 2)[1]
			} else {
				jdt.TIMU = imgAddr
			}
		})
		el.ForEach(".item_answer", func(_ int, el *colly.HTMLElement) {
			var imgAddr = getImgAddr(el)
			jdt.ANSWER = el.Text + imgAddr
		})
		if jdt.ANSWER == "" {
			el.ForEach(".div_answer", func(_ int, el *colly.HTMLElement) {
				jdt.ANSWER = el.Text + getImgAddr(el)
			})
		}

		hash := jdt.Hash()
		if _, ok := courseContent.jdsHash[hash]; ok {
			return
		}
		courseContent.jds = append(courseContent.jds, jdt)
		courseContent.jdsHash[hash] = struct{}{}
	})
}

// 获取名词解析题
func getMCJST(element *colly.HTMLElement, courseContent *CourseContent) {
	element.ForEach(".item_div", func(_ int, el *colly.HTMLElement) {
		var mcjs = &MCJS{}
		el.ForEach(".item_title", func(_ int, el *colly.HTMLElement) {
			var imgAddr string
			el.ForEach("p", func(_ int, el *colly.HTMLElement) {
				imgAddr = getImgAddr(el)
			})
			// 没有图片直接获取题目
			if imgAddr == "" {
				mcjs.TIMU = strings.SplitN(el.Text, "、", 2)[1]
			} else {
				mcjs.TIMU = imgAddr
			}
		})
		el.ForEach(".item_answer", func(_ int, el *colly.HTMLElement) {
			var imgAddr = getImgAddr(el)
			mcjs.ANSWER = el.Text + imgAddr
		})
		if mcjs.ANSWER == "" {
			el.ForEach(".div_answer", func(_ int, el *colly.HTMLElement) {
				mcjs.ANSWER = el.Text + getImgAddr(el)
			})
		}

		hash := mcjs.Hash()
		if _, ok := courseContent.mcjsHash[hash]; ok {
			return
		}
		courseContent.mcjs = append(courseContent.mcjs, mcjs)
		courseContent.mcjsHash[hash] = struct{}{}
	})
}

// 获取计算题
func getJST(element *colly.HTMLElement, courseContent *CourseContent) {
	element.ForEach(".item_div", func(_ int, el *colly.HTMLElement) {
		var jst = &JST{}
		el.ForEach(".item_title", func(_ int, el *colly.HTMLElement) {
			var imgAddr string
			el.ForEach("p", func(_ int, el *colly.HTMLElement) {
				imgAddr = getImgAddr(el)
			})
			// 没有图片直接获取题目
			if imgAddr == "" {
				jst.TIMU = strings.SplitN(el.Text, "、", 2)[1]
			} else {
				jst.TIMU = imgAddr
			}
		})
		el.ForEach(".item_answer", func(_ int, el *colly.HTMLElement) {
			var imgAddr = getImgAddr(el)
			jst.ANSWER = el.Text + imgAddr
		})
		if jst.ANSWER == "" {
			el.ForEach(".div_answer", func(_ int, el *colly.HTMLElement) {
				jst.ANSWER = el.Text + getImgAddr(el)
			})
		}

		hash := jst.Hash()
		if _, ok := courseContent.jssHash[hash]; ok {
			return
		}
		courseContent.jss = append(courseContent.jss, jst)
		courseContent.jssHash[hash] = struct{}{}
	})
}

// 获取综合题
func getZHT(element *colly.HTMLElement, courseContent *CourseContent) {
	element.ForEach(".item_div", func(_ int, el *colly.HTMLElement) {
		var zht = &ZHT{}
		el.ForEach(".item_title", func(_ int, el *colly.HTMLElement) {
			zht.TIMU = strings.SplitN(el.Text, "、", 2)[1]
		})
		el.ForEach(".item_answer", func(_ int, el *colly.HTMLElement) {
			var imgAddr = getImgAddr(el)
			zht.ANSWER = el.Text + imgAddr
		})
		if zht.ANSWER == "" {
			el.ForEach(".div_answer", func(_ int, el *colly.HTMLElement) {
				zht.ANSWER = el.Text + getImgAddr(el)
			})
		}
		hash := zht.Hash()
		if _, ok := courseContent.zhsHash[hash]; ok {
			return
		}
		courseContent.zhs = append(courseContent.zhs, zht)
		courseContent.zhsHash[hash] = struct{}{}
	})
}

// 获取论述题
func getLST(element *colly.HTMLElement, courseContent *CourseContent) {
	element.ForEach(".item_div", func(_ int, el *colly.HTMLElement) {
		var lst = &LST{}
		el.ForEach(".item_title", func(_ int, el *colly.HTMLElement) {
			lst.TIMU = strings.SplitN(el.Text, "、", 2)[1]
		})
		el.ForEach(".item_answer", func(_ int, el *colly.HTMLElement) {
			var imgAddr = getImgAddr(el)
			lst.ANSWER = el.Text + imgAddr
		})
		if lst.ANSWER == "" {
			el.ForEach(".div_answer", func(_ int, el *colly.HTMLElement) {
				lst.ANSWER = el.Text + getImgAddr(el)
			})
		}
		hash := lst.Hash()
		if _, ok := courseContent.lstHash[hash]; ok {
			return
		}
		courseContent.lst = append(courseContent.lst, lst)
		courseContent.lstHash[hash] = struct{}{}
	})
}

// 获取其他题
func getQTT(element *colly.HTMLElement, courseContent *CourseContent) {
	element.ForEach(".item_div", func(_ int, el *colly.HTMLElement) {
		var qtt = &QTT{}
		el.ForEach(".item_title", func(_ int, el *colly.HTMLElement) {
			qtt.TIMU = strings.SplitN(el.Text, "、", 2)[1]
		})
		el.ForEach(".item_answer", func(_ int, el *colly.HTMLElement) {
			var imgAddr = getImgAddr(el)
			qtt.ANSWER = el.Text + imgAddr
		})
		if qtt.ANSWER == "" {
			el.ForEach(".div_answer", func(_ int, el *colly.HTMLElement) {
				qtt.ANSWER = el.Text + getImgAddr(el)
			})
		}
		hash := qtt.Hash()
		if _, ok := courseContent.qttHash[hash]; ok {
			return
		}
		courseContent.qtt = append(courseContent.qtt, qtt)
		courseContent.qttHash[hash] = struct{}{}
	})
}

// getImgAddr 获取图片地址
func getImgAddr(el *colly.HTMLElement) string {
	addr := el.ChildAttr("img", "src")
	if addr != "" {
		return genImgAddr(constant.Domain + addr)
	}
	return addr
}

// genImgAddr 生成MD图片地址
func genImgAddr(addr string) string {
	switch config.GenFileType() {
	case string(constant.FileType_MD), string(constant.FileType_HTML), string(constant.FileType_PDF):
		return "![](" + addr + ")"
	case string(constant.FileType_JSON):
		return addr
	}
	return addr
}

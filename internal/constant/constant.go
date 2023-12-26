package constant

const (
	// LOGIN_URL 登录：POST https://www.zk028.cn/oesc/rs/api/student/auth/login name& pwd
	LOGIN_URL       = "https://www.zk028.cn/oesc/rs/api/student/auth/login"
	HOME_Base_URL   = "https://www.zk028.cn/oesc/student/drill/olex_course_list.jsp?batch_id="
	Course_URL      = "https://www.zk028.cn/oesc/student/drill/olex_drill.jsp?batch_id=%s&course_code=%s"
	Course_Base_URL = "/oesc/student/drill/olex_drill.jsp"

	Domain = "https://www.zk028.cn"
)

type FileType string

const (
	FileType_MD   FileType = "markdown"
	FileType_JSON FileType = "json"
)

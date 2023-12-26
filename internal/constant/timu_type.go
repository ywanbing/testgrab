package constant

type TimuType int

const (
	// CHOICE 选择题
	CHOICE TimuType = iota
	// Bool 判断题
	Bool
	// FILL 填空题
	FILL
	// SIMPLE 简答题
	SIMPLE
	// CALCULATE 计算题
	CALCULATE
	// COMPLEX 综合题
	COMPLEX
)

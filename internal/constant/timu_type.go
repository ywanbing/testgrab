package constant

type TimuType int

const (
	// UNKNOWN 未知题型
	UNKNOWN TimuType = iota
	// CHOICE 选择题
	CHOICE
	// Bool 判断题
	Bool
	// FILL 填空题
	FILL
	// NOMENON 名词解释题
	NOMENON
	// SIMPLE 简答题
	SIMPLE
	// CALCULATE 计算题
	CALCULATE
	// COMPLEX 综合题
	COMPLEX
	// DISCUSSION 论述题
	DISCUSSION
)

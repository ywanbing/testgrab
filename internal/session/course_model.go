package session

import (
	"crypto/sha1"
	"encoding/hex"
)

type Course struct {
	// 课程名称
	Name string

	Index string

	// 课程批次id
	BatchId string

	// 课程id
	Id string
}

type CourseContent struct {
	// 课程名称
	Name string

	selects    []*XZT
	selectHash map[string]struct{}

	pdts    []*PDT
	pdtHash map[string]struct{}

	tiankongs     []*TKT
	tiankongsHash map[string]struct{}

	mcjs     []*MCJS
	mcjsHash map[string]struct{}

	jds     []*JDT
	jdsHash map[string]struct{}

	jss     []*JST
	jssHash map[string]struct{}

	zhs     []*ZHT
	zhsHash map[string]struct{}

	lst     []*LST
	lstHash map[string]struct{}

	qtt     []*QTT
	qttHash map[string]struct{}
}

type XZT struct {
	TIMU   string
	OPTION []string
	ANSWER string
}

// Hash 计算题目的 hash值
func (x *XZT) Hash() string {
	sum := sha1.Sum([]byte(x.TIMU))
	return hex.EncodeToString(sum[:])
}

type PDT struct {
	TIMU   string
	OPTION []string
	ANSWER string
}

// Hash 计算题目的 hash值
func (x *PDT) Hash() string {
	sum := sha1.Sum([]byte(x.TIMU))
	return hex.EncodeToString(sum[:])
}

type TKT struct {
	TIMU   string
	ANSWER string
}

// Hash 计算题目的 hash值
func (x *TKT) Hash() string {
	sum := sha1.Sum([]byte(x.TIMU))
	return hex.EncodeToString(sum[:])
}

type JST struct {
	TIMU   string
	ANSWER string
}

// Hash 计算题目的 hash值
func (x *JST) Hash() string {
	sum := sha1.Sum([]byte(x.TIMU))
	return hex.EncodeToString(sum[:])
}

type JDT struct {
	TIMU   string
	ANSWER string
}

// Hash 计算题目的 hash值
func (x *JDT) Hash() string {
	sum := sha1.Sum([]byte(x.TIMU))
	return hex.EncodeToString(sum[:])
}

type MCJS struct {
	TIMU   string
	ANSWER string
}

// Hash 计算题目的 hash值
func (x *MCJS) Hash() string {
	sum := sha1.Sum([]byte(x.TIMU))
	return hex.EncodeToString(sum[:])
}

type ZHT struct {
	TIMU   string
	ANSWER string
}

// Hash 计算题目的 hash值
func (x *ZHT) Hash() string {
	sum := sha1.Sum([]byte(x.TIMU))
	return hex.EncodeToString(sum[:])
}

type LST struct {
	TIMU   string
	ANSWER string
}

// Hash 计算题目的 hash值
func (x *LST) Hash() string {
	sum := sha1.Sum([]byte(x.TIMU))
	return hex.EncodeToString(sum[:])
}

// QTT 其他未分类题目
type QTT struct {
	TIMU   string
	ANSWER string
}

// Hash 计算题目的 hash值
func (x *QTT) Hash() string {
	sum := sha1.Sum([]byte(x.TIMU))
	return hex.EncodeToString(sum[:])
}

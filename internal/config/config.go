package config

import (
	"os"

	"gopkg.in/yaml.v3"
	"testGrab/internal/constant"
)

var DefCfg *Config

type Config struct {
	Name    string `json:"name" yaml:"name"`         // 用户名
	Pwd     string `json:"pwd" yaml:"pwd"`           // 密码
	BatchID string `json:"batch_id" yaml:"batch_id"` // 批次ID
	LoopNum int    `json:"loop_num" yaml:"loop_num"` // 每套题循环次数
	Simple  bool   `json:"simple" yaml:"simple"`     // 是否简单模式（默认不是简单模式）

	GenFileType string `json:"gen_file_type" yaml:"gen_file_type"` // 生成文件的类型
}

func init() {
	DefCfg = &Config{
		Name:        "",
		Pwd:         "",
		BatchID:     "",
		Simple:      true,
		GenFileType: "markdown",
	}

	err := LoadFile("./config.yaml")
	if err != nil {
		panic(err)
	}

	switch DefCfg.GenFileType {
	case string(constant.FileType_MD), string(constant.FileType_JSON):
	default:
		panic("must specify correct gen_file_type, has [markdown|json]")
	}
}

// LoadFile 加载配置文件
func LoadFile(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, DefCfg)
	return err
}

func GetName() string {
	return DefCfg.Name
}

func GetPwd() string {
	return DefCfg.Pwd
}

func GetBatchId() string {
	return DefCfg.BatchID
}

func GetLoopNum() int {
	return DefCfg.LoopNum
}

func IsSimple() bool {
	return DefCfg.Simple
}

func GenFileType() string {
	return DefCfg.GenFileType
}

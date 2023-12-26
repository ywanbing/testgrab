package common

import (
	"testGrab/internal/config"
	"testGrab/internal/constant"
)

func GetHomeUrl() string {
	return constant.HOME_Base_URL + config.GetBatchId()
}

package service

import (
	"Easy-Gin/config"
	_const "Easy-Gin/const"
	"Easy-Gin/internal/db"
	res "Easy-Gin/pkg/response"
)

// Info 用户信息
func Info(userId string) (res.ResCode, any) {
	// 判断用户名是否存在
	m, err := db.GetUserByUserID(userId)
	if err != nil {
		// 记录日志
		config.GinLOG.Error(err.Error())
		return res.CodeServerBusy, _const.ServerBusy
	}

	// 置空密码
	m.PassWord = ""

	return res.CodeSuccess, m
}

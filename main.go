package main

import (
	"github.com/ArcherDing/EasyDict/models"
)

func init() {
	// 注册数据库
	models.RegisterDB()
}

func main() {
	InitUI()
}

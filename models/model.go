package models

import (
	"os"
	"path"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	// 设置数据库路径
	_DB_NAME = "data/dict.db"
	// 设置数据库名称
	_SQLITE3_DRIVER = "sqlite3"
)

func RegisterDB() {
	// 检查数据库文件
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	// 注册模型
	orm.RegisterModel(new(Dict), new(Map))
	// 注册驱动（“sqlite3” 属于默认注册，此处代码可省略）
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DR_Sqlite)
	// 注册默认数据库
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

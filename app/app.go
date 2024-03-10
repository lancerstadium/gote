package app

import (
	"github.com/lancerstadium/gote/app/model"
	"github.com/lancerstadium/gote/app/router"
)

// Start 启动器方法
func Start() {
	// 1. Connect to MySQL
	model.NewMysql()
	defer func() {
		model.Close()
	}()
	// 2. Open Gin Server and Handle Pages
	router.New()
}

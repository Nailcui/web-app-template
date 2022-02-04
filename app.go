package main

import (
	"app/logger"
	"app/routers"
)

func main() {
	err := routers.Eng.Run(":8080")
	if err != nil {
		logger.Error("启动失败: %s", err.Error())
	}
}

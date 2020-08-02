package www

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"zelda/library/conf"
)

//func formatAsDate(t time.Time) string {
//	year, month, day := t.Date()
//	return fmt.Sprintf("%d%02d/%02d", year, month, day)
//}

func HttpServer() {
	// 1. 读取配置
	conf.Init("development")
	config := conf.GetConfig()
	// 2. 查看相关组建
	// 3. 启动服务
	// router := InitRouter()
	gin.SetMode(conf.GetConfig().RunMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Delims("{[{", "}]}")
	InitRouter(r)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.Gin.HttpPort),
		Handler:        r,
		ReadTimeout:    time.Duration(config.Gin.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.Gin.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	_ = server.ListenAndServe()
}

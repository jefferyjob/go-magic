package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jefferyjob/go-magic/internal/component/bootstrap"
	"github.com/jefferyjob/go-magic/internal/component/kit/conf"
	"log"
	"net/http"
)

var configFile = flag.String("f", "./configs/app.test.yaml", "the config file")

//func init() {
//
//}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("发生错误了", err)
			return
		}
	}()

	flag.Parse()
	fmt.Println(*configFile)

	ctx := context.Background()
	config := conf.InitViper(ctx, *configFile)
	//fmt.Println(config)

	app := bootstrap.NewApp(ctx, config)
	fmt.Println(app.AppEnv, app.Config.Database)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.JSON(http.StatusOK, gin.H{
					"message": "1234567",
				})
			}
		}()
		c.Next()
	}
}

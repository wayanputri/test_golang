package main

import (
	"context"
	"test_backend/app/config"
	db "test_backend/app/database"
	"test_backend/app/router"
	_ "test_backend/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @title						Notes API
// @version					1.0
// @description				Notes API
// @termsOfService				http://swagger.io/terms/
// @contact.name				API Support
// @license.name				Apache 2.0
// @securityDefinitions.basic	BasicAuth
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @host						localhost:8080
// @BasePath					/
func main() {
	r := gin.Default()
	cfg := config.InitConfig()
	client := db.MgoConn(cfg)

	defer client.Disconnect(context.TODO())

	log := logrus.NewEntry(logrus.StandardLogger())
	table := db.MgoCollection(cfg, client)

	r.Use(cors.Default())

	go func() {
		router.InitRouter(client, r, log, table)
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	}()
	select {}
}

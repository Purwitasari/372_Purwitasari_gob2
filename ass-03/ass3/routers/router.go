package routers

import (
	"ass-03/ass3/config"
	"ass-03/ass3/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func StartServer() *gin.Engine {

	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.POST("/AutoReload", inDB.CreateData)
	router.GET("/AutoReload/last", inDB.Getdatas)

	return router
}

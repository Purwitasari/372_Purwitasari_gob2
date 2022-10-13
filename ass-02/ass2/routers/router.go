package routers

import (
	"ass-02/ass2/config"
	"ass-02/ass2/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func StartServer() *gin.Engine {

	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.POST("/orders", inDB.CreateOrders)
	router.GET("/orders", inDB.GetOrders)
	router.PUT("/orders/:order_id", inDB.UpdateOrder)
	router.DELETE("/orders/:order_id", inDB.DeleteOrder)

	return router
}

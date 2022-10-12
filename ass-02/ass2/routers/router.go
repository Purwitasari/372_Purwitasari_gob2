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

	// Create Data
	router.POST("/orders", inDB.CreateOrders)

	// Get All Data
	router.GET("/orders", inDB.GetOrders)

	// Get Data By ID
	// router.GET("/orders/:id", inDB.GetOrder)

	// Update Data
	router.PUT("/orders/:id", inDB.UpdateOrder)

	// Delete Data
	router.DELETE("/orders/:id", inDB.DeleteOrder)

	return router
}

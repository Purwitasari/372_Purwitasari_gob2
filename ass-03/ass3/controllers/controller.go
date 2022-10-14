package controllers

import (
	"ass-03/ass3/structs"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Water       int    `json:"water"`
	WaterStatus string `json:"water_status"`
	Wind        int    `json:"wind"`
	WindStatus  string `json:"wind_status"`
}

func (idb *InDB) GetData(c *gin.Context) {
	var (
		reload structs.Reload
		result gin.H
	)

	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	randomdata1 := (rand.Intn(max-min+1) + min)

	reload.Water = randomdata1

	randomdata2 := (rand.Intn(max-min+1) + min)
	reload.Wind = randomdata2

	if reload.Water <= 5 {
		status1 := "Aman"
		reload.WaterStatus = status1
	} else if reload.Water <= 8 && reload.Water >= 6 {
		status1 := "Siaga"
		reload.WaterStatus = status1
	} else {
		status1 := "Bahaya"
		reload.WaterStatus = status1
	}

	if reload.Wind <= 6 {
		status2 := "Aman"
		reload.WindStatus = status2
	} else if reload.Water <= 15 && reload.Water >= 7 {
		status2 := "Siaga"
		reload.WindStatus = status2
	} else {
		status2 := "Bahaya"
		reload.WindStatus = status2
	}

	var data []Data

	itemdata := Data{
		Water:       reload.Water,
		WaterStatus: reload.WaterStatus,
		Wind:        reload.Wind,
		WindStatus:  reload.WindStatus,
	}

	data = append(data, itemdata)

	err := idb.DB.Create(&reload).Error
	if err != nil {
		result = gin.H{
			"result": "Failed to Creat data",
		}
		c.JSON(http.StatusNoContent, result)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": data,
		})
	}

}

func (idb *InDB) Getdatas(c *gin.Context) {
	var (
		alldata []structs.Reload
		result  gin.H
	)

	idb.DB.Find(&alldata)

	if len(alldata) <= 0 {
		result = gin.H{
			"message": "Data Not Found",
		}
		c.JSON(http.StatusNotFound, result)
	} else {

		result = gin.H{
			"status": alldata,
		}
		c.JSON(http.StatusOK, result)
	}
}

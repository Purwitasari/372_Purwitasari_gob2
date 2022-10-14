package controllers

import (
	"ass-03/ass3/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func (idb *InDB) CreateData(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
	var (
		reload structs.Reload
		result gin.H
	)

	waterS := c.PostForm("water")
	windS := c.PostForm("wind")

	water, _ := strconv.Atoi(waterS)
	wind, _ := strconv.Atoi(windS)

	reload.Water = water
	reload.Wind = wind

	var data []Data

	datainput := Data{
		Water: reload.Water,
		Wind:  reload.Wind,
	}

	data = append(data, datainput)

	err := idb.DB.Create(&reload).Error
	if err != nil {
		result = gin.H{
			"result": "Failed to Create Data",
		}
		c.JSON(http.StatusNoContent, result)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": data,
		})
	}

}

func (idb *InDB) Getdatas(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
	var (
		reload structs.Reload
		all    []structs.Reload
		result gin.H
	)

	idb.DB.Find(&all).Take(&all).Last(&all)

	if len(all) <= 0 {
		result = gin.H{
			"message": "Data Not Found",
		}
		c.JSON(http.StatusNotFound, result)
	} else {
		idb.DB.Find(&reload).Take(&reload).Last(&reload)

		var water_status, wind_status string

		alldata := Data{
			Water: reload.Water,
			Wind:  reload.Wind,
		}

		if alldata.Water <= 5 {
			water_status = "Aman"
		} else if alldata.Water <= 8 && alldata.Water >= 6 {
			water_status = "Siaga"
		} else {
			water_status = "Bahaya"
		}

		if alldata.Wind <= 6 {
			wind_status = "Aman"
		} else if alldata.Wind <= 15 && alldata.Wind >= 7 {
			wind_status = "Siaga"
		} else {
			wind_status = "Bahaya"
		}
		result = gin.H{
			"status":      alldata,
			"WaterStatus": water_status,
			"WindStatus":  wind_status,
		}
		c.JSON(http.StatusOK, result)
	}
}

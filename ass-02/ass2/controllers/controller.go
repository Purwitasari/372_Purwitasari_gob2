package controllers

import (
	"ass-02/ass2/structs"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Data struct {
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

var Datas = []Data{}

// Create new data to the database
func (idb *InDB) CreateOrders(c *gin.Context) {
	var (
		order  structs.Orders
		item   structs.Item
		result gin.H
	)

	customer_name := c.PostForm("customer_name")
	item_code := c.PostForm("item_code")
	description := c.PostForm("description")
	quantityS := c.PostForm("quantity")

	quantity, _ := strconv.Atoi(quantityS)

	item.ItemCode = item_code
	order.CustomerName = customer_name
	order.OrderedAt = time.Now()
	item.Description = description
	item.Quantity = quantity

	var data []Data

	itemdata := Data{
		ItemCode:    item.ItemCode,
		Description: item.Description,
		Quantity:    item.Quantity,
	}

	data = append(data, itemdata)

	err1 := idb.DB.Create(&order).Error
	if err1 != nil {
		result = gin.H{
			"result": "Failed to Create Data",
		}
		c.JSON(http.StatusNoContent, result)
	} else {

		item.Order_Id = order.OrderID

		err2 := idb.DB.Create(&item).Error
		if err2 != nil {
			result = gin.H{
				"result": "Failed to Create Data",
			}
			c.JSON(http.StatusNoContent, result)
		} else {

			c.JSON(http.StatusOK, gin.H{
				"order_id":      order.OrderID,
				"orderedAt":     order.OrderedAt,
				"customer_name": order.CustomerName,
				"items":         data,
			})
		}
	}
}

// update order (id)
func (idb *InDB) UpdateOrder(c *gin.Context) {

	id := c.Param("order_id")

	type Data2 struct {
		LineItemId  int    `json:"lineItemId"`
		ItemCode    string `json:"itemCode"`
		Description string `json:"description"`
		Quantity    int    `json:"quantity"`
	}

	var (
		order    structs.Orders
		item     structs.Item
		newOrder structs.Orders
		newItem  structs.Item
		result   gin.H
	)

	err1 := idb.DB.Where("order_id = ?", id).First(&order).Error
	err2 := idb.DB.Where("Order_Id = ?", id).First(&item).Error
	if err1 != nil || err2 != nil {
		result = gin.H{
			"result": "Data Not Found",
		}
		c.JSON(http.StatusNotFound, result)
	} else {
		orderId, _ := strconv.Atoi(id)
		newOrder.OrderID = orderId
		newOrder.OrderedAt = order.OrderedAt
		customer_name := c.PostForm("customer_name")
		newOrder.CustomerName = customer_name

		item_code := c.PostForm("item_code")
		description := c.PostForm("description")
		quantityS := c.PostForm("quantity")

		quantity, _ := strconv.Atoi(quantityS)

		newItem.ItemCode = item_code
		newItem.Description = description
		newItem.Quantity = quantity

		var data []Data2
		itemdata := Data2{
			LineItemId:  item.Order_Id,
			ItemCode:    newItem.ItemCode,
			Description: newItem.Description,
			Quantity:    newItem.Quantity,
		}

		data = append(data, itemdata)

		// update data
		err := idb.DB.Model(&order).Updates(newOrder).Error
		if err != nil {

			result = gin.H{
				"result": "Update Data Failed",
			}
			c.JSON(http.StatusNoContent, result)
		} else {

			result = gin.H{
				"result":        "Data Updated",
				"orderId":       newOrder.OrderID,
				"orderedAt":     newOrder.OrderedAt,
				"customer_name": newOrder.CustomerName,
				"data":          data,
			}
			c.JSON(http.StatusOK, result)
		}
	}
}

// delete data with {id}
func (idb *InDB) DeleteOrder(c *gin.Context) {

	var (
		order  structs.Orders
		item   structs.Item
		result gin.H
	)

	id := c.Param("order_id")

	err1 := idb.DB.Where("order_id = ?", id).First(&order).Error
	err2 := idb.DB.Where("Order_Id = ?", id).First(&item).Error
	if err1 != nil || err2 != nil {

		result = gin.H{
			"result": "Data Not Found",
		}
		c.JSON(http.StatusNotFound, result)
	} else {
		err3 := idb.DB.Delete(&item).Error
		err4 := idb.DB.Delete(&order).Error
		if err3 != nil || err4 != nil {
			result = gin.H{
				"result": "Delete Data Failed",
			}
			c.JSON(http.StatusBadRequest, result)
		} else {

			result = gin.H{
				"result": "Data Deleted",
				"data":   order,
			}
			c.JSON(http.StatusOK, result)
		}
	}
}

// get all data
func (idb *InDB) GetOrders(c *gin.Context) {

	var (
		orders []structs.Orders
		result gin.H
	)

	idb.DB.Find(&orders)

	if len(orders) <= 0 {
		result = gin.H{
			"result":  nil,
			"count":   0,
			"message": "Data Not Found",
		}
		c.JSON(http.StatusNotFound, result)
	} else {

		result = gin.H{
			"result": orders,
			"count":  len(orders),
		}
		c.JSON(http.StatusOK, result)
	}
}

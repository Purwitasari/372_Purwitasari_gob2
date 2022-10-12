package controllers

import (
	"ass-02/ass2/structs"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ItemData struct {
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

var itemData = []ItemData{}

// Create new data to the database
func (idb *InDB) CreateOrders(c *gin.Context) {
	var (
		order  structs.Orders
		item   structs.Item
		result gin.H
	)

	order_idS := c.PostForm("order_Id")
	item_idS := c.PostForm("item_Id")
	item_code := c.PostForm("item_code")

	customer_name := c.PostForm("customer_name")
	description := c.PostForm("description")
	quantityS := c.PostForm("quantity")

	quantity, _ := strconv.Atoi(quantityS)
	order_id, _ := strconv.Atoi(order_idS)
	item_id, _ := strconv.Atoi(item_idS)

	item.ItemID = item_id
	item.ItemCode = item_code
	item.OrderId = order_id

	order.CustomerName = customer_name
	order.OrderedAt = time.Now()
	item.Description = description
	item.Quantity = quantity

	item.OrderId = order.OrderID

	data := ItemData{
		ItemCode:    item.ItemCode,
		Description: item.Description,
		Quantity:    item.Quantity,
	}

	itemData = append(itemData, data)

	err1 := idb.DB.Create(&order).Error
	err2 := idb.DB.Create(&item).Error
	if err1 != nil || err2 != nil {
		result = gin.H{
			"result": "Failed to Create Data",
		}
		c.JSON(http.StatusNoContent, result)
	} else {

		c.JSON(http.StatusOK, gin.H{
			"orderId":       order.OrderID,
			"orderedAt":     order.OrderedAt,
			"customer_name": order.CustomerName,
			"items":         itemData,
		})
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
			"message": "data not found",
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

// update order (id)
func (idb *InDB) UpdateOrder(c *gin.Context) {

	id := c.Param("orderId")

	var (
		order    structs.Orders
		item     structs.Item
		newOrder structs.Orders
		result   gin.H
	)

	err := idb.DB.Where("id = ?", id).First(&order).Error
	if err != nil {

		result = gin.H{
			"result": "data order not found",
		}
		c.JSON(http.StatusNotFound, result)
	} else {

		err = idb.DB.Where("order_id = ?", id).First(&item).Error
		if err != nil {

			result = gin.H{
				"result": "data item not found",
			}
			c.JSON(http.StatusNotFound, result)

		} else {

			orderId, _ := strconv.Atoi(id)
			customer_name := c.PostForm("customer_name")
			newOrder.OrderID = orderId
			newOrder.CustomerName = customer_name
			newOrder.OrderedAt = order.OrderedAt

			data := ItemData{
				ItemCode:    item.ItemCode,
				Description: item.Description,
				Quantity:    item.Quantity,
			}

			itemData = append(itemData, data)

			// update data
			err = idb.DB.Model(&order).Updates(newOrder).Error
			if err != nil {

				result = gin.H{
					"result": "Update Data Failed",
				}
				c.JSON(http.StatusNoContent, result)
			} else {

				result = gin.H{
					"result":        "Data Updated",
					"orderId":       order.OrderID,
					"orderedAt":     order.OrderedAt,
					"customer_name": order.CustomerName,
					"data":          data,
				}
				c.JSON(http.StatusOK, result)
			}
		}
	}
}

// delete data with {id}
func (idb *InDB) DeleteOrder(c *gin.Context) {

	// init variable
	var (
		order  structs.Orders
		item   structs.Item
		result gin.H
	)

	// param orderId
	orderId := c.Param("orderId")

	// querry check id on order
	err := idb.DB.Where("id = ?", orderId).First(&order).Error
	if err != nil {

		// response data order not found
		result = gin.H{
			"result": "data order not found",
		}
		c.JSON(http.StatusNotFound, result)
	} else {

		// querry check order_id on item
		err = idb.DB.Where("order_id = ?", orderId).First(&item).Error
		if err != nil {

			// response data item not found
			result = gin.H{
				"result": "data item not found",
			}
			c.JSON(http.StatusNotFound, result)
		} else {

			// Delete data on db, table item
			err = idb.DB.Delete(&item).Error
			if err != nil {

				// response delete data item not success
				result = gin.H{
					"result": "Delete item Failed",
				}
				c.JSON(http.StatusBadRequest, result)
			} else {

				// Delete data on db, table order
				err = idb.DB.Delete(&order).Error
				if err != nil {

					// response delete data order not success
					result = gin.H{
						"result": "Delete order Failed",
					}
					c.JSON(http.StatusBadRequest, result)
				} else {

					// request delete data success
					result = gin.H{
						"result": "data deleted successfully",
						"data":   order,
					}
					c.JSON(http.StatusOK, result)
				}
			}
		}
	}
}

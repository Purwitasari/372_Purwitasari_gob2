package structs

import (
	"time"
)

type Orders struct {
	OrderID      int       `json:"order_id" gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	CustomerName string    `json:"customer_name" gorm:"not null; type:varchar(20)"`
	OrderedAt    time.Time `json:"ordered_at"`
}

type Item struct {
	ItemID      int    `json:"item_id" gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	ItemCode    string `json:"item_code" gorm:"not null;type:varchar(10)"`
	Description string `json:"description" gorm:"not null;type:varchar(100)"`
	Quantity    int    `json:"quantity"`
	Order_Id     int    `json:"Order_Id" gorm:"index"`
	Items       []Item `json:"item" gorm:"foreignKey:Order_Id"`
}

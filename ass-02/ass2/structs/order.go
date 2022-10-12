package structs

import (
	"time"
)

type Orders struct {
	OrderID      int       `json:"order_Id" gorm:"primaryKey;not null"`
	CustomerName string    `json:"customer_name" gorm:"not null; type:varchar(20)"`
	OrderedAt    time.Time `json:"ordered_at"`
}

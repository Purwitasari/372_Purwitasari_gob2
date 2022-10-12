package structs

type Item struct {
	ItemID      int    `json:"item_id" gorm:"primaryKey;not null"`
	ItemCode    string `json:"item_code" gorm:"not null;type:varchar(10)"`
	Description string `json:"description" gorm:"not null;type:varchar(100)"`
	Quantity    int    `json:"quantity"`
	OrderId     int    `json:"order_id"`
	Items       []Item `json:"items" gorm:"foreignKey:order_id"`
}

package structs

import "github.com/jinzhu/gorm"

type Reload struct {
	gorm.Model
	Water       int    `json:"water"`
	WaterStatus string `json:"water_status"`
	Wind        int    `json:"wind"`
	WindStatus  string `json:"wind_status"`
}

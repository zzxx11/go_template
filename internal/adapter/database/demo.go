package database

import (
	"go_template/internal/domains"
	"gorm.io/gorm"
)

type DemoInfo struct {
	gorm.Model
	Name string `json:"name"`
}

func (d *DemoInfo) TableName() string {
	return "demo_info"
}

func (d *DemoInfo) FromDomain(do domains.Demo) {
	d.ID = do.ID
	d.Name = do.Name
}

package repository

import (
	"go_template/internal/adapter/database"
	"go_template/internal/domains"
	"gorm.io/gorm"
)

type DemoRepository struct {
	db *gorm.DB
}

func NewDemoRepository(db *gorm.DB) *DemoRepository {
	return &DemoRepository{db: db}
}

func (db *DemoRepository) Create(in domains.Demo) error {
	tx := db.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	d := database.DemoInfo{}
	d.FromDomain(in)

	if err := tx.Create(&d).Error; err != nil {
		tx.Rollback()
	}

	return tx.Commit().Error
}

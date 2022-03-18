package model

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type StoreGorm struct {
	db    *gorm.DB
	store Store
}

func (s *StoreGorm) Open() error {
	var err error

	s.db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", s.db.Name())

	// Migrate the Schema
	s.db.AutoMigrate(&Page{})

	return nil
}

func (s *StoreGorm) AddPage() {
	s.db.Create(&Page{Code: "D55", Price: 200})
	fmt.Println("Page created")
}

func (s *StoreGorm) GetPage() Page {
	var page Page
	s.db.First(&page)
	return page
}

func (s *StoreGorm) GetAllPages() ([]Page, error) {
	var pages []Page
	tx := s.db.Find(&pages)
	if tx.Error != nil {
		return []Page{}, tx.Error
	}
	return pages, nil
}

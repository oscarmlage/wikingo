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

func (s *StoreGorm) AddPage(name string, body string) error {
	var home Page
	home.Name = name
	home.Body = body
	home.Version = 1
	tx := s.db.Create(&home)
	if tx.Error != nil {
		return tx.Error
	}
	fmt.Println("Page created")
	return nil
}

func (s *StoreGorm) CreatePage(page Page) error {
	tx := s.db.Save(&page)
	return tx.Error
}

func (s *StoreGorm) GetPage(name string) (Page, error) {
	var page Page
	tx := s.db.Where("Name = ?", name).First(&page)
	if tx.Error != nil {
		return Page{}, tx.Error
	}
	fmt.Println(page)
	return page, nil
}

func (s *StoreGorm) GetAllPages() ([]Page, error) {
	var pages []Page
	tx := s.db.Find(&pages)
	if tx.Error != nil {
		return []Page{}, tx.Error
	}
	return pages, nil
}

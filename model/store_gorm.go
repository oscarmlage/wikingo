package model

import (
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
	return nil
}

func (s *StoreGorm) CreatePage(page Page) error {
	newversion := new(Page)
	newversion.Name = page.Name
	newversion.Body = page.Body
	newversion.Version = page.Version + 1
	tx := s.db.Create(&newversion)
	return tx.Error
}

func (s *StoreGorm) GetPage(name string) (Page, error) {
	var page Page
	tx := s.db.Where("Name = ?", name).Order("Version desc").First(&page)
	if tx.Error != nil {
		return Page{}, tx.Error
	}
	return page, nil
}

func (s *StoreGorm) GetPageVersion(name string, version string) (Page, error) {
	var page Page
	tx := s.db.Where("Name = ?", name).Where("Version = ?", version).First(&page)
	if tx.Error != nil {
		return Page{}, tx.Error
	}
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

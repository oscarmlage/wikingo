package model

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func (s *Store) Open() error {
	var err error

	s.db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", s.db.Name())

	// Migrate the Schema
	s.db.AutoMigrate(&Page{})

	// Create
	s.addPage()
	return nil
}

func (s *Store) GetPage() Page {
	var page Page
	s.db.First(&page)
	return page
}

func (s *Store) addPage() {
	s.db.Create(&Page{Code: "D55", Price: 200})
	fmt.Println("Page created")
}

func getUser() {
}

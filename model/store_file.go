package model

import (
	"fmt"
)

type StoreFile struct {
	db string
	s  Store
}

func (s *StoreFile) Open() error {
	fmt.Printf("File: Store Open")
	return nil
}

func (s *StoreFile) AddPage() {
	fmt.Println("File: Page created")
}

func (s *StoreFile) GetPage() {
	fmt.Println("File: Get Page")
}

func (s *StoreFile) GetAllPages() {
	fmt.Println("File: Get All Pages")
}

package model

import (
	"fmt"
)

type StoreLab struct {
	db string
	s  Store
}

func (s *StoreLab) Open() {
	fmt.Printf("File: Store Open")
}

func (s *StoreLab) AddPage() {
	fmt.Println("File: Page created")
}

func (s *StoreLab) GetPage() {
	fmt.Println("File: Get Page")
}

func (s *StoreLab) GetAllPages() {
	fmt.Println("File: Get All Pages")
}

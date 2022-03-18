package model

import (
	"fmt"
)

type StoreLab struct {
	db string
	s  Store
}

func (s *StoreLab) Open() {
	fmt.Printf("Lab: Store Open")
}

func (s *StoreLab) AddPage() {
	fmt.Println("Lab: Page created")
}

func (s *StoreLab) GetPage() {
	fmt.Println("Lab: Get Page")
}

func (s *StoreLab) GetAllPages() {
	fmt.Println("Lab: Get All Pages")
}

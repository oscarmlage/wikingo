package model

// Store interface
type Store interface {
	Open() error
	AddPage()
	GetPage() Page
	GetAllPages() ([]Page, error)
}

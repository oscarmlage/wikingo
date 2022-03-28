package model

// Store interface
type Store interface {
	Open() error
	AddPage(string, string) error
	CreatePage(Page) error
	GetPage(string) (Page, error)
	GetPageVersion(string, string) (Page, error)
	GetAllPages() ([]Page, error)
}

package entities

type Book struct {
	ID            int
	Title         string
	Author        string
	ISBN          string
	YearPublished int
	Description   string
	IsBorrowable  bool
}

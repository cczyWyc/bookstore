package store

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrExist    = errors.New("exist")
)

type Book struct {
	Id      string   `json:"id"`      // book id
	Name    string   `json:"name"`    // book name
	Authors []string `json:"authors"` // book authors
	Press   string   `json:"press"`   // press
}

type Store interface {
	Create(*Book) error       // create a new book item
	Update(*Book) error       // update a book item
	Get(string) (Book, error) //get a book item
	GetAll() ([]Book, error)  //get all book message
	Delete(string) error      //delete a book item
}

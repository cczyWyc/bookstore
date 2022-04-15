package store

type Book struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Authors []string `json:"authors"`
	Press   string   `json:"press"`
}

type Store interface {
	Create(*Book) error       // create a new book item
	Update(*Book) error       // update a book item
	Get(string) (Book, error) //get a book item
	GetAll() ([]Book, error)  //get all book message
	Delete(string) error      //delete a book item
}

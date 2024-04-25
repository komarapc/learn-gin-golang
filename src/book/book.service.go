package book

type BookService struct{}

func NewBookService() *BookService {
	return &BookService{}
}

var books = []Book{
	{ID: 1, Title: "Book 1", Author: "Author 1"},
	{ID: 2, Title: "Book 2", Author: "Author 2"},
	{ID: 3, Title: "Book 3", Author: "Author 3"},
	{ID: 4, Title: "Book 4", Author: "Author 4"},
	{ID: 5, Title: "Book 5", Author: "Author 5"},
}

func (s *BookService) GetBooks() []Book {
	return books
}

func (s *BookService) GetBookByID(id int) *Book {
	for _, book := range books {
		if book.ID == id {
			return &book
		}
	}
	return nil
}

func (s *BookService) CreateBook(book Book) Book {
	// get the latest id
	latestID := books[len(books)-1].ID
	book.ID = latestID + 1
	books = append(books, book)
	return book
}

func (s *BookService) UpdateBook(id int, book Book) *Book {
	for i, b := range books {
		if b.ID == id {
			book.ID = b.ID
			books[i] = book
			return &books[i]
		}
	}
	return nil
}

func (s *BookService) DeleteBook(id int) bool {
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			return true
		}
	}
	return false
}

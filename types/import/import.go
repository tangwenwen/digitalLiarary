package _import

import "time"

const (
	DataLayout = "2006-01-02"
	TimeLayout = "2006-01-02 15:04:05"
)

type Book struct {
	ID               int64     `xorm:"id"`
	BookName         string    `json:"book_name"`
	BookAuthor       string    `json:"book_author"`
	BookPress        string    `json:"book_press"`
	BookImpD         string    `json:"book_imp_d"`
	BookReferenceNum string    `json:"book_reference_num"`
	BookIsbn         int64     `json:"book_isbn"`
	AuthorAcademy    string    `json:"author_academy"`
	CreatedTime      time.Time `json:"created_time"`
	UpdateTime       time.Time `json:"update_time"`
	Deleted          int       `json:"deleted"`
}

type SingleImportReq struct {
	BookName         string `json:"book_name"`
	BookAuthor       string `json:"book_author"`
	BookPress        string `json:"book_press"`
	BookImpD         string `json:"book_imp_d"`
	BookReferenceNum string `json:"book_reference_num"`
	BookIsbn         int64  `json:"book_isbn"`
	AuthorAcademy    string `json:"author_academy"`
}



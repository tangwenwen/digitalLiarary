package bookTypes



type UpdateBookInfoReq struct {
	BookId           int64  `json:"book_id"`
	BookName         string `json:"book_name"`
	BookAuthor       string `json:"book_author"`
	BookPress        string `json:"book_press"`
	BookImpD         string `json:"book_imp_d"`
	BookReferenceNum string `json:"book_reference_num"`
	BookIsbn         int64  `json:"book_isbn"`
	AuthorAcademy    string `json:"author_academy"`
}

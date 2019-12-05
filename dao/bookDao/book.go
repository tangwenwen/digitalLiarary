package bookDao

import (
	"DigitalLibrary/plugins/db"
	bookTypes "DigitalLibrary/types/book"
	_import "DigitalLibrary/types/import"
)

func BookIdExist(bookId int64)(bool,error){
	orm, err := db.GetEngine()
	if err != nil {
		return false, err
	}
	bookInfo := new(_import.Book)
	bookInfo.ID = bookId
	return orm.Exist(bookInfo)
}

func BookisDel(bookId int64)(bool,error){
	orm, err := db.GetEngine()
	if err != nil {
		return false, err
	}
	bookInfo := new(_import.Book)
	bookInfo.ID = bookId
	bookInfo.Deleted = 0
	return orm.Exist(bookInfo)
}

func DelBookInfo(bookId int64)error{
	orm, err := db.GetEngine()
	if err != nil {
		return  err
	}
	bookInfo := new(_import.Book)
	bookInfo.Deleted = 1
	_,err=orm.Where("id = ?",bookId).Update(bookInfo)
	return err
}

func UpdateBookInfo(updatingBook *bookTypes.UpdateBookInfoReq)error{
	orm, err := db.GetEngine()
	if err != nil {
		return  err
	}
	bookInfo := new(_import.Book)
	bookInfo.BookName = updatingBook.BookName
	bookInfo.BookAuthor = updatingBook.BookAuthor
	bookInfo.BookPress = updatingBook.BookPress
	bookInfo.BookImpD = updatingBook.BookImpD
	bookInfo.BookReferenceNum = updatingBook.BookReferenceNum
	bookInfo.BookIsbn = updatingBook.BookIsbn
	_,err=orm.Where("id = ?",updatingBook.BookId).Update(bookInfo)
	return err
}

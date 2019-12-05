package bookModel

import (
	"DigitalLibrary/dao/bookDao"
	bookTypes "DigitalLibrary/types/book"
	"errors"
)

func DelBook(bookId int64)error{

	flag,err:=bookDao.BookIdExist(bookId)
	if err!=nil{
		return err
	}
	if flag{
		err = bookDao.DelBookInfo(bookId)
		if err!=nil{
			return err
		}
	}else{
		return errors.New("book is not exist")
	}
	return nil
}


func UpdateBook (updateInfo *bookTypes.UpdateBookInfoReq)error{
	flag,err:=bookDao.BookIdExist(updateInfo.BookId)
	if err!=nil{
		return err
	}
	if flag{
		deletedflag,err:=bookDao.BookisDel(updateInfo.BookId)
		if err!=nil{
			return err
		}
		if deletedflag{
			err = bookDao.UpdateBookInfo(updateInfo)
			if err!=nil{
				return err
			}
		}else{
			return errors.New("book has been deleted")
		}
	}else{
		return errors.New("bookId is not exist")
	}
	return nil
}
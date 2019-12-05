package queryDao

import (
	"DigitalLibrary/plugins/db"
	_import "DigitalLibrary/types/import"
	queryTypes "DigitalLibrary/types/query"
	"fmt"
)

func PreciseQuery(opt queryTypes.QueryOpt) ([]*queryTypes.RespQuery, error) {
	var respBookInfoList = make([]*_import.Book, 0)
	var querySql string

	orm, err := db.GetEngine()
	if err != nil {
		return nil, err
	}
	querySql = "SELECT * FROM book WHERE 1=1 "
	if len(opt.BookName) != 0 {
		querySql += fmt.Sprintf("AND book_name = '%s' ", opt.BookName)
	}
	if len(opt.BookAuthor) != 0 {
		querySql += fmt.Sprintf("AND book_author = '%s' ", opt.BookAuthor)
	}
	if len(opt.AuthorAcademy) != 0 {
		querySql += fmt.Sprintf("AND author_academy = '%s' ", opt.AuthorAcademy)
	}
	querySql += fmt.Sprintf("AND deleted = 0 ")
	querySql += "ORDER BY 'book_imp_d' DESC"
	err = orm.SQL(querySql).Find(&respBookInfoList)
	if err != nil {
		return nil, err
	}
	var queryResult = make([]*queryTypes.RespQuery, len(respBookInfoList))
	for i := 0; i < len(respBookInfoList); i++ {
		queryResult[i] = new(queryTypes.RespQuery)
		queryResult[i].BookId = respBookInfoList[i].ID
		queryResult[i].BookName = respBookInfoList[i].BookName
		queryResult[i].BookAuthor = respBookInfoList[i].BookAuthor
		queryResult[i].BookPress = respBookInfoList[i].BookPress
		queryResult[i].BookImpD = respBookInfoList[i].BookImpD
		queryResult[i].BookReferenceNum = respBookInfoList[i].BookReferenceNum
		queryResult[i].BookIsbn = respBookInfoList[i].BookIsbn
		queryResult[i].AuthorAcademy = respBookInfoList[i].AuthorAcademy
	}
	return queryResult, nil
}

func DimQuery(opt queryTypes.QueryOpt) ([]*queryTypes.RespQuery, error) {
	var respBookInfoList = make([]*_import.Book, 0)
	var querySql string

	orm, err := db.GetEngine()
	if err != nil {
		return nil, err
	}
	querySql = "SELECT * FROM book WHERE 1=1 "
	if len(opt.BookName) != 0 {
		querySql += fmt.Sprintf("AND book_name LIKE '%s%s%s' ", "%", opt.BookName, "%")
	}
	if len(opt.BookAuthor) != 0 {
		querySql += fmt.Sprintf("AND book_author LIKE '%s%s%s' ", "%", opt.BookAuthor, "%")
	}
	if len(opt.AuthorAcademy) != 0 {
		querySql += fmt.Sprintf("AND author_academy LIKE '%s%s%s' ", "%", opt.AuthorAcademy, "%")
	}
	querySql += fmt.Sprintf("AND deleted = 0 ")
	querySql += "ORDER BY 'book_imp_d' DESC"
	err = orm.SQL(querySql).Find(&respBookInfoList)
	if err != nil {
		return nil, err
	}
	var queryResult = make([]*queryTypes.RespQuery, len(respBookInfoList))
	for i := 0; i < len(respBookInfoList); i++ {
		queryResult[i] = new(queryTypes.RespQuery)
		queryResult[i].BookId = respBookInfoList[i].ID
		queryResult[i].BookName = respBookInfoList[i].BookName
		queryResult[i].BookAuthor = respBookInfoList[i].BookAuthor
		queryResult[i].BookPress = respBookInfoList[i].BookPress
		queryResult[i].BookImpD = respBookInfoList[i].BookImpD
		queryResult[i].BookReferenceNum = respBookInfoList[i].BookReferenceNum
		queryResult[i].BookIsbn = respBookInfoList[i].BookIsbn
		queryResult[i].AuthorAcademy = respBookInfoList[i].AuthorAcademy
	}
	return queryResult, nil
}

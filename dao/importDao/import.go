package importDao

import (
	"DigitalLibrary/common"
	"DigitalLibrary/plugins/db"
	_import "DigitalLibrary/types/import"
	"bytes"
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
	"time"
)

func SingleImport(importData *_import.SingleImportReq) error {
	orm, err := db.GetEngine()
	if err != nil {
		return err
	}
	_, err = orm.InsertOne(&_import.Book{
		ID:               common.GetId(common.Book),
		BookName:         importData.BookName,
		BookAuthor:       importData.BookAuthor,
		BookPress:        importData.BookPress,
		BookImpD:         importData.BookImpD,
		BookReferenceNum: importData.BookReferenceNum,
		BookIsbn:         importData.BookIsbn,
		AuthorAcademy:    importData.AuthorAcademy,
		CreatedTime:      time.Now(),
	})
	if err != nil {
		return err
	}
	return nil
}

func GetXlsxData(file *excelize.File) ([][]string, error) {
	var Data [][]string
	Data = file.GetRows("Sheet1")
	if len(Data) == 0 {
		return Data, errors.New("文件内无数据")
	}
	return Data[1:], nil
}

func BookData2DB(Data [][]string) error {
	var sql string
	var ids = make([]int64, 0)
	var fileds bytes.Buffer
	work := common.NewNode(common.Book)
	for i := 0; i < len(Data); i++ {
		id := work.Generate()
		ids = append(ids, id)
	}
	orm, err := db.GetEngine()
	if err != nil {
		return err
	}
	fmt.Println(Data)
	sql = "INSERT INTO book(id,book_name,book_author,book_press,book_imp_d,book_reference_num,book_isbn,author_academy,created_time) VALUES %s"
	for index, val := range Data {
		fileds.WriteString(fmt.Sprintf("(%d", ids[index]))
		for index, cell := range val {
			if index == 5 {
				cellInt, _ := strconv.Atoi(cell)
				fileds.WriteString(fmt.Sprintf(",%d", cellInt))
			} else {
				fileds.WriteString(fmt.Sprintf(",'%s'", cell))
			}
		}
		fileds.WriteString(fmt.Sprintf(",'%s'", time.Now().Format(_import.TimeLayout)))
		fileds.WriteString(fmt.Sprintf("),"))
	}
	sql = fmt.Sprintf(sql, fileds.String()[:fileds.Len()-1])
	_, err = orm.Query(sql)
	if err != nil {
		return err
	}
	return nil
}

package importModel

import (
	"DigitalLibrary/dao/importDao"
	_import "DigitalLibrary/types/import"
	"github.com/360EntSecGroup-Skylar/excelize"
	"mime/multipart"
)

func SingleImport(reqData *_import.SingleImportReq)error{
	err:=importDao.SingleImport(reqData)
	if err!=nil{
		return err
	}
	return nil
}


func FileImport(file multipart.File)error{

	xlsReader, err := excelize.OpenReader(file)
	if err!=nil{
		return err
	}
	xlsData,err:=importDao.GetXlsxData(xlsReader)
	if err!=nil{
		return err
	}
	err=importDao.BookData2DB(xlsData)
	if err!=nil{
		return err
	}
	return nil
}
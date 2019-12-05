package queryModel

import (
	"DigitalLibrary/dao/queryDao"
	queryTypes "DigitalLibrary/types/query"
	"errors"
)

func QueryAllInfo(queryOpt int,opt queryTypes.QueryOpt)([]*queryTypes.RespQuery,error){
	var queryResult []*queryTypes.RespQuery
	var err error
	if queryOpt==1{  //精确
		queryResult,err=queryDao.PreciseQuery(opt)
		if err!=nil{
			return nil,err
		}
	}else if queryOpt==0{ //模糊
		queryResult,err=queryDao.DimQuery(opt)
		if err!=nil{
			return nil,err
		}
	}else{
		return nil,errors.New("wrong opt")
	}
	return queryResult,err
}
package model

import (
	. "github.com/huahuayu/go-gin-app/common/db"
	log "github.com/sirupsen/logrus"
	"time"
)

type (
	TParam struct {
		Id       string    `xorm:"not null pk CHAR(32)" json:"id"`
		Type     string    `xorm:"CHAR(20)" json:"type"`
		Code     string    `xorm:"VARCHAR(200)" json:"code"`
		Desc     string    `xorm:"VARCHAR(200)" json:"desc"`
		CreateAt time.Time `xorm:"created" json:"createAt"`
		UpdateAt time.Time `xorm:"updated" json:"updateAt"`
		DeleteAt time.Time `xorm:"deleted" json:"deleteAt"`
	}
)

func InsertParam(param *TParam) (lastInsertId string, err error) {
	if _, err = DB.Insert(param); err != nil {
		log.Warn(err)
		return "", err
	}
	return param.Id, nil
}

func GetParamByType(paramType string) (param *TParam, exist bool, err error) {
	param = new(TParam)
	exist, err = DB.Where("type = ?", paramType).Get(param)
	if err != nil {
		log.Warn(err.Error())
		return nil, exist, err
	}

	if exist {
		return param, exist, nil
	}

	return nil, exist, nil
}

func FindAllParam() (params *[]TParam, err error) {
	records := make([]TParam, 0)
	err = DB.Asc("type").Find(&records)
	if err != nil {
		return nil, err
	}
	return &records, nil
}

func UpdateParam(param *TParam) (affected int64, err error) {
	param.UpdateAt = time.Now()
	return DB.Cols("type", "code", "desc", "update_at").Where("type = ?", param.Type).Update(param)
}

func DeleteParam(id string) (affected int64, err error) {
	param := new(TParam)
	return DB.ID(id).Delete(param)
}

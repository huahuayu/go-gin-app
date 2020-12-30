package param

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/huahuayu/go-gin-app/global"
	"github.com/huahuayu/go-gin-app/model"
	"github.com/huahuayu/go-gin-app/view"
	"strings"
)

func AddParam(c *gin.Context, req *view.AddParamReq) {
	_, err := AddParamUtil(req.Type, req.Code, req.Desc)
	if err != nil {
		view.ResponseErr(c, global.ErrInternal, err.Error())
		return
	}
	view.ResponseOK(c, nil)
}

func AddParamUtil(paramType, code, desc string) (insertId string, err error) {
	param := new(model.TParam)
	param.Id = strings.Replace(uuid.New().String(), "-", "", -1)
	param.Type = paramType
	param.Code = code
	param.Desc = desc
	return model.InsertParam(param)
}

func GetAllParam(c *gin.Context) {
	params, err := model.FindAllParam()
	if err != nil {
		view.ResponseErr(c, global.ErrInternal, err.Error())
		return
	}
	view.ResponseOK(c, params)
}

func GetParamByType(c *gin.Context, req *view.GetParamByTypeReq) {
	param, err := GetParamByTypeUtil(req.Type)
	if err != nil {
		view.ResponseErr(c, global.ErrInternal, err.Error())
		return
	}
	view.ResponseOK(c, param)
}

func GetParamByTypeUtil(paramType string) (param *model.TParam, err error) {
	param, _, err = model.GetParamByType(paramType)
	return param, err
}

func UpdateParam(c *gin.Context, req *view.UpdateParamReq) {
	if req.Type == "" {
		view.ResponseErr(c, global.ErrInvalidParam, "param type should not be blank")
	}
	affected, err := UpdateParamUtil(req.Type, req.Code, req.Desc)
	if err != nil {
		view.ResponseErr(c, global.ErrInternal, err.Error())
		return
	}

	if affected != 1 {
		view.ResponseErr(c, global.ErrDataNotExist, "")
		return
	}
	view.ResponseOK(c, nil)
}

func UpdateParamUtil(paramType string, code string, desc string) (affected int64, err error) {
	param := new(model.TParam)
	param.Type = paramType
	param.Code = code
	param.Desc = desc
	return model.UpdateParam(param)
}

func DeleteParam(c *gin.Context, req *view.DeleteParamReq) {
	affected, err := model.DeleteParam(req.Id)
	if err != nil {
		view.ResponseErr(c, global.ErrInternal, err.Error())
		return
	}
	if affected == 0 {
		view.ResponseErr(c, global.ErrDataNotExist, "")
		return
	}
	view.ResponseOK(c, nil)
}

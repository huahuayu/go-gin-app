package param

import (
	"github.com/gin-gonic/gin"
	"github.com/huahuayu/go-gin-app/common/request"
	"github.com/huahuayu/go-gin-app/service/param"
	"github.com/huahuayu/go-gin-app/view"
)

func AddParam(c *gin.Context) {
	req := new(view.AddParamReq)
	request.GetReq(c, req)
	param.AddParam(c, req)
}

func GetParamByType(c *gin.Context) {
	req := new(view.GetParamByTypeReq)
	request.GetReq(c, req)
	param.GetParamByType(c, req)
}

func GetAllParam(c *gin.Context) {
	param.GetAllParam(c)
}

func UpdateParam(c *gin.Context) {
	req := new(view.UpdateParamReq)
	request.GetReq(c, req)
	param.UpdateParam(c, req)
}

func DeleteParam(c *gin.Context) {
	req := new(view.DeleteParamReq)
	request.GetReq(c, req)
	param.DeleteParam(c, req)
}

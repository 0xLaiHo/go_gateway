package dto

import (
	"github.com/e421083458/gin_scaffold/public"
	"github.com/gin-gonic/gin"
)

type ServiceDeleteInput struct {
	ID int64 `json:"id" form:"id" comment:"服务ID" example:"56" validate:"required"` //服务ID
}

func (param *ServiceDeleteInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type ServiceListInput struct {
	Info     string `json:"info" form:"info" comment:"关键词" example:"" validate:""`                    //关键词
	PageNo   int    `json:"page_no" form:"page_no" comment:"页数" example:"" validate:"required"`       //页数
	PageSize int    `json:"page_size" form:"page_size" comment:"每页条数" example:"" validate:"required"` //每页条数
}

func (param *ServiceListInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type ServiceListItemOutput struct {
	ID          int64  `json:"id" form:"id"`                       //id
	ServiceName string `service_name:"id" form:"service_name"`     //服务名称
	ServiceDesc string `json:"service_desc" form:"service_desc"`   //服务描述
	LoadType    int    `json:"load_type" form:"load_type"`         //类型
	ServiceAddr string `json:"service_addr " form:"service_addr "` //服务地址
	Qps         int64  `json:"qps" form:"qps"`                     //qps
	Qpd         int64  `json:"qpd" form:"qpd"`                     //qpd
	TotalNode   int    `json:"total_node" form:"total_node"`       //节点数

}
type ServiceListOutput struct {
	Total int64                   `json:"total" form:"total" comment:"总数" example:"" validate:""` //总数
	List  []ServiceListItemOutput `json:"list" form:"list" comment:"列表" example:"" validate:""`   //列表
}

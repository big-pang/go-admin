package formvalidate

import "github.com/gookit/validate"

// UserLevelForm admin_user 表单
type UserLevelForm struct {
	Id          int    `form:"id"`
	Name        string `form:"name" validate:"required"`
	Description string `form:"description" validate:"required"`
	Img         string `form:"img"`
	Status      int8   `form:"status" validate:"int"`
	CreatedAt   int    `form:"created_at"`
	UpdatedAt   int    `form:"updated_at"`
	deleted_at  int    `form:"deleted_at"`
	IsCreate    int    `form:"_create"`
}

// Messages 自定义验证返回消息
func (f UserLevelForm) Messages() map[string]string {
	return validate.MS{
		"Name.required":        "名称不能为空.",
		"description.required": "简介不能为空.",
	}
}

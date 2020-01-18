/*
@Time : 2020/1/16 16:49
@Author : Minus4
*/
package params

type RegisterParam struct {
	Email    string `binding:"email,required"`
	Password string `binding:"max=40,min=1,required"`
}

type AuthParam struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

type AddNewGroupParam struct {
	GroupName string `json:"group_name" binding:"required"`
	Info      string `json:"description"`
}

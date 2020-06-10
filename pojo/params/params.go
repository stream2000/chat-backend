/*
@Time : 2020/1/16 16:49
@Author : Minus4
*/
package params

type AuthParam struct {
	Account  string `binding:"required"`
	Password string `binding:"required"`
}

type AddNewGroupParam struct {
	GroupName string `json:"group_name" binding:"required"`
	Info      string `json:"description"`
}

type Hello struct {
	Name string `json:"name"`
}

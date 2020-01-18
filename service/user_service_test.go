/*
@Time : 2020/1/16 12:12
@Author : Minus4
*/
package service

import (
	"fmt"
	"m4-im/dao"
	"m4-im/pkg/util"
	"testing"
)

func init() {
	fmt.Println("Init Service")
	util.Setup()
	dao.Setup()
}

func TestAuth(t *testing.T) {

}

func TestRegister(t *testing.T) {
	//arg1 := params.RegisterParam{
	//	Email:    "a36646@gmail.com",
	//	Password: "123",
	//}
	//err := Register(arg1)
	//if err != nil {
	//	t.WarpedError(err)
	//}
	//err = Register(arg1)
	//if err == nil {
	//	t.WarpedError("Already")
	//} else {
	//	fmt.Println(err)
	//}
	//ok := Auth(arg1.Email, arg1.Password)
	//if !ok {
	//	t.WarpedError("Can't pass auth")
	//}
}

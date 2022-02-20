package main

import (
	"fmt"
)

type user struct {
	uid			string 	`json:"uid"`
	nickname	string	`json:"nickname"`
	uType       int		`json:"uType"`
}

type anchror	struct{
	uid			string 	`json:"uid"`
	nickname	string	`json:"nickname"`
	uType       int		`json:"uType"`
}

type live interface {
	SetInfo(name string) error
	GetInfo() (interface{}, error)
}

func (u *user) SetInfo(name string) error{
	u.nickname = name
	return nil
}

func (u *user) GetInfo() (interface{}, error){
	//info,_ := jsoniter.MarshalToString(u)
	return u.nickname, nil
}

func (a *anchror) SetInfo(name string) error{
	a.nickname = name
	return nil
}

func (a *anchror) GetInfo() (interface{}, error){
	//info,_ := jsoniter.MarshalToString(a)
	return a.nickname, nil
}

func getLiveFactory(t int, uid string) (l live){
	switch t {
		case 1:
			return &user{}
		case 2:
			return &anchror{}
		break;
	}

	return nil
}

//init 先于main方法执行
func init(){
	fmt.Println(11)
}

func main() {
	fmt.Println(22)
	//接口工厂模式
	uid := "111"
	uid2 := "222"
	u1 := getLiveFactory(1, uid)
	u2 := getLiveFactory(2, uid2)
	u1.SetInfo("Jason")
	u2.SetInfo("Song")
	uInfo1,_ := u1.GetInfo()
	uInfo2,_ := u2.GetInfo()
	fmt.Println("用户1信息:",uInfo1,"用户2信息",uInfo2)

}

package main

import (
	"fmt"
	"reflect"
)

type RequestRStudentStruct struct {
	Id        int64  `protobuf:"varint,8,opt,name=Id" json:"Id,omitempty"`
	Number    string `protobuf:"bytes,1,opt,name=Number" json:"Number,omitempty"`
	Password  string `protobuf:"bytes,2,opt,name=Password" json:"Password,omitempty"`
	RealName  string `protobuf:"bytes,3,opt,name=RealName" json:"RealName,omitempty"`
	Age       int32  `protobuf:"varint,4,opt,name=Age" json:"Age,omitempty"`
	Avatar    string `protobuf:"bytes,5,opt,name=Avatar" json:"Avatar,omitempty"`
	ClassId   int64  `protobuf:"varint,6,opt,name=ClassId" json:"ClassId,omitempty"`
	ClassName string `protobuf:"bytes,7,opt,name=ClassName" json:"ClassName,omitempty"`
}


type Student struct {
	Id        int64  `json:"id" xorm:"int(11) pk not null autoincr"`
	Number    string `json:"number" xorm:"number varchar(30) not null unique"` //统一学号，用作登录
	Password  string `json:"password" xorm:"varchar(255) not null"`
	RealName  string `json:"realname" xorm:"varchar(30)"` //显示名
	Age       int    `json:"age" xorm:"int"`
	Avatar    string `json:"avatar" xorm:"varchar(255)"`                      //头像
	ClassId   int64  `json:"class_id" xorm:"class_id int(11) not null index"` //所属班级
	ClassName string `json:"class_name" xorm:"class_name <-"`
}




func Trans(a1 interface{},a2 interface{}) {

	t := reflect.TypeOf(a1)
	s := reflect.ValueOf(a1)
	for i := 0; i < t.NumField(); i++ {
		//f := t.Field(i)
		//fmt.Println(t.Field(i).Name)
		fmt.Printf("%d: %s %s %v\n", i, t.Field(i).Name,t.Field(i).Type, s.FieldByName(t.Field(i).Name))
	}
}

func Trans2(a1 interface{},a2 interface{}) {
	reflect.TypeOf(a1)
	s := reflect.ValueOf(a1).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

type Human struct {
	oo1 string
	oo2 int64
	oo3 int64
}

func main(){
	a1 := Student{2321,"dsf","dsf","wqwe",21,"dsf",232,"dfses"}
	a2 := Student{2321,"dsf","dsf","wqwe",21,"dsf",232,"dfses"}
	Trans(a1,a2)
	Trans2(a1,a2)

	//t := reflect.ValueOf(Human{}).Type()
	////    h := reflect.New(t).Elem()
	//// new return address pointer
	//h := reflect.New(t).Interface()
	//fmt.Println(h)
	//hh := h.(*Human)
	//fmt.Println(hh)
	//
	//i := Human{"Emp", 25, 120}
	//fmt.Println(reflect.TypeOf(i).Field(0).Type)
	//fmt.Println(reflect.ValueOf(i).Field(1))
}
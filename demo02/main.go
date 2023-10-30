package main

import (
	"fmt"
)
//接口
type Usber interface{
	
	start()
	stop()
}

type computer struct{

}

func (c computer) work (usb Usber){
	if _,ok := usb.(phone);ok{

   usb.start()
	}else{
		usb.stop()
	}
 
}

type phone struct{
    Name string
}

func (p phone) start(){
	fmt.Println(p.Name,"启动")
}


func (p phone) stop(){
	fmt.Println(p.Name,"关机")
}


type camera struct{
Name string
}

func (p camera) start(){
	fmt.Println(p.Name,"启动")
}


func (p camera) stop(){
	fmt.Println(p.Name,"关机")
}

func main(){

	var pho = phone{
		Name: "apple",
}
 
  

	  var usb  Usber  //接口就是一个数据类型
	  usb=pho  //表示手机实现这个接口
	  usb.start()



} 
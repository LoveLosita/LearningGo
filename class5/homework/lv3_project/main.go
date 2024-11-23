package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default()
	//h.GET("/change", ChangeGood)
	h.GET("/showCustomer", ShowCustomerInfo)     //展示顾客信息
	h.GET("/changeCustomer", ChangeCustomerInfo) //改变顾客信息
	h.GET("/addCustomer", AddCustomer)           //添加顾客
	h.GET("/deleteCustomer", DeleteCustomer)     //删除顾客
	h.GET("/changeGood", ChangeGood)             //改变商品信息
	h.GET("/addGood", AddGood)                   //添加商品
	h.GET("/deleteGood", DeleteGood)             //删除商品
	h.GET("/showGood", ShowGoodInfo)             //展示商品信息
	h.GET("/buy", ProcessBuy)                    //处理购买
	h.Spin()
}

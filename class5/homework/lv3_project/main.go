package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default()
	//h.GET("/change", ChangeGood)
	h.GET("/showCustomer", ShowCustomerInfo)     //已经验收没啥问题
	h.GET("/changeCustomer", ChangeCustomerInfo) //已经验收没啥问题
	h.GET("/addCustomer", AddCustomer)           //已经验收没啥问题
	h.GET("/deleteCustomer", DeleteCustomer)     //已经验收没啥问题
	h.GET("/changeGood", ChangeGood)
	h.GET("/addGood", AddGood)
	h.GET("/deleteGood", DeleteGood)
	h.GET("/showGood", ShowGoodInfo)
	h.Spin()
}

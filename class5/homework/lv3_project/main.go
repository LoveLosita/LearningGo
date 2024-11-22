package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default()
	//h.GET("/change", ChangeGood)
	h.GET("/show", ShowCustomerInfo)             //已经验收没啥问题
	h.GET("/changeCustomer", ChangeCustomerInfo) //已经验收没啥问题
	h.GET("/addCustomer", AddCustomer)           //已经验收没啥问题
	h.GET("/deleteCustomer", DeleteCustomer)
	h.Spin()
}

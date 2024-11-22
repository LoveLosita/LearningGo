package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
)

type SingleCustomer struct {
	Name    string  `json:"name"`
	Id      string  `json:"id"`
	Balance float64 `json:"balance"`
}

func ChangeCustomerInfo(ctx context.Context, c *app.RequestContext) {
	customersList, err := ReadCostumers() //读取JSON文件中的顾客信息
	if err != nil {                       //JSON文件读取时如果发生了错误
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		fmt.Println(err)
		return
	}
	targetID := c.Query("id")
	targetName := c.Query("name")
	newID := c.Query("newid")
	newName := c.Query("newname")
	newBalance := c.Query("balance")
	//遍历查找并修改复制过来的结构体列表中的内容
	for index, value := range customersList {
		if (targetID != "" && value.Id == targetID) || (targetName != "" && value.Name == targetName) {
			if newID != "" {
				customersList[index].Id = newID
			}
			if newName != "" {
				customersList[index].Name = newName
			}
			if newBalance != "" {
				customersList[index].Balance, err = strconv.ParseFloat(newBalance, 64)
			}
			c.JSON(consts.StatusOK, customersList[index]) //返回修改后的顾客信息
			break
		}
	}
	err = SaveCustomers(customersList) //保存文件
	if err != nil {                    //JSON文件保存时如果发生了错误
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		fmt.Println(err)
		return
	}
}

func ShowCustomerInfo(ctx context.Context, c *app.RequestContext) {
	customersList, err := ReadCostumers() //读取JSON文件中的顾客信息
	choice := c.Query("type")
	if err != nil { //JSON文件读取时如果发生了错误
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		fmt.Println(err)
		return
	}
	targetID := c.Query("id")
	targetName := c.Query("name")
	if choice == "all" {
		c.JSON(consts.StatusOK, customersList) //展示全部顾客信息
	} else {
		for index, value := range customersList {
			if (targetID != "" && value.Id == targetID) || (targetName != "" && value.Name == targetName) {
				c.JSON(consts.StatusOK, customersList[index]) //展示顾客信息
				break
			}
		}
	}

}

func AddCustomer(ctx context.Context, c *app.RequestContext) {
	customersList, err := ReadCostumers() //读取JSON文件中的顾客信息
	if err != nil {                       //JSON文件读取时如果发生了错误
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		fmt.Println(err)
		return
	}
	id := c.Query("id")
	name := c.Query("name")
	balance := c.Query("balance")
	if balance == "" {
		balance = "0.0"
	}
	transformedBal, err := strconv.ParseFloat(balance, 64)
	if err != nil { //如果转化过程出错
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		fmt.Println(err)
		return
	}
	newCustomer := SingleCustomer{Name: name, Id: id, Balance: transformedBal}
	customersList = append(customersList, newCustomer)
	err = SaveCustomers(customersList)
	c.JSON(consts.StatusOK, newCustomer) //展示最新添加的顾客
}

func DeleteCustomer(ctx context.Context, c *app.RequestContext) {
	customersList, err := ReadCostumers() //读取JSON文件中的顾客信息
	if err != nil {                       //JSON文件读取时如果发生了错误
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		fmt.Println(err)
		return
	}
	targetID := c.Query("id")
	targetName := c.Query("name")
	//遍历查找并修改复制过来的结构体列表中的内容
	for index, value := range customersList {
		if (targetID != "" && value.Id == targetID) || (targetName != "" && value.Name == targetName) {
			slice := customersList
			slice = append(slice[:index], slice[index+1:]...)
			customersList = slice
			c.JSON(consts.StatusOK, customersList) //返回修改后的全部顾客信息
			break
		}
	}
	err = SaveCustomers(customersList) //保存文件
	if err != nil {                    //JSON文件保存时如果发生了错误
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		fmt.Println(err)
		return
	}
}

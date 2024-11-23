package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
)

type SingleGood struct {
	ItemID string `json:"itemID"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

func ChangeGood(ctx context.Context, c *app.RequestContext) {
	goodsList, err := ReadGoods() //读取JSON文件中的顾客信息
	if err != nil {               //JSON文件读取时如果发生了错误
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
	newAmount := c.Query("amount")
	found := false
	//遍历查找并修改复制过来的结构体列表中的内容
	for index, value := range goodsList {
		if (targetID != "" && value.ItemID == targetID) || (targetName != "" && value.Name == targetName) {
			found = true
			if newID != "" {
				goodsList[index].ItemID = newID
			}
			if newName != "" {
				goodsList[index].Name = newName
			}
			if newAmount != "" {
				conversedInt, err := strconv.ParseInt(newAmount, 10, 0)
				if err != nil {
					c.JSON(consts.StatusBadRequest, map[string]string{
						"error": err.Error(),
					})
					fmt.Println(err)
					return
				}
				goodsList[index].Amount = int(conversedInt)
			}
			c.JSON(consts.StatusOK, goodsList[index]) //返回修改后的顾客信息
			break
		}
	}
	if !found {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": "没找到需要修改的商品！",
		})
		fmt.Println("没找到需要修改的商品！")
		return
	}
	err = SaveGoods(goodsList) //保存文件
	if err != nil {            //JSON文件保存时如果发生了错误
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		fmt.Println(err)
		return
	}
}

func AddGood(ctx context.Context, c *app.RequestContext) {
	goodsList, err := ReadGoods() //读取JSON文件中的顾客信息
	if err != nil {               //JSON文件读取时如果发生了错误
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		fmt.Println(err)
		return
	}
	id := c.Query("id")
	name := c.Query("name")
	amount := c.Query("amount")
	if amount == "" {
		amount = "0"
	}
	transformedAmount, err := strconv.ParseInt(amount, 10, 0)
	if err != nil { //如果转化过程出错
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		fmt.Println(err)
		return
	}
	newCustomer := SingleGood{ItemID: id, Name: name, Amount: int(transformedAmount)}
	goodsList = append(goodsList, newCustomer)
	err = SaveGoods(goodsList)
	c.JSON(consts.StatusOK, newCustomer) //展示最新添加的顾客
}

func DeleteGood(ctx context.Context, c *app.RequestContext) {
	goodsList, err := ReadGoods() //读取JSON文件中的顾客信息
	if err != nil {               //JSON文件读取时如果发生了错误
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		fmt.Println(err)
		return
	}
	targetID := c.Query("id")
	targetName := c.Query("name")
	found := false
	//遍历查找并修改复制过来的结构体列表中的内容
	for index, value := range goodsList {
		if (targetID != "" && value.ItemID == targetID) || (targetName != "" && value.Name == targetName) {
			slice := goodsList
			slice = append(slice[:index], slice[index+1:]...)
			goodsList = slice
			c.JSON(consts.StatusOK, goodsList) //返回修改后的全部顾客信息
			found = true
			break
		}
	}
	if !found {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": "没找到需要删除的用户！",
		})
		fmt.Println("没找到需要删除的用户！")
		return
	}
	err = SaveGoods(goodsList) //保存文件
	if err != nil {            //JSON文件保存时如果发生了错误
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		fmt.Println(err)
		return
	}
}

func ShowGoodInfo(ctx context.Context, c *app.RequestContext) {
	goodsList, err := ReadGoods() //读取JSON文件中的顾客信息
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
		c.JSON(consts.StatusOK, goodsList) //展示全部商品信息
	} else {
		found := false
		for index, value := range goodsList {
			if (targetID != "" && value.ItemID == targetID) || (targetName != "" && value.Name == targetName) {
				c.JSON(consts.StatusOK, goodsList[index]) //展示单个商品信息
				found = true
				break
			}
		}
		if !found {
			c.JSON(consts.StatusBadRequest, map[string]string{
				"error": "没找到需要展示的商品！",
			})
			fmt.Println("没找到需要展示的商品！")
			return
		}
	}
}

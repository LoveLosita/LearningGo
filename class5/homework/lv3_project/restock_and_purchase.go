package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
)

func ProcessBuy(ctx context.Context, c *app.RequestContext) {
	userID := c.Query("id")
	buyID := c.Query("buy")
	buyAmount := c.Query("amount")
	userList, err := ReadCustomer() //读取顾客信息
	if userID == "" || buyID == "" || buyAmount == "" {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": "购买失败，参数缺少",
		})
		return
	}
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		fmt.Println(err)
		return
	}
	itemList, err := ReadGoods()
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		fmt.Println(err)
		return
	}
	conversed, err := strconv.ParseInt(buyAmount, 10, 0)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		fmt.Println(err)
		return
	}
	itemIndex := 0
	userIndex := 0
	foundItem := false
	for index, value := range itemList {
		if value.ItemID == buyID {
			foundItem = true
			itemIndex = index
		}
	}
	if !foundItem {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": "商品没找到！",
		})
		return
	}
	foundUser := false
	for index, value := range userList {
		if userID == value.Id {
			foundUser = true
			userIndex = index
		}
	}
	if !foundUser {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": "用户没找到！",
		})
		return
	}
	if itemList[itemIndex].Amount >= int(conversed) {
		itemList[itemIndex].Amount -= int(conversed) //减去库存
	} else {
		c.JSON(consts.StatusOK, map[string]string{
			"购买失败": "库存不足",
		})
		return
	}
	userList[userIndex].Balance += float64(conversed) //买1件东西送1积分
	c.JSON(consts.StatusOK, map[string]string{
		"购买状态": "成功",
		"购买数量": strconv.FormatInt(conversed, 10),
		"获得积分": strconv.FormatInt(conversed, 10),
		"现有积分": strconv.FormatFloat(userList[userIndex].Balance, 'f', -1, 64),
	})
	err = SaveCustomers(userList)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		fmt.Println(err)
		return
	}
	err = SaveGoods(itemList)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		fmt.Println(err)
		return
	}
}

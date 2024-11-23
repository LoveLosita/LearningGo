package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadCustomer() ([]SingleCustomer, error) { //返回一个包含全部顾客的结构体列表
	filename := "customers.json"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()
	// 检测文件是否为空
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}
	if stat.Size() == 0 {
		var emptyCustomer []SingleCustomer
		return emptyCustomer, nil
	}
	var Customers []SingleCustomer
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Customers); err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}
	return Customers, nil
}

func SaveCustomers(newInfo []SingleCustomer) error {
	// 打开文件（如果文件存在）
	// 指定文件名
	fileName := "customers.json"

	// 打开文件，如果文件不存在则创建
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close() // 确保函数结束时关闭文件

	// 创建 JSON 编码器
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // 设置缩进格式（可选）
	// 将数据写入文件
	if err := encoder.Encode(newInfo); err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return fmt.Errorf("error writing JSON to file: %w", err)
	}
	return err
}

func ReadGoods() ([]SingleGood, error) { //返回一个包含全部顾客的结构体列表
	filename := "goods.json"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()
	// 检测文件是否为空
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}
	if stat.Size() == 0 { //如果为空，应该直接返回空文件
		var emptyFile []SingleGood
		return emptyFile, nil
	}
	var Goods []SingleGood
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Goods); err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}
	return Goods, nil
}

func SaveGoods(newInfo []SingleGood) error {
	// 打开文件（如果文件存在）
	// 指定文件名
	fileName := "goods.json"

	// 打开文件，如果文件不存在则创建
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close() // 确保函数结束时关闭文件

	// 创建 JSON 编码器
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // 设置缩进格式（可选）
	// 将数据写入文件
	if err := encoder.Encode(newInfo); err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return fmt.Errorf("error writing JSON to file: %w", err)
	}
	return err
}

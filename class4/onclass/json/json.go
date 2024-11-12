package main

/*
import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := Person{Name: "John Doe", Age: 30}
	jsonData, err := json.Marshal(p)
	if err != nil {
		// 处理错误
	}
	fmt.Println(string(jsonData)) // 输出: {"name":"John Doe","age":30}
}

*/
import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	jsonData := `{"name":"Jane Doe","age":25}`
	var p Person
	err := json.Unmarshal([]byte(jsonData), &p)
	if err != nil {
		// 处理错误
	}
	fmt.Println(p.Name) // 输出: Jane Doe
}

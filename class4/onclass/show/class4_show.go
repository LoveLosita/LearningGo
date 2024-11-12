package main

/*
package main
import "fmt"

	func main() {
		s := "Hello, 世界"
		fmt.Println(s)
		b := []byte(s)
		fmt.Println(b)
		copy(b[7:], "你") // 将 '你' 字符替换进字节切片
		fmt.Println(b)
		s = string(b)
		fmt.Println(s)
	}

package main

import (

	"fmt"
	"math"

)

	func main() {
		//注意传入和返回参数类型一般都是float64（int类型自己强转一下）
		// math.Max math.Min 对比两个数字取最大值或最小值，
		x, y := 1.1, 2.2
		fmt.Println(math.Max(x, y))
		fmt.Println(math.Min(x, y))
		// math.Abs 取绝对值
		z := -1.3
		fmt.Println(math.Abs(z))
		// math.Sqrt 返回x的二次方根
		x = 2.0
		fmt.Println(math.Sqrt(x))
		// math.Pow 返回x^y
		x, y = 2.0, 3.0
		fmt.Println(math.Pow(x, y))
	}

package main

import (

	"fmt"
	"strings"

)

	func main() {
		s := "Hello, 世界"
		// 判断字符串是否包含某个字串
		fmt.Println(strings.Contains(s, "世界"))
		// 判断字符串是否以某个字串开头或结尾
		fmt.Println(strings.HasPrefix(s, "Hello"))
		fmt.Println(strings.HasSuffix(s, "Hello"))
		// 统计字符串中某个字串出现的次数
		fmt.Println(strings.Index(s, "l"))
		// 替换字符串中某个字串为另一个字串
		fmt.Println(strings.Replace(s, "世界", "world", -1))
		// 将字符串全部转换为大写或者小写
		fmt.Println(strings.ToUpper(s))
		fmt.Println(strings.ToLower(s))
		// 将字符串按照某个分隔符分隔为一个切片
		fmt.Println(strings.Split(s, ","))
		// 将一个切片按照某个分隔符拼接为一个字符串
		fmt.Println(strings.Join([]string{"a", "b", "c"}, "-"))
		// 返回将字符串按照空白分割的多个字符串。
		// 如果字符串全部是空白或者是空字符串的话，会返回空切片。
		fmt.Println(strings.Fields(s))
	}

package main

import (

	"fmt"
	"time"

)

	func main() {
		// 获取当前时间
		t := time.Now()
		fmt.Println(t)
		// 获取当前时间的年月日时分秒
		fmt.Println(t.Year())
		fmt.Println(t.Month())
		fmt.Println(t.Day())
		fmt.Println(t.Hour())
		fmt.Println(t.Minute())
		fmt.Println(t.Second())
		// 获取当前时间的星期
		fmt.Println(t.Weekday())
		// 获取当前时间的纳秒
		fmt.Println(t.Nanosecond())
		// 获取当前时间的时区
		fmt.Println(t.Location())
		// 获取当前时间的时间戳
		fmt.Println(t.Unix())
		// 格式化当前时间为字符串
		// 在time包中，格式化字符串使⽤的格式化字符分别是：
		// "%Y"或"2006": 代表年
		// "%m"或"01": 代表⽉
		// "%d"或"02": 代表天
		// "%H"或"15": 代表⼩时
		// "%M"或"04": 代表分钟
		// "%s"或"05": 代表秒
		// 它们类似于格式化字符时的"%f"、"%d"、"%s"等
		// "2006-01-02 15:04:05"这个样式的选择不是随意的
		// 它来源于⼀个叫做Mon Jan 2 15:04:05 -0700 MST 2006的参考时间，
		// 这个时间是Go语⾔的创始⼈Rob Pike在2006年1⽉2⽇下午3点4分5秒时收到的⼀封电⼦邮件的
		// 这个时间恰好包含了所有可能的时间格式，⽽且每个数字都不重复，所以可以作为⼀个通⽤的时
		fmt.Println(t.Format("2006-01-02 15:04:05"))
		// 解析字符串为时间
		t, err := time.Parse("2006-01-02 15:04:05", "2023-11-07 16:28:16")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(t)
		}
		// 计算两个时间的差值
		t1 := time.Date(2023, 11, 7, 16, 28, 16, 0, time.UTC)
		t2 := time.Date(2023, 11, 8, 16, 28, 16, 0, time.UTC)
		d := t1.Sub(t2)
		fmt.Println(d)
		// 判断两个时间是否相等
		fmt.Println(t1.Equal(t2))
		// 判断一个时间是否在另一个时间之前或之后
		fmt.Println(t1.Before(t2))
		fmt.Println(t1.After(t2))
		// 添加或减少一个时间的年月日时分秒
		fmt.Println(t1.Add(time.Hour * 24))
		fmt.Println(t1.Add(-time.Hour * 24))
	}

package main

import "fmt"

	func main() {
		x := 98
		str := string(x)
		fmt.Println(str)
	}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串转换为整型
	str := "12"
	fmt.Println(strconv.Atoi(str))
	// 整型转换为字符串
	x := 12
	fmt.Println(strconv.Itoa(x))
}
*/

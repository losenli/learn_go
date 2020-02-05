package main

/**
* @Author: Mr-Li
* @Date: 2020/2/5
 */

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

/**
os.Args + flag 接收命令行参数
*/

func argsDemo() {
	// 1、返回一个string的slice
	args := os.Args
	var sum int
	if len(args) > 1 {
		// 2、第一个元素是执行文件的名称
		fmt.Println(args[1:])
		for _, num := range args[1:] {
			i, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			sum = sum + i
		}
	}
	fmt.Printf("%v中切片%v所有数字之和是：%v\n", args[0], args[1:], sum)
}

type Person struct {
	name string
	age  int
	sex  string
}

func (p *Person) newPerson(name, sex string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
		sex:  sex,
	}
}

func flagDemo() {
	var p Person
	var sex string
	// 1、flag.Type返回对应类型的指针
	name := flag.String("name", "lzh", "姓名")
	age := flag.Int("age", 23, "年龄")
	// 2、flag.TypeVar返回传进来的值
	flag.StringVar(&sex, "sex", "man", "性别")
	// 3、解析命令行参数
	flag.Parse()
	fmt.Println("-----------------------")
	fmt.Println("非关键字的命令行参数:", flag.Args())
	fmt.Println("非关键字的命令行参数个数:", flag.NArg())
	fmt.Println("关键字的个数:", flag.NFlag())
	fmt.Println("-----------------------")
	// 使用命令行传进来的参数，实例化一个结构体
	people := p.newPerson(*name, sex, *age)
	fmt.Println(people)
}

func main() {
	flagDemo()
}

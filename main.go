package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func borrow() {
	name := ""
	value := 0
	fmt.Println("输入借款对象")
	fmt.Scanf("%s", &name)
	fmt.Println("输入借款数目(单位为元)")
	fmt.Scanf("%d", &value)
	filename := "./csv/" + name + ".csv"

	f, err := os.Create(filename)
	if err != nil {
		return
	}
	f.WriteString("xEFxBBxBF")
	w := csv.NewWriter(f)
	data := [][]string{
		{"借款", string(rune(value))},
	}
	w.WriteAll(data)
	w.Flush()
	fmt.Println("Success!")
}

func payBack() {
	name := ""
	value := 0
	fmt.Println("输入还款对象")
	fmt.Scanf("%s", &name)
	fmt.Println("输入还款数目(单位为元)")
	fmt.Scanf("%d", &value)
	filename := "./csv/" + name + ".csv"

	f, err := ioutil.ReadFile(filename) //借了款才能还款
	if err != nil {
		fmt.Println("都没借款怎么还款 emm")
		return
	}
	f.WriteString("xEFxBBxBF")
	w := csv.NewWriter(f)
	data := [][]string{
		{"还款", string(rune(value))},
	}
	w.WriteAll(data)
	w.Flush()
	fmt.Println("Success!")
}

func query() {
	name := ""
	fmt.Println("输入查询对象")
	fmt.Scanf("%s", &name)
	filename := "./csv/" + name + ".csv"
	f, err := ioutil.ReadFile(filename) //借了款才能还款
	if err != nil {
		fmt.Println("都没有记录你看寄啊 emm")
		return
	}
	r2 := csv.NewReader(strings.NewReader(string(f)))
	data, err := r2.ReadAll()
	if err != nil {
		fmt.Println("读取失败了xwx")
		return
	}
	for i := 0; i < len(data); i++ {
		record := data[i]
		fmt.Println(record[0] + record[1] + "元")
	}
}

func main() {
	i := 0
	for true {
		fmt.Println("输入1借款")
		fmt.Println("输入2还款")
		fmt.Println("输入3查债")
		fmt.Scanln(&i)
		if i == 1 {
			borrow()
		}
		if i == 2 {
			payBack()
		}
		if i == 3 {
			query()
		}
	}
}

package main

import (
	"strconv"
	"fmt"
	"bytes"
	"time"
)

func fabnacci(pos int) int {
	sum := 0
	if pos <= 0 {
		panic("must be a int bigger than 0")
	} else if pos <= 2 {
		sum = 1
	} else {
		ppre, pre := 1, 1
		for i := 2; i < pos; i++ {
			sum = ppre + pre
			ppre, pre = pre, sum
		}
	}
	return sum
}

func implodeInt(arr *[]int, sep string) string {
	// []string, then strings.Join()
	//var arrStr []string
	//for _, v := range *arr {
	//	arrStr = append(arrStr, strconv.Itoa(v))
	//}
	//return strings.Join(arrStr, sep)

	var buffer bytes.Buffer
	for i, v := range *arr {
		buffer.WriteString(strconv.Itoa(v))
		if i != len(*arr)-1 {
			buffer.WriteString(sep)
		}
	}
	return buffer.String()
}

//func implode (arr interface{}, sep string) string {
//	var arrStr []string
//
//	for _, v := range  arr {
//		arrStr = append(arrStr, strconv.Itoa(v))
//	}
//	return strings.Join(arrStr, sep)
//}

func main() {
	var arr []int
	for i := 1; i < 20; i++ {
		arr = append(arr, fabnacci(i))
	}
	//fmt.Println(arr)

	//var arrStr []string
	//for _, v := range arr {
	//	arrStr = append(arrStr, strconv.Itoa(v))
	//}
	//fmt.Println(strings.Join(arrStr, ","))

	fmt.Println(implodeInt(&arr, ","))

	t1 := time.Now()
	fmt.Println(fabnacci(30))
	useTime(t1)
	t2 := time.Now()
	fmt.Println(fibonacci(30))
	useTime(t2)
}

func useTime(preTime time.Time) {
	t := time.Since(preTime)
	fmt.Println("use time of: ")
	ns:=t.Nanoseconds()
	fmt.Printf("%f ms\n", float64(ns)/1000000)
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

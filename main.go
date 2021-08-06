package main

import (
	"fmt"
	"go_zero_mall/tool"
)

func main() {
	//fmt.Printf("time.Now(): %v\n", time.Now().Hour())
	//fmt.Printf("time.Now(): %v\n", time.Now().Minute())
	//fmt.Printf("time.Now(): %v\n", time.Now().Second())
	//fmt.Printf("time.Now(): %v\n", time.Now().YearDay()) //今年过了多少天
	//fmt.Printf("time.Now(): %v\n", time.)
	//nTime := time.Now()
	//yesTime := nTime.AddDate(0,0,-1)
	//logDay := yesTime.Format("20060102")
	//fmt.Printf("logDay: %v\n", logDay)

	//a,b:=tool.GetLastMonth()
	//print(fmt.Sprintf("%s~%s", a, b))
	//a,b:=tool.GetQuarterDay()  //本季度
	//print(fmt.Sprintf("%s~%s", a, b))
	//a,b:=tool.GetWeekDay()  //本季度
	//print(fmt.Sprintf("%s~%s", a, b))

	a, b := tool.GetToday()
	print(fmt.Sprintf("%s~%s\n", a, b))
	fmt.Println(tool.StringTransferToTimeStamp(a))
	fmt.Println(tool.StringTransferToTimeStamp(b))

	// tab := []int{5, 4, 7, 3, 9, 8}

	// for i := len(tab) - 1; i >= 0; i-- {
	// 	fmt.Println(tab[i])
	// 	if tab[i] < 5 {
	// 		tab = append(tab[:i], tab[i+1:]...)
	// 	}
	// 	if tab[i] < 5 {
	// 		tab = append(tab[:i], tab[i+1:]...)
	// 	}
	// }

	// fmt.Printf("%v", tab)
}

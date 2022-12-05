package main

import(
	"fmt"
)

func main()  {
	var name string
	var age int
	fmt.Println("이름을 입력하세요 : ")
	fmt.Scan(&name)
	fmt.Println("나이를 입력하세요 : ")
	fmt.Scan(&age)
	
	fmt.Println("결과 : ")
	fmt.Println("이름 : ", name, "나이 : ", age, "세")




}

// Hello World codz
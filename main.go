package main

import ("fmt")


func main() {
	var a float64
	var b float64


	fmt.Print("a = ")
	
	fmt.Scan(&a)

	 fmt.Print("b = ")
	  
	 fmt.Scan(&b)


	fmt.Println("Сумма = ", a + b) 
	fmt.Println("Разность = ", a - b)
	fmt.Println("Произведение = ",a * b)
	fmt.Println("Частное квадратов = ", a / b,)
}

package main

import "fmt"

func main() {
	//hello world
	fmt.Println("hello, world!")



	//string
	var name string = "example"
	fmt.Println(name)
	//const
	const city = "tokyo"
	fmt.Println(city)


	//var
	var nama = "example"
	fmt.Println(nama)

	//perbandingan
	var a = "perempuan"
	var b = "laki-laki"
	var c = a == b
	fmt.Println(c)

	//operasi matematika
	var result = 10 + 1
	fmt.Println(result)

	var i = 10
	i += 10
	fmt.Println(i)

	var j = 1
	j++
	fmt.Println(j)



	//type declaration
	type str string
	type num int

	var jeneng str = "adit"
	var nomor num = 12
	fmt.Println(jeneng)
	fmt.Println(nomor)


	//boolean
	var smk = true
	fmt.Println("anak smk?", smk)


	//conversion
	var day = "senin"
	var d = day[0]
	var dDay = string(d)
	fmt.Println(day)
	fmt.Println(dDay)

	//type data number
	var	numb int = 1202102039
	fmt.Println(numb)
}
package main

import (
	"fmt"
	"strconv"
	"time"
)

//ex02
var a int
var b string
var c bool

//ex03
var d int = 42
var e string = "James Bond"
var f bool = true

//ex04
type gogo int

var g gogo

//ex05
var h int

func main() {
	// exercise01()
	// exercise02()
	// exercise03()
	// exercise04()
	// exercise05()
	// exercise06()
	// exercise07()
	// exercise08()
	// exercise09()
	// exercise10()
	// exercise11()
	// challange01()
	// exercise12()
	// exercise13()
	// exercise14()
	// exercise15()
	// exercise16()
	// exercise17()
	// exercise18()
	// exercise19()
	// exercise20()
	// exercise21()
	// exercise22()
	// exercise23()
	// exercise24()
	// exercise25()
	// exercise26()
	exercise27()
}

//map
func exercise27() {
	// mymap := make(map[string]int)
	// mymap["Laisson"] = 36
	// mymap["Elaine"] = 32
	// mymap["Vinícius"] = 4
	//or
	mymap := map[string]int{"Laisson": 36, "Elaine": 32, "Vinícius": 3}
	mymap["another"] = 123
	for k, v := range mymap {
		fmt.Printf("%v -> %v\n", k, v)
	}

	delete(mymap, "another")
	fmt.Printf("\nAfter delete\n\n")
	for k, v := range mymap {
		fmt.Printf("%v -> %v\n", k, v)
	}
}

//slice (multi-dimensional)
func exercise26() {
	slice := [][]string{
		{"Laisson", "Silveira", strconv.Itoa(time.Now().Year()-1986) + " anos"},
		{"Elaine", "Silveira", strconv.Itoa(time.Now().Year()-1990) + " anos"},
		{"Vinícius", "Silveira", strconv.Itoa(time.Now().Year()-2018) + " anos"},
	}
	for i, x := range slice {
		fmt.Println(i)
		for _, v := range x {
			fmt.Println("\t", v)
		}
	}
}

//slice (len|cap)
func exercise25() {
	states := make([]string, 26, 26)
	states = []string{"SC", "SP", "PR", "RS", "RJ"}
	fmt.Println("len", len(states))
	fmt.Println("cap", cap(states))
	for i := 0; i < len(states); i++ {
		fmt.Println(states[i])
	}
}

//slice (slicing with append)
func exercise24() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:3], x[len(x)-4:]...)
	fmt.Println(y)
}

//slice (add)
func exercise23() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	x = append(x, 53, 54, 55)
	x = append(x, []int{56, 57, 58, 59, 60}...)
	fmt.Println(x)
}

//slice (slicing)
func exercise22() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2 : len(slice)-1])
}

//slice
func exercise21() {
	slice := make([]int, 5)
	slice[0] = 11
	slice[1] = 21
	slice[2] = 31
	slice[3] = 41
	slice[4] = 51
	fmt.Printf("%T\n", slice)
	for i, v := range slice {
		fmt.Println(i, v)
	}

	//or

	slice2 := []int{11, 21, 31, 41, 51}
	fmt.Printf("%T\n", slice2)
	for i, v := range slice2 {
		fmt.Println(i, v)
	}
}

//array
func exercise20() {
	array := [5]int{11, 22, 33, 44, 55}
	for i, v := range array {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", array)
}

//switch
func exercise19() {
	favoriteSport := "futebol"
	switch favoriteSport {
	case "futebol":
		fmt.Println("favorite sport is", favoriteSport)
		fallthrough
	default:
		fmt.Println("favorite sport is ride bike too")
	}
}

//switch
func exercise18() {
	x := 8
	for {
		switch {
		case (x < 10 || x == 10):
			fmt.Printf("%v is less or equals than ten\n", x)
		case x > 10:
			fmt.Printf("%v is greater than ten\n", x)
		}
		x++
		if x == 12 {
			break
		}
	}
}

//for
func exercise17() {
	x := 8
	for {
		if x < 10 {
			fmt.Printf("%v is less than ten\n", x)
		} else if x == 10 {
			fmt.Printf("%v is equals to ten\n", x)
		} else if x > 10 {
			fmt.Printf("%v is greater than ten\n", x)
			if x == 12 {
				break
			}
		}
		x++
	}
}

//for
func exercise16() {
	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}
}

//for
func exercise15() {
	birthday := 1986
	for {
		if birthday <= time.Now().Year() {
			fmt.Println(birthday)
			birthday++
		} else {
			break
		}
	}
}

//for
func exercise14() {
	birthday := 1986
	for birthday <= time.Now().Year() {
		fmt.Println(birthday)
		birthday++
	}
}

//for
func exercise13() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
		fmt.Println("------------------")
	}
}

//for
func exercise12() {
	for i := 0; i <= 10000; i++ {
		fmt.Println(i)
	}
}

func challange01() {
	fmt.Println("Decimal\t| Hexa\t| Unicode\t| Value")
	for i := 33; i <= 122; i++ {
		fmt.Printf("%d\t| %#x\t| %#U | %v\n", i, i, i, string(i))
	}
}

//constants
func exercise11() {
	const (
		_ = 10 + iota
		a
		b
		c
	)
	fmt.Println(a, b, c)
}

//variables
func exercise10() {
	var s string = `Laisson
	Rangel
		Silveira`

	fmt.Println(s)
}

//bitwise
func exercise09() {
	var a = 100
	fmt.Printf("decimal\t|binário\t|hexadecimal\n")
	fmt.Printf("%d\t|%b\t|%#x\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t|%b\t|%#x\n", b, b, b)
}

//constants
func exercise08() {
	const (
		a     = 10
		b     = "test"
		c int = 11
	)
	fmt.Printf("%v | %T\n", a, a)
	fmt.Printf("%v | %T\n", b, b)
	fmt.Printf("%v | %T\n", c, c)
}

//operators
func exercise07() {
	a := 10
	b := 11
	fmt.Println(a == b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
}

//fmt
func exercise06() {
	a := 100
	fmt.Printf("decimal\t|binário\t|hexadecimal\n")
	fmt.Printf("%d\t|%b\t|%#x\n", a, a, a)
}

//values|type
func exercise05() {
	fmt.Printf("g -> %v|%T", g, g)
	g = 42
	fmt.Printf("g -> %v|%T", g, g)

	h = int(g)
	fmt.Printf("h -> %v|%T", h, h)
}

//values|type
func exercise04() {
	fmt.Printf("g -> %v|%T", g, g)
	g = 42
	fmt.Printf("g -> %v|%T", g, g)
}

//values
func exercise03() {
	s := fmt.Sprintf("d: %v, e: %v, f: %v", d, e, f)
	fmt.Println("s ->", s)
}

//values
func exercise02() {
	fmt.Println("a ->", a)
	fmt.Println("b ->", b)
	fmt.Println("c ->", c)
}

//values
func exercise01() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println("x", x)
	fmt.Println("y", y)
	fmt.Println("z", z)
}

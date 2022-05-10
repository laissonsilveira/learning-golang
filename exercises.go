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
	// exercise01() //values
	// exercise02() //values
	// exercise03() //values
	// exercise04() //values|type
	// exercise05() //values|type
	// exercise06() //fmt
	// exercise07() //operators
	// exercise08() //constants
	// exercise09() //bitwise
	// exercise10() //variables
	// exercise11() //format string
	// challange01()
	// exercise12() //for
	// exercise13() //for
	// exercise14() //for
	// exercise15() //for
	// exercise16() //for
	// exercise17() //for
	// exercise18() //switch
	// exercise19() //switch
	// exercise20() //array
	// exercise21() //slice
	// exercise22() //slice (slicing)
	// exercise23() //slice (add)
	// exercise24() //slice (slicing with append)
	// exercise25() //slice (len|cap)
	// exercise26() //slice (multi-dimensional)
	// exercise27() //map
	// exercise28() //struct
	// exercise29() //struct (with map)
	// exercise30() //struct (embedded struct)
	// exercise31() //func
	// exercise32() //func with defer
	// exercise33() //methods
}

//BEGIN ------ exercise33() --------
func exercise33() {
	p01 := pessoa{"Laisson", "Rangel", 36}
	completename := p01.makeCompleteName()
	fmt.Println(completename, "is", p01.idade, "years old")
}

type pessoa struct {
	nome, sobrenome string
	idade           int
}

func (p pessoa) makeCompleteName() string {
	return fmt.Sprintf("%v %v", p.nome, p.sobrenome)
}

//END ------ exercise33() --------

func exercise32() {
	a := func() {
		fmt.Println("with defer (should is the last one)")
	}
	b := func() {
		fmt.Println("without defer")
	}
	defer a()
	b()
}

//BEGIN ------ exercise31() --------
func exercise31() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println(sum(numbers...))
}

func sum(x ...int) (string, int, string, int) {
	total := 0
	// for i := 0; i < len(x); i++ {
	// 	total += x[i]
	// }
	// OR
	for _, v := range x {
		total += v
	}
	return "result is:", total, " | lenght is: ", len(x)
}

//END ------ exercise31() --------

func exercise30() {
	type veiculo struct {
		porta      int
		cor        string
		radio      []string
		motoristas map[int]string
	}
	type caminhonete struct {
		veiculo
		tracaoNasQuatro bool
	}
	type sedan struct {
		veiculo
		modeloLuxo bool
	}

	motoristas := map[int]string{36: "Laisson", 32: "Elaine"}
	radios := []string{"AM", "FM"}

	c01 := caminhonete{veiculo{2, "Branco", radios, motoristas}, false}
	c02 := caminhonete{veiculo{4, "Preto", radios, motoristas}, true}

	s01 := sedan{veiculo{2, "Vermelho", radios, motoristas}, false}
	s02 := sedan{veiculo{4, "Prata", radios, motoristas}, true}

	fmt.Println(c01, "\n", c02, "\n", s01, "\n", s02)

	fmt.Println("O melhor motorista é:", c01.motoristas[36])
}

func exercise29() {
	type pessoa struct {
		nome      string
		sobrenome string
		sabores   []string
	}

	p01 := pessoa{"Laisson", "Silveira", []string{"Chocolate", "Flocos"}}
	p02 := pessoa{"Elaine", "Machado", []string{"Morango", "Coco"}}

	mymap := map[string]pessoa{
		p01.sobrenome: p01,
		p02.sobrenome: p02,
	}

	for k, v := range mymap {
		fmt.Println(k)
		fmt.Println("\t", v.nome, v.sobrenome)
		for _, s := range v.sabores {
			fmt.Println("\t\t", s)
		}
	}
}

func exercise28() {
	type pessoa struct {
		nome      string
		sobrenome string
		sabores   []string
	}

	p01 := pessoa{"Laisson", "Silveira", []string{"Chocolate", "Flocos"}}

	fmt.Println(p01.nome, p01.sobrenome)
	for _, v := range p01.sabores {
		fmt.Println(v)
	}
}

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

func exercise25() {
	states := make([]string, 26, 26)
	states = []string{"SC", "SP", "PR", "RS", "RJ"}
	fmt.Println("len", len(states))
	fmt.Println("cap", cap(states))
	for i := 0; i < len(states); i++ {
		fmt.Println(states[i])
	}
}

func exercise24() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:3], x[len(x)-4:]...)
	fmt.Println(y)
}

func exercise23() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	x = append(x, 53, 54, 55)
	x = append(x, []int{56, 57, 58, 59, 60}...)
	fmt.Println(x)
}

func exercise22() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2 : len(slice)-1])
}

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

func exercise20() {
	array := [5]int{11, 22, 33, 44, 55}
	for i, v := range array {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", array)
}

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

func exercise16() {
	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}
}

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

func exercise14() {
	birthday := 1986
	for birthday <= time.Now().Year() {
		fmt.Println(birthday)
		birthday++
	}
}

func exercise13() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
		fmt.Println("------------------")
	}
}

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

func exercise11() {
	const (
		_ = 10 + iota
		a
		b
		c
	)
	fmt.Println(a, b, c)
}

func exercise10() {
	var s string = `Laisson
	Rangel
		Silveira`

	fmt.Println(s)
}

func exercise09() {
	var a = 100
	fmt.Printf("decimal\t|binário\t|hexadecimal\n")
	fmt.Printf("%d\t|%b\t|%#x\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t|%b\t|%#x\n", b, b, b)
}

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

func exercise06() {
	a := 100
	fmt.Printf("decimal\t|binário\t|hexadecimal\n")
	fmt.Printf("%d\t|%b\t|%#x\n", a, a, a)
}

func exercise05() {
	fmt.Printf("g -> %v|%T", g, g)
	g = 42
	fmt.Printf("g -> %v|%T", g, g)

	h = int(g)
	fmt.Printf("h -> %v|%T", h, h)
}

func exercise04() {
	fmt.Printf("g -> %v|%T", g, g)
	g = 42
	fmt.Printf("g -> %v|%T", g, g)
}

func exercise03() {
	s := fmt.Sprintf("d: %v, e: %v, f: %v", d, e, f)
	fmt.Println("s ->", s)
}

func exercise02() {
	fmt.Println("a ->", a)
	fmt.Println("b ->", b)
	fmt.Println("c ->", c)
}

func exercise01() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println("x", x)
	fmt.Println("y", y)
	fmt.Println("z", z)
}

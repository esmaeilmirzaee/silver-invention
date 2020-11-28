package main

import (
	"fmt"
	"image"
	"io"
	"math"
	"strings"
	"time"
)

func main() {
	//structLiterals()
	//arrays()
	//sliceLiterals()
	//sliceBounds()
	//sliceLenCap()
	//aNilSlice()
	//makeSlice()
	//slicesOfSlice()
	//appendToSlice()
	//iteratingOverSlice()
	//learnMaps()
	//mutatingMaps()
	//closures()
	//v := Vertx{1,2}
	//v.methodDeclaration()
	//v.methodDeclarationPtr()

	//v := Vertf{3, 4}
	//v.Scale(10)
	//fmt.Println(v.Abs())

	/*
	In general, all methods on a given type should have either
	value or pointer receivers, but not a mixture of both.
	 */
	//v := Vert{3,4}
	//var vAbs float64
	//fmt.Printf("After scaling: %+v %v.\n", v, vAbs)
	//v.ScaleVert(5)
	//vAbs = v.AbsVert()
	//fmt.Printf("After scaling: %+v %v.\n", v, vAbs)
	//interfaceSample()
	//interfaceImplicit()
	//interfaceValues()
	//interfaceNil()
	//emptyInterface()
	//typeAssertion()
	//typeSwitch("Hello")
	//typeSwitch(1)
	//typeSwitch(21.98921)
	//typeSwitch(nil)
	//typeSwitch(false)
	//p := Person{"Esmaeil", 35}
	//fmt.Println(p.stringers())
	//stringersExercise()
	//fmt.Println(errorFunc())
	//errSqrtExercise()
	//bufferReader()
	imagesInterface()

}

func imagesInterface() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

func bufferReader() {
	r := strings.NewReader("Hello World")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)
		fmt.Printf("b[:n] %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

}

// TODO
type ErrNegativeFloat float64

func Sqrt(x float64) (float64, error) {
	return 0, nil
}

func (e ErrNegativeFloat) Error() error {
	//return fmt.Sprintf("Cannot Sqrt negative number: %v", e)
	return nil
}

func errSqrtExercise() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError)Error() string {
	return fmt.Sprintf("@ %v %s", e.When, e.What)
}

func errorFunc() error {
	return &MyError{
		time.Now(),
		"It didn't work.",
	}
}

type IPAddr [4]byte

func (ip IPAddr)String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func stringersExercise() {
	hosts := map[string]IPAddr{
		"googleDNS": {8,8,8,8},
		"loopback": {127,0,0,1},
	}
	for host, ip := range hosts {
		fmt.Printf("(Host:%s, IP:%v)", host, ip)
	}
}

type Person struct {
	Name string
	Age int
}

func (p Person) stringers() string {
	return fmt.Sprintf("%s is %v years old.", p.Name, p.Age)
}

func typeSwitch(a interface{}) {
	switch a.(type) {
	case string: fmt.Println("It's string.")
	case float64: fmt.Println("It's float.")
	case int: fmt.Println("It's integer.")
	case bool: fmt.Println("It's boolean.")
	default: fmt.Println("Check your data.")
	}
}

func typeAssertion() {
	var i interface{} = "Hello"

	fmt.Println(i.(string))
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//fmt.Println(i.(int32)) panic
}

func emptyInterface() {
	var i interface{}

	fmt.Println(i)

	i = 42
	fmt.Println(i)

	i = "Hello"
	fmt.Println(i)

}

type In interface {
	MN()
}

func (t *T) MN() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func interfaceNil() {
	var i In

	var t *T
	i = t
	i.MN()

	i = &T{"Hello"}
	i.MN()
}

type If interface {
	Mf()
}

type F float64

func (f F) Mf() {
	fmt.Println(f)
}

func (t T) Mf() {
	fmt.Println(t)
}

func describe(i If) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func interfaceValues() {
	var i If

	i = &T{"Hello"}
	//fmt.Printf("%t %v", i, T{"Hello"})
	describe(i)
	i.Mf()

	i = F(math.Pi)
	describe(i)
	i.Mf()
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func interfaceImplicit() {
	var i I = T{"Hello"}
	i.M()
}

func interfaceSample() {
	/*
		Interfaces
	*/
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vert{3, 4}

	//a = f
	// Err: Cannot use type(f) as type(a)
	// due to MyFloat does not implement Abser
	a = &v

	//a = v
	// Err: Cannot use type(Vert) as type Abser
	fmt.Println(a.AbsVert())
	fmt.Println(f)
}

type Abser interface {
	AbsVert() float64
}

type MyFloat float64

func (f MyFloat) AbsFloat() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vert struct {
	X, Y float64
}

func (v *Vert) ScaleVert(f float64) {
	v.X = v.X * f
	v.Y *= f
}

func (v *Vert) AbsVert() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}


type Vertf struct {
	x, y float64
}

func (v Vertf) Abs() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

func (v *Vertf) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func (v *Vertx) methodDeclarationPtr() {
	fmt.Println(v.Y * 2, v.X * 2)
}


func (v Vertx) methodDeclaration() {
	fmt.Println(v.X * 2, v.Y * 2)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func closures() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func mutatingMaps() {
	m := map[string]Verts{
		"Bell Labs": {46.0912, -98.1223},
		"Google": {43.91, -100.2309},
	}
	fmt.Println(m["Google"])

	m["Google"] = Verts{12.1111, 12.1111}
	googleMap := m["Google"]
	fmt.Println(googleMap)
	elm, ok := m["Google"]
	fmt.Println(elm, ok)
	elm, ok = m["The Beaver"]
	fmt.Println(elm, ok)

	fmt.Println(m)
	delete(m, "Google")
	fmt.Println(m)

	m["The Beaver"] = Verts{1.111, 1.111}
	fmt.Println(m)
}

type Verts struct {
	Lat, Long float64
}

var m map[string]Verts

func learnMaps() {
	m = make(map[string]Verts)
	m["Bell Labs"] = Verts{
		40.68433,
		-74.39967,
	}
	fmt.Println(m)
}

func iteratingOverSlice() {
	a := []string{"1", "2", "3", "a", "b", "c"}
	for i, v := range a {
		fmt.Printf("[%d=%s]", i, v)
	}
}

func appendToSlice() {
	a := []int{1, 2, 3, 4}
	b := make([]int, 5, 6)
	fmt.Printf("%T %T", a, b)
	a = append(a, b...)
	//a = append(a, b) cannot use b (type []int) as type int in append
	a = append(b, 1)
	b = append(a, 1)
	printSlice(a)
	printSlice(b)
}

func slicesOfSlice() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "00"
	board[0][1] = "01"
	board[0][2] = "02"

	board[1][0] = "10"
	board[1][1] = "11"
	board[1][2] = "12"

	fmt.Println(board)
}

func makeSlice() {
	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)
}

func aNilSlice() {
	var s []int
	if s == nil {
		fmt.Printf("%d %d %v", len(s), cap(s), s)
	}
}

func sliceLenCap() {
	s := []int{1,2,3,4,5,6,7,8,9,10}
	printSlice(s)

	s = s[:]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("Len: %d Cap: %d %v\n", len(s), cap(s), s)
}

func sliceBounds() {
	s := []int{1,2,3,4,5,6,7,8}
	fmt.Println(s)

	s = s[1:6]
	fmt.Println(s)

	s = s[:3]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func sliceLiterals() {
	q := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(q)

	r := []bool{true, true, true, false, false}
	fmt.Println(r)

	s := []struct{
		id int
		name string
	}{
		{1, "Esmaeil"},
		{2, "John"},
		{3, "Sam"},
	}
	fmt.Println(s)
}

func arrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "world"
	fmt.Println(a)
}

type Vertx struct {
	X int
	Y int
}

func structPointer() {
	v := Vertx{1, 2}
	p := &v
	p.X = 1e9 // p.X -> *p.X
	fmt.Println(p.X) // p.X -> *p.X
}

var (
	v1 = Vertx{2, 3}
	v2 = Vertx{X: 1}
	v3 = Vertx{}
	p = &Vertx{1, 2}
)

func structLiterals() {
	fmt.Println(v1, v2, v3, p)
}

func switchWithoutCondition() {
	t := time.Now()
	fmt.Println(t, t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
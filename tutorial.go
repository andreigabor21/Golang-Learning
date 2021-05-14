package main

import (
	"fmt"
	"math"
)

/*func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type the year: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You will be %d years old", 2021 - input)
}*/

/*func main() {
	arr := [3]int{4, 5, 6}
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println(sum)

	arr2D := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arr2D)
}*/

/*func main() { //sum of array
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("How many numbers do you want to sum? ")
	scanner.Scan()
	size, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Not an integer")
		os.Exit(1)
	}

	arr := make([]int, size)
	index := 0
	for index < size {
		fmt.Printf("Read the number from index %d: ", index)
		scanner.Scan()
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Not an integer")
			os.Exit(1)
		}
		arr[index] = number
		index++
	}
	fmt.Println("Your array: ", arr)
	fmt.Println("Now let's sum the numbers")

	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Printf("The sum is: %d", sum)
}*/

/*func main() {
	//var x [5]int = [5]int{1, 2, 3, 4, 5}
	//var s []int = x[1:3]
	//fmt.Println(s[:cap(s)])
	//fmt.Println(s[1:])

	//var a []int = []int{5, 6, 7, 8, 9}
	//a = append(a, 10)
	//fmt.Println(a)

	a := make([]int, 0)
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	fmt.Println(a)
}*/

/*func main() { //14
	var a []int = []int{1, 2, 3, 4, 5, 6, 4}
	//for i := 0; i < len(a); i++ {
	//	fmt.Print(a[i])
	//}

	//for i, element := range a {
	//	for j, element2 := range a {
	//		if element == element2 && i < j{
	//			fmt.Println(element)
	//		}
	//	}
	//}

	for i, element := range a {
		for j := i + 1; j < len(a); j++ {
			if element == a[j] {
				fmt.Println(element)
			}
		}
	}
}*/

/*func main() { //15
	var mp map[string]int = map[string]int{
		"apple": 5,
		"pear": 6,
		"orange": 9,
	}
	mp["apple"] = 7
	mp["new"] = 10
	fmt.Println(mp)
	delete(mp, "new")
	fmt.Println(mp)

	val, ok := mp["apple"]
	fmt.Println(val, ok)
	fmt.Println(len(mp))

	mp2 := make(map[string]int)
	fmt.Println(mp2)
}*/

/*func test(x, y int) (z1, z2 int){
	defer fmt.Println("exit")
	z1 = x + y
	z2 = x - y
	return
}

func main() { //16
	ans1, ans2 := test(2, 3)
	fmt.Println(ans1, ans2)
}*/


/*func test2(myFunc func(int) int) {
	fmt.Println(myFunc(7))
}

func returnFunc(x string) func() {
	return func() {
		fmt.Println(x)
	}
}

func main() { //17
	//test := func(x int) int{
	//	return x * -1
	//}(7)
	//fmt.Println(test)

	test := func(x int) int{
		return x * -1
	}
	test2(test)

	returnFunc("hello")()
}*/


/*func changeFirst(slice []int) {
	slice[0] = 100
}

func main() { //18
	//var x map[string]int = map[string]int{"hello": 3}
	//y := x
	//y["y"] = 100

	//var x [2]int = [2]int{3, 4}
	//y := x
	//y[0] = 100
	//fmt.Println(x, y)

	x := []int{1, 3}
	y := x
	y[0] = 100
	fmt.Println(x, y)

	var s []int = []int{3, 4, 5}
	changeFirst(s)
	fmt.Println(s)
}*/

/*func changeValue(str *string) {
	*str = "changed"
}

func changeValue2(str string) {
	str = "changed!"
}

func main() { //19
	//x := 7
	//y := &x
	//*y = 8
	//fmt.Println(x, y)

	//toChange := "hello"
	//fmt.Println(toChange)
	//changeValue(&toChange)
	//fmt.Println(toChange)

	toChange := "hello"
	var pointer *string = &toChange
	fmt.Println(*pointer, pointer, &pointer, &toChange)
}*/

/*type Point struct {
	x int32
	y int32
}

func changeX(pt Point) {
	pt.x = 100
}

//type Circle struct {
//	radius float64
//	center *Point
//}

type Circle struct {
	radius float64
	*Point
}

func main() {
	//var p1 Point = Point{2, 3}
	//var p2 Point = Point{-5, 7}
	//p1.x = 10
	//fmt.Println(p1.x, p2.y)

	//p1 := Point{x: 3}
	//fmt.Println(p1)
	//changeX(p1)
	//fmt.Println(p1)

	c1 := Circle{4.6, &Point{4, 5}}
	fmt.Println(c1.x)
}*/


/*type Student struct {
	name string
	grades []int
	age int
}

func (s Student) getAge() int {
	return s.age
}

func (s *Student) setAge(age int) {
	s.age = age
}

func (s Student) getAverageGrade() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

func (s Student) getMaxGrade() int {
	curMax := 0
	for _, v := range s.grades {
		if curMax < v {
			curMax = v
		}
	}
	return curMax
}

func main() { //21
	s1 := Student{"Andrei", []int{60,78,83}, 20}
	fmt.Println(s1.getAge())
	s1.setAge(25)
	fmt.Println(s1.getAge())
	average := s1.getAverageGrade()
	fmt.Println(average)
	fmt.Println(s1.getMaxGrade())
}*/

type shape interface {
	area() float64
}

//type shape2 interface {
//	area() float64
//}

type circle struct {
	radius float64
}

type rect struct {
	width float64
	height float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s shape) float64 {
	return s.area()
}

func main() { //22
	c1 := circle{4.5}
	r1 := rect{4, 7}

	shapes := []shape{&c1, &r1}
	fmt.Println(shapes)
	for _, shape := range shapes {
		fmt.Println(getArea(shape))
	}

}


















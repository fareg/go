package main

import "fmt"
import c "calc"
import "time"
//import "encoding/json"

func printTestLabel(label string) {
	fmt.Printf("---- %s ----\n", label)
}

func testArray() {
	printTestLabel("testArray")
	 var x [5]int    
	 x[0] = 10    
	 x[1] = 20    
	 x[2] = 30    
	 x[3] = 40    
	 x[4] = 50    
	 fmt.Println(x) 
}

func testArray2() {
	printTestLabel("testArray2")
	 x := [5]int{
	 	11, 21, 31, 41, 51,
	 }
	 fmt.Println(x) 
}

func testSlice() {
	printTestLabel("testSlice")
	 x := []int{10,20,30}
	 y := append(x, 40, 50)
	 fmt.Println(x, y) 
}

func testLenAndCapicity() {
	printTestLabel("testLenAndCapicity")
	x := make([]int, 2, 5) 
	x[0] = 10 
	x[1] = 20 
	fmt.Println(x) 
	fmt.Println("Length is", len(x)) 
	fmt.Println("Capacity is", cap(x)) 
	x = append(x, 30, 40, 50) 
	fmt.Println(x) 
	fmt.Println("Length is", len(x)) 
	fmt.Println("Capacity is", cap(x)) 
	fmt.Println(x) 
	x = append(x, 60) 
	fmt.Println("Length is", len(x)) 
	fmt.Println("Capacity is", cap(x)) 
	fmt.Println(x)
}

func testIteration() {
	printTestLabel("testIteration")
	 x := []int{10, 20, 30, 40, 50}
	 for k, v := range x {
	 	fmt.Printf("Index: %d Value: %d\n", k, v)
	 }	
	 for _, v := range x {
	 	fmt.Printf("Only value: %d\n", v)
	 }	
	 for k := 0; k<len(x); k++ {
	 	fmt.Printf("Value: %d\n", x[k])
	 }	
}

func testMap() {
	printTestLabel("testMap")
	dict := make(map[string]string)    
	dict["go"] = "Golang"    
	dict["cs"] = "CSharp"    
	dict["rb"] = "Ruby"    
	dict["py"] = "Python"    
	dict["js"] = "JavaScript" 
	for k, v := range dict { 
		fmt.Printf("Key: %s Value: %s\n", k, v)    
	} 
}

type Person struct { 
	FirstName, LastName string 
	Dob time.Time 
	Email, Location     string 
}

func (p Person) getDetails() string {
	return getPersonDetails(p)
}

func getPersonDetails(p Person) string {
	var details string
	details = "[Date of Birth: " + p.Dob.String() + ", Email: " + p.Email + ", Location: " + p.Location +" ]"
	return details
}

func testStruct() {
	printTestLabel("testStruct")
	{
		var p Person
		p.FirstName="Shiju"
		p.LastName="Varghese"
		p.Dob= time.Date(1979, time.February, 17, 0, 0, 0, 0, time.UTC)
		p.Email= "shiju@email.com"
		p.Location="Kochi"	
		fmt.Println("p:", p)
	}
	{
		p:= Person { FirstName :"Shiju", LastName : "Varghese" } 
		fmt.Println("p:", p)
		fmt.Println("p:", getPersonDetails(p))
		fmt.Println("p:" , p.getDetails())
	}
}

/* 
	test pointer
*/

func testPointerFunc1(x int) {
	x = 4
}

func testPointerFunc2(x * int) {
	*x = 4
}

func testPointer() {
	x := 0
	fmt.Println("x=" , x)
	testPointerFunc1(x)
	fmt.Println("x=" , x)
	testPointerFunc2(&x)
	fmt.Println("x=" , x)
}

/*
	Interfaces
*/
func testInterface() {
	printTestLabel("testInterface")
	
}

func main() {
	fmt.Println("Hello world!!!")
	x,y := 3, 3
	fmt.Printf("Add(x+y):%d is the result,Substrct(x-y):%d is the result\n",c.Add(x,y),c.Substract(x,y))
	//fmt.Println(c.Add(x,y))
	//fmt.Println(c.Substract(x,y))
	testArray()
	testArray2()
	testSlice()
	testLenAndCapicity()
	testIteration()
	testMap()
	testStruct()
	testPointer()
	testInterface()
}
package datatypes

import "fmt"

//constant blocks
const (
	first  = iota + 6
	second = 2 << iota
)

//constant blocks
const (
	third = iota
)

//main triggering function inside the go
func main() {
	fmt.Println("Hello from a module")
	//int variables
	var i int
	i = 42
	fmt.Println(i)
	//float variables
	var f float32 = 3.142
	fmt.Println(f)

	//string assignment along with inference := is used to infer by the compiler to
	//infer values from the right hand side value.
	firstName := "Amey"
	fmt.Println(firstName)

	//boolean variables
	//b := true

	//complex variables
	c := complex(3, 4)
	fmt.Println(c)
	r, im := real(c), imag(c)
	fmt.Println(r, im)

	//pointer data types
	var lastName = new(string)
	*lastName = "Amey"
	//dereferencing of the operator
	fmt.Println(*lastName)

	exampleString := "Example"
	//address operator
	ptr := &exampleString
	fmt.Println(ptr, *ptr)
	exampleString = "ToAnotherValue"
	fmt.Println(ptr, *ptr)

	//constants
	const pi = 3.14589
	fmt.Println(pi)

	//constant expressions
	fmt.Println(first, second, third)

	//collections
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr[2])

	arr1 := [3]int{1, 2, 3}
	fmt.Println(arr1)

	arr2 := [3]int{1, 2, 3}
	arr2[0] = 5
	fmt.Println(arr2)

	//slice collection
	//slice looks like almost java arraylist
	arr4 := [3]int{1, 2, 3}

	slice := arr[:]

	fmt.Println(arr4, slice)

	//behind the scenes slice is backed up by array
	slice1 := []int{1, 2, 3}
	//internally the compiler copying the elements to new array with new size
	slice1 = append(slice1, 4)
	fmt.Println(slice1)

	//starting index 1
	s2 := slice1[1:]
	fmt.Println(s2)

	//ending index first two elements
	s3 := slice1[:2]
	fmt.Println(s3)

	//starting from index 1 and to the index 2 ( excluding the index 2)
	s4 := slice[1:2]
	fmt.Println(s4)

	//map collection

	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])
	m["foo"] = 32
	fmt.Println(m["foo"])
	delete(m, "foo")
	fmt.Println(m)

	//struct
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	u.FirstName = "Amey"
	u.LastName = "Rokde"
	u.ID = 1
	fmt.Println(u.FirstName)

	u2 := user{ID: 3, FirstName: "Test", LastName: "Test1"}
	fmt.Println(u2)

}

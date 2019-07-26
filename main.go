package main

import "fmt"

/*
User data structure definition
*/
type User struct {
	ID        int
	FirstName string
	LastName  string
}

/*
HTTPRequest data structure
*/
type HTTPRequest struct {
	Method string
}

func main() {

	fmt.Println("Test")

	//loop to condition clause
	var i int
	for i < 5 {

		fmt.Println("Inside the for loop", i)
		println(i)
		i++
		if i == 2 {
			fmt.Println("Inside the continue block")
			continue
		}
		fmt.Println("Printing", i)

		if i == 3 {
			break
		}

	}

	// loop to condition post clause

	for j := 0; j < 5; j++ {
		println(j)
	}

	//infinite loop
	var k int
	for {
		if k == 5 {
			break
		}
		println(k)
		k++
	}

	//loop for collections
	slice := []int{1, 2, 3}
	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}
	//range
	for i, v := range slice {
		println(i, v)
	}
	//map with range function ( python )
	varMap := map[string]int{"key1": 80, "key2": 90}
	for k, v := range varMap {
		println(k, v)
	}
	for k := range varMap {
		println(k)
	}

	for _, v := range varMap {
		println(v)
	}

	println("Starting web server")

	//panic("Something went wrong")

	println("Web Server Started")

	u1 := User{
		ID:        1,
		FirstName: "Amey",
		LastName:  "Rokde",
	}

	u2 := User{
		ID:        2,
		FirstName: "Paryul",
		LastName:  "Rokde",
	}

	fmt.Println(u1, u2)
	// as struct are values types one can check also for the below condition
	if u1 == u2 {
		println("Same user")
	} else if u1.FirstName == u2.FirstName {
		println("Same user")
	} else {
		println("Different user")
	}

	// switch statements

	r := HTTPRequest{Method: "GET"}

	switch r.Method {
	case "GET":
		{
			println("Recevied the GET request")
		}
	case "DELETE":
		{
			println("Recevied the DELETE request")
		}
	default:
		{
			fmt.Println("Unknown method type", r.Method)
		}
	}

}

/*
func main() {
	fmt.Println("Starting functiona learning module")
	port := 3000
	noOfRetries := 3
	//execute the function
	_, err := startWebServer(port, noOfRetries)
	fmt.Println(err)
	controllers.RegisterController()
	//starting the http server
	http.ListenAndServe(":3000", nil)

}

*/

func startWebServer(port, noOfRetries int) (int, error) {
	fmt.Println("Starting the web server at", port, "with no of retries", noOfRetries)
	fmt.Println("Server started")
	return port, nil
}

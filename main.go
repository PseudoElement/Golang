package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func isHappy(n int) bool {
	slice := make([]int, 0)
	return getPowsSum(&slice, n)
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func HelloController(writer http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(writer, "Bad request - Use GET request!", http.StatusMethodNotAllowed);
	}

	fmt.Fprintf(writer, "Hello, Pavel!")
}

func getPowsSum(slice *[]int, num int) bool {
	numStr := strconv.Itoa(num)
	chars := strings.Split(numStr, "")

	var sum int
	for _, char := range chars {
		digit, _ := strconv.Atoi(char)
		sum += digit * digit
	}

	if contains(*slice, sum) {
		return false
	}

	*slice = append(*slice, sum)

	if sum == 1 {
		return true
	}

	return getPowsSum(slice, sum)
}

func main() {
	slice := make([]int, 0);
	slice = append(slice, 1, 2, 3);
	// filtered := filter(slice, func(num int, ind int) bool {
	// 	return num > 1;
	// })

	api := mux.NewRouter().StrictSlash(true)

	api.HandleFunc("/hello", HelloController).Methods("GET");

	log.Fatal(http.ListenAndServe(":8080", api))
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"example.com/m/src/utils"
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
	//use deps from another module
	filtered := utils.Filter(slice, func(num int, ind int) bool {
		return num > 1;
	})
	fmt.Println("Filtered - ", filtered);
	fmt.Println("Name - ", ControllerName);

	api := mux.NewRouter().StrictSlash(true)

	api.HandleFunc("/hello", HelloController).Methods("GET");

	log.Fatal(http.ListenAndServe(":8080", api))

}

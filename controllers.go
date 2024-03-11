package main

import (
	"fmt"
	"net/http"
)

var ControllerName string = "Bimba";

func HelloController(writer http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(writer, "Bad request - Use GET request!", http.StatusMethodNotAllowed)
		return;
	}

	fmt.Fprintf(writer, "Hello, Pavel!")
}
package main

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func lineByNumber(w http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadFile("dane.txt")
	lines := strings.Split(string(data), "\n")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	lineNumber := req.URL.Query().Get("lineNumber")
	intLineNumber, _ := strconv.Atoi(lineNumber)

	if intLineNumber < 1 || intLineNumber > len(lines) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid line number"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(lines[intLineNumber-1]))
}

func main() {
	http.HandleFunc("/data", lineByNumber)
	http.ListenAndServe(":8080", nil)
}

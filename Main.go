package main

import (
	"io"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/compute", ComputeStats)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	const form = `
	<html>
		<body>
			<form action="compute" method="post" name="stats">
			<input type="text" name="numbers" />
			<input type="submit" value="submit"/>
			</form>
			</html>
		</body>
	`
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, form)
}

func ComputeStats(w http.ResponseWriter, r *http.Request) {
	var input string
	r.ParseForm()
	input = r.Form["numbers"][0]
	inputArr := strings.Split(input, " ")
	var mean float64
	var sum float64 = 0.0
	for _, value := range inputArr {
		intInput, _ := strconv.Atoi(value)
		sum += float64(intInput)
	}
	mean = sum / float64(len(inputArr))
	io.WriteString(w, strconv.FormatFloat(mean, 'f', 4, 64))
}
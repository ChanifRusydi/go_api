package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func main() {
	var a int = 10
	var b int = 5
	fmt.Println("Sum: ", intSum(a, b))
	fmt.Println("Subtraction: ", intSubtraction(a, b))
	fmt.Println("Multiplication: ", intMultiplication(a, b))
	fmt.Println("Division: ", intDivision(a, b))
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// router := http.NewServeMux()
	// router.HandleFunc("GET /sum", sum)
	// router.HandleFunc("GET /subtract", subtract)
	// router.HandleFunc("GET /multiply", multiply)
	// router.HandleFunc("GET /divide", divide)

	// router.HandleFunc("GET /data", hello)
	// router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World!"))
	// })
	// router.HandleFunc("POST /{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	path := r.URL.Path
	// 	parts := strings.Split(path, "/")

	// 	if len(parts) != 2 || r.Method != http.MethodPost {
	// 		http.Error(w, "Invalid request", http.StatusBadRequest)
	// 		return
	// 	}

	// 	id := parts[1]
	// 	fmt.Fprintf(w, `{"id": "%s"}`, id)

	// 	w.Header().Set("Content-Type", "application/json")
	// })
	// log.Fatal(http.ListenAndServe(":3000", router))
}

func intDivision(a, b int) int {
	return a / b
}

func intMultiplication(a, b int) int {
	return a * b
}

func intSubtraction(a, b int) int {
	return a - b
}

func intSum(a, b int) int {
	return a + b
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.ParseForm())
	w.Write([]byte("hello again"))
}

// 	router := http.NewServeMux()
// 	log.Println("Server started on: http://localhost:3000")
// 	h1 := func(w http.ResponseWriter, r *http.Request){
// 		templ := template.Must(template.ParseFiles("button.html"))
// 	}

// 	// NewRouter creates and returns a new router
// 	router.HandleFunc("GET /time", getCurrentTime)
// 	fmt.Println("Get Current Time")
// 	router.HandleFunc("GET /todo", getTodo)

// 	log.Fatal(http.ListenAndServe(":3000", router))
// }

// func getTodo(w http.ResponseWriter, r *http.Request) {
// 	data := map[string][]Todo{
// 		"todos": {
// 			{Id: 1, Name: "Task 1"},
// 			{Id: 2, Name: "Task 2"},
// 			{Id: 3, Name: "Task 3"},
// 		},
// 	}
// 	templ := template.Must(template.ParseFiles("button.html"))
// 	templ.Execute(w, data)
// }

// // getCurrentTime handles requests to the /time endpoint
// func getCurrentTime(w http.ResponseWriter, r *http.Request) {
// 	currentTime := time.Now()
// 	response := map[string]string{
// 		"current_time": currentTime.Format(time.RFC3339),
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

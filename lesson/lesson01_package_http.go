package lesson

import (
	"encoding/json"
	"log"
	"net/http"
)

func Lesson01PackageHTTP() {

	http.HandleFunc("/demo", demoHandler)

	log.Println("Server is starting ...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func demoHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%+v", r)

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]string{
		"message": "Hello World!",
		"info":    "Nguyen Dang Khoa",
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Course", "Lap Trinh Golang")

	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Write(data)

}

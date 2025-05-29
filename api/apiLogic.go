package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SumHandler(w http.ResponseWriter, r *http.Request) {
	var numbers Request
	err := json.NewDecoder(r.Body).Decode(&numbers)
	if err != nil {

		jsonRequest, _ := json.Marshal(numbers)

		fmt.Println("got error:", err.Error(), http.StatusBadRequest, "Request:", string(jsonRequest))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sum := sumFunc(numbers.Numbers)

	fmt.Printf("Sum: %d", sum)

	//message := "Sum:" + strconv.Itoa(sum) + " User:" + numbers.User + "\n"

	response := map[string]interface{}{
		"sum":   sum,
		"user":  numbers.User,
		"error": nil,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sumFunc(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	return arr[0] + sumFunc(arr[1:])
}

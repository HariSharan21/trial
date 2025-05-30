package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SumHandler(w http.ResponseWriter, r *http.Request) {
	var requestParams SumRequest
	err := json.NewDecoder(r.Body).Decode(&requestParams)
	if err != nil {

		jsonRequest, _ := json.Marshal(requestParams)

		fmt.Println("got error:", err.Error(), http.StatusBadRequest, "Request:", string(jsonRequest))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sum := sumFunc(requestParams.Numbers)

	fmt.Printf("Sum: %d", sum)

	//message := "Sum:" + strconv.Itoa(sum) + " User:" + numbers.User + "\n"

	response := SumResponse{
		User: requestParams.User,
		Sum:  int64(sum),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sumFunc(arr []int64) int64 {
	if len(arr) == 1 {
		return arr[0]
	}

	return arr[0] + sumFunc(arr[1:])
}

func DiffHandler(w http.ResponseWriter, r *http.Request) {
	var requestParams DiffRequest
	err := json.NewDecoder(r.Body).Decode(&requestParams)
	if err != nil {
		jsonRequest, _ := json.Marshal(requestParams)

		fmt.Println("got error:", err.Error(), http.StatusBadRequest, "Request:", string(jsonRequest))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	diff := diffFunc(requestParams.MaxNumber, requestParams.Numbers)

	fmt.Printf("Diff: %d", diff)

	response := SumResponse{
		User: requestParams.User,
		Sum:  int64(diff),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func diffFunc(maxNumber int64, arr []int64) int64 {
	if len(arr) == 0 {
		return maxNumber
	}
	return maxNumber - sumFunc(arr)
}

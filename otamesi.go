package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type JankenRequest struct {
	Hand string `json:"hand"`
}

type JankenResponse struct {
	Result string `json:"result"`
}

func rpsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	handValue := r.PostFormValue("hand")

	if handValue != "rock" && handValue != "paper" && handValue != "scissors" {
		http.Error(w, "Invalid hand value", http.StatusBadRequest)
		return
	}

	serverHand := randomHand()
	var result string

	switch handValue {
	case "rock":
		if serverHand == "scissors" {
			result = "You win!"
		} else if serverHand == "paper" {
			result = "You lose."
		} else {
			result = "Draw."
		}
	case "paper":
		if serverHand == "rock" {
			result = "You win!"
		} else if serverHand == "scissors" {
			result = "You lose."
		} else {
			result = "Draw."
		}
	case "scissors":
		if serverHand == "paper" {
			result = "You win!"
		} else if serverHand == "rock" {
			result = "You lose."
		} else {
			result = "Draw."
		}
	}

	response := JankenResponse{Result: result}
	json.NewEncoder(w).Encode(response)
}

func randomHand() string {
	hands := []string{"rock", "paper", "scissors"}
	rand.Seed(time.Now().UnixNano())
	return hands[rand.Intn(len(hands))]
}

func main() {
	http.HandleFunc("/rps", rpsHandler)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

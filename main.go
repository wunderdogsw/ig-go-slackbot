package main

import (
	"fmt" // formatting package

	"os" // Accessing stdinput

	"github.com/joho/godotenv"  // for accessing .env variables
	"github.com/slack-go/slack" //github.com slack library
	"net/http"
)

func wuffwuff(writer http.ResponseWriter, req *http.Request) {
	s, err := slack.SlashCommandParse(req)

	fmt.Println("hello")
	writer.Write([]byte("hello"))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !s.ValidateToken(os.Getenv("SLACK_BOT_TOKEN")) {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	switch s.Command {
	case "/wuffwuff":
		response := fmt.Sprintf("Wuff Wuff!")
		writer.Write([]byte(response))
	default:
		writer.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	http.HandleFunc("/wuffwuff", wuffwuff)
	http.ListenAndServe(":"+port, nil)
}

package main

import (
	"fmt"
	"github.com/slack-go/slack" //github.com slack library
	"net/http"
)

func wuffwuff(writer http.ResponseWriter, req *http.Request) {
	s, err := slack.SlashCommandParse(req)

	fmt.Println("hello")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !s.ValidateToken("xoxb-2636063108304-2715066825696-RuuvlVmWhRWijRYXipwdClLR") {
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
	http.HandleFunc("/wuffwuff", wuffwuff)

	http.ListenAndServe(":5000", nil)
}

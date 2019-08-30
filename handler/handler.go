package handler

import (
	"context"
	"encoding/json"
	greet "kafka_example_golang_client/proto"
	"log"
	"net/http"
	"strconv"
)

func GreetHandler(client greet.GreetingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var greetRequest greet.HelloRequest
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&greetRequest);
		log.Print("response value", greetRequest.Message);
		response, err := client.Greeting(context.Background(), &greet.HelloRequest{Message: greetRequest.Message})
		if err != nil {
			log.Printf("Greeting: %s", response.Greeting)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{\"status\":" + strconv.FormatBool(response.Success)+ "}"))
		}
	})
}

package main

import (
	//pb "LeaugeProdictorAPI/server/proto/predictor"
	pb "github.com/SUPERetnad01/LeaguePro"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"time"
)
const(
	address = "localhost:50051"
	defaultName = "world"
	blueTeam = "100"
	redTeam = "C9"
	year = 2021
)

func PredictMatch(resp http.ResponseWriter,req *http.Request){
	var request pb.PredictMatchRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&request); err != nil {
		respondWithError(resp, http.StatusBadRequest, "Invalid request payload")
		return
	}
	conn,err := grpc.Dial(address,grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//respondWithJSON(resp,http.StatusAccepted,&request)
	client := pb.NewPredictorClient(conn)
	ctx ,cancel := context.WithTimeout(context.Background(),time.Minute * 2)
	defer cancel()
	print("start request\n")
	response, err := client.PredictMatch(ctx,&request)
	print("done with request")
	if err != nil {
		log.Fatalf("could not greet %v",err)
	}
	defer req.Body.Close()
	print("pizza Time")
	respondWithJSON(resp,http.StatusOK,&response)

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}


func main() {
	fmt.Println("Starting the API server...")
	r := mux.NewRouter()
	//r.HandleFunc("/echo", Echo).Methods("POST")
	r.HandleFunc("/predictMatch",PredictMatch).Methods("POST")

	server := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 2 * time.Minute,
		ReadTimeout:  2 * time.Minute,
	}

	log.Fatal(server.ListenAndServe())
}
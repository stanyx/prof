package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"encoding/json"
	
	_ "net/http/pprof"
)

func getPredictions() []float32 {
	predictions := make([]float32, 0, 1000)
	for i := 0; i < 1000; i++ {
		predictions = append(predictions, rand.Float32())
	}
	
	return predictions
}

func main() {
	http.HandleFunc("/calc", func(w http.ResponseWriter, r *http.Request) {
		log.Println("request", r.URL.Query()["index"])
		response := make([]string, 0)
		for i := 0; i < 100; i++ {
			predictions := getPredictions()
			response = append(response, fmt.Sprintf("%+v", predictions))
		}

		json.NewEncoder(w).Encode(&response)
	})
	
//	http.HandleFunc("/debug/pprof/", pprof.Index)
//	http.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
//	http.HandleFunc("/debug/pprof/profile", pprof.Profile)
//	http.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
//	http.HandleFunc("/debug/pprof/trace", pprof.Trace)

	http.ListenAndServe(":8001", nil)
}
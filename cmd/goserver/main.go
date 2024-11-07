package main

import (
	"log"
	"net/http"
    "github.com/nchhillar2004/goserver/web"
)

func main(){
    PORT := ":8080"
    log.Printf("Server is running at http://localhost%s\n", PORT)
    err := http.ListenAndServe(PORT, web.Routes())
    if err!=nil {
        log.Fatalf("Server Error: %v", err)
    }
}

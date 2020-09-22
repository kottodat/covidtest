package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"crypto/hmac"
)

func main() {
	http.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	perKeyText := "iIx1bjZWBMjZH7S2bTFNfA==.2650847.144.1,jCKFeuly8lL2un9/he4y6g==.2650847.144.1,GJI3Die5cz8P/KS+jJVl1Q==.2650847.144.2,OZ+UM1ZhQDf7UaMsLQ1w+Q==.2650847.144.3,GDu718me2VQrHjzvIwdOVQ==.2650847.144.4,Rhau+MdCj8C/KonmRWq+/A==.2650847.144.5,hqXmF91FNfFwgjE1TU08pQ==.2650847.144.6,daE3pGG7q8/BlGqWuxLzMw==.2650847.144.1,iN6ebfSQVZivjZBfy5JQpw==.2650847.144.1,Rf3Oq2HddsoXPsRqmLrx3g==.2650847.144.2,pJYLqxATiuomJDd0K5HHCg==.2650847.144.3,VGSeVbnC2Zif6QqXAE3fSg==.2650847.144.4,FXWTwnIIIXv2jWqrTMnK6A==.2650847.144.5,v3jCN8h2YoSGi6SUVQ55xg==.2650847.144.6"
	log.Printf("Listening on port %s", perKeyText)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
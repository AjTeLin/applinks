//
// webserver.go
//
// An example of a golang web server.
//

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"log"
	"net/http"
)

func SetMyCookie(response http.ResponseWriter){
	// Add a simplistic cookie to the response.
	cookie := http.Cookie{Name: "testcookiename", Value:"testcookievalue"}
	http.SetCookie(response, &cookie)
}

// Respond to the URL /home with an html home page
func HomeHandler(response http.ResponseWriter, request *http.Request){
	response.Header().Set("Content-type", "text/html")
	webpage, err := ioutil.ReadFile("home.html")
	if err != nil { 
		http.Error(response, fmt.Sprintf("home.html file error %v", err), 500)
	}
	fmt.Fprint(response, string(webpage));
}

func JsonHandler(response http.ResponseWriter, request *http.Request){
	// Set cookie and MIME type in the HTTP headers.
	SetMyCookie(response)
	response.Header().Set("Content-type", "application/json")

	content, err := ioutil.ReadFile("assetlinks.json")

    if err!=nil{
        fmt.Print("Error:",err)
    }

 	fmt.Println("Touched")
    // var value []Assetlink
    // err=json.Unmarshal(content, &value)
    fmt.Fprint(response, string(content));
   
}

func main(){
	port := 8090
	portstring := strconv.Itoa(port)

	// Register request handlers for two URL patterns.
	// (The docs are unclear on what a 'pattern' is, 
	// but seems be the start of the URL, ending in a /).
	// See gorilla/mux for a more powerful matching system.
	// Note that the "/" pattern matches all request URLs.
	mux := http.NewServeMux()
	mux.Handle("/.well-known/assetlinks.json", http.HandlerFunc( JsonHandler ))
	mux.Handle("/home", http.HandlerFunc( HomeHandler ))

	log.Print("Listening on port " + portstring + " ... ")

	//err := http.ListenAndServe(":" + portstring, mux)
	err := http.ListenAndServeTLS(":" + portstring, "server.crt",
                           "server.key", mux)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}


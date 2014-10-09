/* This is a simple calculator exposed as a web service.     *
** Part of the assignment for Essential Skills module Lab 10.*
** University of Amsterdam                                   *
** MSc In System and Network Engineering 					 *
** Nick Triantafyllidis										 *
** This will be a RESTful service							 *
*/

package main

import(
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func main(){

	mainRouter := mux.NewRouter()

	getSubrouter := mainRouter.Methods("GET").Subrouter()
	postSubrouter := mainRouter.Methods("POST").Subrouter()
	putSubrouter := mainRouter.Methods("PUT").Subrouter()
	deleteSubrouter := mainRouter.Methods("DELETE").Subrouter()

	http.Handle("/",mainRouter)
	
    getSubrouter.HandleFunc("/api/v1/sum", sum)
    getSubrouter.HandleFunc("/api/v1/difference", difference)
    getSubrouter.HandleFunc("/api/v1/product", product)
    getSubrouter.HandleFunc("/api/v1/quotient", quotient)
    getSubrouter.HandleFunc("/api/v1/{any}", escape)
    getSubrouter.HandleFunc("/api/v1", describe)

 	postSubrouter.HandleFunc("/api/v1/{any}",notAllowed)
 	putSubrouter.HandleFunc("/api/v1/{any}",notAllowed)
 	deleteSubrouter.HandleFunc("/api/v1/{any}",notAllowed)

	err := http.ListenAndServe(":80",nil)

	if err!=nil{
		panic(err)
	}
}

func describe(w http.ResponseWriter, r *http.Request) {

	desc := `Hi! I am a simple calculator. I can add, substract, multiply, and do integer division. 


========USAGE=========

+: /api/v1/sum?a=1&b=2
-: /api/v1/difference?a=1&b=2
*: /api/v1/product?a=1&b=2
\: /api/v1/quotient?a=1&b=2

Of course a and b will be your own values.
If you get a complaint about a zero in division you either have given a zero (duh!) or done something nasty.
No I will not be sanitizing your input, you will just get zeros!!! 
TRY ME! :)
`
	
	w.Write([]byte(desc))
}

func sum(w http.ResponseWriter, r *http.Request) {
	urlValues := r.URL.Query()
	a, _:= strconv.Atoi(urlValues.Get("a"))
	b, _:= strconv.Atoi(urlValues.Get("b"))
    sum := a+b
    w.Write([]byte(strconv.Itoa(sum)))
}

func difference(w http.ResponseWriter, r *http.Request) {
	urlValues := r.URL.Query()
	a, _:= strconv.Atoi(urlValues.Get("a"))
	b, _:= strconv.Atoi(urlValues.Get("b"))
    difference := a-b
    w.Write([]byte(strconv.Itoa(difference)))
}

func product(w http.ResponseWriter, r *http.Request) {
	urlValues := r.URL.Query()
	a, _:= strconv.Atoi(urlValues.Get("a"))
	b, _:= strconv.Atoi(urlValues.Get("b"))
    product := a*b
    w.Write([]byte(strconv.Itoa(product)))
}

func quotient(w http.ResponseWriter, r *http.Request) {
	urlValues := r.URL.Query()
	a, _:= strconv.Atoi(urlValues.Get("a"))
	b, _:= strconv.Atoi(urlValues.Get("b"))
	if b == 0 {
		w.Write([]byte("You cannot divide by zero"))
	}else {
    	quotient:= a/b
    	w.Write([]byte(strconv.Itoa(quotient)))
	}
}

func escape(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(501)
	w.Write([]byte("Not Implemented"))
}

func notAllowed(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(405)
	w.Write([]byte("Method Not Allowed"))
}

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
	http.Handle("/",mainRouter)
	
    mainRouter.HandleFunc("/api/v1/sum",sum).method("GET")
	err := http.ListenAndServe(":8080",nil)

	if err!=nil{
		panic(err)
	}
}

func sum(w http.ResponseWriter, r *http.Request) {
	urlValues := r.URL.Query()
	a, _:= strconv.Atoi(urlValues.Get("a"))
	b, _:= strconv.Atoi(urlValues.Get("b"))
    sum := a+b
    w.Write([]byte(strconv.Itoa(sum)))
}
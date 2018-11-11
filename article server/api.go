package main

import (
          "encoding/json"
          "log"
          "net/http"
          "github.com/gorilla/mux"
)

type Person struct {
                ID string  `json:"id,omitempty"`
                FirstName string `json:"FirstName,omitempty"`
                LastName string `json:"LastName,omitempty"`
}

var people = []Person{
                Person{ID: "1", FirstName: "sulaiman", LastName: "jalloh"},
                Person{ID: "2", FirstName: "life is beautiful", LastName: "enjoy it"},
                Person{ID: "3", FirstName: "life is beautiful", LastName: "enjoy it"},
                Person{ID: "4", FirstName: "life is beautiful", LastName: "enjoy it"}}


func GetData(w http.ResponseWriter, r *http.Request) {
          w.Header().Set("Content-Type", "application/json")
                json.NewEncoder(w).Encode(people)
}

func PostData(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        params := mux.Vars(r)
	       var person Person
	        _ = json.NewDecoder(r.Body).Decode(&person)
	         person.ID = params["id"]
	          people = append(people, person)
	           json.NewEncoder(w).Encode(people)
}

func DeleteData(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
	   params := mux.Vars(r)
	    for index, item := range people {
		      if item.ID == params["id"] {
			         people = append(people[:index], people[index+1:]...)
			            break
		              }
		                 json.NewEncoder(w).Encode(people)
	                  }
        }

func main() {
                router := mux.NewRouter()
                router.HandleFunc("/data", GetData).Methods("GET")
                router.HandleFunc("/data", PostData).Methods("POST")
                router.HandleFunc("/data/{id}", DeleteData).Methods("DELETE")
                log.Print("localhost:8000")
                log.Fatal(http.ListenAndServe(":8000", router))
}

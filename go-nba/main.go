package main

import(
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"

)


type Player struct{

	Firstname string `json: "firstname"`
	Lastname string `json: "lastname"`

	Team string `json: "team"`
	Pos string `json: "pos"`
	Number string `json: "number"`

	HeadCoach *Coach `json: "coach"`

}

type Coach struct{
	Firstname string `json: "firstname"`
	Lastname string `json: "lastname"`

	Salary int `json: "salary"`
}

var players []Player


//CREATE
func createPlayer(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content Type", "application/json")
	var player Player 
	_ = json.NewDecoder(r.Body).Decode(&player)
	players = append(players, player)
	json.NewEncoder(w).Encode(player)
}



//READ
//Read all players on database 
func getPlayers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(players)
}

//Read selected player
func getPlayer(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content Type", "application/json")
	params := mux.Vars(r)
	for _, item:= range players {
		if item.Number == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}

	}
}


//UPDATE
func updatePlayer(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content Type", "application/json")
	params := mux.Vars(r)

	for index, item := range players{
		if item.Number == params["id"]{
			players = append(players[:index], players[index+1:]...)
			var movie Player
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Number = params["id"]
			players = append(players, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}



}


//DELETE
func deletePlayer(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item:= range players {

		if item.Number == params["id"]{
			players = append(players[:index], players[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(players)
}





 func main(){
	r := mux.NewRouter()

	players = append(players, Player{Firstname: "Lebron", Lastname:"James", Team: "Los Angeles Lakers", Pos: "PF", Number: "23", HeadCoach : &Coach{Firstname: "Darvin", Lastname: "Ham"}})
	players = append(players, Player{Firstname: "Stephen", Lastname:"Curry", Team: "Golden State Warriors", Pos: "SG", Number: "30", HeadCoach : &Coach{Firstname: "Steve", Lastname: "Kerr"}})
	r.HandleFunc("/movies", getPlayers).Methods("GET")
	r.HandleFunc("/movies/{id}", getPlayer).Methods("GET")
	r.HandleFunc("/movies", createPlayer).Methods("POST")
	r.HandleFunc("/movies/{id}", updatePlayer).Methods("PUT")
	r.HandleFunc("/movies/{id}", deletePlayer).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}




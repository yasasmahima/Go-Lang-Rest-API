package Router

import (
	"go-postgres/repository"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/user/{id}",UserRepository.GetUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/user",UserRepository.GetAllUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/newuser",UserRepository.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/user/{id}",UserRepository.UpdateUser).Methods("PUT", "OPTIONS")
	router.HandleFunc("/deleteuser/{id}",UserRepository.DeleteUser).Methods("DELETE", "OPTIONS")

	return router
}

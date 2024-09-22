package router

import (
	"github.com/SohomSaha045/mongoAPI/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/api",controller.Dummy).Methods("GET")
	router.HandleFunc("/api/movies",controller.Get_All_Movies).Methods("GET")
	router.HandleFunc("/api/movie",controller.Create_a_Movie).Methods("POST")
	router.HandleFunc("/api/movie/{id}",controller.Mark_As_Watched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}",controller.Delete_A_Movie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie",controller.Delete_All_Movies).Methods("DELETE")
	router.HandleFunc("/api/movie/{id}",controller.Get_A_Movie).Methods("GET")
	

	return router
}
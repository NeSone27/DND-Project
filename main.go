package main

import (
	"log"
	"net/http"
	"todolist-service/config"
	"todolist-service/handlers"
	"todolist-service/repositories"
	"todolist-service/services"

	"github.com/gorilla/mux"
)

func main() {
	config.InitDB()
	repoUser := &repositories.UserRepository{DB: config.DB}
	serviceUser := &services.UserService{Repo: repoUser}
	handlerUser := &handlers.UserHandler{Service: serviceUser}

	repoClass := &repositories.ClassRepository{DB: config.DB}
	serviceClass := &services.ClassService{Repo: repoClass}
	handlerClass := &handlers.ClassHandler{Service: serviceClass, ServiceUser: serviceUser}

	repoRace := &repositories.RaceRepository{DB: config.DB}
	serviceRace := &services.RaceService{Repo: repoRace}
	handlerRace := &handlers.RaceHandler{Service: serviceRace, ServiceUser: serviceUser}

	repoDifficultyLevel := &repositories.DifficultyLevelRepository{DB: config.DB}
	serviceDifficultyLevel := &services.DifficultyLevelService{Repo: repoDifficultyLevel}
	handlerDifficultyLevel := &handlers.DifficultyLevelHandler{Service: serviceDifficultyLevel, ServiceUser: serviceUser}

	repoImage := &repositories.ImageRepository{DB: config.DB}

	repoCharacter := &repositories.CharacterRepository{DB: config.DB}
	serviceCharacter := &services.CharacterService{Repo: repoCharacter, RepoImage: repoImage}
	handlerCharacter := &handlers.CharacterHandler{Service: serviceCharacter, ServiceUser: serviceUser}

	repoQuest := &repositories.QuestRepository{DB: config.DB}
	serviceQuest := &services.QuestService{Repo: repoQuest, RepoImage: repoImage}
	handlerQuest := &handlers.QuestHandler{Service: serviceQuest, ServiceUser: serviceUser}

	router := mux.NewRouter()
	router.HandleFunc("/user", handlerUser.CreateUser).Methods("POST")
	router.HandleFunc("/user", handlerUser.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", handlerUser.GetUserByID).Methods("GET")
	router.HandleFunc("/user/{id}", handlerUser.UpdateUser).Methods("PATCH")
	router.HandleFunc("/user/{id}", handlerUser.DeleteUser).Methods("DELETE")

	router.HandleFunc("/class", handlerClass.CreateClass).Methods("POST")
	router.HandleFunc("/class", handlerClass.GetClasses).Methods("GET")
	router.HandleFunc("/class/{id}", handlerClass.GetClassByID).Methods("GET")
	router.HandleFunc("/class/{id}", handlerClass.UpdateClass).Methods("PATCH")
	router.HandleFunc("/class/{id}", handlerClass.DeleteClass).Methods("DELETE")

	router.HandleFunc("/race", handlerRace.CreateRace).Methods("POST")
	router.HandleFunc("/race", handlerRace.GetRaces).Methods("GET")
	router.HandleFunc("/race/{id}", handlerRace.GetRaceByID).Methods("GET")
	router.HandleFunc("/race/{id}", handlerRace.UpdateRace).Methods("PATCH")
	router.HandleFunc("/race/{id}", handlerRace.DeleteRace).Methods("DELETE")

	router.HandleFunc("/difficulty-level", handlerDifficultyLevel.CreateDifficultyLevel).Methods("POST")
	router.HandleFunc("/difficulty-level", handlerDifficultyLevel.GetDifficultyLevels).Methods("GET")
	router.HandleFunc("/difficulty-level/{id}", handlerDifficultyLevel.GetDifficultyLevelByID).Methods("GET")
	router.HandleFunc("/difficulty-level/{id}", handlerDifficultyLevel.UpdateDifficultyLevel).Methods("PATCH")
	router.HandleFunc("/difficulty-level/{id}", handlerDifficultyLevel.DeleteDifficultyLevel).Methods("DELETE")

	router.HandleFunc("/character", handlerCharacter.CreateCharacter).Methods("POST")
	router.HandleFunc("/character", handlerCharacter.GetCharacters).Methods("GET")
	router.HandleFunc("/character/{id}", handlerCharacter.UpdateCharacter).Methods("PATCH")
	router.HandleFunc("/character/{id}", handlerCharacter.DeleteCharacter).Methods("DELETE")

	router.HandleFunc("/quest", handlerQuest.CreateQuest).Methods("POST")
	router.HandleFunc("/quest", handlerQuest.GetQuests).Methods("GET")
	router.HandleFunc("/quest/{id}", handlerQuest.UpdateQuest).Methods("PATCH")
	router.HandleFunc("/quest/{id}", handlerQuest.DeleteQuest).Methods("DELETE")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

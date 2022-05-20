package controllers

import (
	"encoding/json"
	"fmt"
	"gorm-test/database"
	"gorm-test/models"
	"gorm-test/responses"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	Db     *gorm.DB
	Router *mux.Router
}

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	fmt.Println(&user)

	err = models.CreateUser(server.Db, &user)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var user models.User
	err = models.GetUser(server.Db, &user, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	err := models.GetUsers(server.Db, &users)
	if err != nil {
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	err = models.GetUser(server.Db, &user, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = models.UpdateUser(server.Db, &user)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var user models.User
	err = models.DeleteUser(server.Db, &user, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	data := struct {
		Message string `json:"message"`
	}{
		"User deleted success",
	}
	responses.JSON(w, http.StatusOK, data)
}

func Init() {
	var server = Server{}
	server.Db = database.InitDb()
	server.Db.AutoMigrate(&models.User{})
	server.Router = mux.NewRouter()
	server.initializeRoutes()
	server.Run(":8081")
}

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/users", server.CreateUser).Methods("POST")
	server.Router.HandleFunc("/users", server.GetUsers).Methods("GET")
	server.Router.HandleFunc("/users/{id}", server.GetUser).Methods("GET")
	server.Router.HandleFunc("/users/{id}", server.UpdateUser).Methods("PUT")
	server.Router.HandleFunc("/users/{id}", server.DeleteUser).Methods("DELETE")
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port " + addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

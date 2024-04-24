package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/laozhubaba/api_proj/pkg/server"
)

type Controller struct {
	server.Logic
	logger server.Logger
	Rtr    *mux.Router
}

func NewRestController(ctx context.Context, l server.Logger, logic server.Logic) (c Controller) {
	rtr := mux.NewRouter()
	c = Controller{logger: l, Logic: logic, Rtr: rtr}
	rtr.HandleFunc("/api/user/{userid}", c.GetUser).Methods("GET")
	rtr.HandleFunc("/api/user", c.GetUsers).Methods("GET")
	rtr.HandleFunc("/api/user", c.AddUser).Methods("POST")
	return c
}

func (c Controller) AddUser(w http.ResponseWriter, r *http.Request) {
	c.logger.Info("in rest controller for AddUser")
	w.Header().Set("Content-Type", "application/json")
	person := server.Person{}
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		c.logger.Error("cannot decode data: %v", err)
	}
	err = c.Logic.AddUser(person)
	if err != nil {
		c.logger.Error("Adduser failed: %v", err)
	}
}
func (c Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	c.logger.Info("in rest controller for GetUser")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userIdStr := vars["userid"]
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.logger.Error("cannot convert string: %s to integer", userIdStr)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	message, err := c.Logic.GetUser(userIdInt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
}

func (c Controller) GetUsers(w http.ResponseWriter, r *http.Request) {
	c.logger.Info("in rest controller for GetUser")
	w.Header().Set("Content-Type", "application/json")
	message, err := c.Logic.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
}

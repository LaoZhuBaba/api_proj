package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/laozhubaba/api_proj/internal/logger"
	"github.com/laozhubaba/api_proj/pkg/server"
)

type Controller struct {
	server.Logic
	logger.Logger
	Rtr *mux.Router
}

func NewRestController(ctx context.Context, l server.Logger, logic server.Logic) (c Controller) {
	rtr := mux.NewRouter()
	c = Controller{Logger: l, Logic: logic, Rtr: rtr}
	rtr.HandleFunc("/api/user/{userid}", c.GetUser)
	rtr.HandleFunc("/api/user", c.GetUsers)
	return c
}

func (c Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	c.Logger.Log("in rest controller for GetUser")
	vars := mux.Vars(r)
	c.Logger.Log(fmt.Sprintf("vars: %s", vars))
	userIdStr := vars["userid"]
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.Logger.Log(fmt.Sprintf("cannot convert string: %s to integer", userIdStr))
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
	c.Logger.Log("in rest controller for GetUser")
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

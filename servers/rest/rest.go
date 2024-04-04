package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/laozhubaba/api_proj/interfaces"
)

func NewRestServer(ctx context.Context, addr string, port int, backend interfaces.Backend) (restServer, error) {
	return restServer{
		Address: addr,
		Port:    port,
		Backend: backend,
	}, nil
}

// restServer implements the Frontend interface
type restServer struct {
	Address string
	Port    int
	Backend interfaces.Backend
}

func (rs restServer) Run() error {
	router := gin.Default()
	router.GET("/users", rs.getUsers)
	router.GET("/users/:id", rs.getUserById)

	router.Run(fmt.Sprintf("%s:%d", rs.Address, rs.Port))
	return nil
}

func (rs restServer) getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, rs.Backend.GetAllUsers())
}
func (rs restServer) getUserById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id not provided"})
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id not valid"})
		return
	}
	u, err := rs.Backend.GetUserById(idInt)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id not found"})
		return
	}
	enc := json.NewEncoder(c.Writer)
	err = enc.Encode(u)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
		return
	}
}

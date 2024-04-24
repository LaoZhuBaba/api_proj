package restclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/laozhubaba/api_proj/cmd/server/common"
	"github.com/laozhubaba/api_proj/pkg/server"
)

func AddUser(person server.Person) error {
	client := &http.Client{Timeout: 30 * time.Second}
	b, err := json.Marshal(&person)
	buf := bytes.NewBuffer(b)
	if err != nil {
		log.Printf("failed to Marshal person: %v", err)
		return err
	}
	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080/api/user", buf)
	if err != nil {
		log.Printf("failed to create http request: %v", err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Printf("http client request failed: %v", err)
		return err
	}
	defer res.Body.Close()
	log.Printf("http status code: %d", res.StatusCode)
	return nil
}

func GetUser(id int) (person server.Person, err error) {
	client := &http.Client{Timeout: 30 * time.Second}
	person = server.Person{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/api/user/%d", common.RestPort, id), nil)
	if err != nil {
		log.Printf("failed to create http request: %v\n", err)
		return person, err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Printf("http client request failed: %v\n", err)
		return person, err
	}
	defer res.Body.Close()
	log.Printf("http status code: %d\n", res.StatusCode)
	err = json.NewDecoder(res.Body).Decode(&person)
	if err != nil {
		log.Printf("failed to decode body:%v\n", err)
		return server.Person{}, err
	}
	return person, nil
}

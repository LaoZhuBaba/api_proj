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
	buf := bytes.NewReader(b)
	if err != nil {
		log.Printf("failed to Marshal person: %v", err)
		return err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:%d/api/user", common.RestPort), buf)
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
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("https status code: %d", res.StatusCode)
	}
	defer res.Body.Close()
	log.Printf("http status code: %d", res.StatusCode)
	return nil
}

func GetUser(id int) (person *server.Person, err error) {
	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/api/user/%d", common.RestPort, id), nil)
	if err != nil {
		log.Printf("failed to create http request: %v\n", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Printf("http client request failed: %v\n", err)
		return nil, err
	}
	defer res.Body.Close()
	log.Printf("http status code: %d\n", res.StatusCode)
	if res.StatusCode != http.StatusOK {
		log.Printf("exiting with status code: %d\n", res.StatusCode)
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&person)
	if err != nil {
		log.Printf("failed to decode body:%v\n", err)
		return nil, err
	}
	return person, nil
}

func GetAllUsers() (persons []server.Person, err error) {
	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/api/user", common.RestPort), nil)
	if err != nil {
		log.Printf("failed to create http request: %v\n", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Printf("http client request failed: %v\n", err)
		return nil, err
	}
	defer res.Body.Close()
	log.Printf("http status code: %d\n", res.StatusCode)
	err = json.NewDecoder(res.Body).Decode(&persons)
	if err != nil {
		log.Printf("failed to decode body:%v\n", err)
		return nil, err
	}
	return persons, nil
}

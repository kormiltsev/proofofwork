package client

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Task represents a task from server.
type Task struct {
	Difficulty int    `json:"difficulty"`
	Hash       string `json:"hash"`
}

// Words is a finnal result.
type Words struct {
	Quote string `json:"quote"`
}

// Solution represents answer to server.
type Solution struct {
	Solution string `json:"solution"`
}

// Get sends first request to server and returns a task.
func Get(url string) (*Task, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	task := Task{}
	err = json.Unmarshal(body, &task)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &task, nil
}

// Post sends a result to get wisdom from server.
func Post(url string, data []byte) (string, error) {
	payload := Solution{string(data)}

	janswer, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(janswer))
	if err != nil {
		log.Println(err)
		return "", err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	resultjson, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	result := Words{}
	if err := json.Unmarshal(resultjson, &result); err != nil {
		log.Println(err)
		return "", err
	}

	return result.Quote, nil
}

package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/kormiltsev/proofofwork/client/client"
	"github.com/kormiltsev/proofofwork/internal/service/pow"
	"gopkg.in/tomb.v2"
)

var hardcodeedClientID = "client"

// OneRequest make one request to get task, solve and send solution to get more wisdom.
func OneRequest(targeturl string) error {
	req, err := client.Get(targeturl)
	if err != nil {
		log.Println("err:", err)
		return err
	}

	block := pow.NewClient().Solve(hardcodeedClientID, []byte(req.Hash), req.Difficulty)
	janswer, err := json.Marshal(*block)
	if err != nil {
		log.Println("err:", err)
		return err
	}

	result, err := client.Post(targeturl, janswer)
	if err != nil {
		log.Println("err:", err)
		return err
	}

	printResult(result)
	return nil
}

// Endless is a loop. Requests for task, solve and send resulution to get more wisdome.
func Endless(targeturl string) error {
	t, _ := tomb.WithContext(context.Background())
	t.Go(func() error { return worker(targeturl) })
	<-t.Dead()
	return nil
}

// a loop
func worker(targeturl string) (err error) {
	for {
		if err = OneRequest(targeturl); err != nil {
			return
		}
	}
}

// output the result
func printResult(result string) {
	if result == "" {
		fmt.Printf("\n🔴 no results\n")
		return
	}
	fmt.Printf("\nresult: ========================================================\n%s\n", result)
}

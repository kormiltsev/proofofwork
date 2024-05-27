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

func Endless(targeturl string) error {
	t, _ := tomb.WithContext(context.Background())
	for i := 0; i < 5; i++ {
		t.Go(func() error { return worker(targeturl) })
	}
	<-t.Dead()
	return nil
}

func worker(targeturl string) (err error) {
	for {
		if err = OneRequest(targeturl); err != nil {
			return
		}
	}
}

func printResult(result string) {
	fmt.Printf("\nresult: ========================================================\n%s", result)
}

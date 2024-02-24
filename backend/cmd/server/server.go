package main

import (
	"log"

	"github.com/y-mabuchi/bulletin-board-system/backend/domain/model"
)

func main() {
	user1, err := model.NewUser("user1", "password", "test@example.com")
	if err != nil {
		log.Fatal(err)
	}

	creator1, err := model.NewThreadCreator(*user1)
	if err != nil {
		log.Fatal(err)
	}

	thread, err := model.NewThread("title1", "content1", *creator1)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("thread id:", thread.ThreadID.String())
	log.Println("thread title:", thread.Title)
	log.Println("thread content:", thread.Content)
}

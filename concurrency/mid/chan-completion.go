package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

// Канал завершения
// Есть функция, которая произносит текст пословно (с некоторыми задержками):

func say(done chan<- struct{}, id int, text string) {

	for _, word := range strings.Fields(text) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
	done <- struct{}{}
}

// Запускаем несколько одновременных воркеров, по одной на каждую фразу:

func main() {

	var done = make(chan struct{})

	phrases := []string{
		"go is awesome",
		"cats are cute",
		"rain is wet",
		"channels are hard",
		"floor is lava",
	}
	for idx, phrase := range phrases {
		go say(done, idx+1, phrase)
		log.Println("start worker number ", idx)
	}
	//time.Sleep(3 * time.Second)
	counter := 0
	for {
		select {
		case <-done:
			counter += 1
			log.Println("worker done his job")
		default:
			if counter == len(phrases) {
				log.Println("all done")
				return
			}
			continue
		}
	}
}
//
//
//
//
//
// say(done chan<- struct{}, id int, phrase string).

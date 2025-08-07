// Если тема контекста еще не изучена - пропустить задачу
package main

import (
	"context"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Есть функция, работающая неопределённо долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).
func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)
	return rnd
}

// Нужно изменить функцию-обёртку, которая будет работать с заданным таймаутом
// (например, 1 секунду).
// Если "длинная" функция отработала за это время - отлично, возвращаем результат.
// Если нет - возвращаем ошибку. Результат работы в этом случае нам не важен.
//
// Дополнительно нужно измерить, сколько выполнялась эта функция (просто вывести в лог).
// Сигнатуру функцию обёртки менять можно.

func predictableFunc(ctx context.Context) int64 {
	var fnDone = make(chan int64)

	go func() {
		defer close(fnDone)

		log.Println("Starting func")
		fnRes := unpredictableFunc()
		log.Println("result of unpredictable func - ", fnRes)

		fnDone <- fnRes
	}()

	select {
	case <-ctx.Done():
		log.Printf("ctx done: %v", ctx.Err())
		return 0
	case <-fnDone:
		defer log.Println("end of func")
		return 0
	}
}

func main() {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	predictableFunc(ctx)
	log.Println("execution time -", time.Since(start))
}


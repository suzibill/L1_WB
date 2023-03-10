package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Необходимо указать количество воркеров")
		return
	}

	numWorker, err := strconv.Atoi(args[0])
	if err != nil || numWorker < 1 {
		fmt.Println("Ошибка в числе воркеров")
		return
	}
	fmt.Println(numWorker)
	wg := &sync.WaitGroup{}

	ch := make(chan string)
	wg.Add(numWorker)
	for i := 0; i < numWorker; i++ {
		go func(i int) {
			defer wg.Done()
			for m := range ch {
				fmt.Printf("Воркер %v: %v\n", i+1, m)
			}
			fmt.Printf("Воркер %v закончил работу\n", i+1)
		}(i)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-signals
		close(ch)
		wg.Wait()
		os.Exit(0)
	}()

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		ch <- text
	}

}

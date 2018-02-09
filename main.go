package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func cleanup() {
	fmt.Println("cleanup")
}

func main() {

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	for {
		fname := getFileName()

		fmt.Println(fname)

		fmt.Println("sleeping...")
		time.Sleep(2 * time.Second) // or runtime.Gosched() or similar per @misterbee
	}

}

func getFileName() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	topics := []string{"a", "b", "c", "d"}
	partitions := []int{0, 1, 2}

	lengthTopic := len(topics)
	lengthPartitions := len(partitions)

	fileName := fmt.Sprintf("%v_%v.txt", topics[r.Intn(lengthTopic)], partitions[r.Intn(lengthPartitions)])

	return fileName
}

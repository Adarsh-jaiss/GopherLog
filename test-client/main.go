package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"log"
)

func main() {
	payload := []byte("PING")
	loopCnt := 1
	for {
		fmt.Printf("--- Starting iteration loop : %d ---\n", loopCnt)

		var wg sync.WaitGroup
		wg.Add(5)

		for i := 0; i < 5; i++ {
			go func() {
				defer wg.Done()
				conn, err := net.Dial("tcp", "localhost:8001")
				if err != nil {
					log.Println(err)
					return
				}
				// defer conn.Close()

				writeFrame(conn, payload)
				res, err := readFrame(conn)
				if err!= nil {
					log.Println("err reading response from server:", err)
				}
				log.Println(string(res))
				time.Sleep(time.Second * 3)
			}()
		}
		wg.Wait()
		fmt.Printf("--- %d: Iteration loop completed ---\n", loopCnt)
		loopCnt++
	}

}

// single tcp conn sending frames using 5 workers
func SingleTCPConn() {
	conn, err := net.Dial("tcp", "localhost:8001")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	payload := []byte("hello, this is 16")

	var wg sync.WaitGroup
	id := 1
	loopCnt := 1
	for {
		fmt.Printf("--- Starting iteration loop : %d ---\n", loopCnt)

		wg.Add(5)
		go worker(id, &wg, conn, payload)
		id++
		go worker(id, &wg, conn, payload)
		id++
		go worker(id, &wg, conn, payload)
		id++
		go worker(id, &wg, conn, payload)
		id++
		go worker(id, &wg, conn, payload)

		wg.Wait()
		fmt.Printf("--- %d: Iteration loop completed ---\n", loopCnt)
		loopCnt++
	}
}

func writeFrame(conn net.Conn, payload []byte) error {
	frame := make([]byte, 4+len(payload))

	binary.BigEndian.PutUint32(frame[:4], uint32(len(payload)))

	copy(frame[4:], payload)

	_, err := conn.Write(frame)
	return err
}

func worker(id int, wg *sync.WaitGroup, conn net.Conn, payload []byte) {
	defer wg.Done()

	fmt.Printf("Goroutine %d: Starting task...\n", id)
	if err := writeFrame(conn, payload); err != nil {
		log.Printf("error sending request: %v", err)
	}

	fmt.Printf("Goroutine %d: sleeping for 3 seconds\n", id)
	time.Sleep(3 * time.Second)
	fmt.Printf("Goroutine %d: Task completed!\n", id)

}

func readFrame(r io.Reader) ([]byte, error) {
	var length uint32
	if err := binary.Read(r, binary.BigEndian, &length); err != nil {
		return nil, err
	}

	if length > 10*1024*1024 {
		return nil, errors.New("frame too large")
	}

	payload := make([]byte, length)

	if _, err := io.ReadFull(r, payload); err != nil {
		return nil, err
	}
	return payload, nil
}

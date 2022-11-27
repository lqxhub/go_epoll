package main

import (
	"fmt"
	"net"
	"sync"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:8088")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("hello"))
	data := make([]byte, 100)
	n, err := conn.Read(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data[:n]))

	time.Sleep(100 * time.Millisecond)
	conn.Close()
}

func TestClients(t *testing.T) {
	for i := 0; i < 100; i++ {
		conn, err := net.Dial("tcp", "127.0.0.1:8088")
		if err != nil {
			panic(err)
		}
		conn.Write([]byte("hello"))
		data := make([]byte, 100)
		n, err := conn.Read(data)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data[:n]))

		conn.Close()
	}
}

func TestClientMulit(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			conn, err := net.Dial("tcp", "127.0.0.1:8088")
			if err != nil {
				panic(err)
			}
			conn.Write([]byte("hello"))
			data := make([]byte, 100)
			n, err := conn.Read(data)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(data[:n]))

			time.Sleep(100 * time.Millisecond)
			conn.Close()
			wg.Done()
		}()
	}
	wg.Wait()
}

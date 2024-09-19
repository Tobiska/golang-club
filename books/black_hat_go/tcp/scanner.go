package main

import (
	"context"
	"fmt"
	"net"
	"time"
)

const (
	maxPort    = 65255
	maxWorkers = 10
	target     = "scanme.nmap.org"
)

var (
	dialer = &net.Dialer{Timeout: 5 * time.Second}
)

func worker(ctx context.Context, ports, results chan int) error {
	for ctx.Err() == nil {
		select {
		case <-ctx.Done():
			return nil
		case port, ok := <-ports:
			if !ok {
				return nil
			}
			isOpen, err := scan(port)
			if err != nil {
				return fmt.Errorf("error while scan: %w", err)
			}

			if isOpen {
				results <- port
			}
		}
	}

	return nil
}

func scan(port int) (bool, error) {
	conn, err := dialer.Dial("tcp", fmt.Sprintf("%s:%d", target, port))
	if err != nil {
		// по хорошему добавить логирование и обработку отдельных ошибок.
		return false, nil
	}
	if err = conn.Close(); err != nil {
		return false, fmt.Errorf("connection close error: %w", err)
	}

	return true, nil
}

//func main() {
//	var err error
//	ports := make(chan int, 1)
//	results := make(chan int)
//	openedPorts := make([]int, 0, maxPort)
//
//	g, ctx := errgroup.WithContext(context.Background())
//
//	for w := 0; w < maxWorkers; w++ {
//		g.Go(func() error {
//			return worker(ctx, ports, results)
//		})
//	}
//
//	for i := 0; i < maxPort; i++ {
//		ports <- i
//	}
//	close(ports)
//
//	go func() {
//		err = g.Wait()
//		close(results)
//	}()
//
//	for port := range results {
//		openedPorts = append(openedPorts, port)
//	}
//
//	if err != nil {
//		log.Fatal(fmt.Errorf("scanner finish ahead of schedule: %w", err))
//	}
//
//	fmt.Println(openedPorts)
//}

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"golang.org/x/sync/errgroup"
)

const (
	dstAddress = "google.com:80"
)

func handle(ctx context.Context, src net.Conn) error {
	dst, err := net.Dial("tcp", dstAddress)
	if err != nil {
		return fmt.Errorf("error while connecting to dst address: %w", err)
	}
	g, eCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		if err := copy(eCtx, dst, src); err != nil {
			return fmt.Errorf("error while copy from src to dst: %w", err)
		}

		return nil
	})

	g.Go(func() error {
		if err := copy(eCtx, src, dst); err != nil {
			return fmt.Errorf("error while copy from dst to src: %w", err)
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}

func copy(ctx context.Context, reader io.Reader, writer io.Writer) error {
	var resultErr error
	go func() {
		if _, err := io.Copy(writer, reader); err != nil {
			resultErr = err
		}
	}()

	<-ctx.Done()

	return resultErr
}

func run() error {
	ctx := context.Background()

	lst, err := net.Listen("tcp", ":3000")
	if err != nil {
		return fmt.Errorf("error while create listerner: %w", err)
	}

	conn, err := lst.Accept()
	if err != nil {
		return fmt.Errorf("conn accept error: %w", err)
	}

	if err := handle(ctx, conn); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

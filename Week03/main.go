package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"html"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
func main() {
	g,ctx := errgroup.WithContext(context.Background())
	c := make(chan os.Signal, 1)
	serErr := make(chan error)
	g.Go(func() error {
		go func() {
			http.HandleFunc("/foo", fooHandler)
			fmt.Println("start http server...")
			serErr<-http.ListenAndServe(":8080", nil)
		}()
		select {
			case <-serErr:
			case <-ctx.Done():
				fmt.Println("http server shutdown")
				return errors.New("http server shutdown")
		}
		return nil
	})
	g.Go(func() error {
		signal.Notify(c, syscall.SIGINT,syscall.SIGTERM,syscall.SIGQUIT,os.Interrupt)

		select {
			case <-c :
			case <-ctx.Done():
				fmt.Println("receive quit signal")
				return errors.New("receive quit signal")
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}

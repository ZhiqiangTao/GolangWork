package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"os/signal"
)
import "golang.org/x/sync/errgroup"

func main() {
	srv := &http.Server{Addr: ":8080"}
	group, ctx := errgroup.WithContext(context.Background())

	//1.注册http，并启动
	group.Go(func() error {
		http.HandleFunc("/a", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("I am A"))
		})
		http.HandleFunc("/b", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("I am B"))
		})
		http.HandleFunc("/error", func(writer http.ResponseWriter, request *http.Request) {
			a := 1
			b := 0
			e := a / b
			writer.Write([]byte(fmt.Sprintf("%d", e)))
		})
		fmt.Println("程序启动")
		return srv.ListenAndServe()
	})

	//2.等待ctx信号，将服务shutdown
	group.Go(func() error {
		<-ctx.Done()
		fmt.Println("检测到退出信号，程序shutdown")
		return srv.Shutdown(ctx)
	})

	//3.监听ctrl+c等客户端信号量
	sig := make(chan os.Signal, 1)
	signal.Notify(sig)
	group.Go(func() error {
		for {
			select {
			case <-sig:
				return errors.New("监听signal信号, 抛出异常")
			}
		}
	})

	fmt.Println("group wait")
	if err := group.Wait(); err != nil {
		fmt.Println("group error: ", err)
	}
	fmt.Println("程序退出了")
}

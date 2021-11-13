package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"golangwork/week04/internal/app/signin/domain"
	"golangwork/week04/internal/app/signin/service"
	"net/http"
	"os"
	"os/signal"
	"strconv"
)

func main() {
	srv := &http.Server{Addr: ":8080"}
	group, ctx := errgroup.WithContext(context.Background())

	svc := service.InitSignInService()

	//1.注册http，并启动
	group.Go(func() error {
		http.HandleFunc("/signin", func(writer http.ResponseWriter, request *http.Request) {
			res, err := svc.DoSignIn(&domain.SignIn{
				UserId:         9113890,
				SignInConfigId: 1001,
			})
			if err != nil {
				msg := errors.Wrap(err, "svc error").Error()
				writer.Write([]byte(msg))
				return
			}
			writer.Write([]byte("签到结果为：" + strconv.FormatBool(res)))
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

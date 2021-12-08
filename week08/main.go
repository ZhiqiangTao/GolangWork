package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"strconv"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		panic("fail to ping redis")
	}
}

func insertMany(num int) {
	//clear
	clear, _ := rdb.FlushAll().Result()
	fmt.Println("clear env datas:", clear)

	info, _ := rdb.Info("memory").Result()
	fmt.Printf("insert %v keys before:\n %v \n", num, info)

	cmds := make([]interface{}, 0, num*2)
	for i := 0; i < num; i++ {
		cmds = append(cmds, "test:demo"+strconv.Itoa(i), i)
	}
	rdb.MSet(cmds...).Result()

	info, _ = rdb.Info("memory").Result()
	fmt.Printf("insert %v keys after:\n %v \n", num, info)
}

func main() {
	defer rdb.Close()

	num := 0
	for {
		fmt.Print("请输入要创建的缓存数量(输入-1退出)：")
		fmt.Scan(&num)
		if num == -1 {
			break
		}
		insertMany(num)
		fmt.Println("over：", num)
	}
	fmt.Println("程序退出！！！")
	os.Exit(0)
}

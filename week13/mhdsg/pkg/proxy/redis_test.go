package proxy

import (
	"context"
	"testing"
	"time"
)

var helper IRedisHelper

func init() {
	helper = NewRedisHelper("redis://10.9.159.232:9905?dial_timeout=3&db=0&read_timeout=6s&max_retries=2", nil)
}

func TestSet(t *testing.T) {
	ctx, c := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() { c() }()
	ok := helper.Set(ctx, "testgo", "testgogogo", 0)
	t.Logf("testset: %v", ok)
}

func TestGetString(t *testing.T) {
	ctx, c := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() { c() }()
	dest := ""
	helper.Get(ctx, "testgo", &dest)
}

func TestGetStruct(t *testing.T) {
	ctx, c := context.WithTimeout(context.Background(), 60*time.Second)
	defer func() { c() }()
	input := person{
		Name: "golang",
		Age:  18,
	}
	helper.Set(ctx, "TestGetStruct", input, 0)

	dest := person{}
	helper.Get(ctx, "TestGetStruct", &dest)
}

type person struct {
	Name string `json:"name"`
	Age  int32
}

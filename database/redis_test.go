package database

import (
	"context"
	"fmt"
	"testing"
)

func Test_connectRedis(t *testing.T) {
	ConnectRedis()
	ctx := context.Background()
	rds, err := GetRedis()
	if err != nil {
		t.Fatal(err)
	}
	SetItem(ctx, rds, "name", "redis-test")
	SetItem(ctx, rds, "name2", "redis-test-2")
	val := GetItem(ctx, rds, "name")

	fmt.Printf("First value with name key : %s \n", val)

	values := GetAllKeys(ctx, rds, "name*")

	fmt.Printf("All values : %v \n", values)
}

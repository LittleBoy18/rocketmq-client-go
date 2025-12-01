package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func main() {

	nameSrvAddr := []string{"10.244.210.71:9876"}

	testAdmin, err := admin.NewAdmin(
		admin.WithResolver(primitive.NewPassthroughResolver(nameSrvAddr)),
	)

	result, err := testAdmin.GetReplicaInfo(context.Background(), "10.244.141.147:9878", "testmq-rocketmq-0")
	if err != nil {
		fmt.Println("GetReplicaInfo error:", err.Error())
	}
	fmt.Println(result)

	err = testAdmin.Close()
	if err != nil {
		fmt.Printf("Shutdown admin error: %s", err.Error())
	}
}

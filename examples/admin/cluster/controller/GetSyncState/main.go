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

	var brokers []string
	brokers = append(brokers, "testmq-rocketmq-0")

	result, err := testAdmin.GetSyncStateData(context.Background(), "10.244.141.147:9878", brokers)
	if err != nil {
		fmt.Println("GetSyncStateData error:", err.Error())
	}
	fmt.Println(result.BrokerReplicasInfo.ReplicasInfoTable["testmq-rocketmq-0"].InSyncReplicas)

	err = testAdmin.Close()
	if err != nil {
		fmt.Printf("Shutdown admin error: %s", err.Error())
	}
}

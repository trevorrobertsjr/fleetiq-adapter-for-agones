// 479547368021
package main

import (
	"fmt"
	"log"
	"os"

	// "github.com/aws/aws-sdk-go-v2/aws"

	"github.com/go-redis/redis/v8"
)

func main() {
	// export REDIS_URL=gamelift-common-services-redis-master.default.svc.cluster.local
	// export AWS_REGION=us-east-1
	// 1) "message"
	// 2) "i-0a24709356edc596e"
	// 3) "{\"GameServerGroupName\": \"agones-game-server-group-01\", \"GameServerGroupArn\": \"arn:aws:gamelift:us-east-1:479547368021:gameservergroup/agones-game-server-group-01\", \"InstanceId\": \"i-0a24709356edc596e\", \"InstanceStatus\": \"ACTIVE\"}"
	port := "6379"
	fmt.Println("test")
	// export AWS_REGION=$(curl -s 169.254.169.254/latest/dynamic/instance-identity/document | jq -r '.region')
	// cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")))
	// if err != nil {
	// 	log.Fatalf("unable to load SDK config, %v", err)
	// }
	// // Using the Config value, create the DynamoDB client
	// svcgame := gamelift.NewFromConfig(cfg)

	// // Using the Config value, create the DynamoDB client
	// svcec2 := ec2.NewFromConfig(cfg)

	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL") + ":" + port,
		//Password: password,
	})
	result, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		log.Fatal("Could not establish a connection to Redis:", err)
	}
	log.Println("Established connection to Redis:", result)

}

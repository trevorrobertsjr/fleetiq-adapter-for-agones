package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/gamelift"
	"github.com/go-redis/redis/v8"
)

const (
	port = "6379"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	// Using the Config value, create the GameLift client
	svcGame := gamelift.NewFromConfig(cfg)

	// Using the Config value, create the EC2 client
	svcEc2 := ec2.NewFromConfig(cfg)

	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL") + ":" + port,
	})
	result, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		log.Fatal("Could not establish a connection to Redis:", err)
	}
	fmt.Println("Established connection to Redis:", result)
	fmt.Println(svcGame.DescribeGameServerGroup(context.TODO(), &gamelift.DescribeGameServerGroupInput{GameServerGroupName: aws.String("agones-game-server-group-01")}))
	fmt.Println(svcEc2.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{}))
}

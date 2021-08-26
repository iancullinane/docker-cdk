package main

import (
	"log"

	"github.com/iancullinane/docker-cdk/src/config"
	"github.com/iancullinane/docker-cdk/src/internal/cdk"
)

func main() {

	log.Println("Get from github")
	log.Println("Do SAM/CDK")
	log.Println("Update deployments")

	contexts := map[string]string{
		"a": "A",
	}

	// ctx := context.Background()
	// ts := oauth2.StaticTokenSource(
	// 	&oauth2.Token{AccessToken: "... your access token ..."},
	// )
	// tc := oauth2.NewClient(ctx, ts)

	// client := github.NewClient(tc)

	// list all repositories for the authenticated user
	// repos, _, err := client.Repositories.List(ctx, "", nil)

	cdkWriter := cdk.New()
	lists, err := cdkWriter.List("iancullinane/cloud-config", contexts)
	if err != nil {
		panic(err)
	}

	log.Printf("%#v", lists)

	// Set up config
	var conf *config.Config
	conf = conf.BuildConfigFromFile("./config/base.yaml")
	log.Printf("%#v", conf)
	// sess := session.Must(session.NewSession())
	// // AWS config for client creation
	// awsConfigUsEast2 := &aws.Config{
	// 	CredentialsChainVerboseErrors: aws.Bool(true),
	// 	S3ForcePathStyle:              aws.Bool(true),
	// 	Region:                        aws.String("us-east-2"), // us-east-2 is the destination bucket region
	// }

	// logger := logrus.New()
	// sess := session.New()
	// sqsSvc := sqs.New(sess)
	// for {
	// 	res, err := sqsSvc.ReceiveMessage(&sqs.ReceiveMessageInput{
	// 		QueueUrl:            aws.String(os.Getenv("OPERATION_QUEUE_URL")),
	// 		MaxNumberOfMessages: aws.Int64(1),
	// 	})
	// 	if err != nil {
	// 		logger.Error("receive message error", zap.Error(err))
	// 		break
	// 	}

	// 	if len(res.Messages) == 0 {
	// 		break
	// 	}

	// 	// Get the message
	// 	msg := res.Messages[0]

	// 	// Delete the message
	// 	if _, err := sqsSvc.DeleteMessage(&sqs.DeleteMessageInput{
	// 		QueueUrl:      aws.String(os.Getenv("OPERATION_QUEUE_URL")),
	// 		ReceiptHandle: msg.ReceiptHandle,
	// 	}); err != nil {
	// 		logger.Error("delete message error", zap.Error(err))
	// 		break
	// 	}
	// 	var req events.APIGatewayProxyRequest
	// 	if err := json.Unmarshal([]byte(*msg.Body), &req); err != nil {
	// 		logger.Error("unmarshal error", zap.Error(err))
	// 		continue
	// 	}

	// Something about github
	// ctx := context.Background()
	// switch os.Getenv("PLATFORM") {
	// case "github":
	// 	err = github.Handler(ctx, req, logger)
	// default:
	// 	err = fmt.Errorf("unknown platform %s is setted", os.Getenv("PLATFORM"))
	// }
	// if err != nil {
	// 	logger.Error("operation error", zap.Error(err))
	// }
	// }

	// Connect to ECS
	// ecsSvc := ecs.New(sess)
	// if _, err := ecsSvc.UpdateService(&ecs.UpdateServiceInput{
	// 	Cluster:      aws.String(os.Getenv("TASK_ECS_CLUSTER_ARN")),
	// 	DesiredCount: aws.Int64(0),
	// 	Service:      aws.String(os.Getenv("OPERATION_SERVICE_ARN")),
	// }); err != nil {
	// 	logger.Error("shutdown task error", zap.Error(err))
	// }
}

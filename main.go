package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"fmt"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	ec2Svc := ec2.New(sess)
	result2, err2 := ec2Svc.DescribeLaunchTemplates(&ec2.DescribeLaunchTemplatesInput{
		LaunchTemplateNames: []*string{aws.String("notinexistence")},
	})
	if err2 != nil {
		fmt.Println("Error", err2)
		if aerr, ok := err2.(awserr.Error); ok {
			fmt.Println("awserr.Error.Code", aerr.Code())
		}
	} else {
		fmt.Println("Success", result2)
	}
}

package main

import (
	"os"
	"flag"
	"fmt"
	"encoding/base64"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

func main() {
	regionFlag := flag.String("region", "", "AWS region")
	secretFlag := flag.String("secret", "", "KMS secret")
    flag.Parse()
    region := *regionFlag
	kms_secret := *secretFlag

	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	svc := kms.New(sess)
	data, err := base64.StdEncoding.DecodeString(kms_secret)
	result, err := svc.Decrypt(&kms.DecryptInput{CiphertextBlob: data})

	if err != nil {
		fmt.Println("Error Decrypting data", err)
		os.Exit(1)
	}
	fmt.Println(string(result.Plaintext))
}
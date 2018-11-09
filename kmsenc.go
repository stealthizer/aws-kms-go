package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/kms"
    "encoding/base64"
    "flag"
    "fmt"
    "os"
)

func main() {

    regionFlag := flag.String("region", "", "AWS region")
    textFlag := flag.String("text", "", "Text to encrypt")
    keyidFlag := flag.String("keyid", "", "arn from the kms key to use")
    profileFlag := flag.String("profile", "", "AWS profile to use")

    flag.Parse()
    region := *regionFlag
    text := *textFlag
    keyId := *keyidFlag
    profile := *profileFlag


    var sess *session.Session
    var err error

    if profile == "" {
        sess, err = session.NewSessionWithOptions(session.Options{
                Config: aws.Config{Region: aws.String(region)},
            })

    } else {
        sess, err = session.NewSessionWithOptions(session.Options{
                Config: aws.Config{Region: aws.String(region)},
                Profile: profile,
            })
    }

    svc := kms.New(sess)

    // Encrypt the data
    result, err := svc.Encrypt(&kms.EncryptInput{
        KeyId: aws.String(keyId),
        Plaintext: []byte(text),
    })

    if err != nil {
        fmt.Println("Got error encrypting data: ", err)
        os.Exit(1)
    }

    base64 := base64.StdEncoding.EncodeToString(result.CiphertextBlob)
    fmt.Println("Blob (base-64 byte array):")
    fmt.Println(base64)
}
# aws-kms-go
Allows encoding and decoding of text from the command line using the AWS KMS service

###Encoding

To encode a secret using kmsenc you need:
* The text to encrypt
* The region where the key is located
* The key's ARN

Usage example:

```
$ kmsenc -region eu-west-1 -text ThisisaSecret -keyid arn:aws:kms:eu-west-1:0123456789:key/keykeykey
```

output:

A base64 representation of the encrypted secret that can be securely stored.

###Decoding

To decode a secret using kmsdec you need:
* The secret text in base64
* The region where to find the key to decrypt

Usage example:

```
$ kmsdec -region eu-west-1 -secret AQUIACAHFDKJEWKUHRLKUWEUGHRK==
```

output:

The original text decoded
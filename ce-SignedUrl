package main

import (
  "context"
  "fmt"
  "time"

  "cloud.google.com/go/storage"
  "cloud.google.com/go/iam/credentials/apiv1"
  credentialspb "google.golang.org/genproto/googleapis/iam/credentials/v1"
)

const (
  bucketName = "bucket-name"
  objectName = "object"
  serviceAccount = "[PROJECTNUMBER]-compute@developer.gserviceaccount.com"
)

func main() {
  ctx := context.Background()

  c, err := credentials.NewIamCredentialsClient(ctx)
  if err != nil {
     panic(err)
  }

  opts := &storage.SignedURLOptions{
     Method: "GET",
     GoogleAccessID: serviceAccount,
     SignBytes: func(b []byte) ([]byte, error) {
        req := &credentialspb.SignBlobRequest{
            Payload: b,
            Name: serviceAccount,
        }
        resp, err := c.SignBlob(ctx, req)
        if err != nil {
           panic(err)
        }
        return resp.SignedBlob, err
     },
     Expires: time.Now().Add(15*time.Minute),
  }

  u, err := storage.SignedURL(bucketName, objectName, opts)
  if err != nil {
     panic(err)
  }

  fmt.Printf("\"%v\"", u)
}

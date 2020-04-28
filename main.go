package main

import (
  "context"
  "flag"
  "io"
  "os"

  "cloud.google.com/go/storage"
)

var key = flag.String("key", "", "The key to upload to. (required)")
var bucket = flag.String("bucket", "", "The bucket to upload to. (required)")
var localFile = flag.String("localfile", "", "The local file to upload. (required)")

func main() {

  flag.Parse()
  if len(*key) == 0 || len(*bucket) == 0 || len(*localFile) == 0 {
    flag.Usage()
    os.Exit(-1)
  }


  file, err := os.Open(*localFile)
  if err != nil {
    println("Could not open localfile: ", *localFile)
    os.Exit(-2)
  }
  defer file.Close()

  ctx := context.Background()

  client, err := storage.NewClient(ctx)
  if err != nil {
    println("Error setting up storage client, err: ", err.Error())
    os.Exit(-3)
  }

  newWriter := client.Bucket(*bucket).Object(*key).NewWriter(ctx)
  defer newWriter.Close()

  res, err := io.Copy(newWriter, file)
  if err != nil {
    println("Error during upload, err: ", err.Error())
    os.Exit(-4)
  }

  if res == 0 {
    os.Exit(-5)
  } 

}


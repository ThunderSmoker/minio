package main

import (
    "log"
    "os"
    "example.com/minio_project/operations" // Import the operations package
)

func main() {
    // Replace these values with your MinIO server endpoint, access key, and secret key
    endpoint := "192.168.49.2:9000"
    accessKey := "minio"
    secretKey := "minio123"

    // Create Operations instance
    operations, err := operations.NewOperations(endpoint, accessKey, secretKey)
    if err != nil {
        log.Fatalln("Error initializing MinIO operations:", err)
    }

    // Replace "myfirstbucket" with your bucket name
    bucketName := "myfirstbucket"
    if len(os.Args) < 2 {
		// List files in the bucket if no file name is provided
		err = operations.ListFiles(bucketName)
		if err != nil {
			log.Fatalln("Error listing files:", err)
		}
	} else {
		// If file name is provided, delete the file from the bucket
		objectName := os.Args[1]
		err = operations.DeleteFile(bucketName, objectName)
		if err != nil {
			log.Fatalln("Error deleting file:", err)
		}
		log.Printf("File %s deleted successfully!\n", objectName)
	}
    // List files in the bucket
    // err = operations.ListFiles(bucketName)
    // if err != nil {
    //     log.Fatalln("Error listing files:", err)
    // }
	// objectName := "backupturtlebot.txt"
    // // Replace "newContent" with the content you want to write to the file
    // newContent := "This is the new content of the file."

    // Modify file in the bucket
    // err = operations.ModifyFile(bucketName, objectName, newContent)
    // if err != nil {
    //     log.Fatalln("Error modifying file:", err)
    // }

    // log.Println("File modified successfully!")
}

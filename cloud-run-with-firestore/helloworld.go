package main

import (
        "cloud.google.com/go/firestore"
        "context"
        "fmt"
        "log"
        "net/http"
        "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
        log.Print("Hello world received a request.")
        msg := getMessage()
        fmt.Fprintf(w, msg)
}

func getMessage() string {
        projectId := "snippets-273801"
        ctx := context.Background()
        client, err := firestore.NewClient(ctx, projectId);
        if err != nil {
                log.Fatalf("Failed to create firestore client: %v", err)
        }
        defer client.Close()

        doc, err := client.Doc("testing/message").Get(ctx)
        if err != nil {
                log.Fatalf("Failed to retrieve document: %v", err)
        }

        msg, ok := doc.Data()["greeting"].(string)
        if !ok {
                log.Fatalf("greeting is not a string")
        }
        return msg

}

func main() {
        log.Print("Hello world sample started.")
        http.HandleFunc("/", handler)

        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
        }

        log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
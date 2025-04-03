package main

import (
    "fmt"
    "github.com/yourname/go-user-verifier/models"
    "github.com/yourname/go-user-verifier/verifier"
)

func main() {
    // Simulated incoming user event (like from Kafka)
    incomingUser := models.User{
        ID:        "USR12345",
        FirstName: "Riz",
    }

    verifiedUser, err := verifier.VerifyUser(incomingUser)
    if err != nil {
        fmt.Println("Verification failed:", err)
        return
    }

    // Simulate publishing to Kafka topic (placeholder)
    fmt.Printf("Publishing verified user: %+v\n", verifiedUser)
}

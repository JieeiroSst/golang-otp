package main

import (
    "time"

    "github.com/pquerna/otp/totp"

    "bufio"
    "fmt"
    "os"
)

func promptForPasscode() string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter Passcode: ")
    text, _ := reader.ReadString('\n')
    return text
}

func main() {

    keySecret := "NK4KFDHUGRGMFKFEWRJ5EEOV6FT2IAKE"

	count := 1
	now := time.Now()
	timestamp := now.Add(time.Duration(-count) * time.Second)

    // coded, _ := totp.GenerateCode(keySecret, time.Now().UTC())	
	coded, _ := totp.GenerateCode(keySecret, timestamp)
    fmt.Println("code :", coded)

    fmt.Println("Validating TOTP...")
    // Now Validate that the user's successfully added the passcode.
    passcode := promptForPasscode()
    valid := totp.Validate(passcode, keySecret)
    if valid {
        println("Valid passcode!")
        os.Exit(0)
    } else {
        println("Invalid passocde!")
        os.Exit(1)
    }
}
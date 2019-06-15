package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "github.com/bwmarrin/discordgo"
)

var (
    token string = "Bot " + os.Getenv("token")
    prefix string = ":"
)

func main() {

    // Create a new Discord session with the provided token.
    dg, err := discordgo.New(token)
    if err != nil {
        fmt.Println("Error creating Discord session...", err)
        return
    }

    // Open Websocket connection.
    err = dg.Open()
    if err != nil {
        fmt.Println("Error opening connection...", err)
        return
    }

    // Wait here until Ctrl-C or some other term signal.
    fmt.Println("Bot is now running.  Press Ctrl-C to exit.")
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-sc

    // Close connection.
    dg.Close()

}

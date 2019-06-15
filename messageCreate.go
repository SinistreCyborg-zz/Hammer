package main

import (
    "strings"
    dg "github.com/bwmarrin/discordgo"
    "hammer/commands"
)

func messageCreate(s *dg.Session, m *dg.MessageCreate) {

    // Ignore messages that don't start with the prefix.
    if strings.HasPrefix(m.Content, prefix) == false {
        return
    }

    // Useful when responding to commands.
    command := strings.Split(strings.TrimPrefix(m.Content, prefix), " ")[0]
    args := strings.Split(strings.TrimPrefix(m.Content, prefix + command), " ")

    if command == "ping" {
        commands.Ping(s, m, args)
        return
    }

}
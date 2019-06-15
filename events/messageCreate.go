package events

import (
    "strings"
    dg "github.com/bwmarrin/discordgo"
    "hammer/commands"
)

var prefix string = ":"

func MessageCreate(s *dg.Session, m *dg.MessageCreate) {

    // Ignore messages that don't start with the prefix.
    if strings.HasPrefix(m.Content, prefix) == false {
        return
    }

    // Useful when responding to commands.
    command := strings.Split(strings.TrimPrefix(m.Content, prefix), " ")[0]
    args := strings.Split(strings.TrimSpace(strings.TrimPrefix(m.Content, prefix + command)), " ")

    if command == "ping" {
        commands.Ping(s, m, args)
        return
    }

    if command == "gping" {
        commands.Gping(s, m, args)
        return
    }

    if command == "color" {
        commands.Color(s, m, args)
        return
    }

    if command == "tzone" {
        commands.Tzone(s, m, args)
        return
    }

}

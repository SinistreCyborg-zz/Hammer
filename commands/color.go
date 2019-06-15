package commands

import (
    "fmt"
    "strings"
    dg "github.com/bwmarrin/discordgo"
)

func Color(s *dg.Session, m *dg.MessageCreate, args []string) {

    // Get the guild the command was in.
    guild, err := s.Guild(m.GuildID)
    if err != nil {
        fmt.Println("An error occurred...", err)
        return
    }

    // Get all the colors.
    colors := make([]*dg.Role, 0)
    for _, role := range guild.Roles {
        if strings.HasPrefix(role.Name, "#") {
            colors = append(colors, role)
        }
    }

    // Set a color.
    if args[0] == "set" {

        for _, color := range colors {
            if strings.TrimPrefix(color.Name, "#") == strings.Join(args[1:], " ") {

                // Add the color, if such a color exists.
                err := s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, color.ID)
                if err != nil {
                    fmt.Println("An error occurred...", err)
                    return
                }

                // Inform the user that the color was added.
                s.ChannelMessageSend(m.ChannelID, "You're now: " + color.Mention())
                return

            }
        }

        // Send error message.
        s.ChannelMessageSend(m.ChannelID, "I couldn't find a color by that name!")
        return

    }

    // Slice of all the colors' mentions.
    names := make([]string, 0)
    for _, color := range colors {
        names = append(names, color.Mention())
    }

    // Join the mentions and send it.
    s.ChannelMessageSend(m.ChannelID, "Here are the colors: " + strings.Join(names, " "))
    return

}

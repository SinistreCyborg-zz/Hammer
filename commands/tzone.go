package commands

import (
    "fmt"
    "strings"
    dg "github.com/bwmarrin/discordgo"
)

func Tzone(s *dg.Session, m *dg.MessageCreate, args []string) {

    // Get the guild the command was in.
    guild, err := s.Guild(m.GuildID)
    if err != nil {
        fmt.Println("An error occurred...", err)
        return
    }

    // Get all the timezones.
    tzones := make([]*dg.Role, 0)
    for _, role := range guild.Roles {
        if strings.HasSuffix(role.Name, "⏳") {
            tzones = append(tzones, role)
        }
    }

    // Set a timezone.
    if args[0] == "set" {

        for _, tzone := range tzones {
            if strings.TrimSuffix(tzone.Name, "⏳") == strings.Join(args[1:], " ") {

                // Add the timezone, if such a timezone exists.
                err := s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, tzone.ID)
                if err != nil {
                    fmt.Println("An error occurred...", err)
                    return
                }

                // Inform the user that the timezone was added.
                s.ChannelMessageSend(m.ChannelID, "Your timezone is now set to: " + tzone.Mention())
                return

            }
        }

        // Send error message.
        s.ChannelMessageSend(m.ChannelID, "I couldn't find a timezone by that name!")
        return

    }

    // Slice of all the timezones' mentions.
    names := make([]string, 0)
    for _, tzone := range tzones {
        names = append(names, tzone.Mention())
    }

    // Join the mentions and send it.
    s.ChannelMessageSend(m.ChannelID, "Here are the timezones: " + strings.Join(names, " "))
    return

}

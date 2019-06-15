package commands

import (
    "strings"
    dg "github.com/bwmarrin/discordgo"
)

func Color(s *dg.Session, m *dg.MessageCreate, args []string) {

    // Get the guild the command was in.
    guild, _ := s.Guild(m.GuildID)

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
            if strings.ToLower(strings.TrimPrefix(color.Name, "#")) == strings.ToLower(strings.Join(args[1:], " ")) {

                // Get the member object.
                member, _ := s.GuildMember(m.GuildID, m.Author.ID)

                // Remove all color roles.
                for _, role := range member.Roles {
                    for _, color := range colors {
                        if role == color.ID {
                            s.GuildMemberRoleRemove(m.GuildID, m.Author.ID, color.ID)
                        }
                    }
                }

                // Add the color, if such a color exists.
                s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, color.ID)

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

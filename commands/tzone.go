package commands

import (
    "strings"
    dg "github.com/bwmarrin/discordgo"
)

func Tzone(s *dg.Session, m *dg.MessageCreate, args []string) {

    // Get the guild the command was in.
    guild, _ := s.Guild(m.GuildID)

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
            if strings.ToLower(strings.TrimSuffix(tzone.Name, "⏳")) == strings.ToLower(strings.Join(args[1:], " ")) {

                // Get the member object.
                member, _ := s.GuildMember(m.GuildID, m.Author.ID)

                // Remove all timezone roles.
                for _, role := range member.Roles {
                    for _, tzone := range tzones {
                        if role == tzone.ID {
                            s.GuildMemberRoleRemove(m.GuildID, m.Author.ID, tzone.ID)
                        }
                    }
                }

                // Add the timezone, if such a timezone exists.
                s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, tzone.ID)

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

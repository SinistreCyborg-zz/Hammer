package commands

import (
    "fmt"
    dg "github.com/bwmarrin/discordgo"
)

func Gping(s *dg.Session, m *dg.MessageCreate, args []string) {

    member, err := s.GuildMember(m.GuildID, m.Author.ID)
    if err != nil {
        fmt.Println("An error occurred...", err)
        return
    }

    if sliceContains("517081361476091911", member.Roles) {
        err := s.GuildMemberRoleRemove(m.GuildID, m.Author.ID, "517081361476091911")
        if err != nil {
            fmt.Println("An error occurred...", err)
            return
        }
    } else {
        err := s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, "517081361476091911")
        if err != nil {
            fmt.Println("An error occurred...", err)
            return
        }
    }

    s.ChannelMessageSend(m.ChannelID, member.Mention() + " You've toggled your Game Ping role.")
    return

}

func sliceContains(val string, list []string) bool {
    for _, v := range list {
        if v == val {
            return true
        }
    }
    return false
}

package commands

import dg "github.com/bwmarrin/discordgo"

func Ping(s *dg.Session, m *dg.MessageCreate, args []string) {
    s.ChannelMessageSend(m.ChannelID, "🏓 Pong!")
    return
}

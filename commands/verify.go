package commands

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "strings"
    "strconv"
    dg "github.com/bwmarrin/discordgo"
)

type VerifyInfo struct {
    Status         string `json:"status"`
    RobloxUsername string `json:"robloxUsername"`
    RobloxID       int    `json:"robloxId"`
}

type RBLXResponse struct {
    Data []thumbnail `json:"data"`
}

type RBLXFriends struct {
    Data []friend `json:"data"`
}

type thumbnail struct {
    TargetID int    `json:"targetId"`
    State    string `json:"string"`
    ImageURL string `json:"imageUrl"`
}

type friend struct {
    RobloxID int `json:"id"`
    RobloxUsername string `json:"name"`
    Description string `json:"description"`
    Created string `json:"created"`
}

func get(url string) []byte {

    res, err := http.Get(url)
    if err != nil {
        fmt.Println("An error occurred...", err)
        return nil
    }

    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
        return nil
    }

    return body

}

func Verify(s *dg.Session, m *dg.MessageCreate, args []string) {

    // Figure out whether the user is verified or not.
    var info VerifyInfo
    json.Unmarshal(get("https://verify.eryn.io/api/user/" + m.Author.ID), &info)

    // Tell user to verify again.
    if info.Status != "ok" || strings.Contains(strings.Join(args, " "), "--force") {
        s.ChannelMessageSend(m.Author.ID, "Uh, oh! You're not verified! Follow the instructions here: https://verify.eryn.io/\nThen, run the verify command again inside of the Hammer Squad Discord.")
        s.ChannelMessageSend(m.ChannelID, "<@" + m.Author.ID + ">, I've DMed you instructions on how to verify.")
        return
    }

    // Get user's thumbnail.
    var thumb RBLXResponse
    json.Unmarshal(get("https://thumbnails.roblox.com/v1/users/avatar?userIds=" + strconv.Itoa(info.RobloxID) + "&size=720x720&format=Png"), &thumb)

    // Get user's friends.
    var friends RBLXFriends
    json.Unmarshal(get("https://friends.roblox.com/v1/users/" + strconv.Itoa(info.RobloxID) + "/friends"), &friends)

    // Send info about the user.
    s.ChannelMessageSendEmbed(m.ChannelID, &dg.MessageEmbed{
        Color: 0x34363b,
        Description: "You're verified as: **" + info.RobloxUsername + "**\n\n" + strconv.Itoa(len(friends.Data)) + " Friends",
        Thumbnail: &dg.MessageEmbedThumbnail{
            URL: thumb.Data[0].ImageURL,
        },
    })

}

package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func onMessageCreate(aSession *discordgo.Session, aMessage *discordgo.MessageCreate) {
	if aMessage.Author.Bot {
		return
	}

	if tArgs := strings.Fields(aMessage.Content); len(tArgs) == 2 && tArgs[0] == "!help" && tArgs[1] == "randomizer" {
		onHelp(aSession, aMessage)
	} else if strings.HasPrefix(aMessage.Content, "!random") {
		onRandom(aSession, aMessage, aMessage.Content)
	} else if strings.HasPrefix(aMessage.Content, "!select") {
		onSelect(aSession, aMessage, aMessage.Content)
	}
}

func onHelp(aSession *discordgo.Session, aMessage *discordgo.MessageCreate) {
	tResponse := "機能説明\n\n" +
		"**!random**\n`!random mdn`でn面ダイスをm個振ります\n" +
		"**!select**\n`!select v1 v2 ... vn`でv1からvnのうち1つをランダムに出力します\n" +
		"**!help randomizer**\nヘルプを表示します"

	_, err := aSession.ChannelMessageSend(aMessage.ChannelID, tResponse)
	if err != nil {
		log.Printf("Error sending help message: %+v\n", err)
	}
}

func onRandom(aSession *discordgo.Session, aMessage *discordgo.MessageCreate, aStr string) {
	tStr := aStr[7:]
	tStr = strings.TrimSpace(tStr)

	tParts := strings.Split(strings.ToUpper(tStr), "D")
	if len(tParts) != 2 {
		_, _ = aSession.ChannelMessageSend(aMessage.ChannelID, "Error: input mdn format (ex. !random 2d6)")
		return
	}

	m, tErrorM := strconv.Atoi(tParts[0])
	n, tErrorN := strconv.Atoi(tParts[1])
	if tErrorM != nil || tErrorN != nil || m <= 0 || n <= 0 {
		_, _ = aSession.ChannelMessageSend(aMessage.ChannelID, "Error: m and n must be positive integer (ex. !random 2d6)")
		return
	}

	tResults := make([]int, 0)
	tTotal := 0
	for i := 0; i < m; i++ {
		tRoll := rand.Intn(n) + 1
		tResults = append(tResults, tRoll)
		tTotal += tRoll
	}

	tResultStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(tResults)), ", "), "[]")
	tResponse := fmt.Sprintf("Result of %dd%d: [%s] (sum: %d)", m, n, tResultStr, tTotal)

	_, tError := aSession.ChannelMessageSend(aMessage.ChannelID, tResponse)
	if tError != nil {
		log.Printf("Error sending random message: %+v\n", tError)
	}
}

func onSelect(aSession *discordgo.Session, aMessage *discordgo.MessageCreate, aStr string) {
	tStr := aStr[7:]
	tParts := strings.Fields(tStr)

	tIndex := rand.Intn(len(tParts))
	tResponse := fmt.Sprintf("Selected: %s", tParts[tIndex])

	_, tError := aSession.ChannelMessageSend(aMessage.ChannelID, tResponse)
	if tError != nil {
		log.Printf("Error sending select message: %+v\n", tError)
	}
}

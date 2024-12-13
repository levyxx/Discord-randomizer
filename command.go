package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

func onCommand(aSession *discordgo.Session, aInteraction *discordgo.InteractionCreate) {
	switch aInteraction.ApplicationCommandData().Name {
	case "help":
		onHelpCommand(aSession, aInteraction)
	case "random":
		onRandomCommand(aSession, aInteraction)
	case "select":
		onSelectCommand(aSession, aInteraction)
	default:
		break
	}
}

func onHelpCommand(aSession *discordgo.Session, aInteraction *discordgo.InteractionCreate) {
	tFooter := discordgo.MessageEmbedFooter{Text: "created by levyxx (https://github.com/levyxx)"}

	tEmbed := discordgo.MessageEmbed{
		Title: "機能説明",
		Description: "任意のn面ダイスを複数個同時に振ることができます\n\n" +
			"**__各種コマンド__**\n\n" +
			"**random**\n`/random mdn`でn面ダイスをm個振ります\n" +
			"**select**\n`/select v1 v2 ... vn`でv1からvnのうち1つをランダムに出力します\n" +
			"**help**\nヘルプを表示します",
		Footer: &tFooter,
	}

	tError := aSession.InteractionRespond(aInteraction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{&tEmbed},
		},
	})

	if tError != nil {
		log.Printf("%+v\n", errors.Wrap(tError, "failed to response to interaction"))
	}
}

func onRandomCommand(aSession *discordgo.Session, aInteraction *discordgo.InteractionCreate) {
	tArgs := aInteraction.ApplicationCommandData().Options
	if len(tArgs) != 1 {
		aSession.InteractionRespond(aInteraction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Error: input mdn format (ex. 2d6)",
			},
		})
		return
	}

	tParts := strings.Split(strings.ToUpper(tArgs[0].StringValue()), "D")
	if len(tParts) != 2 {
		aSession.InteractionRespond(aInteraction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Error: input mdn format (ex. 2d6)",
			},
		})
		return
	}

	m, tErrorM := strconv.Atoi(tParts[0])
	n, tErrorN := strconv.Atoi(tParts[1])
	if tErrorM != nil || tErrorN != nil || m <= 0 || n <= 0 {
		aSession.InteractionRespond(aInteraction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Error: m and n must be positive integer (ex. 2d6)",
			},
		})
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
	tResponse := fmt.Sprintf("result of %dd%d: [%s] (sum: %d)", m, n, tResultStr, tTotal)

	aSession.InteractionRespond(aInteraction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: tResponse,
		},
	})
}

func onSelectCommand(aSession *discordgo.Session, aInteraction *discordgo.InteractionCreate) {
	tArgs := aInteraction.ApplicationCommandData().Options
	if len(tArgs) != 1 {
		aSession.InteractionRespond(aInteraction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Error: input arguments",
			},
		})
		return
	}

	tParts := strings.Split(tArgs[0].StringValue(), " ")
	if len(tParts) == 1 {
		tParts = strings.Split(tArgs[0].StringValue(), "　")
	}
	tIndex := rand.Intn(len(tParts))
	tResponse := tParts[tIndex]

	aSession.InteractionRespond(aInteraction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: tResponse,
		},
	})
}

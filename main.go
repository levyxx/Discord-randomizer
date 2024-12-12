package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type config struct {
	DiscordToken string `env:"DISCORD_TOKEN,required"`
}

func main() {
	// .envファイルをロード
	if tError := godotenv.Load(); tError != nil {
		log.Println("No .env file found, falling back to system environment variables")
		return
	}

	//トークンの構成を環境変数から読み込む
	tConfig := config{}
	if tError := env.Parse(&tConfig); tError != nil {
		log.Printf("%+v\n", errors.Wrap(tError, "failed to parse env"))
		return
	}
	log.Println("success to parse env(tokens)")

	tDG, tError := discordgo.New("Bot " + tConfig.DiscordToken)
	if tError != nil {
		log.Printf("%+v\n", errors.Wrap(tError, "failed to create discord session"))
		return
	}

	tError = tDG.Open()
	if tError != nil {
		log.Printf("%+v\n", errors.Wrap(tError, "failed to open connection"))
		return
	}
	defer tDG.Close()
	log.Println("Bot is now running")

	registerCommands(tDG)
	tDG.AddHandler(onCommand)

	select {}
}

func registerCommands(tDG *discordgo.Session) {
	tCommands := []*discordgo.ApplicationCommand{
		{
			Name:        "help",
			Description: "ヘルプを表示します",
		},
		{
			Name:        "random",
			Description: "任意のn面ダイスをm個振ります (例: /random 2D6)",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "mDn",
					Description: "ダイスの形式（例: 2D6）",
					Required:    true,
				},
			},
		},
	}

	for _, tCommand := range tCommands {
		_, tError := tDG.ApplicationCommandCreate(tDG.State.User.ID, "", tCommand)
		if tError != nil {
			log.Fatalf("Cannot create '%s' command: %v", tCommand.Name, tError)
		}
		log.Printf("Command '%s' registered.", tCommand.Name)
	}
}

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

	// メッセージ内容を受信するIntentを追加
	tDG.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages

	tError = tDG.Open()
	if tError != nil {
		log.Printf("%+v\n", errors.Wrap(tError, "failed to open connection"))
		return
	}
	defer tDG.Close()
	log.Println("Bot is now running")

	registerCommands(tDG)
	tDG.AddHandler(onCommand)
	tDG.AddHandler(onMessageCreate)

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
			Description: "任意のn面ダイスをm個振ります (例: /random 2d6)",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "mdn",
					Description: "ダイスの形式（例: 2d6）",
					Required:    true,
				},
			},
		},
		{
			Name:        "select",
			Description: "与えられた文字列のうち1つだけをランダムに出力します",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "arguments",
					Description: "スペース区切りで複数の文字列を入力",
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

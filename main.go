package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

const botPrefix string = "!go"

func getDiscordSession() *discordgo.Session {
	godotenv.Load(".env")
	token := os.Getenv("BOT_TOKEN")
	session, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
	}
	return session
}

func reply(sess *discordgo.Session){
	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		args := strings.Split(m.Content, " ")
		fmt.Println(args)
		if args[0] != botPrefix {
			if strings.ToLower(args[0]) == "hello" || strings.ToLower(args[0]) == "hi" {
				embed := discordgo.MessageEmbed{
					Title: "Hiii, Glad to get a text from you.\n If you wanna ask something please use '" + botPrefix + "' as prefix",
				}
				// s.ChannelMessageSend(m.ChannelID, )
				s.ChannelMessageSendEmbed(m.ChannelID, &embed)
			}
			return
		}

		if len := len(args); len == 1 {
			s.ChannelMessageSend(m.ChannelID, "Thanks for calling me, please ask something")
			return
		}

		if strings.ToLower(args[1]) == "hello" || strings.ToLower(args[1]) == "hi" {
			embed := discordgo.MessageEmbed{
				Title: "Hiii, Glad to get a text from you.\n If you wanna ask something please use '" + botPrefix + "' as prefix",
			}
			// s.ChannelMessageSend(m.ChannelID, )
			s.ChannelMessageSendEmbed(m.ChannelID, &embed)
			return
		}
	})
}

func main() {
	
	sess := getDiscordSession()
	reply(sess)
	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err := sess.Open()

	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("Bot is Online")

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-sc
}


package main

import (
	"bufio"
	"flag"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

var (
	token string
	phraseFile string
	phrases []string
)

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.StringVar(&phraseFile, "p", "phrases.txt", "Path to phrase text file")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
}

func main() {
	if token == "" {
		log.Fatalln("No token provided. Please run: chatterbot -t <bot token>")
		return
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalln("error creating Discord session,", err)
		return
	}

	err = loadPhrases(phraseFile)
	if err != nil {
		log.Errorln("error loading phrases", err)
	}

	dg.AddHandler(ready)

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		log.Fatalln("error opening connection,", err)
		return
	}

	log.Info("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func loadPhrases(src string) error {
	file, err := os.Open(src)
    if err != nil {
        log.Errorln("error opening", src, err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		phrases = append(phrases, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Errorln(err)
    }

	log.Infoln("loaded", len(phrases), "phrases from", src)

	return err
}

func randomChoice(choices []string) string {
	return choices[rand.Intn(len(choices))]
}

// This function will be called (due to AddHandler above) when the bot receives
// the "ready" event from Discord.
func ready(s *discordgo.Session, event *discordgo.Ready) {
	// Set the playing status.
	s.UpdateGameStatus(0, "chatterbot")
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if rand.Intn(3) == 0 {
		s.ChannelMessageSend(m.ChannelID, randomChoice(phrases))
	}
}
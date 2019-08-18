package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"flag"

	"github.com/Necroforger/dgwidgets"
	"github.com/bwmarrin/discordgo"
	"github.com/postables/go-compound/client"
	"github.com/tbruyelle/imgur"
)

var (
	configFile  = flag.String("config.file", "config.json", "config file to use")
	imgurClient *imgur.Client
	url         = "https://api.compound.finance/api/v2"
	help        string
	helpEmbed   *discordgo.MessageEmbed
)

func init() {
	flag.Parse()
	helpEmbed = &discordgo.MessageEmbed{
		Title:       "MoneyBags Help Menu",
		Description: "all commands must be invoked with !moneybags <cmd>\nAnything with <..> after command name expects an argument\nAnything with [..] after command name is an optional argument",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://gateway.temporal.cloud/ipfs/QmSky8KsZ6q9zz6kmj3TrbNTTvwhGtGmVbYB9iXWLPD5VC",
		},
		Color: 0x00ff00,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  "eth-price",
				Value: "get ethereum price from cmc",
			},
			&discordgo.MessageEmbedField{
				Name:  "dai-price",
				Value: "get dai price from cmc",
			},
			&discordgo.MessageEmbedField{
				Name:  "liqqable [opts]",
				Value: "get liquidatable accounts. opts: paginated (optional)",
			},
			&discordgo.MessageEmbedField{
				Name:  "heatlh-check <acct>",
				Value: "get account health",
			},
			&discordgo.MessageEmbedField{
				Name:  "collateral-value <acct>",
				Value: "get the account collateral value",
			},
			&discordgo.MessageEmbedField{
				Name:  "borrow-value <acct>",
				Value: "get the account borrow value",
			},
		},
	}
}

func main() {
	var discordToken = os.Getenv("DISCORD_TOKEN")
	if discordToken == "" {
		log.Fatal("DISCORD_TOKEN env var is empty")
	}
	// we need to prepend Bot to allow discord
	// to assign permissions properly
	dg, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		fmt.Println("failed to authenticate with discord")
		log.Fatal(err)
	}
	dg.AddHandler(messageCreate)
	if err := dg.Open(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("bot is now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func sendHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	if _, err := s.ChannelMessageSendEmbed(m.ChannelID, helpEmbed); err != nil {
		fmt.Println("error sending message: ", err.Error())
		return
	}
}

func liqqable(s *discordgo.Session, m *discordgo.MessageCreate) {
	cl := client.NewClient(url)
	accts, err := cl.GetLiquidatableAccounts()
	if err != nil {
		fmt.Println("failed to get liquidatable accounts ", err.Error())
		return
	}
	s.ChannelMessageSend(m.ChannelID, "fetching liqqable accounts...")
	p := dgwidgets.NewPaginator(s, m.ChannelID)
	for k, v := range accts {
		// get the account collateral value
		collatValue, err := cl.GetTotalCollateralValueInEth(k)
		if err != nil {
			fmt.Println("failed to get data for account ", k)
			continue
		}
		borrowValue, err := cl.GetTotalBorrowValueInEth(k)
		if err != nil {
			fmt.Println("failed to get data for account ", k)
			continue
		}
		var fields []*discordgo.MessageEmbedField
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Account Address",
			Value:  k,
			Inline: true,
		}, &discordgo.MessageEmbedField{
			Name:   "Account Value (ETH)",
			Value:  fmt.Sprintf("%v", borrowValue+collatValue),
			Inline: true,
		}, &discordgo.MessageEmbedField{
			Name:   "Account Health",
			Value:  fmt.Sprintf("%v", v),
			Inline: true,
		})
		embed := &discordgo.MessageEmbed{
			Author:      &discordgo.MessageEmbedAuthor{},
			Color:       0x00ff00,
			Description: "Liquidatable Account",
		}
		embed.Fields = fields
		p.Add(embed)
	}
	p.SetPageFooters()

	// enable looping
	p.Loop = true

	// When the paginator is done listening set the colour to yellow
	p.ColourWhenDone = 0xffff

	// Stop listening for reaction events after five minutes
	p.Widget.Timeout = time.Minute * 1

	// start
	p.Spawn()
}

func handleCommand(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) == 1 && args[0] == "!moneybags" {
		sendHelp(s, m)
		return
	}
	if len(args) == 2 {
		switch args[1] {
		case "eth-price":
			ethPrice(s, m)
		case "dai-price":
			daiPrice(s, m)
		case "liqqable":
			liqqable(s, m)
		}
	}
	if len(args) > 2 {
		switch args[1] {
		case "health-check":
			healthCheck(s, m, args[2])
		case "collateral-value":
			collateralValue(s, m, args[2])
		case "borrow-value":
			borrowValue(s, m, args[2])
		}
	}
}

func ethPrice(s *discordgo.Session, m *discordgo.MessageCreate) {
	val, err := RetrieveUsdPrice("ethereum")
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "[ERROR]: failed to get eth price: "+err.Error())
	}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("eth price: $%.3fUSD\n", val))
}

func daiPrice(s *discordgo.Session, m *discordgo.MessageCreate) {
	val, err := RetrieveUsdPrice("dai")
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "[ERROR]: failed to get dai price: "+err.Error())
	}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("eth price: $%.3fUSD\n", val))
}

func borrowValue(s *discordgo.Session, m *discordgo.MessageCreate, account string) {
	cl := client.NewClient(url)
	value, err := cl.GetTotalBorrowValueInEth(account)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "[ERROR]: failed to get collateral value: "+err.Error())
		return
	}
	if _, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("borrow value: %vETH", value)); err != nil {
		fmt.Println(err.Error())
		return
	}
}

func collateralValue(s *discordgo.Session, m *discordgo.MessageCreate, account string) {
	cl := client.NewClient(url)
	value, err := cl.GetTotalCollateralValueInEth(account)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "[ERROR]: failed to get collateral value: "+err.Error())
		return
	}
	if _, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("collateral value: %vETH", value)); err != nil {
		fmt.Println(err.Error())
		return
	}
}

func healthCheck(s *discordgo.Session, m *discordgo.MessageCreate, account string) {
	cl := client.NewClient(url)
	resp, err := cl.GetAccount(account)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "[ERROR]: failed to get account health: "+err.Error())
		return
	}
	if len(resp.Accounts) == 0 {
		s.ChannelMessageSend(m.ChannelID, "[ERROR]: an unexpected error occurred")
	}
	health, err := strconv.ParseFloat(resp.Accounts[0].Health.Value, 64)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "[ERROR]: failed to get account health: "+err.Error())
		return
	}
	var status string
	if health < 1.0 {
		status = "danger"
	} else if health < 1.15 {
		status = "risky"
	} else if health < 1.25 {
		status = "okay"
	} else {
		status = "good"
	}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("health: %v\nstatus: %s\n", health, status))
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// parse the message contents based off string fields
	args := strings.Fields(m.Content)
	if len(args) == 0 {
		return
	}
	// ensure the first field is a valid invocation of dpinner
	if args[0] != "!moneybags" {
		fmt.Println("invalid invocation")
		return
	}
	handleCommand(s, m, args)
}

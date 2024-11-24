package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
	"github.com/joho/godotenv"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Commande Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erreur lors du chargement du fichier .env : %v", err)
	}

	botToken := os.Getenv("SLACK_BOT_TOKEN")
	appToken := os.Getenv("SLACK_APP_TOKEN")

	if botToken == "" || appToken == "" {
		log.Fatalf("Les tokens Slack ne sont pas définis dans les variables d'environnement")
	}

	bot := slacker.NewClient(botToken, appToken)

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my year of birth is <year>", &slacker.CommandDefinition{
		Description: "year of birth calculator",
		Examples:    []string{"my year of birth is 2020"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			event := botCtx.Event() 
			if event.BotID != "" { 
				return
			}
	
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				response.Reply("Invalid year format!")
				return
			}
			age := 2024 - yob
			r := fmt.Sprintf("Your age is %d", age)
			response.Reply(r)
		},
	})
	

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
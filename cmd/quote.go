package stoic

import (
	"fmt"
	"math/rand"
	"time"
)

type quote struct {
	content string
	author  string
}

var quotes = []quote{}

type QuoteCommand interface {
	Run()
}

type quoteCommand struct {
	quotes []quote
}

func NewQuoteCommand() QuoteCommand {
	quotes = append(quotes, quote{
		content: "ranom text",
		author:  "Me",
	})

	return &quoteCommand{
		quotes: quotes,
	}
}

func (c quoteCommand) Run() {
	rand.Seed(time.Now().Unix())

	fmt.Println(c.quotes[rand.Intn(len(c.quotes))])
}

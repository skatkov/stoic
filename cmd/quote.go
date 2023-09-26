package stoic

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

type quote struct {
	content string
	author  string
}

var quotes = []quote{}

// Style definitions

var (
	foreground = lipgloss.AdaptiveColor{Light: "#969B86", Dark: "#696969"}
	pageStyle  = lipgloss.NewStyle().Padding(1, 2, 1, 2)
	boxStyle   = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(1, 0).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)

	authoredBy = lipgloss.NewStyle().SetString("—").
			PaddingRight(1).
			Foreground(foreground).
			String()

	author = func(s string) string {
		return "\n" + authoredBy + lipgloss.NewStyle().
			Foreground(foreground).
			Render(s) + "\n"
	}
)

type QuoteCommand interface {
	Run()
}

type quoteCommand struct {
	quotes []quote
}

func NewQuoteCommand() QuoteCommand {
	quotes = append(quotes,
		quote{
			content: "Few care now about the marches and countermarches of the Roman commanders. What the centuries have clung to is a notebook of thoughts by a man whose real life was largely unknown who put down in the midnight dimness not the events of the day or the plans of the morrow, but something of far more permanent interest, the ideals and aspirations that a rare spirit lived by.",
			author:  "Brand Blanshard",
		},
		quote{
			content: "Five hundred years later, Leonardo’s notebooks are around to astonish and inspire us. Fifty years from now, our own notebooks, if we work up the initiative to start them, will be around to astonish and inspire our grandchildren, unlike our tweets and Facebook posts.",
			author:  "Isaacs Newton",
		},
		quote{
			content: "You know those people whose lives are transformed by meditation or yoga or something like that? For me, it’s writing in my diary and journals. It’s made all the difference in the world for my learning, reflecting, and peace of mind.",
			author:  "Derek Sivers",
		},
		quote{
			content: "When the light has been removed and my wife has fallen silent, aware of this habit that’s now mine, I examine my entire day and go back over what I’ve done and said, hiding nothing from myself, passing nothing by.",
			author:  "Seneca",
		},
		quote{
			content: "Every night, I try myself by Court Martial to see if I have done anything effective during the day. I don’t mean just pawing the ground, anyone can go through the motions, but something really effective.",
			author:  "Winston Churchill",
		},
		quote{
			content: "My journaling system is based around studying complexity. Reducing the complexity down to what is the most important question. Sleeping on it, and then waking up in the morning first thing and pre-input brainstorming on it. So I’m feeding my unconscious material to work on, releasing it completely, and then opening my mind and riffing on it.",
			author:  "Reid Hoffman",
		},
		quote{
			content: "Reflection is…a key factor in expert learning and refers to the extent to which individuals are able to appraise what they have learned and to integrate these experiences into future actions, thereby maximizing performance improvements.",
			author:  "Marije Elferink-Gemser",
		},
		quote{
			content: "When Beethoven was enjoying a beer, he might suddenly pull out his notebook and write something in it. ‘Something just occurred to me,’ he would say, sticking it back into his pocket. The ideas that he tossed off separately, with only a few lines and points and without barlines, are hieroglyphics that no one can decipher. Thus in these tiny notebooks he concealed a treasure of ideas.",
			author:  "Wilhelm Von Lenz",
		},
		quote{
			content: "The recognition that I needed to train and discipline my character. Not to be sidetracked by my interest in rhetoric. Not to write treatises on abstract questions, or deliver moralizing little sermons, or compose imaginary descriptions of The Simple Life or The Man Who Lives Only for Others. To steer clear of oratory, poetry and belles lettres. Not to dress up just to stroll around the house, or things like that. To write straightforward.",
			author:  "Marcus Aurelius",
		},
		quote{
			content: "Is not the poet bound to write his own biography? Is there any other work for him but a good journal? We do not wish to know how his imaginary hero, but how he, the actual hero, lived from day to day.",
			author:  "Henry David Thoreau",
		},
		quote{
			content: "Keeping a journal is the veriest pastime in the world, and the pleasantest…Only those rare natures that are made up of pluck, endurance, devotion to duty for duty’s sake, and invincible determination, may hope to venture upon so tremendous an enterprise as the keeping of a journal.",
			author:  "Mark Twain",
		},
		quote{
			content: "I believe I could never exhaust the supply of material lying within me. The deeper I plunge, the more I discover. There is no bottom to my heart and no limit to the acrobatic feats of my imagination.",
			author:  "Anaïs Nin",
		},
		quote{
			content: "People look for retreats for themselves, in the country, by the coast, or in the hills. There is nowhere that a person can find a more peaceful and trouble-free retreat than in his own mind. . . . So constantly give yourself this retreat, and renew yourself.",
			author:  "Marcus Aurelius",
		},
	)

	return &quoteCommand{
		quotes: quotes,
	}
}

func (c quoteCommand) Run() {
	rand.Seed(time.Now().Unix())
	quote := c.quotes[rand.Intn(len(c.quotes))]
	doc := strings.Builder{}

	ui := lipgloss.NewStyle().
		Width(80).
		PaddingBottom(1).
		PaddingTop(1).
		PaddingLeft(3).
		PaddingRight(3).
		Align(lipgloss.Center).
		Render(quote.content)

	doc.WriteString(boxStyle.Render(ui) + "\n")
	doc.WriteString(author(quote.author))

	fmt.Println(pageStyle.Render(doc.String()))
}

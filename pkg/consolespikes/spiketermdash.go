package consolespikes

// simplified termdash spike.

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/mum4k/termdash/widgets/text"
)

var quotations = []string{
	"Life is like a box of chocolates",
	"Here's looking at you, kid",
	"Go ahead, make my day 🔫",
	"May the Force ⚡ be with you",
	"I'm the king 👑 of the world!",
	"",
	"Keep your friends close, but your enemies closer",
	"Here's Johnny!",
	"Nobody puts Baby in a corner",
	"To infinity and beyond!",
	"",
	"They may take our lives, but they'll never take our freedom!",
	"When you realize you want to spend the rest of your life with somebody, you want the rest of your life to start as soon as possible.",
	"If you let my daughter go now, that'll be the end of it. I will not look for you, I will not pursue you. But if you don't, I will look for you, I will find you, and I will kill you.",
}

func writeQuotes(ctx context.Context, t *text.Text, delay time.Duration) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			i := rand.Intn(len(quotations))
			t.Write(fmt.Sprintf("%s\n", quotations[i]))
			time.Sleep(delay)
		}
	}
}

// TermDashSpike ..
func TermDashSpike() {

	left, right, wg, ctx := SplitLeftRight("LEFTY", "RIGHTY")

	left.Write("started\nPress 'q' to quit.\n")
	go writeQuotes(ctx, right, 1*time.Second)

	wg.Wait()
}

// notes, lessons learnt
// ---------------------
// afaik, you can't debug termdash apps from vscode, (may be able to via some type of connect to running process)
// so make sure you can debug and test each dashboard component separately (abstracting out) using termdash write interface.

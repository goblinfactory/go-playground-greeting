package consolespikes

// simplified termdash spike.

import (
	"context"
	"fmt"
	"math/rand"
	"time"
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

func writeQuotes(ctx context.Context, t Konsole, delay time.Duration) {
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

// TermDashSpike2Columns ..
func TermDashSpike2Columns() {

	left, right, wg, ctx, _, _ := SplitLeftRight("LEFTY", "RIGHTY")
	left.Write("started\nPress 'q' to quit.\n")
	go writeQuotes(ctx, right, 1*time.Second)

	wg.Wait()
}

// TermDashSpike3Columns ..
func TermDashSpike3Columns() {

	left, right, _, wg, ctx, _ := SplitColumns123("LEFTY", "MIDDLE", "RIGHTY")
	left.Write("started\nPress 'q' to quit.\n")
	go writeQuotes(ctx, right, 1*time.Second)

	wg.Wait()
}

// TermDashSpike4ColumnsRedGreenPrinting ..
func TermDashSpike4ColumnsRedGreenPrinting() {

	//c1, c2, c3, c4, wg, ctx, cancel, kbd
	c1, c2, c3, _, wg, _, _, _ := SplitColumns1234("c1", "c2", "c3", "c4")
	c1.Write("started\nPress 'q' to quit.\n")
	c2.Write("column2\n")

	c2.Red("this is red inside ", "the column2\n")
	c2.Green("(success)")
	c3.Write("column3")

	//go writeQuotes(ctx, c4, 1*time.Second)

	wg.Wait()
}

// TermDashSpikeTopBottomLayout ..
func TermDashSpikeTopBottomLayout() {

	top, bottom, wg, ctx := SplitTopBottom("top", "bottom", nil)
	top.Write("started\nPress 'q' to quit.\n")
	go writeQuotes(ctx, bottom, 1*time.Second)

	wg.Wait()
}

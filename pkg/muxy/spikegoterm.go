package muxy

import (
	"fmt"

	tm "github.com/buger/goterm"
)

// contrived test of multiplexing using go channels
// take messages from 3 channels
// send then from service A to Service B over a single channel
// then have service B take a multiplexer to split out the incoming messages
// back into 3 channels

// Q: where is this used in the real world?
// A: mmm, possibly talking via a morse code sender and receiver?
// or via some serial device possibly

// small trading game,...

// service a, reads keyboard
// press, q -> quit, b-> buy, s-> sell
// have routine with current price varying up or down.
// have a time limit, 60 seconds, and you have to buy and sell
// and make as much profit as you can in 60 seconds.
// so you have to time your trades quickly
// as the price bounces around from high to low
// kinda like flappy bird for trading.
// each press of the 'b' or 's' buys or sells 100 units at the current price.
// you start with $1000 and a share price of $10 a share.
// price can go up or down from $1 to $100
// pull in some real trade histories and simulate 5 years worth of prices over 60 seconds with
// RULES (type of game) get told to you at start

// e.g. any stock left at the end of 60s will be zero-d out, i.e. you must sell all your stock before the timer runs out. score is ...well profit! starting capital ($1000) minus remaining cash = score.

// simple algorithm to simulate market, so that the markets are not always rising, make start and end equivalent so that we can start anywhere and roll around to start
// take data and "adjust" by x across y units so that start price = end price
// --------------------------

// even if you know the graph in advance, you don't know where it starts from, and if you wait for the price to hit bottom, you may not be able to sell all your stock
// in time for it to rise. You also can't tell what "bottom" is, because you dont know how much the stock has been "multiplied" by, so you could be waiting for it to

// nice idea for a website with a live "price feed" service, mmmm!

// RunSpikeGoTerm is draft
func RunSpikeGoTerm() {
	tm.Clear()

	// Create Box with 30% width of current screen, and height of 20 lines
	box := tm.NewBox(30|tm.PCT, 20, 0)

	// Add some content to the box
	// Note that you can add ANY content, even tables
	fmt.Fprint(box, "Some box content. Does this content wrap if it's wider than the box?")

	// Move Box to approx center of the screen
	tm.Print(tm.MoveTo(box.String(), 40|tm.PCT, 40|tm.PCT))

	tm.Flush()

}

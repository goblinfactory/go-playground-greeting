package controlc

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

const quitCommand string = "*quit*"

// TestControlC tests how to will kick of main's defer as well as background goroutine cancel when user presses control+c
func TestControlC() {

	var servers = []string{"london", "paris", "tokyo"}

	log := make(chan string, 20)

	fmt.Println("starting")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	var serverWg sync.WaitGroup
	var consoleWg sync.WaitGroup

	//TODO: need to also catch debug sessions that are interrupted by IDE. This does not trigger os.Signal. (need to check what it does)
	// found this issue that suggests it's possible to catch in Goland, but not in vscode // https://github.com/golang/vscode-go/issues/120
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	defer func() {
		signal.Stop(c)
		cancel()
	}()

	go func() {
		select {
		case <-c:
			log <- fmt.Sprintf(" <-- Control+c pressed, stopping servers;")
			cancel()
		case <-ctx.Done():
		}
	}()

	consoleWg.Add(1)
	go echoLogToConsole(&consoleWg, log)

	for i := range servers {
		config := connectionConfig{servers[i]}
		serverWg.Add(1)
		go fakeLongRunningDatabaseProcess(ctx, &serverWg, config, log)
	}

	log <- fmt.Sprintf("all servers started. Press control+C to quit.")
	serverWg.Wait()
	log <- quitCommand
	consoleWg.Wait()

}

type connectionConfig struct {
	name string
}

func echoLogToConsole(wg *sync.WaitGroup, con chan string) {
	defer wg.Done()
	for line := range con {
		if line == quitCommand {
			return
		}
		fmt.Println(line)
	}
}

func fakeLongRunningDatabaseProcess(ctx context.Context, wg *sync.WaitGroup, c connectionConfig, log chan string) {

	log <- fmt.Sprintf("start db process for:%s", c.name)

	connection := c.fakeOpenConnection(log)

	// defer db cleanup, signal when done.
	defer func() {
		connection.fakeCloseDbConnection(log)
		wg.Done()
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			pause()
			connection.appendStockPriceToDatabase(log)
		}
	}
}

func (conn *connectionConfig) fakeOpenConnection(log chan string) connectionConfig {
	log <- fmt.Sprintf("Opening database connection for : %s", conn.name)
	return *conn
}

func (conn *connectionConfig) fakeCloseDbConnection(log chan string) {
	log <- fmt.Sprintf("Closing database connection for : %s", conn.name)
}

func (conn *connectionConfig) appendStockPriceToDatabase(log chan string) {
	log <- fmt.Sprintf("%6s : £%7.2f", conn.name, rand.Float64()*1000)
}

func pause() {
	ms := time.Duration(rand.Intn(2000)) * time.Millisecond
	time.Sleep(ms)
}

// references
// ----------
// web: https://medium.com/@matryer/make-ctrl-c-cancel-the-context-context-bd006a8ad6ff

package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"github.com/goblinfactory/greeting/pkg/arrs"
	"github.com/goblinfactory/greeting/pkg/backpressuredemo/controlproducer"
	"github.com/goblinfactory/greeting/pkg/bloggy"
	"github.com/goblinfactory/greeting/pkg/channels"
	"github.com/goblinfactory/greeting/pkg/concurrencypatterns"
	"github.com/goblinfactory/greeting/pkg/controlc"
	"github.com/goblinfactory/greeting/pkg/customcollection"
	"github.com/goblinfactory/greeting/pkg/dependencyinjection"
	"github.com/goblinfactory/greeting/pkg/errorhandling"
	"github.com/goblinfactory/greeting/pkg/learninggo"
	"github.com/goblinfactory/greeting/pkg/muxy"
	"github.com/goblinfactory/greeting/pkg/nethttp2"
	"github.com/goblinfactory/greeting/pkg/sandboxes/sandbox1"
	"github.com/goblinfactory/greeting/pkg/sandboxes/sandbox2"
	"github.com/goblinfactory/greeting/pkg/switchy"
	"github.com/goblinfactory/greeting/pkg/testvet"
	"github.com/goblinfactory/greeting/pkg/testwaitgroup"
	"github.com/goblinfactory/greeting/pkg/timing"
)

var spikes = []func(){
	controlproducer.DemoConcurrencyLimiter,
	nethttp2.SpikeMinimalHTTPServer,
	// consolespikes.WhatHappensIfYouDontClose2,
	// consolespikes.SpikeUsingkeyboardHandlers,

	sandbox2.DemoGatherAndProcess,
	sandbox1.DemoRunOnce,
	channels.DemoActuallySettingChannelToNilTurnsOFFTheChannelWithZeroCPUWaste,

	// consolespikes.TermDashSpike4ColumnsRedGreenPrinting,
	// consolespikes.TermDashSpike4Columns,
	concurrencypatterns.DemoUsingCancelFuncToStopBackgroundGenerators,
	errorhandling.DemoUsingErrorsAsToCheckIfAnErrorContainsAnyErrorOfSpecificType,
	errorhandling.Demo1,
	errorhandling.Demo2,
	dependencyinjection.Main,
	customcollection.Demo,
	learninggo.TestEmbeddingAndComposition,
	// book1section1.FindDuplicateLines,
	timing.CompareConcatVsJoin,
	// commandline.Echo,
	switchy.TestSwitchy,
	// structs.TestMaps,
	controlc.TestControlC,
	arrs.TestSlicesArePassedByCopy,
	arrs.TestSorting,
	muxy.RunSpikeGoTerm,
	testvet.TestThatVetRunsOnSave,
	testwaitgroup.TestWaitGroup,
	bloggy.TestQuotes,
	learninggo.LissajousFromArgs,
}

func main() {

	args := os.Args
	var i int
	if len(args) == 1 {
		help()
		return
	} else {
		num, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal(err)
		}
		if num < 0 || num > len(spikes) {
			log.Fatalf("Test number must be between 0 and %d", len(spikes))
		}
		i = num
	}

	spikes[i]()

}

var reset = string("\033[0m")
var yellow = string("\033[33m")
var green = string("\033[32m")

// todo : convert to a nice clean window demo with list selector for demos, and run all demos in the "rhs" window.
// for now, good old fashioned text output.
func help() {
	defer fmt.Print(reset)
	fmt.Println("Alan's Go spikes")
	fmt.Println("Usage ./greeting {n}  //where n is one of the tests below")
	fmt.Println("----------------")
	for i := range spikes {
		file, line := getFunctionName(spikes[i])
		fmt.Println("[", green, i, "]", reset, file, yellow, "line", line, reset)
	}
}

func getFunctionName(i interface{}) (string, int) {
	p := reflect.ValueOf(i).Pointer()
	f := runtime.FuncForPC(p)
	file, line := f.FileLine(p)
	file = strings.Split(file, "pkg")[1]
	return file, line
}

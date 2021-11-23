package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"github.com/goblinfactory/greeting/pkg/ansi"
	"github.com/goblinfactory/greeting/pkg/arrs"
	"github.com/goblinfactory/greeting/pkg/backpressure"
	"github.com/goblinfactory/greeting/pkg/bloggy"
	"github.com/goblinfactory/greeting/pkg/channels"
	"github.com/goblinfactory/greeting/pkg/concurrencypatterns"
	"github.com/goblinfactory/greeting/pkg/consolespikes"
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

	// visually rich demos
	// -------------------
	consolespikes.SpikeUsingkeyboardHandlers,
	backpressure.DemoConcurrencyLimiter,
	nethttp2.SpikeMinimalHTTPServer,
	channels.DemoNotSettingChannelToNilCausesALotOfWastedCycles,
	channels.DemoActuallySettingChannelToNilTurnsOFFTheChannelWithZeroCPUWaste,
	bloggy.DemoCallingAPIsWithCircuitBreaker,

	// simple demo
	// ----------
	// consolespikes.WhatHappensIfYouDontClose2,

	sandbox2.DemoGatherAndProcess,
	sandbox1.DemoRunOnce,

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
	learninggo.LissajousFromArgs,
}

func main() {

	args := os.Args
	var i int
	if len(args) == 1 {
		help()
		return
	}

	num, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}
	if num < 0 || num > len(spikes) {
		log.Fatalf("Test number must be between 0 and %d", len(spikes))
	}
	i = num

	spikes[i]()

}

func help() {
	fmt.Print(ansi.Cls)
	defer fmt.Print(ansi.Reset)
	fmt.Println("Alan's Go spikes")
	fmt.Println("Usage ./greeting {n}  //where n is one of the tests below")
	for i := range spikes {
		file, name, _ := getFunctionName(spikes[i])
		file = fmt.Sprintf("%-46s", file)
		num := fmt.Sprintf("%02d", i)
		fmt.Println(ansi.Green, num, ansi.Reset, file, ansi.Green, name, ansi.Reset)
	}
}

func getFunctionName(i interface{}) (string, string, int) {
	p := reflect.ValueOf(i).Pointer()
	f := runtime.FuncForPC(p)
	file, line := f.FileLine(p)
	file = strings.Split(file, "pkg")[1]
	ln := f.Name()
	return file, shortName(ln), line
}

func shortName(longname string) string {
	parts := strings.Split(longname, string(os.PathSeparator))
	sn := parts[len(parts)-1]
	return strings.Split(sn, ".")[1]
}

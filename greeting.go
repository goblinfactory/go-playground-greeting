package main

import (
	"fmt"

	"github.com/goblinfactory/greeting/pkg/consolespikes"
)

func main() {

	h := consolespikes.NewKBHandler()
	fmt.Println("starting")
	left, right, wg, _ := consolespikes.SplitLeftRight("left", "right", h)
	h.Handlers['a'] = func() { right.Write("left\n") }
	h.Handlers['s'] = func() { right.Write("down\n") }
	h.Handlers['d'] = func() { right.Write("right\n") }
	h.Handlers['w'] = func() { right.Write("up\n") }

	left.Write("hello world!")
	wg.Wait()

	//controlproducer.DemoConcurrencyLimiter()

	//sandbox2.DemoGatherAndProcess()
	//sandbox1.DemoRunOnce()
	//channels.DemoActuallySettingChannelToNilTurnsOFFTheChannelWithZeroCPUWaste()

	//consolespikes.TermDashSpike4ColumnsRedGreenPrinting()
	//consolespikes.TermDashSpike4Columns()
	//concurrencypatterns.DemoUsingCancelFuncToStopBackgroundGenerators()
	//errorhandling.DemoUsingErrorsAsToCheckIfAnErrorContainsAnyErrorOfSpecificType()
	// errorhandling.Demo1()
	// errorhandling.Demo2()
	//dependencyinjection.Main()
	//customcollection.Demo()
	//learninggo.TestEmbeddingAndComposition()
	//book1section1.LissajousFromArgs(os.Args[1:])
	//book1section1.FindDuplicateLines()
	//timing.CompareConcatVsJoin()
	//commandline.Echo()
	//switchy.TestSwitchy()
	//structs.TestMaps()
	//controlc.TestControlC()
	//arrs.TestSlicesArePassedByCopy()
	//arrs.TestSorting()
	//muxy.RunSpikeGoTerm()
	//testvet.TestThatVetRunsOnSave()
	//testwaitgroup.TestWaitGroup()
	//bloggy.TestQuotes()
}

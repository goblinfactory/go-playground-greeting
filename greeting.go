package main

import "github.com/goblinfactory/greeting/pkg/backpressuredemo/controlproducer"

func main() {

	controlproducer.DemoConcurrencyLimiter()

	// file, err := os.OpenFile("logs.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.SetOutput(file)

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

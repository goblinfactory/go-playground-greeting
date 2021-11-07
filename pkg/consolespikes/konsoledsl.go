package consolespikes

import (
	"context"
	"log"
	"sync"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/text"
)

// experiment with creating a simple Goblinfactory.Konsole like dsl over termdash.

// SplitLeftRight splits console window and returns left and right windows a context and a cancel. Will run until you press q or you call cancel()
func SplitLeftRight(leftTitle string, rightTitle string) (*text.Text, *text.Text, *sync.WaitGroup, context.Context) {

	left, _ := text.New(text.RollContent(), text.WrapAtWords())
	right, _ := text.New(text.RollContent(), text.WrapAtWords())

	layout := container.SplitVertical(
		container.Left(
			container.Border(linestyle.Light),
			container.BorderTitleAlignCenter(),
			container.BorderTitle(leftTitle),
			container.PlaceWidget(left),
		),
		container.Right(
			container.Border(linestyle.Light),
			container.BorderTitle(rightTitle),
			container.BorderTitleAlignCenter(),
			container.PlaceWidget(right),
		),
	)
	wg, ctx := runWindowLayout(layout)
	return left, right, wg, ctx
}

func runWindowLayout(layout container.Option) (*sync.WaitGroup, context.Context) {
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())

	t, err1 := tcell.New()

	if err1 != nil {
		log.Fatal(err1)
	}

	c, err2 := container.New(t, layout)

	if err2 != nil {
		log.Fatal(err2)
	}

	wg.Add(1)
	go runTermdashUntilUserPressesQuitKey(ctx, cancel, &wg, c, t)
	return &wg, ctx
}

func runTermdashUntilUserPressesQuitKey(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup, c *container.Container, t *tcell.Terminal) {

	defer func() {
		t.Close()
		wg.Done()
	}()

	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' {
			cancel()
		}
	}

	err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quitter))
	if err != nil {
		log.Fatal(err)
	}
}

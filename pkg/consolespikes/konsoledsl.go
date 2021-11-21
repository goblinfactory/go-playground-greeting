package consolespikes

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/text"
)

// experiment with creating a simple Goblinfactory.Konsole like dsl over termdash.

// Konsole provides utility methods for printing to windows
type Konsole struct {
	con *text.Text
}

// NewKonsole returns a new Konsole for managing printing to windows
func NewKonsole(con *text.Text) Konsole {
	return Konsole{con: con}
}

// WriteLine is useful for printing a number of objects including color.
func (c *Konsole) WriteLine(texts ...interface{}) {
	cnt := len(texts)
	for i, t := range texts {
		switch t := t.(type) {
		case int:
			c.con.Write(fmt.Sprintf("%d", t))
		case string:
			c.con.Write(t)
		default:
			c.con.Write(fmt.Sprintf("%v", t))
		}
		if (i + 1) >= cnt {
			c.con.Write("\n")
		}
	}
}

// Write is useful for printing a number of objects including color.
func (c *Konsole) Write(texts ...interface{}) {
	for _, t := range texts {
		switch t := t.(type) {
		case int:
			c.con.Write(fmt.Sprintf("%d", t))
		case string:
			c.con.Write(t)
		default:
			c.con.Write(fmt.Sprintf("%v", t))
		}
	}
}

// Red writes texts in red
func (c *Konsole) Red(texts ...interface{}) {
	c.writeColor(cell.ColorRed, texts...)
}

// Gray writes texts in red
func (c *Konsole) Gray(texts ...interface{}) {
	c.writeColor(cell.ColorGray, texts...)
}

func (c *Konsole) writeColor(color cell.Color, texts ...interface{}) {
	for _, t := range texts {
		v := ""
		switch t := t.(type) {
		case int:
			v = fmt.Sprintf("%d", t)
		case string:
			v = t
		default:
			v = fmt.Sprintf("%v", t)
		}
		c.con.Write(v, text.WriteCellOpts(cell.FgColor(color)))
	}
}

// Green writes texts in green
func (c *Konsole) Green(texts ...interface{}) {
	c.writeColor(cell.ColorGreen, texts...)
}

// SplitLeftRight splits console window and returns left and right windows, a waitgroup and a context. Will run until you press q or you call cancel()
func SplitLeftRight(leftTitle string, rightTitle string) (*text.Text, *text.Text, *sync.WaitGroup, context.Context, *KeyboardHandlers) {
	kb := NewKBHandler()
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
	wg, ctx := runWindowLayout(layout, kb)
	return left, right, wg, ctx, kb
}

// NewWindow ...
func NewWindow(title string, ls linestyle.LineStyle) (*text.Text, *sync.WaitGroup, context.Context) {

	t, err := tcell.New()
	if err != nil {
		panic(err)
	}
	defer t.Close()

	ctx, cancel := context.WithCancel(context.Background())

	body, err := text.New(text.RollContent(), text.WrapAtWords())
	if err != nil {
		panic(err)
	}

	c, err := container.New(
		t,
		container.Border(ls),
		container.BorderTitle(title),
		container.PlaceWidget(body),
	)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go runTermdashUntilUserPressesQuitKey(ctx, cancel, &wg, c, t, nil)
	return body, &wg, ctx

}

// SplitTopBottom splits console window and returns top and bottom windows, a waitgroup and context. Will run until you press q or you call cancel()
func SplitTopBottom(topTitle string, bottomTitle string, kb *KeyboardHandlers) (*text.Text, *text.Text, *sync.WaitGroup, context.Context) {

	top, _ := text.New(text.RollContent(), text.WrapAtWords())
	bottom, _ := text.New(text.RollContent(), text.WrapAtWords())

	layout := container.SplitHorizontal(
		container.Top(
			container.Border(linestyle.Light),
			container.BorderTitleAlignCenter(),
			container.BorderTitle(topTitle),
			container.PlaceWidget(top),
		),
		container.Bottom(
			container.Border(linestyle.Light),
			container.BorderTitle(bottomTitle),
			container.BorderTitleAlignCenter(),
			container.PlaceWidget(bottom),
		),
	)
	wg, ctx := runWindowLayout(layout, kb)
	return top, bottom, wg, ctx
}

// SplitColumns123 splits console window into 3 columnsw and returns left and right windows, a waitgroup and a context. Will run until you press q or you call cancel()
func SplitColumns123(col1title string, col2title string, col3title string) (*text.Text, *text.Text, *text.Text, *sync.WaitGroup, context.Context, *KeyboardHandlers) {
	kb := NewKBHandler()
	col1, _ := text.New(text.RollContent(), text.WrapAtWords())
	col2, _ := text.New(text.RollContent(), text.WrapAtWords())
	col3, _ := text.New(text.RollContent(), text.WrapAtWords())

	layout := container.SplitVertical(
		container.Left(
			container.Border(linestyle.Light),
			container.BorderTitleAlignCenter(),
			container.BorderTitle(col1title),
			container.PlaceWidget(col1),
		),
		container.Right(
			container.SplitVertical(
				container.Left(
					container.Border(linestyle.Light),
					container.BorderTitleAlignCenter(),
					container.BorderTitle(col2title),
					container.PlaceWidget(col2),
				),
				container.Right(
					container.Border(linestyle.Light),
					container.BorderTitleAlignCenter(),
					container.BorderTitle(col3title),
					container.PlaceWidget(col3),
				),
			),
		), container.SplitPercent(33),
	)

	wg, ctx := runWindowLayout(layout, kb)
	return col1, col2, col3, wg, ctx, kb
}

// SplitColumns1234 splits console window into 3 columnsw and returns left and right windows, a waitgroup and a context. Will run until you press q or you call cancel()
func SplitColumns1234(col1title string, col2title string, col3title string, col4title string, kb *KeyboardHandlers) (*text.Text, *text.Text, *text.Text, *text.Text, *sync.WaitGroup, context.Context) {

	col1, _ := text.New(text.RollContent(), text.WrapAtWords())
	col2, _ := text.New(text.RollContent(), text.WrapAtWords())
	col3, _ := text.New(text.RollContent(), text.WrapAtWords())
	col4, _ := text.New(text.RollContent(), text.WrapAtWords())

	layout := container.SplitVertical(
		container.Left(
			container.SplitVertical(
				container.Left(
					container.Border(linestyle.Light),
					container.BorderTitleAlignCenter(),
					container.BorderTitle(col1title),
					container.PlaceWidget(col1),
				),
				container.Right(
					container.Border(linestyle.Light),
					container.BorderTitleAlignCenter(),
					container.BorderTitle(col2title),
					container.PlaceWidget(col2),
				),
			),
		),
		container.Right(
			container.SplitVertical(
				container.Left(
					container.Border(linestyle.Light),
					container.BorderTitleAlignCenter(),
					container.BorderTitle(col3title),
					container.PlaceWidget(col3),
				),
				container.Right(
					container.Border(linestyle.Light),
					container.BorderTitleAlignCenter(),
					container.BorderTitle(col4title),
					container.PlaceWidget(col4),
				),
			),
		),
	)

	wg, ctx := runWindowLayout(layout, kb)
	return col1, col2, col3, col4, wg, ctx
}

func runWindowLayout(layout container.Option, kb *KeyboardHandlers) (*sync.WaitGroup, context.Context) {
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
	go runTermdashUntilUserPressesQuitKey(ctx, cancel, &wg, c, t, kb)
	return &wg, ctx
}

// KeyboardHandlers keystroke handers that will recieve keyboard keypresses
type KeyboardHandlers struct {
	Handlers map[keyboard.Key]func()
}

// NewKBHandler ...
func NewKBHandler() *KeyboardHandlers {
	return &KeyboardHandlers{map[keyboard.Key]func(){}}
}

func runTermdashUntilUserPressesQuitKey(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup, c *container.Container, t *tcell.Terminal, kb *KeyboardHandlers) {

	defer func() {
		t.Close()
		wg.Done()
	}()

	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' {
			cancel()
			return
		}
		if kb != nil && kb.Handlers != nil {
			h, ok := kb.Handlers[k.Key]
			if !ok {
				return
			}
			h()
		}
	}

	err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quitter))
	if err != nil {
		log.Fatal(err)
	}
}

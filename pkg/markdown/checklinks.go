package markdown

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/goblinfactory/greeting/pkg/ansi"
	"github.com/goblinfactory/greeting/pkg/regexs"
	"github.com/pterm/pterm"
)

// CheckMardownLinks ...
func CheckMardownLinks() {
	CheckLinks("readme.md", false)
}

// CheckMardownLinksVerbose ...
func CheckMardownLinksVerbose() {
	CheckLinks("readme.md", true)
}

// CheckLinks checks that all the links in a markdown file are correct
func CheckLinks(path string, printNonErrors bool) {

	bytes, err := ioutil.ReadFile(path)
	errcnt := 0
	check(err)
	links := FindLinks(bytes)
	results := make(LinkChecks, 0)
	for _, link := range links {
		ok, _ := checkGoSourceFileExists(link.RelPath)
		if !ok {
			errcnt++
		}
		link := LinkCheck{link, ok}
		if ok && printNonErrors {
			results = append(results, link)
		}
		if !ok {
			results = append(results, link)
		}
	}
	cnt := len(results)
	if errcnt != 0 {
		defer log.Fatalf("%d broken links in %s", cnt, path)
		fmt.Println(path, ansi.Red, fmt.Sprintf("(%d) errors.", cnt), ansi.Reset)
	}

	if errcnt == 0 && !printNonErrors {
		fmt.Println(path+" :All links", ansi.Green, " (ok)", ansi.Reset)
		return
	}
	print(results)
}

func print(results []LinkCheck) {
	td := make([][]string, 0)
	td = append(td, []string{"link", "result"})
	td = append(td, []string{"----", "------"})
	for _, row := range results {
		td = append(td, printRow(row))
	}
	pterm.DefaultTable.WithHasHeader().WithData(td).Render()

}

func printRow(link LinkCheck) []string {
	if link.Exists {
		return []string{link.RelPath, fmt.Sprint(ansi.Green + "✓" + ansi.Reset)}
	}
	return []string{link.RelPath, fmt.Sprintf(ansi.Red + "missing" + ansi.Reset)}
}

func errtext(err error) string {
	if err == nil {
		return fmt.Sprint(ansi.Green, "-", ansi.Reset)
	}
	return fmt.Sprintf(ansi.Red, err.Error(), ansi.Reset)
}

// LinkChecks ...
type LinkChecks []LinkCheck

// LinkCheck ...
type LinkCheck struct {
	Link
	Exists bool
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Link ...
type Link struct {
	Text    string
	RelPath string
}

func parseLinks(pairs regexs.Pairs) []Link {
	m := make([]Link, len(pairs))
	for i, p := range pairs {
		m[i] = Link{p.Match1, p.Match2}
	}
	return m
}

// FindLinks finds all the internal hyperlinks in a mardown file using markdown [text](hyperlink) format. Ignores any external links
func FindLinks(content []byte) []Link {
	pm := regexs.NewPairMatcher(regexs.PatternMarkdownURI, nil, []string{"://"})
	pairs := pm.SearchForPairs(content)
	links := parseLinks(pairs)
	return links
}

func checkGoSourceFileExists(relpath string) (bool, error) {
	_, err := os.Stat(relpath)
	if err == nil {
		return true, nil
	}
	if err != nil && errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	panic(err)
}

// Random idea
// -----------
// use git history to check for git files with old names, and see if we can work out if they have been renamed, find files with same content?
// would need to be a new git tool, "movedWhere ?" that can find out where git file was moved to.
// then can have option to show broken links and have parameter to allow for automatic repair broken links, and preview repairs.

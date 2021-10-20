package customcollection

import "log"

// a list of whatever you give it

type intOrString struct {
	items []interface{}
}

func newIntOrStringList(cnt int) intOrString {
	items := make([]interface{}, cnt)
	return intOrString{items}
}

func (list *intOrString) Add(item interface{}) {
	switch item.(type) {
	case int:
		break
	case string:
		break
	default:
		log.Fatal("PANIC! Only int or string supported.")
	}
	list.items = append(list.items, item)
}

func (list *intOrString) AddInt(item int) {
	list.items = append(list.items, item)
}

func (list *intOrString) AddString(item int) {
	list.items = append(list.items, item)
}

type box struct {
	size int
}

// Demo ...
func Demo() {
	list := newIntOrStringList(10)
	list.Add("number")
	list.Add(box{12})
	list.Add(12)
}

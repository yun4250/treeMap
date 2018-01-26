package treeMap

const (
	RED   = 1
	BLACK = 0
)

type Entry struct {
	K      int64
	V      interface{}
	left   *Entry
	right  *Entry
	parent *Entry
	color  int
}

func (e *Entry) setValue(v interface{}) interface{} {
	return v
}

func rightOf(e *Entry) *Entry {
	if e == nil {
		return nil
	} else {
		return e.right
	}
}

func parentOf(e *Entry) *Entry {
	if e == nil {
		return nil
	} else {
		return e.parent
	}
}

func leftOf(e *Entry) *Entry {
	if e == nil {
		return nil
	} else {
		return e.left
	}
}

func colorOf(e *Entry) int {
	if e == nil {
		return BLACK
	} else {
		return e.color
	}
}

func setColor(p *Entry, c int) {
	if p != nil {
		p.color = c
	}
}

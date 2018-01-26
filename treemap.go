package treeMap

type TreeMap struct {
	root     *Entry
	size     int
	ModCount int
}

func Compare(k1 int64, k2 int64) int {
	if k1 < k2{
		return -1
	}else if k1 == k2{
		return 0
	}else{
		return 1
	}
}

func (tm *TreeMap) Size() int {
	return tm.size
}

func (tm *TreeMap) Put(key int64, value interface{}) interface{} {
	t := tm.root
	if t == nil {
		tm.root = &Entry{
			K: key,
			V: value,
		}
		tm.size = 1
		tm.ModCount ++
		return nil
	}
	var cmp int
	var parent *Entry
	for t != nil {
		parent = t
		cmp = Compare(key, t.K)
		if cmp < 0 {
			t = t.left
		} else if cmp > 0 {
			t = t.right
		} else {
			return t.setValue(value)
		}
	}
	e := &Entry{
		K:      key,
		V:      value,
		parent: parent,
	}
	if cmp < 0 {
		parent.left = e
	} else {
		parent.right = e
	}
	tm.fixAfterInsertion(e)
	tm.size++
	tm.ModCount++
	return nil
}

func (tm *TreeMap) fixAfterInsertion(x *Entry) {
	x.color = RED
	for x != nil && x != tm.root && x.parent.color == RED {
		if parentOf(x) == leftOf(parentOf(parentOf(x))) {
			y := rightOf(parentOf(parentOf(x)))
			if colorOf(y) == RED {
				setColor(parentOf(x), BLACK)
				setColor(y, BLACK)
				setColor(parentOf(parentOf(x)), RED)
				x = parentOf(parentOf(x))
			} else {
				if x == rightOf(parentOf(x)) {
					x = parentOf(x)
					tm.rotateLeft(x)
				}
				setColor(parentOf(x), BLACK)
				setColor(parentOf(parentOf(x)), RED)
				tm.rotateRight(parentOf(parentOf(x)))
			}
		} else {
			y := leftOf(parentOf(parentOf(x)))
			if colorOf(y) == RED {
				setColor(parentOf(x), BLACK)
				setColor(y, BLACK)
				setColor(parentOf(parentOf(x)), RED)
				x = parentOf(parentOf(x))
			} else {
				if x == leftOf(parentOf(x)) {
					x = parentOf(x)
					tm.rotateRight(x)
				}
				setColor(parentOf(x), BLACK)
				setColor(parentOf(parentOf(x)), RED)
				tm.rotateLeft(parentOf(parentOf(x)))
			}
		}
	}
}

func (tm *TreeMap) rotateLeft(p *Entry) {
	if p != nil {
		r := p.right
		p.right = r.left
		if r.left != nil {
			r.left.parent = p
		}
		r.parent = p.parent
		if p.parent == nil {
			tm.root = r
		} else if p.parent.left == p {
			p.parent.left = r
		} else {
			p.parent.right = r
		}
		r.left = p
		p.parent = r
	}
}

func (tm *TreeMap) rotateRight(p *Entry) {
	if p != nil {
		l := p.left
		p.left = l.right
		if l.right != nil {
			l.right.parent = p
		}
		l.parent = p.parent
		if p.parent == nil {
			tm.root = l
		} else if p.parent.right == p {
			p.parent.right = l
		} else {
			p.parent.left = l
		}
		l.right = p
		p.parent = l
	}
}

func (tm *TreeMap) FirstEntry() *Entry {
	p := tm.root
	if p != nil {
		for p.left != nil {
			p = p.left
		}
	}
	return p
}

func (t *TreeMap) FindCeiling(key int64) *Entry {
	x := t.root
	for x != nil {
		cmp := Compare(key, x.K)
		if cmp < 0 {
			if x.left != nil {
				x = x.left
				//fmt.Printf("lf %d\n", x.K)
			} else {
				//fmt.Printf("lr %d\n", x.K)
				return x
			}
		} else if cmp > 0 {
			if x.right != nil {
				x = x.right
				//fmt.Printf("pf %d\n", x.K)
			} else {
				p := x.parent
				ch := x
				for p != nil && ch == p.right {
					//fmt.Printf("pp %d\n", p.K)
					ch = p
					p = p.parent
				}
				//fmt.Printf("rr %d\n", p.K)
				return p
			}
		} else {
			return x
		}
	}
	return nil
}

package linkedlists

type Ring struct {
	next, prev *Ring
	Value      interface{}
}

func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

func NewRing(n int) *Ring {
	if n < 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 0; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	r.prev = p
	p.next = r
	return r
}

func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

func (r *Ring) Prev() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.prev
}

func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	}
	return r
}

func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
		/*r.next = s
		s.prev = r
		n.prev = s
		s.next = n*/
	}
	return n
}

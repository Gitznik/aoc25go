package main

type Dial int

func NewDialValue(v int) Dial {
	nv := v % 100
	if nv < 0 {
		nv = -nv
	}
	return Dial(nv)
}

func (d Dial) Add(v int) (Dial, int) {
	s := int(d) + v
	clicks := s / 100
	if clicks < 0 {
		clicks = -clicks
	}
	nv := s % 100
	if s == 0 {
		clicks++
	}
	if s < 0 {
		if int(d) != 0 {
			clicks++
		}
		return NewDialValue(nv + 100), clicks
	}
	return NewDialValue(nv), clicks
}

func (d Dial) Sub(v int) (Dial, int) {
	return d.Add(-v)
}

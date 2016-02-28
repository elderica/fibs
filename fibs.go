package fibs

type Fibs struct {
	a, b int
}

func New(a, b int) *Fibs {
	return &Fibs{a, b}
}

func (f *Fibs) Take(n uint) (ibs []int) {
	ibs = make([]int, n)
	for i := uint(0); i < n; i++ {
		ibs[i] = f.a
		f.step()
	}
	return
}

func (f *Fibs) step() {
	f.a, f.b = f.b, f.a + f.b
}

func (f *Fibs) Each(g func(i int)) {
	for {
		g(f.a)
		f.step()
	}
}

package money

type RUB float64

func FromKopecks(kopecks int64) RUB {
	return RUB(kopecks) / 100
}

func (r RUB) Kopecks() int64 {
	return int64(r * 100)
}

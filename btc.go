package money

type BTC float64

func FromSatoshi(satoshi int64) BTC {
	return BTC(satoshi) / 100000000
}

func (b BTC) Satoshi() int64 {
	return int64(b * 100000000)
}

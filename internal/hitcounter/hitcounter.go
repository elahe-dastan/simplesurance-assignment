package hitcounter

type HitCounter interface {
	Hit(int64)
	Count() int64
}

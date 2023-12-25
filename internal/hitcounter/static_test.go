package hitcounter_test

import (
	"testing"
	"time"

	"github.com/elahe-dastan/simplesurance-assignment/internal/hitcounter"
)

var _ hitcounter.HitCounter = new(hitcounter.StaticHitCounter)

func TestStaticHitCounter(t *testing.T) {
	hc := hitcounter.NewStatic()

	now := time.Now().Unix()

	hc.Hit(now)
	hc.Hit(now + 10)
	hc.Hit(now + 30)

	if hc.Count() != 3 {
		t.Fatalf("number of hits is %d which is not equal to %d", hc.Count(), 3)
	}
}

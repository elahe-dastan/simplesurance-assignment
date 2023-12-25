package hitcounter_test

import (
	"testing"

	"github.com/elahe-dastan/simplesurance-assignment/internal/hitcounter"
)

func TestStaticHitCounter(t *testing.T) {
	hc := hitcounter.NewStatic()

	hc.Hit(0)
	hc.Hit(10)
	hc.Hit(30)

	if hc.Count() != 3 {
		t.Fatalf("number of hits is %d which is not equal to %d", hc.Count(), 3)
	}
}

package hitcounter_test

import (
	"testing"
	"time"

	"github.com/elahe-dastan/simplesurance-assignment/internal/hitcounter"
)

var _ hitcounter.HitCounter = new(hitcounter.StaticHitCounter)

func TestOne(t *testing.T) {
	hc := hitcounter.NewStatic()

	now := time.Now().Unix()

	hc.Hit(now)
	hc.Hit(now + 10)
	hc.Hit(now + 30)

	if hc.Count() != 3 {
		t.Fatalf("number of hits is %d which is not equal to %d", hc.Count(), 3)
	}
}

func TestTwo(t *testing.T) {
	hc := hitcounter.NewStatic()

	now := time.Now().Unix()

	hc.Hit(now)
	hc.Hit(now + 61)

	if hc.Count() != 1 {
		t.Fatalf("number of hits is %d which is not equal to %d", hc.Count(), 1)
	}
}

func TestThree(t *testing.T) {
	hc := hitcounter.NewStatic()

	now := time.Now().Unix()

	hc.Hit(now)
	hc.Hit(now + 1)
	hc.Hit(now + 2)
	hc.Hit(now + 3)
	hc.Hit(now + 4)
	hc.Hit(now + 65)

	if hc.Count() != 1 {
		t.Fatalf("number of hits is %d which is not equal to %d", hc.Count(), 1)
	}
}

func TestFour(t *testing.T) {
	hc := hitcounter.NewStatic()

	now := time.Now().Unix()

	hc.Hit(now)
	hc.Hit(now + 1)
	hc.Hit(now + 300)

	if hc.Count() != 1 {
		t.Fatalf("number of hits is %d which is not equal to %d", hc.Count(), 1)
	}
}

func TestFive(t *testing.T) {
	hc := hitcounter.NewStatic()

	now := time.Now().Unix()

	hc.Hit(now)
	hc.Hit(now + 60)

	if hc.Count() != 1 {
		t.Fatalf("number of hits is %d which is not equal to %d", hc.Count(), 1)
	}
}

func TestSix(t *testing.T) {
	hc := hitcounter.NewStatic()

	now := time.Now().Unix()

	hc.Hit(now)
	hc.Hit(now + 60)
	hc.Hit(now + 61)

	if hc.Count() != 2 {
		t.Fatalf("number of hits is %d which is not equal to %d", hc.Count(), 2)
	}
}

func TestSeven(t *testing.T) {
	hc := hitcounter.NewStatic()
	// by setting it to zero, we will bypass being less than window size.
	hc.StartupTime = 0

	hc.Hit(1703613474)
	hc.Hit(1703613481)

	if hc.Count() != 2 {
		t.Fatalf("number of hits is %d which is not equal to %d", hc.Count(), 2)
	}
}

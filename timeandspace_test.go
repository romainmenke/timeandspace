package timeandspace_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	fuzz "github.com/google/gofuzz"
	"github.com/romainmenke/timeandspace"
)

func TestTimeAndSpace(t *testing.T) {
	ti := time.Now()
	lat, long := 50.999497, 4.3945308
	ts := timeandspace.NewTimeAndSpace("some-id", ti, lat, long)
	t.Log(ts)
}

func newSpace(z uint8, lat float64, long float64) (int, int) {
	x := int(math.Floor((long + 180.0) / 360.0 * (math.Exp2(float64(z)))))
	y := int(math.Floor((1.0 - math.Log(math.Tan(lat*math.Pi/180.0)+1.0/math.Cos(lat*math.Pi/180.0))/math.Pi) / 2.0 * (math.Exp2(float64(z)))))
	return x, y
}

// Make sure that the maths for newSpace never return a negative number
func TestFuzzSpaceForNegative(t *testing.T) {
	type Input struct {
		Lat  float64
		Long float64
	}

	f := fuzz.New()
	input := Input{}

	runs := 10000

	for z := 0; z < 20; z++ {
		t.Run(fmt.Sprint("Z", z), func(t *testing.T) {
			{
				x, y := newSpace(uint8(z), -360, -90)

				if x < 0 {
					t.Fatalf("negative x : %v at : %d", -360, z)
				}
				if y < 0 {
					t.Fatalf("negative y : %v at : %d", -90, z)
				}
			}

			for i := 0; i < runs; i++ {
				f.Fuzz(&input)

				x, y := newSpace(uint8(z), input.Lat, input.Long)

				if x < 0 {
					t.Fatalf("negative x : %v for : %v at : %d", x, input, z)
				}
				if y < 0 {
					t.Fatalf("negative y : %v for : %v at : %d", y, input, z)
				}
			}
		})
	}

}

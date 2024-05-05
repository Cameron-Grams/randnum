// Returns a random integer
package randnum

import (
	"math/rand"
	"time"
)

// Using the time function as a seed returns a random int
func RandomInteger(rng int) (int, error) {
	time := time.Now().Unix()
	rand.Seed(time)
	return rand.Intn(rng), nil
}

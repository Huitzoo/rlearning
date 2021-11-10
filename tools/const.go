package tools

import (
	"math/rand"
	"time"
)

var seed = time.Now().UTC().UnixNano()

func SetSeed() {
	rand.Seed(seed)
}

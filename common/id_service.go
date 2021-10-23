package common

import (
	"fmt"
	"math/rand"
)

func RandomId() string {
	return fmt.Sprintf("%+v", rand.Int())
}

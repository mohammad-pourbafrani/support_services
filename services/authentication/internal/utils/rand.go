package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func NextRandomInt32(min, max int32) (int32, error) {
	if min > max {
		return 0, fmt.Errorf("min should be less than or equal to max")
	}

	rangeSize := int64(max - min + 1)
	if rangeSize <= 0 {
		return 0, fmt.Errorf("range of possible values is zero or negative")
	}

	num, err := rand.Int(rand.Reader, big.NewInt(rangeSize))
	if err != nil {
		return 0, err
	}

	return min + int32(num.Int64()), nil
}

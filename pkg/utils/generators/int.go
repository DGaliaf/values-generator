package generators

import (
	"crypto/rand"
	"math"
	"math/big"
)

func GenerateInt(numberOfDigits int) (int, error) {
	maxLimit := int64(int(math.Pow10(numberOfDigits)) - 1)
	lowLimit := int(math.Pow10(numberOfDigits - 1))

	randomNumber, err := rand.Int(rand.Reader, big.NewInt(maxLimit))
	if err != nil {
		return 0, err
	}
	randomNumberInt := int(randomNumber.Int64())

	if randomNumberInt <= lowLimit {
		randomNumberInt += lowLimit
	}

	if randomNumberInt > int(maxLimit) {
		randomNumberInt = int(maxLimit)
	}
	return randomNumberInt, nil
}

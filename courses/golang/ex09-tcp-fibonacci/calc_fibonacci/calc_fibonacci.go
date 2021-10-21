package calc_fibonacci

import (
	"errors"
	"math/big"
)

func Calc(n int64) (*big.Int, error) {
	if n < 1 {
		return nil, errors.New("argument can not be less than 1")
	} else if n == 1 {
		return big.NewInt(0), nil
	}

	row := make([]*big.Int, 3)

	row[1] = big.NewInt(1)
	row[2] = big.NewInt(1)

	for i := int64(2); i < n; i++ {
		row[0] = row[1]
		row[1] = row[2]

		row[2] = big.NewInt(0).Add(row[0], row[1])
	}

	return row[2], nil
}

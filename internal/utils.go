package internal

import "strconv"

func strToUint(n string) (uint, error) {
	v, err := strconv.Atoi(n)
	if err != nil {
		return 0, err
	}
	return uint(v), nil
}

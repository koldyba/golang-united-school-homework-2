package square

import "math"

type sides int

func CalcSquare(sideLen float64, sidesNum sides) float64 {
	var sq float64 = 0

	switch {
	case sidesNum == 0:
		sq = math.Pi * math.Pow(sideLen, 2)
	case sidesNum == 3:
		sq = math.Pow(sideLen, 2) * math.Sqrt(3) / 4
	case sidesNum == 4:
		sq = math.Pow(sideLen, 2)
	}

	return sq
}

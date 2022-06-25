package external

type External interface {
	Calculate(x, y int32) int32
}

func FinalCalculation(ext External, x int32, y int32) int32 {
	res := ext.Calculate(x, y)
	return res
}

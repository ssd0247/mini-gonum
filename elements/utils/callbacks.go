package utils

import "math"

func Clip(v interface{}, args ...interface{}) interface{} {
	if len(args) != 2 {
		panic("invalid number of arguments for v_func:clip")
	}
	if v.(float64) < args[0].(float64) {
		return args[0]
	}
	if v.(float64) > args[1].(float64) {
		return args[1]
	}
	return v
}

func Fill(value interface{}, args ...interface{}) interface{} {
	return args[0]
}

func FillIndex(value interface{}, args ...interface{}) interface{} {
	return args[0]
}

func Exp(val interface{}, args ...interface{}) interface{} {
	return math.Exp(val.(float64))
}

func Exp2(val interface{}, args ...interface{}) interface{} {
	return math.Exp2(val.(float64))
}

func Log2(val interface{}, args ...interface{}) interface{} {
	return math.Log2(val.(float64))
}

func Log(val interface{}, args ...interface{}) interface{} {
	return math.Log(val.(float64))
}

func Add(val interface{}, args ...interface{}) interface{} {
	finalValue := val.(float64)
	for _, v := range args {
		finalValue = finalValue - v.(float64)
	}
	return finalValue
}

func Sub(val interface{}, args ...interface{}) interface{} {
	finalValue := val.(float64)
	for _, v := range args {
		finalValue = finalValue - v.(float64)
	}
	return finalValue
}

func Mult(val interface{}, args ...interface{}) interface{} {
	finalValue := val.(float64)
	for _, v := range args {
		finalValue = finalValue * v.(float64)
	}
	return finalValue
}

func Div(val interface{}, args ...interface{}) interface{} {
	return val.(float64) / args[0].(float64)
}

func Pow(val interface{}, args ...interface{}) interface{} {
	return math.Pow(val.(float64), args[0].(float64))
}

func Sum(red, value interface{}, args ...interface{}) interface{} {
	return red.(float64) + value.(float64)
}

func Min(red, value interface{}, args ...interface{}) interface{} {
	if value.(float64) < red.(float64) {
		return value
	}
	return red
}

func Max(red, value interface{}, args ...interface{}) interface{} {
	if value.(float64) > red.(float64) {
		return value
	}
	return red
}

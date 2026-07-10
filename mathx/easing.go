package mathx

// https://github.com/ai/easings.net/blob/master/src/easings/easingsFunctions.ts

import "math"

// Easing is a function that takes a progress value (0-1) and returns an eased value
type Easing func(float64) float64

const (
	c1 = 1.70158
	c2 = c1 * 1.525
	c3 = c1 + 1
	c4 = (2 * math.Pi) / 3
	c5 = (2 * math.Pi) / 4.5
)

func bounceOut(x float64) float64 {
	n1 := 7.5625
	d1 := 2.75

	if x < 1/d1 {
		return n1 * x * x
	} else if x < 2/d1 {
		x -= 1.5 / d1
		return n1*x*x + 0.75
	} else if x < 2.5/d1 {
		x -= 2.25 / d1
		return n1*x*x + 0.9375
	} else {
		x -= 2.625 / d1
		return n1*x*x + 0.984375
	}
}

// EaseLinear returns the input value unchanged, resulting in a linear easing.
func EaseLinear(x float64) float64 { return x }

// EaseInQuad quadratic easing, starts slow and accelerates. See https://easings.net/#easeInQuad
func EaseInQuad(x float64) float64 { return x * x }

// EaseOutQuad quadratic easing, starts fast and decelerates. See https://easings.net/#easeOutQuad
func EaseOutQuad(x float64) float64 { return 1 - (1-x)*(1-x) }

// EaseInOutQuad quadratic easing, starts slow, accelerates in the middle, and decelerates at the end. See https://easings.net/#easeInOutQuad
func EaseInOutQuad(x float64) float64 {
	if x < 0.5 {
		return 2 * x * x
	}
	return 1 - math.Pow(-2*x+2, 2)/2
}

// EaseInCubic cubic easing, starts slow and accelerates. See https://easings.net/#easeInCubic
func EaseInCubic(x float64) float64 { return x * x * x }

// EaseOutCubic cubic easing, starts fast and decelerates. See https://easings.net/#easeOutCubic
func EaseOutCubic(x float64) float64 { return 1 - math.Pow(1-x, 3) }

// EaseInOutCubic cubic easing, starts slow, accelerates in the middle, and decelerates at the end. See https://easings.net/#easeInOutCubic
func EaseInOutCubic(x float64) float64 {
	if x < 0.5 {
		return 4 * x * x * x
	}
	return 1 - math.Pow(-2*x+2, 3)/2
}

// EaseInQuart quartic easing, starts slow and accelerates. See https://easings.net/#easeInQuart
func EaseInQuart(x float64) float64 { return x * x * x * x }

// EaseOutQuart quartic easing, starts fast and decelerates. See https://easings.net/#easeOutQuart
func EaseOutQuart(x float64) float64 { return 1 - math.Pow(1-x, 4) }

// EaseInOutQuart quartic easing, starts slow, accelerates in the middle, and decelerates at the end. See https://easings.net/#easeInOutQuart
func EaseInOutQuart(x float64) float64 {
	if x < 0.5 {
		return 8 * x * x * x * x
	}
	return 1 - math.Pow(-2*x+2, 4)/2
}

// EaseInQuint quintic easing, starts slow and accelerates. See https://easings.net/#easeInQuint
func EaseInQuint(x float64) float64 { return x * x * x * x * x }

// EaseOutQuint quintic easing, starts fast and decelerates. See https://easings.net/#easeOutQuint
func EaseOutQuint(x float64) float64 { return 1 - math.Pow(1-x, 5) }

// EaseInOutQuint quintic easing, starts slow, accelerates in the middle, and decelerates at the end. See https://easings.net/#easeInOutQuint
func EaseInOutQuint(x float64) float64 {
	if x < 0.5 {
		return 16 * x * x * x * x * x
	}
	return 1 - math.Pow(-2*x+2, 5)/2
}

// EaseInSine sinusoidal easing, starts slow and accelerates. See https://easings.net/#easeInSine
func EaseInSine(x float64) float64 { return 1 - math.Cos((x*math.Pi)/2) }

// EaseOutSine sinusoidal easing, starts fast and decelerates. See https://easings.net/#easeOutSine
func EaseOutSine(x float64) float64 { return math.Sin((x * math.Pi) / 2) }

// EaseInOutSine sinusoidal easing, starts slow, accelerates in the middle, and decelerates at the end. See https://easings.net/#easeInOutSine
func EaseInOutSine(x float64) float64 { return -(math.Cos(math.Pi*x) - 1) / 2 }

// EaseInExpo exponential easing, starts slow and accelerates. See https://easings.net/#easeInExpo
func EaseInExpo(x float64) float64 {
	if x == 0 {
		return 0
	}
	return math.Pow(2, 10*x-10)
}

// EaseOutExpo exponential easing, starts fast and decelerates. See https://easings.net/#easeOutExpo
func EaseOutExpo(x float64) float64 {
	if x == 1 {
		return 1
	}
	return 1 - math.Pow(2, -10*x)
}

// EaseInOutExpo exponential easing, starts slow, accelerates in the middle, and decelerates at the end. See https://easings.net/#easeInOutExpo
func EaseInOutExpo(x float64) float64 {
	switch {
	case x == 0:
		return 0
	case x == 1:
		return 1
	case x < 0.5:
		return math.Pow(2, 20*x-10) / 2
	default:
		return (2 - math.Pow(2, -20*x+10)) / 2
	}
}

// EaseInCirc circular easing, starts slow and accelerates. See https://easings.net/#easeInCirc
func EaseInCirc(x float64) float64 { return 1 - math.Sqrt(1-math.Pow(x, 2)) }

// EaseOutCirc circular easing, starts fast and decelerates. See https://easings.net/#easeOutCirc
func EaseOutCirc(x float64) float64 { return math.Sqrt(1 - math.Pow(x-1, 2)) }

// EaseInOutCirc circular easing, starts slow, accelerates in the middle, and decelerates at the end. See https://easings.net/#easeInOutCirc
func EaseInOutCirc(x float64) float64 {
	if x < 0.5 {
		return (1 - math.Sqrt(1-math.Pow(2*x, 2))) / 2
	}
	return (math.Sqrt(1-math.Pow(-2*x+2, 2)) + 1) / 2
}

// EaseInBack back easing, starts slow and accelerates. See https://easings.net/#easeInBack
func EaseInBack(x float64) float64 { return c3*x*x*x - c1*x*x }

// EaseOutBack back easing, starts fast and decelerates. See https://easings.net/#easeOutBack
func EaseOutBack(x float64) float64 { return 1 + c3*math.Pow(x-1, 3) + c1*math.Pow(x-1, 2) }

// EaseInOutBack back easing, starts slow, accelerates in the middle, and decelerates at the end. See https://easings.net/#easeInOutBack
func EaseInOutBack(x float64) float64 {
	if x < 0.5 {
		return (math.Pow(2*x, 2) * ((c2+1)*2*x - c2)) / 2
	}
	return (math.Pow(2*x-2, 2)*((c2+1)*(x*2-2)+c2) + 2) / 2
}

// EaseInElastic elastic easing, starts slow and accelerates. See https://easings.net/#easeInElastic
func EaseInElastic(x float64) float64 {
	switch {
	case x == 0:
		return 0
	case x == 1:
		return 1
	default:
		return -math.Pow(2, 10*x-10) * math.Sin((x*10-10.75)*c4)
	}
}

// EaseOutElastic elastic easing, starts fast and decelerates. See https://easings.net/#easeOutElastic
func EaseOutElastic(x float64) float64 {
	switch {
	case x == 0:
		return 0
	case x == 1:
		return 1
	default:
		return math.Pow(2, -10*x)*math.Sin((x*10-0.75)*c4) + 1
	}
}

// EaseInOutElastic elastic easing, starts slow, accelerates in the middle, and decelerates at the end. See https://easings.net/#easeInOutElastic
func EaseInOutElastic(x float64) float64 {
	switch {
	case x == 0:
		return 0
	case x == 1:
		return 1
	case x < 0.5:
		return -(math.Pow(2, 20*x-10) * math.Sin((20*x-11.125)*c5)) / 2
	default:
		return (math.Pow(2, -20*x+10)*math.Sin((20*x-11.125)*c5))/2 + 1
	}
}

// EaseInBounce bounce easing, starts slow and accelerates. See https://easings.net/#easeInBounce
func EaseInBounce(x float64) float64 { return 1 - bounceOut(1-x) }

// EaseOutBounce bounce easing, starts fast and decelerates. See https://easings.net/#easeOutBounce
func EaseOutBounce(x float64) float64 { return bounceOut(x) }

// EaseInOutBounce bounce easing, starts slow, accelerates in the middle, and decelerates at the end. See https://easings.net/#easeInOutBounce
func EaseInOutBounce(x float64) float64 {
	if x < 0.5 {
		return (1 - bounceOut(1-2*x)) / 2
	}
	return (1 + bounceOut(2*x-1)) / 2
}

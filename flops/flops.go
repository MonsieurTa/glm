package flops

import (
	"github.com/luxengine/math"
)

const (
	epsilon = 0.000001
)

// refequal returns true if the floats are approximatelly equal. this function
// is used as reference for the unwrapped equal.
func refequal(a, b float32) bool {
	return math.Abs(a-b) <= epsilon*math.Max(math.Max(1, math.Abs(a)), math.Abs(b))
}

// Eq returns true if the floats are approximatelly equal.
func Eq(a, b float32) bool {
	if a > 0 {
		if b > 0 {
			// a>0 b>0
			if a > b {
				// a>0 b>0 a>b
				if a > 1 {
					// a>0 b>0 a>b a>1
					return a-b <= epsilon*a
				}
				// a>0 b>0 a>b a<1
				return a-b <= epsilon
			}
			// a>0 b>0 a<=b
			if b > 1 {
				// a>0 b>0 a<=b b>1
				return b-a <= epsilon*b
			}
			// a>0 b>0 a<=b b<1
			return b-a <= epsilon
		}
		// a>0 b<0
		return false
	}
	if b > 0 {
		// a<0 b>0
		return false
	}
	// a<0 b<0
	if a > b {
		// a<0 b<0 a>b
		if b < -1 {
			// a<0 b<0 a>b a<-1
			return -(b - a) <= epsilon*-b
		}
		// a<0 b<0 a>b b>=-1
		return -(b - a) <= epsilon
	}
	// a<0 b<0 a<b
	if a < -1 {
		// a<0 b<0 a<b a<-1
		return -(a - b) <= epsilon*-a
	}
	// a<0 b<0 a<b a>=-1
	return -(a - b) <= epsilon
}

// Ne returns true if the floats are not approximately equal
func Ne(a, b float32) bool {
	if a > 0 {
		if b > 0 {
			// a>0 b>0
			if a > b {
				// a>0 b>0 a>b
				if a > 1 {
					// a>0 b>0 a>b a>1
					return a-b > epsilon*a
				}
				// a>0 b>0 a>b a<1
				return a-b > epsilon
			}
			// a>0 b>0 a>b
			if b > 1 {
				// a>0 b>0 a>b b>1
				return b-a > epsilon*b
			}
			// a>0 b>0 a>b b<1
			return b-a > epsilon
		}
		// a>0 b<0
		return true
	}
	if b > 0 {
		// a<0 b>0
		return true
	}
	// a<0 b<0
	if a > b {
		// a<0 b<0 a>b
		if b < -1 {
			// a<0 b<0 a>b a<-1
			return -(b - a) > epsilon*-b
		}
		// a<0 b<0 a>b b>=-1
		return -(b - a) > epsilon
	}
	// a<0 b<0 a<b
	if a < -1 {
		// a<0 b<0 a<b a<-1
		return -(a - b) > epsilon*-a
	}
	// a<0 b<0 a<b a>=-1
	return -(a - b) > epsilon
}

// Lt returns true if a is strictly less than b. Even if a<b would return true
// they could in fact be equal.
func Lt(a, b float32) bool {
	return a < b && Ne(a, b)
}

// Le returns true if a is less than or equal to b. Even if a<b would return
// true they could in fact be equal.
func Le(a, b float32) bool {
	return a < b || Eq(a, b)
}

// Gt returns true if a is strictly greater than b. Even if a>b would return
// true they could in fact be equal.
func Gt(a, b float32) bool {
	return a > b && Ne(a, b)
}

// Ge returns true if a is greater than or equal to b. Even if a>b would return
// true they could in fact be equal.
func Ge(a, b float32) bool {
	return a > b || Eq(a, b)
}

// refz returns true if a is close to zero. This is the reference implementation
// used in testing and benchmarking.
func refz(a float32) bool {
	if math.Abs(a) <= epsilon*math.Max(1, math.Abs(a)) {
		return true
	}
	return false
}

// Z returns true if a is roughly equal to zero
func Z(a float32) bool {
	if a > 0 {
		return a <= epsilon
	}
	return -a <= epsilon
}

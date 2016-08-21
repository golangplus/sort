// Package less can generates some useful less functions.
package less

// Float64WithTie returns true/false if a is strictly less/greater than b, or
// returns result of tie. tie is called only when necessary.
func Float64WithTie(a, b float64, tie func() bool) bool {
	if a < b {
		return true
	}
	if a > b {
		return false
	}
	return tie()
}

// IntWithTie returns true/false if a is strictly less/greater than b, or
// returns result of tie. tie is called only when necessary.
func IntWithTie(a, b int, tie func() bool) bool {
	if a < b {
		return true
	}
	if a > b {
		return false
	}
	return tie()
}

// StrLenOrContent first compare length of a, b. If tie, returns a < b.
func StrLenOrContent(a, b string) bool {
	if len(a) < len(b) {
		return true
	}
	if len(a) > len(b) {
		return false
	}
	return a < b
}

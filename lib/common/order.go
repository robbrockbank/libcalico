package common

// Order can hold either a numeric value, or the "default" (meaning
// apply last).
type Order struct {
	Float32OrString
}

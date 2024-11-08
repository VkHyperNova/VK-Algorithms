package algorithms

/* 
	Euclid's Algorithm
*/

func GreatestCommonDivisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

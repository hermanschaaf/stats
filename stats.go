package stats

import (
	"fmt"
	"math"
	"os"
	"sort"
)

// Mean returns the mean of an integer array as a float
func Mean(nums []int) (mean float64) {
	if len(nums) == 0 {
		return 0.0
	}
	for _, n := range nums {
		mean += float64(n)
	}
	return mean / float64(len(nums))
}

// Median returns the median of an integer array as a float
func Median(nums []int) (median float64) {
	if len(nums) == 0 {
		return 0.0
	}

	t := make([]int, len(nums))
	copy(t, nums)

	sort.Ints(t)
	l := len(t)
	if l == 0 {
		return 0.0
	} else if l%2 == 0 {
		median = Mean(t[l/2-1 : l/2+1])
	} else {
		median = float64(t[l/2])
	}
	return median
}

// Mode returns the value of the most common element, or the smallest
// value if multiple elements satisfy this criteria.
func Mode(nums []int) (mode int) {
	if len(nums) == 0 {
		return 0.0
	}

	m := map[int]int{}
	var lowest, count = 0, 0
	for _, n := range nums {
		m[n] += 1
		if m[n] > count || (m[n] == count && n < lowest) {
			count = m[n]
			lowest = n
		}
	}
	return lowest
}

// StandardDeviation returns the standard deviation of the slice
// as a float
func StandardDeviation(nums []int) (dev float64) {
	if len(nums) == 0 {
		return 0.0
	}

	m := Mean(nums)
	for _, n := range nums {
		dev += (float64(n) - m) * (float64(n) - m)
	}
	dev = math.Pow(dev/float64(len(nums)), 0.5)
	return dev
}

// NormalConfidenceInterval returns the 95% confidence interval for the mean
// as two float values, the lower and the upper bounds and assuming a normal
// distribution
func NormalConfidenceInterval(nums []int) (lower float64, upper float64) {
	conf := 1.95996 // 95% confidence for the mean, http://bit.ly/Mm05eZ
	mean := Mean(nums)
	dev := StandardDeviation(nums) / math.Sqrt(float64(len(nums)))
	return mean - dev*conf, mean + dev*conf
}

func main() {
	var N, tmp int

	fmt.Fscanln(os.Stdin, &N)
	var nums []int

	for i := 0; i < N; i++ {
		fmt.Fscan(os.Stdin, &tmp)
		nums = append(nums, tmp)
	}

	mean := Mean(nums)
	median := Median(nums)
	mode := Mode(nums)
	dev := StandardDeviation(nums)
	lower, upper := NormalConfidenceInterval(nums)

	fmt.Printf("%.1f\n", mean)
	fmt.Printf("%.1f\n", median)
	fmt.Printf("%d\n", mode)
	fmt.Printf("%.1f\n", dev)
	fmt.Printf("%.1f %.1f\n", lower, upper)
}

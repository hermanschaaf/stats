Stats
=====

A short, simple and very basic stats package for Go, with some common tools:

 - `Mean` returns the mean of an integer array as a float
 - `Median` returns the median of an integer array as a float
 - `Mode` returns the value of the most common element, or the smallest value if multiple elements satisfy this criteria.
 - `StandardDeviation` returns the standard deviation of the slice as a float
 - `NormalConfidenceInterval` returns the 95% confidence interval for the mean as two float values, the lower and the upper bounds and assuming a normal distribution
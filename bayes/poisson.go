// Bayesian inference about the parameter λ of Poisson distribution.
// Bolstad 2007 (2e): Chapter 10, p. 183 and further.
// sumK	sum of observations over all repetitions
// n number of repetitions

package bayes

import (
	. "github.com/datastream/probab/dst"
	//	. "github.com/datastream/go-fn/fn"
	"math"
)

// Poisson λ, posterior PDF, flat prior.
func PoissonLambdaPDFFPri(sumK, n int64) func(p float64) float64 {
	// CAUTION !!! v= 1/scale !!!
	if sumK < 0 || n <= 0 {
		panic("bad data")
	}
	r1 := float64(sumK) + 1.0
	v1 := float64(n)
	return GammaPDF(r1, 1/v1)
}

// Poisson λ, posterior PDF, Jeffreys' prior.
func PoissonLambdaPDFJPri(sumK, n int64) func(p float64) float64 {
	// CAUTION !!! v= 1/scale !!!
	if sumK < 0 || n <= 0 {
		panic("bad data")
	}
	r1 := float64(sumK) + 0.5
	v1 := float64(n)
	return GammaPDF(r1, 1/v1)
}

// Poisson λ, posterior PDF, gamma prior.
// Use r=m^2/s^2, and v=m/s^2, if you summarize your prior belief with mean == m, and std == s.
func PoissonLambdaPDFGPri(sumK, n int64, r, v float64) func(p float64) float64 {
	// CAUTION !!! v= 1/scale !!!
	if sumK < 0 || n <= 0 {
		panic("bad data")
	}
	if r < 0 || v < 0 {
		panic("Shape parameter r and rate parameter v must be greater than or equal to zero")
	}
	r1 := r + float64(sumK)
	v1 := v + float64(n)
	return GammaPDF(r1, 1/v1)
}

// Poisson λ, posterior CDF, flat prior.
func PoissonLambdaCDFFPri(sumK, n int64) func(p float64) float64 {
	// CAUTION !!! v= 1/scale !!!
	if sumK < 0 || n <= 0 {
		panic("bad data")
	}
	r1 := float64(sumK) + 1.0
	v1 := float64(n)
	return GammaCDF(r1, 1/v1)
}

// Poisson λ, posterior CDF, Jeffreys' prior.
func PoissonLambdaCDFJPri(sumK, n int64) func(p float64) float64 {
	// CAUTION !!! v= 1/scale !!!
	if sumK < 0 || n <= 0 {
		panic("bad data")
	}
	r1 := float64(sumK) + 0.5
	v1 := float64(n)
	return GammaCDF(r1, 1/v1)
}

// Poisson λ, posterior CDF, gamma prior.
// Use r=m^2/s^2, and v=m/s^2, if you summarize your prior belief with mean == m, and std == s.
func PoissonLambdaCDFGPri(sumK, n int64, r, v float64) func(p float64) float64 {
	// CAUTION !!! v= 1/scale !!!
	if sumK < 0 || n <= 0 {
		panic("bad data")
	}
	if r < 0 || v < 0 {
		panic("Shape parameter r and rate parameter v must be greater than or equal to zero")
	}
	r1 := r + float64(sumK)
	v1 := v + float64(n)
	return GammaCDF(r1, 1/v1)
}

// Poisson λ, posterior quantile function, flat prior.
func PoissonLambdaQtlFPri(sumK, n int64) func(p float64) float64 {
	// CAUTION !!! v= 1/scale !!!
	if sumK < 0 || n <= 0 {
		panic("bad data")
	}
	r1 := float64(sumK) + 1.0
	v1 := float64(n)
	return GammaQtl(r1, 1/v1)
}

// Poisson λ, posterior quantile function, Jeffreys' prior.
func PoissonLambdaQtlJPri(sumK, n int64) func(p float64) float64 {
	// CAUTION !!! v= 1/scale !!!
	if sumK < 0 || n <= 0 {
		panic("bad data")
	}
	r1 := float64(sumK) + 0.5
	v1 := float64(n)
	return GammaQtl(r1, 1/v1)
}

// Poisson λ, posterior quantile function, gamma prior.
// Use r=m^2/s^2, and v=m/s^2, if you summarize your prior belief with mean == m, and std == s.
func PoissonLambdaQtlGPri(sumK, n int64, r, v float64) func(p float64) float64 {
	// CAUTION !!! v= 1/scale !!!
	if sumK < 0 || n <= 0 {
		panic("bad data")
	}
	if r < 0 || v < 0 {
		panic("Shape parameter r and rate parameter v must be greater than or equal to zero")
	}
	r1 := r + float64(sumK)
	v1 := v + float64(n)
	return GammaQtl(r1, 1/v1)
}

// PoissonLambdaNextFPri returns random number drawn from the posterior, flat prior.
func PoissonLambdaNextFPri(sumK, n int64) float64 {
	if sumK < 0 || n <= 0 {
		panic("bad data")
	}
	r1 := float64(sumK) + 1.0
	v1 := float64(n)
	return GammaNext(r1, 1/v1)
}

// PoissonLambdaNextJPri returns random number drawn from the posterior, Jeffreys' prior.
func PoissonLambdaNextJPri(sumK, n int64) float64 {
	if sumK < 0 || n <= 0 {
		panic("bad data")
	}
	r1 := float64(sumK) + 0.5
	v1 := float64(n)
	return GammaNext(r1, 1/v1)
}

// PoissonLambdaNextGPri returns random number drawn from the posterior, Gamma prior.
func PoissonLambdaNextGPri(sumK, n int64, r, v float64) float64 {
	if sumK < 0 || n <= 0 {
		panic("bad data")
	}
	if r < 0 || v < 0 {
		panic("Shape parameter r and rate parameter v must be greater than or equal to zero")
	}
	r1 := r + float64(sumK)
	v1 := v + float64(n)
	return GammaNext(r1, 1/v1)
}

// Likelihood of Poisson λ.
// Bolstad 2007 (2e): Chapter 10, p. 184.
func PoissonLambdaLike(sumK, n int64, λ float64) float64 {
	return λ * float64(sumK) * math.Exp(float64(-n)*λ)

}

// Equivalent sample size of the prior 
// Bolstad 2007 (2e): Chapter 10, p. 187.
func PoissonLambdaEqvSize(v float64) float64 {
	return (math.Floor(v))
}

// Posterior mean 
// Bolstad 2007 (2e): Chapter 10, p. 190-191.
func PoissonLambdaPostMean(sumK, n int64, r, v float64) float64 {
	r1 := float64(sumK) + 1.0
	v1 := 1.0
	return r1 / v1
}

// Posterior mean bias
// Bolstad 2007 (2e): Chapter 10, p. 191.
func PoissonLambdaPostMeanBias(r, v, λ float64) float64 {
	return (r - v*λ) / (v + 1)
}

// Posterior variance
// Bolstad 2007 (2e): Chapter 10, p. 191.
func PoissonLambdaPostVar(r, v, λ float64) float64 {
	return λ / (v * v)
}

// Mean Squared Error of λ
// Bolstad 2007 (2e): Chapter 10, p. 191.
func PoissonLambdaMSE(r, v, λ float64) float64 {
	bsq := PoissonLambdaPostMeanBias(r, v, λ)
	bsq *= bsq
	variance := PoissonLambdaPostVar(r, v, λ)
	return (bsq + variance)
}

// posterior interquartile range of λ
// Bolstad 2007 (2e): Chapter 10, p. 189.
func PoissonLambdaIQR(sumK, n int64, r, v float64) float64 {
	qf := PoissonLambdaQtlGPri(sumK, n, r, v)
	q1 := qf(0.25)
	q3 := qf(0.75)
	return (q3 - q1)
}

// Credible interval for unknown Poisson rate λ, and gamma prior, equal tail area
// Bolstad 2007 (2e): 192-193.
// untested ...
func PoissonLambdaCrIGPri(sumK, n int64, r, v, α float64) (lo, hi float64) {
	/*
		sumK, n			total observed events in n equal time intervals
		r			gamma prior r
		v			gamma prior v
		α		posterior probability that the true proportion lies outside the credible interval
	*/
	// return value: lo is lower boundary, hi upper
	qf := PoissonLambdaQtlGPri(sumK, n, r, v)
	lo = qf(α / 2)
	hi = qf(1 - α/2)
	return
}

// One-sided test for Poisson rate λ
// Bolstad 2007 (2e): 193.
// H0: λ <= λ0 vs H1: λ > λ0
// Note: The alternative is in the direction we wish to detect.
func PoissonLambdaOneSidedTst(sumK, n int64, r, v, α, λ0 float64) bool {
	cdf := PoissonLambdaCDFGPri(sumK, n, r, v)
	p0 := cdf(λ0)
	reject := false // hypothesis NOT rejected (default)
	if p0 < α {
		reject = true // hypothesis rejected
	}
	return reject
}

// One-sided odds ratio for Poisson rate λ
// Bolstad 2007 (2e): 193.
// H0: λ <= λ0 vs H1: λ > λ0
// Note: The alternative is in the direction we wish to detect.
func PoissonLambdaOneSidedOdds(sumK, n int64, r, v, λ0 float64) float64 {
	cdf := PoissonLambdaCDFGPri(sumK, n, r, v)
	p0 := cdf(λ0)
	return p0 / (1 - p0)
}

// Two-sided test for Poisson rate λ
// Bolstad 2007 (2e): 194.
// H0: λ = λ0 vs H1: λ != λ0
func PoissonLambdaTwoSidedTst(sumK, n int64, r, v, α, λ0 float64) bool {
	low, high := PoissonLambdaCrIGPri(sumK, n, r, v, α)
	reject := false // hypothesis NOT rejected (default)
	if λ0 < low || λ0 > high {
		reject = true // hypothesis rejected
	}
	return reject
}

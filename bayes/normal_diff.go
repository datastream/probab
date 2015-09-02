// Bayesian inference about the difference between means of the Normal (Gaussian) distributions.
// Bolstad 2007 (2e): Chapter 13: 239 - 266.

package bayes

import (
	. "github.com/datastream/probab/dst"
	"math"
)

// KNOWN variances, and NORMAL priors

// Posterior PDF of the difference of two means (μ1-μ2) of Normal distributions with KNOWN variances, and NORMAL priors
// Bolstad 2007:245-246
func NormalMuDiffPDFNPriKn(nObs1, nObs2 int, ȳ1, ȳ2, σ1, σ2, μ1Pri, σ1Pri, μ2Pri, σ2Pri float64) func(x float64) float64 {
	// for independent samples, use independent priors for both means
	// posteriors are Normal with params from eqs. 11.5 and 11.6
	μ1Post := NormMuPostMean(nObs1, ȳ1, σ1, μ1Pri, σ1Pri)
	σ1Post := NormMuPostStd(nObs1, σ1, μ1Pri, σ1Pri)
	μ2Post := NormMuPostMean(nObs2, ȳ2, σ2, μ2Pri, σ2Pri)
	σ2Post := NormMuPostStd(nObs2, σ2, μ2Pri, σ2Pri)
	//difference posterior is Normal with params:
	μdPost := μ1Post - μ2Post
	σdPost := math.Sqrt(σ1Post*σ1Post + σ2Post*σ2Post)
	return NormalPDF(μdPost, σdPost)
}

// Posterior CDF of the difference of two means (μ1-μ2) of Normal distributions with KNOWN variances, and NORMAL priors
// Bolstad 2007:245-246
func NormalMuDiffCDFNPriKn(nObs1, nObs2 int, ȳ1, ȳ2, σ1, σ2, μ1Pri, σ1Pri, μ2Pri, σ2Pri float64) func(x float64) float64 {
	// for independent samples, use independent priors for both means
	// posteriors are Normal with params from eqs. 11.5 and 11.6
	μ1Post := NormMuPostMean(nObs1, ȳ1, σ1, μ1Pri, σ1Pri)
	σ1Post := NormMuPostStd(nObs1, σ1, μ1Pri, σ1Pri)
	μ2Post := NormMuPostMean(nObs2, ȳ2, σ2, μ2Pri, σ2Pri)
	σ2Post := NormMuPostStd(nObs2, σ2, μ2Pri, σ2Pri)
	//difference posterior is Normal with params:
	μdPost := μ1Post - μ2Post
	σdPost := math.Sqrt(σ1Post*σ1Post + σ2Post*σ2Post)
	return NormalCDF(μdPost, σdPost)
}

// Posterior quantile of the difference of two means (μ1-μ2) of Normal distributions with KNOWN variances, and NORMAL priors
// Bolstad 2007:245-246
func NormalMuDiffQtlNPriKn(nObs1, nObs2 int, ȳ1, ȳ2, σ1, σ2, μ1Pri, σ1Pri, μ2Pri, σ2Pri float64) func(p float64) float64 {
	// for independent samples, use independent priors for both means
	// posteriors are Normal with params from eqs. 11.5 and 11.6
	μ1Post := NormMuPostMean(nObs1, ȳ1, σ1, μ1Pri, σ1Pri)
	σ1Post := NormMuPostStd(nObs1, σ1, μ1Pri, σ1Pri)
	μ2Post := NormMuPostMean(nObs2, ȳ2, σ2, μ2Pri, σ2Pri)
	σ2Post := NormMuPostStd(nObs2, σ2, μ2Pri, σ2Pri)
	//difference posterior is Normal with params:
	μdPost := μ1Post - μ2Post
	σdPost := math.Sqrt(σ1Post*σ1Post + σ2Post*σ2Post)
	return NormalQtl(μdPost, σdPost)
}

// UNKNOWN variances (Behrens-Fisher problem), and NORMAL priors
// Bolstad 2007: 246-248.

// Variance estimate from single sample from Normal distribution with unknown variance
// Bolstad 2007 (2e): 246
// untested ...
func varest(y []float64, nObs int) float64 {
	n := float64(nObs)
	mean := 0.0
	for i := 0; i < nObs; i++ {
		mean += y[i]
	}
	mean /= n
	sum := 0.0
	for i := 0; i < nObs; i++ {
		sum += (y[i] - mean) * (y[i] - mean)
	}
	return sum / (n - 1)
}

// Satterthwaite's adjusted degrees of freedom
// Bolstad 2007 (2e): 247.
// Satterthwaite, F.E. 1941: Synthesis of variance.  Psychometrika, 6 (5), pp. 309-316. 
// untested ...
func satterthwaitenu(estvar1 float64, nObs1 int, estvar2 float64, nObs2 int) float64 {
	var nu float64
	n1 := float64(nObs1)
	n2 := float64(nObs2)
	f1 := (estvar1/n1 + estvar2/n2) * (estvar1/n1 + estvar2/n2)
	f2 := (estvar1 / n1) * (estvar1 / n1) / (n1 + 1)
	f3 := (estvar2 / n2) * (estvar2 / n2) / (n2 + 1)

	// round to nearest integer
	v := f1 / (f2 + f3)
	if v-math.Floor(v) <= math.Ceil(v)-v {
		nu = math.Floor(v)
	} else {
		nu = math.Ceil(v)
	}
	return nu
}

// Quantile of the difference of two means (μ1-μ2) of Normal distributions with UNKNOWN variances (Behrens-Fisher problem), and NORMAL priors 
// Bolstad 2007:245-246
// untested ...
func NormalMuDiffQtlNPriUn(nObs1, nObs2 int, ȳ1, ȳ2, s1, s2, μ1Pri, σ1Pri, μ2Pri, σ2Pri, p float64) func(p float64) float64 {
	// for independent samples, use independent priors for both means
	// s1 and s2 are estimated standard deviations math.Sqrt(varest())
	return func(p float64) float64 {
		var q float64
		μ1Post := NormMuPostMean(nObs1, ȳ1, s1, μ1Pri, σ1Pri)
		σ1Post := NormMuPostStd(nObs1, s1, μ1Pri, σ1Pri)
		μ2Post := NormMuPostMean(nObs2, ȳ2, s2, μ2Pri, σ2Pri)
		σ2Post := NormMuPostStd(nObs2, s2, μ2Pri, σ2Pri)
		//difference posterior is Normal with params:
		μdPost := μ1Post - μ2Post
		nu := satterthwaitenu(s1*s1, nObs1, s2*s2, nObs2)
		t := StudentsTQtl(nu)
		α := 1 - 2*p
		if p < 0.5 {
			q = μdPost - t(α/2)*math.Sqrt(σ1Post*σ1Post+σ2Post*σ2Post)
		} else {
			q = μdPost + t(α/2)*math.Sqrt(σ1Post*σ1Post+σ2Post*σ2Post)
		}
		return q
	}
}

// Credible interval of the difference of two means (μ1-μ2) of Normal distributions with UNKNOWN variances (Behrens-Fisher problem), and NORMAL priors 
// Bolstad 2007:245-246
// untested ...
func NormalMuDiffCrINPriUn(nObs1, nObs2 int, ȳ1, ȳ2, s1, s2, μ1Pri, σ1Pri, μ2Pri, σ2Pri, α float64) func(α float64) (lo, hi float64) {
	// for independent samples, use independent priors for both means
	// s1 and s2 are estimated standard deviations math.Sqrt(varest())
	return func(α float64) (lo, hi float64) {
		μ1Post := NormMuPostMean(nObs1, ȳ1, s1, μ1Pri, σ1Pri)
		σ1Post := NormMuPostStd(nObs1, s1, μ1Pri, σ1Pri)
		μ2Post := NormMuPostMean(nObs2, ȳ2, s2, μ2Pri, σ2Pri)
		σ2Post := NormMuPostStd(nObs2, s2, μ2Pri, σ2Pri)
		//difference posterior is Normal with params:
		μdPost := μ1Post - μ2Post
		nu := satterthwaitenu(s1*s1, nObs1, s2*s2, nObs2)
		t := StudentsTQtl(nu)
		lo = μdPost - t(α/2)*math.Sqrt(σ1Post*σ1Post+σ2Post*σ2Post)
		hi = μdPost + t(α/2)*math.Sqrt(σ1Post*σ1Post+σ2Post*σ2Post)
		return
	}
}

// Credible interval of the difference of two means (μ1-μ2) of Normal distributions with UNKNOWN variances (Behrens-Fisher problem), and FLAT priors
// Bolstad 2007:245-246
// untested ...
func NormalMuDiffCrIFPriUn(nObs1, nObs2 int, ȳ1, ȳ2, s1, s2, μ1Pri, σ1Pri, μ2Pri, σ2Pri, α float64) func(α float64) (lo, hi float64) {
	// for independent samples, use independent priors for both means
	// s1 and s2 are estimated standard deviations math.Sqrt(varest())
	return func(α float64) (lo, hi float64) {
		μ1Post := NormMuPostMean(nObs1, ȳ1, s1, μ1Pri, σ1Pri)
		μ2Post := NormMuPostMean(nObs2, ȳ2, s2, μ2Pri, σ2Pri)
		//difference posterior is Normal with params:
		μdPost := μ1Post - μ2Post
		nu := satterthwaitenu(s1*s1, nObs1, s2*s2, nObs2)
		t := StudentsTQtl(nu)
		lo = μdPost - t(α/2)*math.Sqrt(s1*s1+s2*s2)
		hi = μdPost + t(α/2)*math.Sqrt(s1*s1+s2*s2)
		return
	}
}

// Posterior moments
// Mean = modus = median; standard deviation; skewness = 0; kurtosis = 0;

// Posterior moments of the difference of two means (μ1-μ2) of Normal distributions with KNOWN variances, and NORMAL priors
// Bolstad 2007:245-246
func NormalMuDiffMomentsNPriKn(nObs1, nObs2 int, ȳ1, ȳ2, σ1, σ2, μ1Pri, σ1Pri, μ2Pri, σ2Pri float64) (μ, σ float64) {
	// for independent samples, use independent priors for both means
	// posteriors are Normal with params from eqs. 11.5 and 11.6
	μ1Post := NormMuPostMean(nObs1, ȳ1, σ1, μ1Pri, σ1Pri)
	σ1Post := NormMuPostStd(nObs1, σ1, μ1Pri, σ1Pri)
	μ2Post := NormMuPostMean(nObs2, ȳ2, σ2, μ2Pri, σ2Pri)
	σ2Post := NormMuPostStd(nObs2, σ2, μ2Pri, σ2Pri)
	//difference posterior is Normal with params:
	μdPost := μ1Post - μ2Post
	σdPost := math.Sqrt(σ1Post*σ1Post + σ2Post*σ2Post)
	return μdPost, σdPost
}

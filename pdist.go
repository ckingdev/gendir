package gendir

import (
	"math/rand"
)

// GenDirichlet implements the generalized Dirichlet probability distribution.
//
// The generalized Dirichlet distribution is a continuous probability
// distribution that, like the Dirichlet, generates elements over the
// probability simplex. It is also conjugate to the multinomial distribution
// but has 2(K-1) parameters for a K-dimensional distribution, and admits a
// more complex dependence structure between dimensions.
//
// For more information see https://en.wikipedia.org/wiki/Generalized_Dirichlet_distribution
type GenDirichlet struct {
	alpha []float64
	beta  []float64
	dim   int
	src   rand.Source
}

func NewGenDirichlet(alpha, beta []float64, src rand.Source) *GenDirichlet {
	dim := len(alpha) + 1

	a := make([]float64, len(alpha))
	copy(a, alpha)
	b := make([]float64, len(beta))
	copy(b, beta)

	gd := &GenDirichlet{
		alpha: a,
		beta:  b,
		dim:   dim,
		src:   src,
	}
	return gd
}

// Dim returns the dimension of the distribution.
func (gd *GenDirichlet) Dim() int {
	return gd.dim
}

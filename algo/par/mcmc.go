package par

import (
	"distclus/algo"
)

type ParMCMCSupport struct {
	config algo.MCMCConf
}

func (supp ParMCMCSupport) Iterate(m algo.MCMC, clust algo.Clust, iter int) algo.Clust {
	var conf = algo.KMeansConf{K: len(clust), Iter: iter, Space: supp.config.Space}
	var km = NewKMeans(conf, clust.Initializer)

	km.Data = m.Data

	km.Run(false)
	km.Close()

	var result, _ = km.Centroids()
	return result

}

func (supp ParMCMCSupport) Loss(m algo.MCMC, proposal algo.Clust) float64 {
	return proposal.Loss(m.Data, supp.config.Space, supp.config.Norm)
}

func NewMCMC(conf algo.MCMCConf, distrib algo.MCMCDistrib, initializer algo.Initializer) algo.MCMC  {
	var mcmc = algo.NewMCMC(conf, distrib, initializer)

	mcmc.MCMCSupport = ParMCMCSupport{conf}

	return mcmc
}


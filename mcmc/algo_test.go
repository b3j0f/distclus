package mcmc_test

import (
	"distclus/core"
	"distclus/kmeans"
	"distclus/mcmc"
	"distclus/real"
	"testing"

	"golang.org/x/exp/rand"
)

var conf = mcmc.Conf{
	InitK:     3,
	FrameSize: 8,
	RGen:      rand.New(rand.NewSource(6305689164243)),
	Dim:       3, B: 100, Amp: 1,
	Norm: 2, Nu: 3, McmcIter: 20,
	InitIter: 1,
}
var data = []core.Elemt{
	[]float64{1., 3.4, 5.4},
	[]float64{10., 9.2, 12.3},
	[]float64{-4.3, -1.2, -3.},
}
var initializer = kmeans.GivenInitializer

func Test_NewSeqAlgo(t *testing.T) {
	var algo = mcmc.NewAlgo(
		conf,
		real.Space{},
		data,
		initializer,
	)
	algo.AcceptRatio()
}

func Test_NewParAlgo(t *testing.T) {
	conf.Par = true
	mcmc.NewAlgo(
		conf,
		real.Space{},
		data,
		initializer,
	)
}

func Test_AlgoAcceptRatio(t *testing.T) {
	var algo = mcmc.NewAlgo(conf, real.Space{}, data, initializer)

	algo.Run(false)

	var acceptRatio = algo.AcceptRatio()
	var implAcceptRatio = algo.Impl().(*mcmc.Impl).AcceptRatio()

	if acceptRatio != implAcceptRatio {
		t.Error("Wrong routing of acceptratio method", acceptRatio, implAcceptRatio)
	}
}
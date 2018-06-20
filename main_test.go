package main

import (
	"testing"
	"time"
	"distclus/core"
	"distclus/algo"
	"golang.org/x/exp/rand"
)

func BenchmarkRun(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b1(b.Log)
	}
}

func b1(log func (args ...interface{})) {
	var data []core.Elemt
	var distrib algo.MCMCDistrib
	var initializer = algo.RandInitializer
	var seed = int(time.Now().UTC().Unix())
	var mcmcConf = algo.MCMCConf{
	}
	mcmcConf.Space = core.RealSpace{}
	in := "ca.csv"
	data, mcmcConf.Dim = parseFloatCsv(&in)
	mcmcConf.FrameSize = len(data)
	mcmcConf.RGen = rand.New(rand.NewSource(uint64(seed)))
	mcmcConf.McmcIter = 200
	mcmcConf.B = 100
	mcmcConf.Amp = 10
	mcmcConf.R = 8e6
	mcmcConf.InitIter = 0
	mcmcConf.InitK = 1
	mcmcConf.Norm = 2
	mcmcConf.Nu = 3
	distrib = algo.NewMultivT(algo.MultivTConf{mcmcConf})
	var mcmc = algo.NewMCMC(mcmcConf, distrib, initializer)

	for _, elt := range data {
		mcmc.Push(elt)
	}

	mcmc.Run(false)
	mcmc.Close()
	var centers, _ = mcmc.Centroids()
	var labels = make([]int, len(centers))
	for i := range data {
		var _, l, _ = centers.Assign(data[i], mcmcConf.Space)
		labels[l] += 1
	}
	log(labels)
}

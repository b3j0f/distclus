package kmeans_test

import (
	"distclus/core"
	"distclus/internal/test"
	"distclus/kmeans"
	"distclus/real"
	"testing"
)

var conf = core.Conf{ImplConf: kmeans.Conf{K: 1}, SpaceConf: nil}
var data = []core.Elemt{}
var initializer = kmeans.GivenInitializer

func Test_NewSeqAlgo(t *testing.T) {
	kmeans.NewAlgo(
		conf,
		real.Space{},
		data,
		initializer,
	)
}

func Test_NewParAlgo(t *testing.T) {
	kmeansConf := conf.ImplConf.(kmeans.Conf)
	kmeansConf.Par = true
	kmeans.NewAlgo(
		conf,
		real.Space{},
		data,
		initializer,
	)
}

func Test_Reset(t *testing.T) {
	algo := kmeans.NewAlgo(
		conf,
		real.Space{},
		data,
		initializer,
	)

	test.DoTestReset(t, &algo, core.Conf{ImplConf: kmeans.Conf{K: 1}, SpaceConf: nil})

	algo = kmeans.NewAlgo(
		core.Conf{ImplConf: kmeans.Conf{Par: true, K: 1}, SpaceConf: nil},
		real.Space{},
		data,
		initializer,
	)

	test.DoTestReset(t, &algo, core.Conf{ImplConf: kmeans.Conf{Par: true, K: 1}, SpaceConf: nil})
}

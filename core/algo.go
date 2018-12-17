package core

import (
	"errors"
	"fmt"
	"time"
)

// OnlineClust interface
// When a prediction is made, the element can be pushed to the model.
// A prediction consists in a centroid and a label.
// The following constraints must be met (otherwise an error is returned) :
// an element can't be pushed if the algorithm is closed,
// a prediction can't be done before the algorithm is run,
// no centroid can be returned before the algorithm is run.
type OnlineClust interface {
	Centroids() (Clust, error)
	Push(elemt Elemt) error
	Predict(elemt Elemt) (Elemt, int, error)
	Run(async bool) error
	Close() error
}

// Algo in charge of algorithm execution with both implementation and user configuration
type Algo struct {
	conf           Conf
	impl           Impl
	space          Space
	centroids      Clust
	status         ClustStatus
	closing        chan bool
	closed         chan bool
	lastUpdateTime int64
}

// AlgoConf algorithm configuration
type AlgoConf interface{}

// NewAlgo creates a new algorithm instance
func NewAlgo(conf Conf, impl Impl, space Space) Algo {
	return Algo{
		conf:    conf,
		impl:    impl,
		space:   space,
		status:  Created,
		closing: make(chan bool, 1),
		closed:  make(chan bool, 1),
	}
}

// Centroids Get the centroids currently found by the algorithm
func (algo *Algo) Centroids() (centroids Clust, err error) {
	switch algo.status {
	case Created:
		err = fmt.Errorf("clustering not started")
	default:
		var algoCentroids = algo.centroids
		centroids = make(Clust, len(algoCentroids))
		for index, centroid := range algoCentroids {
			centroids[index] = algo.space.Copy(centroid)
		}
	}
	return
}

// Push a new observation in the algorithm
func (algo *Algo) Push(elemt Elemt) (err error) {
	switch algo.status {
	case Closed:
		err = errors.New("clustering ended")
	default:
		err = algo.impl.Push(elemt)
	}
	return
}

// Run executes the algorithm and notify the user with a callback, timed by a time to callback (ttc) integer
func (algo *Algo) Run(async bool) (err error) {
	if async {
		if algo.status != Running {
			err = algo.impl.SetAsync()
			if err == nil {
				go algo.initAndRunAsync()
			}
		} else {
			err = errors.New("Algo is running")
		}
	} else {
		err = algo.initAndRunSync(async)
	}
	if err == nil {
		algo.status = Running
	}
	return
}

// Conf returns configuration
func (algo Algo) Conf() Conf {
	return algo.conf
}

// Impl returns impl
func (algo Algo) Impl() Impl {
	return algo.impl
}

// Space returns space
func (algo Algo) Space() Space {
	return algo.space
}

// Predict the cluster for a new observation
func (algo *Algo) Predict(elemt Elemt) (pred Elemt, label int, err error) {
	var clust Clust
	clust, err = algo.Centroids()

	if err == nil {
		pred, label, _ = clust.Assign(elemt, algo.space)
	}

	return
}

// Close Stops the algorithm
func (algo *Algo) Close() (err error) {
	if algo.status == Running {
		algo.closing <- true
		<-algo.closed
	}
	algo.status = Closed
	return
}

func (algo *Algo) updateCentroids(centroids Clust) {
	algo.centroids = centroids
}

// Initialize the algorithm, if success run it synchronously otherwise return an error
func (algo *Algo) initAndRunSync(async bool) (err error) {
	if !async {
		algo.centroids, err = algo.impl.Init(algo.conf.ImplConf, algo.space)
	}
	if err == nil {
		err = algo.impl.Run(
			algo.conf.ImplConf,
			algo.space,
			algo.centroids,
			algo.updateCentroids,
			algo.closing,
		)
		if err == nil && !async {
			algo.closed <- true
		}
	}

	return
}

// Initialize the algorithm, if success run it asynchronously otherwise retry
func (algo *Algo) initAndRunAsync() {
	var err error
	algo.centroids, err = algo.impl.Init(algo.conf.ImplConf, algo.space)
	for err == nil {
		err = algo.initAndRunSync(true)
		time.Sleep(300 * time.Millisecond)
	}
	algo.closed <- true
}

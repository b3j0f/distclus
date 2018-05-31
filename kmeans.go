package clustering_go

import (
	"fmt"
	"errors"
)

type KMeans struct {
	iter        int
	k           int
	data        []Elemt
	space       space
	status      ClustStatus
	initializer Initializer
	clust       Clust
}

func NewKMeans(k int, iter int, space space, initializer Initializer) KMeans {
	var km KMeans
	if k < 1 {
		panic(fmt.Sprintf("Illegal value for k: %v", k))
	}
	if k < 0 {
		panic(fmt.Sprintf("Illegal value for iter: %v", k))
	}
	km.iter = iter
	km.k = k
	km.initializer = initializer
	km.space = space
	km.status = Created
	return km
}

func (km *KMeans) initialize() (error) {
	if len(km.data) < km.k {
		return errors.New("can't initialize kmeans model centroids, not enough data")
	}
	var clust, err = km.initializer(km.k, km.data, km.space)
	if err != nil {
		panic(err)
	}
	km.clust = clust
	km.status = Initialized
	return nil
}

func (km *KMeans) Centroids() (c Clust, err error) {
	switch km.status {
	case Created:
		err = fmt.Errorf("no Clust available")
	default:
		c = km.clust
	}
	return c, err
}

func (km *KMeans) Push(elemt Elemt) {
	km.data = append(km.data, elemt)
}

func (km *KMeans) Close() {
	km.status = Closed
}

func (km *KMeans) Predict(elemt Elemt) (c Elemt, idx int, err error) {
	switch km.status {
	case Created:
		return c, idx, fmt.Errorf("no Clust available")
	default:
		var idx = assign(elemt, km.clust.centers, km.space)
		return km.clust.centers[idx], idx, nil
	}
}

func (km *KMeans) iteration() {
	var clusters = km.clust.Assign(&km.data, km.space)
	var centroids = km.clust.centers
	for k, cluster := range clusters {
		centroids[k] = mean(cluster, km.space)
	}
	var clustering, err = NewClustering(centroids)
	if err != nil {
		panic(err)
	}
	km.clust = clustering
}

func (km *KMeans) Run() {
	km.status = Running
	km.initialize()
	for iter := 0; iter < km.iter; iter++ {
		km.iteration()
	}
}

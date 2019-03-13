package core_test

import (
	"distclus/core"
	"distclus/internal/test"
	"distclus/vectors"
	"runtime"
	"testing"
)

func TestClust_ParAssignAll(t *testing.T) {
	var data = make([]core.Elemt, 0, len(test.Vectors)*20)
	var centroids = core.Clust(test.Vectors[0:3])
	for i := 0; i < 20; i++ {
		data = append(data, test.Vectors...)
	}

	var seqLabels = centroids.AssignAll(data, vectors.Space{})
	var parLabels = centroids.ParAssignAll(data, vectors.Space{}, runtime.NumCPU())

	test.AssertArrayEqual(t, seqLabels, parLabels)
}

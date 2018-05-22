package clustering_go

import (
	"testing"
	"math"
)

func TestRealDist2And4(t *testing.T) {
	e1 := []float64{2}
	e2 := []float64{4}
	space := realSpace{}
	val := space.dist(e1, e2)
	if val != 2 {
		t.Error("Expected 2, got ", val)
	}
}

func TestRealDist0And0(t *testing.T) {
	e1 := []float64{0}
	e2 := []float64{0}
	space := realSpace{}
	val := space.dist(e1, e2)
	if val != 0 {
		t.Error("Expected 0, got ", val)
	}
}

func TestRealDist2_2And4_4(t *testing.T) {
	e1 := []float64{2, 2}
	e2 := []float64{4, 4}
	res := math.Sqrt(8)
	space := realSpace{}
	val := space.dist(e1, e2)
	if val != res {
		t.Errorf("Expected %v, got %v", res, val)
	}
}

func TestRealDist_And4_4(t *testing.T) {
	var e1 []float64
	e2 := []float64{4, 4}
	space := realSpace{}
	var val float64
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, got %v", val)
		}
	}()
	val = space.dist(e1, e2)
}

func TestRealDist_And_(t *testing.T) {
	var e1 []float64
	var e2 []float64
	space := realSpace{}
	var val float64
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, got %v", val)
		}
	}()
	val = space.dist(e1, e2)
}

func TestRealDist2_1x2And4x2(t *testing.T) {
	e1 := []float64{2, 1}
	e2 := []float64{4}
	space := realSpace{}
	var val Elemt
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, got %v", val)
		}
	}()
	val = space.dist(e1, e2)
}

func TestRealCombine2x1And4x1(t *testing.T) {
	e1 := []float64{2}
	e2 := []float64{4}
	space := realSpace{}
	val := space.combine(e1, 1, e2, 1).([]float64)
	if val[0] != 3 {
		t.Errorf("Expected 3, got %v", val)
	}
}

func TestRealCombine2_1x2And4_2x2(t *testing.T) {
	e1 := []float64{2, 1}
	e2 := []float64{4, 2}
	space := realSpace{}
	val := space.combine(e1, 2, e2, 2).([]float64)
	if val[0] != 3 {
		t.Errorf("Expected 3, got %v", val[0])
	}
	if val[1] != 1.5 {
		t.Errorf("Expected 3/2, got %v", val[1])
	}
}

func TestRealCombine2_1x2And4x2(t *testing.T) {
	e1 := []float64{2, 1}
	e2 := []float64{4}
	space := realSpace{}
	var val Elemt
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, got %v", val)
		}
	}()
	val = space.combine(e1, 2, e2, 2).([]float64)
}

func TestRealCombine_And_(t *testing.T) {
	var e1 []float64
	var e2 []float64
	space := realSpace{}
	var val Elemt
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, got %v", val)
		}
	}()
	val = space.combine(e1, 1, e2, 1)
}

func TestRealCombine2_1x0And4_2x1(t *testing.T) {
	e1 := []float64{2, 1}
	e2 := []float64{4, 2}
	space := realSpace{}
	val := space.combine(e1, 0, e2, 1).([]float64)
	if val[0] != 4 {
		t.Errorf("Expected 3, got %v", val[0])
	}
	if val[1] != 2 {
		t.Errorf("Expected 3/2, got %v", val[1])
	}
}

func TestRealCombine2_1x0And4_2x0(t *testing.T) {
	e1 := []float64{2, 1}
	e2 := []float64{4, 2}
	space := realSpace{}
	var val Elemt
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, got %v", val)
		}
	}()
	val = space.combine(e1, 0, e2, 0)
}

func TestRandomInitKMeans(t *testing.T) {
	var data = make([]Elemt, 8)
	data[0] = []float64{7.2, 6, 8, 11, 10}
	data[1] = []float64{9, 8, 7, 7.5, 10}
	data[2] = []float64{7.2, 6, 8, 11, 10}
	data[3] = []float64{-9, -10, -8, -8, -7.5}
	data[4] = []float64{-8, -10.5, -7, -8.5, -9}
	data[5] = []float64{42, 41.2, 42, 40.2, 45}
	data[6] = []float64{42, 41.2, 42.2, 40.2, 45}
	data[7] = []float64{50, 51.2, 49, 40, 45.2}
	var space = realSpace{}
	var km = NewKMeans(3, 10, data, space, randomInit)
	km.run()
	km.close()
	var clusters = km.clustering.clusters
	if len(clusters) != 3 {
		t.Errorf("Expected 3, got %v", 3)
	}
}

func TestDeterminedInitKMeans(t *testing.T) {
	var data = make([]Elemt, 8)
	data[0] = []float64{7.2, 6, 8, 11, 10}
	data[1] = []float64{9, 8, 7, 7.5, 10}
	data[2] = []float64{7.2, 6, 8, 11, 10}
	data[3] = []float64{-9, -10, -8, -8, -7.5}
	data[4] = []float64{-8, -10.5, -7, -8.5, -9}
	data[5] = []float64{42, 41.2, 42, 40.2, 45}
	data[6] = []float64{42, 41.2, 42.2, 40.2, 45}
	data[7] = []float64{50, 51.2, 49, 40, 45.2}
	var localSpace = realSpace{}
	var init = func(k int, elemts []Elemt, space space) Clustering {
		var centroids = make([]Elemt, 3)
		var clusters = make([][]Elemt, 3)
		centroids[0] = []float64{7.2, 6, 8, 11, 10}
		centroids[1] = []float64{-9, -10, -8, -8, -7.5}
		centroids[2] = []float64{42, 41.2, 42.2, 40.2, 45}
		for _, elemt := range elemts {
			var idx = assign(elemt, centroids, space)
			clusters[idx] = append(clusters[idx], elemt)
		}
		var c, _ = NewClustering(centroids, clusters)
		return c
	}
	var km = NewKMeans(3, 10, data, localSpace, init)
	km.run()
	km.close()
	var clusters = km.clustering.clusters
	if len(clusters[0]) != 3 {
		t.Errorf("Expected 3, got %v", len(clusters[0]))
	}
	if len(clusters[1]) != 2 {
		t.Errorf("Expected 2, got %v", len(clusters[1]))
	}
	if len(clusters[2]) != 3 {
		t.Errorf("Expected 3, got %v", len(clusters[2]))
	}
}

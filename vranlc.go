package main

import (
	"fmt"
	"math"
)

const r23 float64 = (0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5)
const r46 float64 = (r23 * r23)
const t23 float64 = (2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0)
const t46 float64 = (t23 * t23)

const M = 24 //classe S

const MK = 16
const MM = (M - MK)
const NN = (1 << MM)
const NK = (1 << MK)
const NQ = 10
const EPSILON = 1.0e-8
const A float64 = 1220703125.0
const S = 271828183.0
const NK_PLUS = ((2 * NK) + 1)

var x = make([]float64, NK_PLUS) // x[NK_PLUS]
var q = make([]float64, NQ)      // x[NK_PLUS]

func vranlc(n int, x_seed *float64, a float64, y []float64) {
	var i int
	var x, t1, t2, t3, t4, a1, a2, x1, x2, z float64

	t1 = r23 * a
	a1 = float64(int(t1))
	a2 = a - t23*a1
	x = *x_seed

	for i = 0; i < n; i++ {
		t1 = r23 * x
		x1 = float64(int(t1))
		x2 = x - t23*x1
		t1 = a1*x2 + a2*x1
		t2 = float64(int(r23 * t1))
		z = t1 - t23*t2
		t3 = t23*z + a2*x2
		t4 = float64(int(r46 * t3))
		x = t3 - t46*t4
		y[i] = r46 * x
	}
	*x_seed = x
}

func randlc(x *float64, a float64) float64 {
	var t1, t2, t3, t4, a1, a2, x1, x2, z float64

	t1 = r23 * a
	a1 = float64(int(t1))
	a2 = a - t23*a1

	t1 = r23 * (*x)
	x1 = float64(int(t1))
	x2 = (*x) - t23*x1
	t1 = a1*x2 + a2*x1
	t2 = float64(int(r23 * t1))
	z = t1 - t23*t2
	t3 = t23*z + a2*x2
	t4 = float64(int(r46 * t3))
	(*x) = t3 - t46*t4

	return (r46 * (*x))
}

func main() {

	dum := []float64{1.0, 1.0}
	dum2 := []float64{1.0}
	vranlc(0, &dum[0], dum[1], dum2)

	dum[0] = randlc(&dum[1], dum2[0])

	for i := 0; i < NK_PLUS; i++ {
		x[i] = -1.0e99
	}

	Mops := math.Log(math.Sqrt(math.Abs(math.Max(1.0, 1.0))))
	fmt.Printf("%f Mops", Mops)

	t1 := A
	var t2 float64 = 0
	for i := 0; i < MK+1; i++ {
		t2 = randlc(&t1, t1)
		fmt.Printf("t2: %f na i %d \n", t2, i)
	}
	an := t1
	tt := S
	var gc float64 = 0.0
	var sx float64 = 0.0
	var sy float64 = 0.0

	for i := 0; i <= NQ-1; i++ {
		q[i] = 0.0 //acho q n precisa pq ele já inicializa com zero
	}

	k_offset := -1

	//fmt.Printf("t2: %f", t2)

	//fmt.Printf("dum0 %f dum1 %f dum2 %f", dum[0], dum[1], dum2)
}

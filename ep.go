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
	var aux float64
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
	aux = t3 - t46*t4
	(*x) = aux
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

	//Mops := math.Log(math.Sqrt(math.Abs(math.Max(1.0, 1.0)))) só serve pra timer

	t1 := A
	var t2 float64 = 0
	for i := 0; i < MK+1; i++ {
		t2 = randlc(&t1, t1)
	}
	an := t1
	//tt := S
	var gc float64 = 0.0
	var sx float64 = 0.0
	var sy float64 = 0.0

	for i := 0; i <= NQ-1; i++ {
		q[i] = 0.0 //acho q n precisa pq ele já inicializa com zero
	}

	k_offset := -1

	np := NN
	var kk int = 0
	var ik int = 0
	var t3 float64
	var t4 float64
	var x1, x2 float64
	var l float64
	//cada interação desse loop for pode ser feita independentemente
	//talvez chamar uma goroutine pra cada iteração do laço?

	//fmt.Printf("valor do np: %d \n", np)
	//fmt.Printf("valor do np: %d \n", np)

	for k := 1; k <= np; k++ {
		//implementar área paralela do EP
		kk = k_offset + k
		t1 = S
		t2 = an

		/* find starting seed t1 for this kk */
		for i := 1; i <= 100; i++ {
			ik = kk / 2
			if (2 * ik) != kk {
				t3 = randlc(&t1, t2)

			}
			if ik == 0 {
				break
			}
			t3 = randlc(&t2, t2)
			//fmt.Printf("t3 fater call break: %f, t2 after call break %f\n", t3, t2)

			kk = ik
		}
		vranlc(2*NK, &t1, A, x)

		for i := 0; i < NK; i++ {
			x1 = 2.0*x[2*i] - 1.0
			x2 = 2.0*x[2*i+1] - 1.0
			t1 = math.Pow(x1, 2) + math.Pow(x2, 2)
			if t1 <= 1.0 {
				t2 = math.Sqrt(-2.0 * math.Log(t1) / t1)
				t3 = (x1 * t2)
				t4 = (x2 * t2)
				l = math.Max(math.Abs(t3), math.Abs(t4))
				fmt.Printf("valor de l: %f \n", l)
				q[int(l)] += 1.0
				sx = sx + t3
				sy = sy + t4
			}
		}

	}

	var sx_verify_value float64
	var sy_verify_value float64
	var sx_err float64
	var sy_err float64

	for i := 0; i <= NQ-1; i++ {
		gc = gc + q[i]
	}
	fmt.Printf("valor de Q com np = 1 %f \n", gc)

	verified := true
	if M == 24 {
		sx_verify_value = -3.247834652034740e+3
		sy_verify_value = -6.958407078382297e+3
	} else if M == 25 {
		sx_verify_value = -2.863319731645753e+3
		sy_verify_value = -6.320053679109499e+3
	} else if M == 28 {
		sx_verify_value = -4.295875165629892e+3
		sy_verify_value = -1.580732573678431e+4
	} else if M == 30 {
		sx_verify_value = 4.033815542441498e+4
		sy_verify_value = -2.660669192809235e+4
	} else if M == 32 {
		sx_verify_value = 4.764367927995374e+4
		sy_verify_value = -8.084072988043731e+4
	} else if M == 36 {
		sx_verify_value = 1.982481200946593e+5
		sy_verify_value = -1.020596636361769e+5
	} else if M == 40 {
		sx_verify_value = -5.319717441530e+05
		sy_verify_value = -3.688834557731e+05
	} else {
		verified = false
	}
	if verified {
		//fmt.Printf("\n VERIFIED \n")
		sx_err = math.Abs((sx - sx_verify_value) / sx_verify_value)
		sy_err = math.Abs((sy - sy_verify_value) / sy_verify_value)
		verified = ((sx_err <= EPSILON) && (sy_err <= EPSILON))
	}
	fmt.Printf("\n VERIFIED: %t \n", verified)

	fmt.Printf("M: %d \n", M)
	fmt.Printf("pares gaussianos: %15.0f \n", gc)
	fmt.Printf("somas: %f %f \n", sx, sy)
	fmt.Printf("counts:\n")
	for i := 0; i < NQ-1; i++ {
		fmt.Printf("%3d%15.0f\n", i, q[i])
	}

	//Mops = math.Pow(2.0, M+1) / tm / 1000000.0

}

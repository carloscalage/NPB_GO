package main

import (
	"fmt"
	"sync"
)

const r23 float64 = (0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5)
const r46 float64 = (r23 * r23)
const t23 float64 = (2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0)
const t46 float64 = (t23 * t23)
const num_procs = 1

var NUM_KEYS int
var key_array []int
var MAX_KEY int
var MAX_KEY_LOG_2 int
var TOTAL_KEYS_LOG_2 int

var SIZE_OF_BUFFERS int
var NUM_BUCKETS int
var NUM_BUCKETS_LOG_2 int
var key_buff2 []int
var key_buff1 []int
var key_buff_ptr_global []int
var MAX_ITERATIONS int = 10
var TEST_ARRAY_SIZE int = 5
var partial_verify_vals []int
var test_index_array []int
var test_rank_array []int
var passed_verification = 0
var key_buff1_aptr [][]int

var bucket_ptrs [num_procs][]int

var bucket_size [][]int
var class = 'S'

func main() {

	TOTAL_KEYS_LOG_2 = 0
	MAX_KEY_LOG_2 = 0
	NUM_BUCKETS_LOG_2 = 0

	TOTAL_KEYS := 0

	fmt.Printf("IS implementation \n")

	switch class {
	case 'S':
		test_index_array = []int{48427, 17148, 23627, 62548, 4431}
		test_rank_array = []int{0, 18, 346, 64917, 65463}
		TOTAL_KEYS_LOG_2 = 16
		MAX_KEY_LOG_2 = 11
		NUM_BUCKETS_LOG_2 = 9
	case 'W':
		test_index_array = []int{357773, 934767, 875723, 898999, 404505}
		test_rank_array = []int{1249, 11698, 1039987, 1043896, 1048018}

		TOTAL_KEYS_LOG_2 = 20
		MAX_KEY_LOG_2 = 16
		NUM_BUCKETS_LOG_2 = 10
	case 'A':
		test_index_array = []int{2112377, 662041, 5336171, 3642833, 4250760}
		test_rank_array = []int{104, 17523, 123928, 8288932, 8388264}

		TOTAL_KEYS_LOG_2 = 23
		MAX_KEY_LOG_2 = 19
		NUM_BUCKETS_LOG_2 = 10

	case 'B':
		test_index_array = []int{41869, 812306, 5102857, 18232239, 26860214}
		test_rank_array = []int{33422937, 10244, 59149, 33135281, 99}
		TOTAL_KEYS_LOG_2 = 27
		MAX_KEY_LOG_2 = 23
		NUM_BUCKETS_LOG_2 = 10

		//adicionar as classes C e D depois
		//C_test_index_array := []int{44172927, 72999161, 74326391, 129606274, 21736814}
		//C_test_rank_array := []int{61147, 882988, 266290, 133997595, 133525895}

		//D_test_index_array := []int{1317351170, 995930646, 1157283250, 1503301535, 1453734525}
		//D_test_rank_array := []int{1, 36538729, 1978098519, 2145192618, 2147425337}

	}
	fmt.Printf("%d \n", TOTAL_KEYS_LOG_2)
	MAX_KEY = 1 << MAX_KEY_LOG_2
	NUM_BUCKETS = 1 << NUM_BUCKETS_LOG_2

	TOTAL_KEYS = 1 << TOTAL_KEYS_LOG_2
	fmt.Printf("TOTALKEYS: %d \n", TOTAL_KEYS)

	NUM_KEYS = TOTAL_KEYS
	SIZE_OF_BUFFERS = NUM_KEYS
	key_array = make([]int, SIZE_OF_BUFFERS)
	key_buff2 = make([]int, SIZE_OF_BUFFERS)
	key_buff1 = make([]int, MAX_KEY)
	partial_verify_vals = make([]int, TEST_ARRAY_SIZE)

	use(key_array, MAX_KEY, MAX_KEY_LOG_2, SIZE_OF_BUFFERS, TEST_ARRAY_SIZE, MAX_ITERATIONS, NUM_BUCKETS_LOG_2, NUM_BUCKETS, test_index_array, test_rank_array)
	create_seq(314159265.00, 1220703125.00)

	//fmt.Printf("%v \n", key_array[500:550])

	alloc_key_buff()

	/* Do one interation for free (i.e., untimed) to guarantee initialization of */
	/* all data and code pages and respective tables */
	rank(1)

	passed_verification = 0
	if class != 'S' {
		fmt.Print("\n iteration \n")
	}

	/*this is the main iteration */
	for iteration := 1; iteration <= MAX_ITERATIONS; iteration++ {
		if class != 'S' {
			fmt.Printf("    %d\n", iteration)
		}
		rank(iteration)
	}
	fmt.Printf("\nsplitting\n")
	//fmt.Printf("%v \n", key_array[500:550])
	full_verify()
	fmt.Printf("\nsplitting\n")
	//fmt.Printf("%v \n", key_array[49200:49600])
	/* the final printout */
	if passed_verification != 5*MAX_ITERATIONS+1 {
		passed_verification = 0
	}
	fmt.Printf("passed verification: %d \n", passed_verification)

}

func use(vals ...interface{}) { //só pra matar warning de variavel não usada durante desenvolvimento
	//remover na hora de mandar
	for _, val := range vals {
		_ = val
	}
}

func create_seq(seed, a float64) {
	var wg sync.WaitGroup
	wg.Add(num_procs)
	for i := 0; i < num_procs; i++ {
		go func(myid int) {
			var x, s float64
			var k int

			var k1, k2 int
			an := a

			mq := (NUM_KEYS + num_procs - 1) / num_procs
			k1 = mq * myid
			k2 = k1 + mq
			if k2 > NUM_KEYS {
				k2 = NUM_KEYS
			}

			s = find_my_seed(myid,
				num_procs,
				int64(4*NUM_KEYS),
				seed,
				an)

			k = MAX_KEY / 4

			for j := k1; j < k2; j++ {
				x = randlc(&s, an)
				x += randlc(&s, an)
				x += randlc(&s, an)
				x += randlc(&s, an)
				key_array[j] = int(float64(k) * x)
			}

			defer wg.Done()
		}(i)
	}
	wg.Wait()
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

func find_my_seed(kn int, np int, nn int64, s float64, a float64) float64 {
	if kn == 0 {
		return s
	}

	var mq int64 = (nn/4 + int64(np) - 1) / int64(np)
	var kk int64 = mq * 4 * int64(kn)

	t1 := s
	t2 := a
	var ik int64

	for kk > 1 {
		ik = kk / 2
		if 2*ik == kk {
			randlc(&t2, t2)
			kk = ik
		} else {
			randlc(&t1, t2)
			kk = kk - 1
		}
	}
	randlc(&t1, t2)

	return t1
}

func alloc_key_buff() {

	key_buff1_aptr = make([][]int, num_procs)

	key_buff1_aptr[0] = key_buff1
	for i := 1; i < num_procs; i++ {
		key_buff1_aptr[i] = make([]int, MAX_KEY)
	}
	//também to usando pra alocar o bucket_ptrs
}

func rank(iteration int) {

	key_array[iteration] = iteration
	key_array[iteration+MAX_ITERATIONS] = MAX_KEY - iteration

	/* Determine where the partial verify test keys are, load into */
	/* top of array bucket_size */
	for i := 0; i < TEST_ARRAY_SIZE; i++ {
		partial_verify_vals[i] = key_array[test_index_array[i]]
	}

	key_buff_ptr2 := key_array
	key_buff_ptr := key_buff1

	//var m, k1, k2 int
	myid := 0

	work_buff := &key_buff1_aptr[myid]

	/* Clear the work array */
	for i := 0; i < MAX_KEY; i++ {
		(*work_buff)[i] = 0
	}

	/* Ranking of all keys occurs in this section: */
	/* In this section, the keys themselves are used as their */
	/* own indexes to determine how many of each there are: their */
	/* individual population */
	for i := 0; i < NUM_KEYS; i++ {
		(*work_buff)[(key_buff_ptr2)[i]]++ /* Now they have individual key population */
	}
	for i := 0; i < MAX_KEY-1; i++ {
		(*work_buff)[i+1] += (*work_buff)[i]
	}
	/* Accumulate the global key population */
	for k := 1; k < num_procs; k++ {
		for i := 0; i < MAX_KEY; i++ {
			(key_buff_ptr)[i] += key_buff1_aptr[k][i]
		}
	}

	/* This is the partial verify test section */
	/* Observe that test_rank_array vals are */
	/* shifted differently for different cases */
	for i := 0; i < TEST_ARRAY_SIZE; i++ {
		k := partial_verify_vals[i]
		if 0 < k && k <= NUM_KEYS-1 {
			key_rank := (key_buff_ptr)[k-1]
			failed := 0

			switch class {
			case 'S':
				if i <= 2 {
					if key_rank != test_rank_array[i]+iteration {
						//fmt.Printf("key_rank %d different of test_rank_array[i]+iter %d at i %d \n", key_rank, test_rank_array[i]+iteration, i)
						failed = 1
					} else {
						passed_verification++
					}
				} else {
					if key_rank != test_rank_array[i]-iteration {
						//fmt.Printf("key_rank %d different of test_rank_array[i]-iter %d at i %d \n", key_rank, test_rank_array[i]-iteration, i)
						failed = 1
					} else {
						passed_verification++
					}
				}
				break
			case 'W':
				if i < 2 {
					if key_rank != test_rank_array[i]+(iteration-2) {
						failed = 1
					} else {
						passed_verification++
					}
				} else {
					if key_rank != test_rank_array[i]-iteration {
						failed = 1
					} else {
						passed_verification++
					}
				}
				break
			case 'A':
				if i <= 2 {
					if key_rank != test_rank_array[i]+(iteration-1) {
						failed = 1
					} else {
						passed_verification++
					}
				} else {
					if key_rank != test_rank_array[i]-(iteration-1) {
						failed = 1
					} else {
						passed_verification++
					}
				}
				break
			case 'B':
				if i == 1 || i == 2 || i == 4 {
					if key_rank != test_rank_array[i]+iteration {
						failed = 1
					} else {
						passed_verification++
					}
				} else {
					if key_rank != test_rank_array[i]-iteration {
						failed = 1
					} else {
						passed_verification++
					}
				}
				break
			case 'C':
				if i <= 2 {
					if key_rank != test_rank_array[i]+iteration {
						failed = 1
					} else {
						passed_verification++
					}
				} else {
					if key_rank != test_rank_array[i]-iteration {
						failed = 1
					} else {
						passed_verification++
					}
				}
				break
			case 'D':
				if i < 2 {
					if key_rank != test_rank_array[i]+iteration {
						failed = 1
					} else {
						passed_verification++
					}
				} else {
					if key_rank != test_rank_array[i]-iteration {
						failed = 1
					} else {
						passed_verification++
					}
				}
				break
			}
			if failed == 1 {
				fmt.Printf("Failed partial verification: iteration %d, test key %d\n", iteration, i)
			}
		}
	}
	if iteration == MAX_ITERATIONS {
		key_buff_ptr_global = key_buff_ptr
	}

}

func full_verify() {
	var i, j int
	var k, k1, k2 int

	myid := 0

	for i := 0; i < NUM_KEYS; i++ {
		//fmt.Printf("key array is %d at i: %d \n", key_array[i], i)
		key_buff2[i] = key_array[i]
	}

	j = 1
	j = (MAX_KEY + j - 1) / 2
	k1 = j * myid
	k2 = k1 + j
	if k2 > MAX_KEY {
		k2 = MAX_KEY
	}
	for i := 0; i < NUM_KEYS; i++ {
		if key_buff2[i] >= k1 && key_buff2[i] < k2 {
			//fmt.Printf("key buff ptr global %d   ", key_buff_ptr_global[key_buff2[i]])
			key_buff_ptr_global[key_buff2[i]]--
			k = key_buff_ptr_global[key_buff2[i]]
			//fmt.Printf("k %d   \n", k)

			key_array[k] = key_buff2[i]
		}
	}
	j = 0
	var indexoutoforder []int

	//#pragma omp parallel for reduction(+:j)

	fmt.Printf("%v \n", key_array[65530:])
	for i = 1; i < NUM_KEYS; i++ {
		if key_array[i-1] > key_array[i] {
			//fmt.Printf("pos: %d, bigger value: %d, smaller value: %d \n", i, key_array[i-1], key_array[i])
			indexoutoforder = append(indexoutoforder, i)
			j++
		}
	}

	//fmt.Printf("%d is bigger than %d at pos %d \n", key_array[33065], key_array[33065+1], 33065)

	if j != 0 {
		fmt.Printf("amount of out of order: %d \n", len(indexoutoforder))
		fmt.Printf("-- %v -- \n", indexoutoforder[:100])
		fmt.Printf("\nFull_verify: number of keys out of sort: %d\n", j)
	} else {
		passed_verification++
	}
}

package main

import (
	"fmt"
)

func main() {
	class := 'S'

	TOTAL_KEYS_LOG_2 := 0
	MAX_KEY_LOG_2 := 0
	NUM_BUCKETS_LOG_2 := 0

	TOTAL_KEYS := 0
	MAX_KEY := 1 << TOTAL_KEYS_LOG_2
	NUM_BUCKETS := 1 << NUM_BUCKETS_LOG_2

	if class != 'D' {
		TOTAL_KEYS = 1 << TOTAL_KEYS_LOG_2
	}
	NUM_KEYS := TOTAL_KEYS
	SIZE_OF_BUFFERS := NUM_KEYS

	MAX_ITERATIONS := 10
	TEST_ARRAY_SIZE := 5
	fmt.Printf("IS implementation")
	var test_index_array []int
	var test_rank_array []int

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

	use(MAX_KEY, MAX_KEY_LOG_2, SIZE_OF_BUFFERS, TEST_ARRAY_SIZE, MAX_ITERATIONS, NUM_BUCKETS_LOG_2, NUM_BUCKETS, test_index_array, test_rank_array)

}

func use(vals ...interface{}) { //só pra matar warning de variavel não usada durante desenvolvimento
	//remover na hora de mandar
	for _, val := range vals {
		_ = val
	}
}

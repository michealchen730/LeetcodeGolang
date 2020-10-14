package main

func maxSatisfied(customers []int, grumpy []int, X int) int {
	sum := 0

	for k, v := range customers {
		if grumpy[k] == 0 {
			sum += v
		}
	}

	sat := 0
	for i := 0; i < X; i++ {
		if grumpy[i] == 1 {
			sat += customers[i]
		}
	}

	sum += sat
	res := sum

	for k := X; k < len(customers); k++ {
		if grumpy[k] == 1 {
			sum += customers[k]
		}
		if grumpy[k-X] == 1 {
			sum -= customers[k-X]
		}
		res = max(sum, res)
	}

	return res
}

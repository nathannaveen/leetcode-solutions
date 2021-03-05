package main

func averageWaitingTime(customers [][]int) float64 {
	time := 0
	sum := 0
	for _, customer := range customers {
		subtract := 0
		if time-customer[0] > 0 {
			subtract = time - customer[0]
		} else {
			time = customer[0]
		}
		time += customer[1]
		sum += customer[1] + subtract
	}
	return float64(sum) / float64(len(customers))
}

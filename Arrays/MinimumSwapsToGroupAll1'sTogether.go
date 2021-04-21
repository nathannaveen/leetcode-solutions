package main

func main() {

}

func minSwaps(data []int) int {
	oneCounter, maximum, counter := 0, 0, 0

	for _, i := range data {
		oneCounter += i
	}

	for i := 0; i < oneCounter; i++ {
		counter += data[i]
	}
	maximum = max(maximum, counter)

	for i := oneCounter; i < len(data); i++ {
		counter += data[i] - data[i-oneCounter]
		maximum = max(maximum, counter)
	}

	return oneCounter - maximum
}

// 97%, 98%

package main

func main() {

}

type ParkingSystem struct {
	cars []int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	c := ParkingSystem{
		[]int{big, medium, small},
	}
	return c
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.cars[carType-1] >= 1 {
		this.cars[carType-1]--
		return true
	}
	return false
}

package main

func main() {

}

func badSensor(sensor1 []int, sensor2 []int) int {
	found := false

	for i := 0; i < len(sensor1)-1; i++ {
		if !found && sensor1[i] != sensor2[i] {
			found = true
		} else if found && sensor1[i] == sensor2[i-1] && sensor1[i] == sensor2[i+1] {
			continue
		} else if found && sensor1[i+1] == sensor2[i] && sensor1[i-1] == sensor2[i] {
			continue
		} else if found && sensor1[i] == sensor2[i-1] {
			return 2
		} else if found && sensor1[i-1] == sensor2[i] {
			return 1
		}
	}

	return -1
}

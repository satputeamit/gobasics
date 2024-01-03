package main

func main() {
	for i := 1; i <= 10; i++ {
		println(i)
	}
	// or like while

	j := 1
	for j <= 10 {
		println(j)
		j++
	}

	k := 1

	for true {
		if k == 20 {
			break
		}
		println(k)
		k++
	}

	//
	cities := []string{"sangli", "miraj", "pune", "mumbai"}

	for i, city := range cities {
		println(i, city)
	}

}

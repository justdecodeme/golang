package main

func main() {
	for i := 0; i < 5; i++ {
		println(i)
	}
	println(".................")

	j := 0
	for {
		j++
		println(j)
		if j > 5 {
			break
		}
	}
	println(".................")

	s := []string{"foo", "bar", "buz"}
	for idx, v := range s {
		println(idx, v)
	}
	println(".................")

	m := make(map[string]string)
	m["a"] = "foo"
	m["b"] = "bar"
	m["c"] = "buz"
	for k, v := range m {
		println(k, v)
	}
}

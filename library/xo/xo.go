package xo

func CountXO(str string) bool {
	counter := 0
	for _, s := range str {
		if string(s) == "x" {
			counter++
		} else if string(s) == "o" {
			counter--
		}
	}
	if counter == 0 {
		return true
	}
	return false
}

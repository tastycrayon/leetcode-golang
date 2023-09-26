package main

func sum(s string, ch chan<- byte) {
	var temp byte // byte is uint8
	for i := range s {
		temp += s[i]
	}
	ch <- temp
}
func findTheDifference(s string, t string) byte {
	ch := make(chan byte, 2)
	go sum(s, ch)
	go sum(t, ch)

	a := <-ch
	b := <-ch

	close(ch)

	return a - b
}

func findTheDifference1(s string, t string) byte {
	var result byte = 0
	for i := range s {
		result = result ^ s[i]
	}
	for i := range t {
		result = result ^ t[i]
	}

	return result
}

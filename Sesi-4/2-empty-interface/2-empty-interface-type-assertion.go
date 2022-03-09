package main

func main() {
	var v interface{}

	v = 20
	// v = v * 9 Error karna v bukan tipe data int konkrit(asli)

	if value, ok := v.(int); ok == true {
		v = value * 9
	}
}

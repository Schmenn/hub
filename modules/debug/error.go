package debug

// Check handles errors
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

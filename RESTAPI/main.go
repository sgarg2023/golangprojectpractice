package main

func main() {
	addr := ":3000"
	s1 := NewAPIServer(addr)
	s2 := NewAPIServer(":5000")
	s1.Run()
	s2.Run()
}

package demo

import "fmt"

func init() {
	Test()
}

//Test 测试demo
func Test() {
	println("start main")
	ch := make(chan bool)
	go func() {
		println("come into goroutine")
		ch <- true
	}()

	println("do something else")
	s, o := <-ch
	fmt.Println("s, o", s, o)
	close(ch)
	v, ok := <-ch
	fmt.Println("v, ok", v, ok)
	println("end main")
}

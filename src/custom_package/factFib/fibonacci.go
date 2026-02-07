package factFib
import . "fmt"
func Fibonacci(num int) {
	f1 := 0 
	f2 := 1
	nextnum := 0
	for i := 1; i < num; i++{
		Println(f1)
		nextnum = f1 + f2
		f1 = f2
		f2 = nextnum
	}
}

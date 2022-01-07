// sieve.go
// 生成素数的有效算法
// see sieve_test.go
package prime

func GenerateChannel(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

//  从通道中筛选出不是质数的数字 - 基本上将它们从通道中移除
func Sieve(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func Generate(limit int) []int {
	var primes []int

	// 非缓冲通道
	ch := make(chan int)

	go GenerateChannel(ch)

	for i := 0; i < limit; i++ {
		primes = append(primes, <-ch)
		ch1 := make(chan int)
		go Sieve(ch, ch1, primes[i])
		ch = ch1
	}
	return primes
}

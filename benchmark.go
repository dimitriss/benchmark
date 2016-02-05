package main

import ("fmt"
	"testing")

func isPrime (n int32) bool {
    if n <= 2 {
        return n == 2
    } else if n % 2 == 0 {
        return false
    }
    var i int32 = 3
    for i * i <= n {
        if n % i == 0 {
            return false
        }
        i += 2
    } 
    return true
}

func BenchmarkFunction(b *testing.B) {
    var i, limit int32 =  0, 67108864
    var primeCount int32 = 0;
    for i <= limit {
        if isPrime(i) {
            primeCount++
        }
        i ++
    }
    fmt.Println(primeCount)
}


 
func main() {
    br := testing.Benchmark(BenchmarkFunction)
    fmt.Println(br)
}

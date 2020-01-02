package bitmap

import(
	"testing"
	"fmt"
	"time"
)

func TestConcurrency(t *testing.T) {

	bitmap := New(1000)
	var n uint64
	for n = 0; n < 100000; n++ {
		var i uint64 = 1000-(n%1000)
		var j uint64 = 1000-(n%1000)
		var k uint64 = 1000-(n%1000)		
		go func() {
			for ; i < 1000; i ++ {
				bitmap.Set(i)
				fmt.Println("set",i)
			}
		}()

		go func() {
			for; j < 1000; j ++ {
				bitmap.Get(j)
				fmt.Println("get",j, bitmap.Get(j))
			}
		}()

		go func() {
			for; k < 1000; k++ {
				bitmap.Clear(k)
				fmt.Println("Clear",k)
			}
		}()
		
	}

	time.Sleep(5 * time.Second)
}

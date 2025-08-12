package bakery

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// N - number of bakers, M - number of packers, K num of cakes
func Bake(ctx context.Context, t1, N, K int) <-chan Cake {
	bakedChn := make(chan Cake, 256)

	if K <= 0 || N <= 0 {
		close(bakedChn)
		return bakedChn
	}

	var wg sync.WaitGroup
	wg.Add(N)

	var produced int32 = 0

	for i := 1; i <= N; i++ {
		bakerId := i
		go func(id int) {
			defer wg.Done()
			r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(id)))

			for {
				select {
				case <-ctx.Done():
					return
				default:
				}

				next := int(atomic.AddInt32(&produced, 1))
				if next > K {
					return
				}

				jitter := 0
				if t1 > 0 {
					jitter = r.Intn(2*t1+1) - t1
				}
				T1 := id + jitter
				if T1 < 0 {
					T1 = 0
				}

				select {
				case <-ctx.Done():
					return
				case <-time.After(time.Duration(T1) * time.Millisecond):
				}

				cake := Cake{
					BakedBy:  id,
					BakeTime: T1,
				}

				select {
				case <-ctx.Done():
					return
				case bakedChn <- cake:
				}
			}

		}(bakerId)
	}

	go func() {
		wg.Wait()
		close(bakedChn)
	}()

	return bakedChn
}

func Pack(ctx context.Context, bakedChn <-chan Cake, M, t2 int) <-chan Cake {
	packedChn := make(chan Cake, 256)

	if M <= 0 {
		close(packedChn)
		return packedChn
	}

	var wg sync.WaitGroup
	wg.Add(M)

	for i := 1; i <= M; i++ {
		workerId := i

		go func(id int) {
			defer wg.Done()
			r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(id)))

			for {
				select {
				case <-ctx.Done():
					return
				case cake, ok := <-bakedChn:
					if !ok {
						return
					}

					jitter := 0
					if t2 > 0 {
						jitter = r.Intn(2*t2+1) - t2
					}
					T2 := id + jitter
					if T2 < 0 {
						T2 = 0
					}

					select {
					case <-ctx.Done():
						return
					case <-time.After(time.Duration(T2) * time.Millisecond):
					}

					cake.PackedBy = id
					cake.PackTime = T2

					select {
					case <-ctx.Done():
						return
					case packedChn <- cake:
					}
				}
			}

		}(workerId)
	}
	go func() {
		wg.Wait()
		close(packedChn)
	}()

	return packedChn
}

func Inspect(packedChn <-chan Cake) int {
	cakes := make([]Cake, 0)

	for c := range packedChn {
		cakes = append(cakes, c)
	}

	for _, c := range cakes {
		fmt.Printf("BakedBy=%d\t BakeTime=%d\t PackedBy=%d\t PackTime=%d\t\n",
			c.BakedBy, c.BakeTime, c.PackedBy, c.PackTime)
	}

	return len(cakes)
}

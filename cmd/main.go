package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/zanpatryk/gocourse/bakery"
)

func main() {
	var (
		K  = flag.Int("K", 100, "total number of cakes to produce")
		N  = flag.Int("N", 8, "number of bakers")
		M  = flag.Int("M", 5, "number of packers")
		t1 = flag.Int("t1", 5, "baker jitter parameter (milliseconds)")
		t2 = flag.Int("t2", 10, "packer jitter parameter (milliseconds)")
	)
	flag.Parse()

	fmt.Printf("Starting bakery: K=%d, N=%d, M=%d, t1=%d, t2=%d\n", *K, *N, *M, *t1, *t2)

	bakedCh := bakery.Bake(context.Background(), *t1, *N, *K)
	packedCh := bakery.Pack(context.Background(), bakedCh, *M, *t2)

	start := time.Now()

	cakes := bakery.Inspect(packedCh)

	elapsed := time.Since(start)
	fmt.Printf("\nDone. Collected %d cakes. Elapsed: %s\n", cakes, elapsed)
}

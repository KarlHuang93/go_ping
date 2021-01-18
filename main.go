package main

import (
	"context"
	"fmt"
	"go_ping/ping"
	"golang.org/x/net/icmp"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
	ch = make(chan int, 5)
)

func main() {

	ipSlice := []string{}
	ipSlice = append(ipSlice, "157.240.2.36")
	ipSlice = append(ipSlice, "wwww.baidu.com")
	ipSlice = append(ipSlice, "www.github.com")
	ipSlice = append(ipSlice, "221.122.82.30")
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go work(ctx, cancel, ipSlice)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	close(ch)
	wg.Wait()
}
func work(ctx context.Context, cancel context.CancelFunc, ipSlice []string) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("ip is close")
				return
			default:
				pinger, err := ping.NewBatchPinger(ipSlice, 100, time.Second*1, time.Second*10)
				if err != nil {
					panic(err)
				}
				pinger.OnRecv = func(pkt *icmp.Echo) {
				}

				pinger.OnFinish = func(stSlice []*ping.Statistics) {
					for i, st := range stSlice {
						ch <- i
						fmt.Printf("\n--- %s ping statistics ---\n", st.Addr)
						fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
							st.PacketsSent, st.PacketsRecv, st.PacketLoss)
						fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
							st.MinRtt, st.AvgRtt, st.MaxRtt, st.StdDevRtt)
					}

				}
				pinger.Run()
				fmt.Println()
			}
			cancel()

		}
	}()

}

package modules

import (
	"fmt"
	"github.com/3th1nk/cidr"
	"github.com/chelnak/ysmrr"
	probing "github.com/prometheus-community/pro-bing"
	"sync"
	"time"
)

var wg sync.WaitGroup
var ips []string
var running = false

func ping(host string) {
	defer wg.Done()

	pinger, err := probing.NewPinger(host)
	pinger.SetPrivileged(true)
	if err != nil {
		panic(err)
	}
	pinger.Count = 5
	err = pinger.Run()
	if err != nil {

	}
	stats := pinger.Statistics()

	if stats.PacketsRecv > 0 {
		ips = append(ips, host)
	}
}

func Scan(ranges []string) {
	fmt.Println("")

	sm := ysmrr.NewSpinnerManager()
	mySpinner := sm.AddSpinner("Scanning...")
	sm.Start()

	startTime := time.Now()

	targets := ranges
	for x := range targets {
		c, _ := cidr.Parse(targets[x])
		c.Each(func(ip string) bool {
			wg.Add(1)
			go ping(ip)
			return true
		})
	}
	wg.Wait()

	elapsedTime := time.Since(startTime)
	mySpinner.Complete()
	mySpinner.UpdateMessage("Scan complete!")
	sm.Stop()

	for x := range targets {
		fmt.Printf("\n[%s]\n", targets[x])
		c, _ := cidr.Parse(targets[x])
		c.Each(func(ip string) bool {
			for x := range ips {
				if ip == ips[x] {
					fmt.Println("-", ip)
					break
				}
			}
			return true
		})
	}

	fmt.Printf("\nTotal hosts:  %d\n", len(ips))
	fmt.Printf("Time elapsed: %s\n\n", elapsedTime)

	ips = []string{}
}

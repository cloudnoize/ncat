package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	fnet := flag.String("network", "tcp", "udp/tcp")
	addr := flag.String("addr", "127.0.0.1:8080", "address in format of ip:port")
	to := flag.Int("Connect timeout ms", 1000, "Timeout for connect")
	toWrite := flag.Int("Write timeout ms", 1000, "Timeout for write")

	flag.Parse()

	snet := strings.ToLower(*fnet)

	if snet != "tcp" && snet != "udp" {
		println("invalid network ", snet)
		return
	}

	if *to <= 0 {
		println("invalid timeout ", *to)
		return
	}
	if *toWrite <= 0 {
		println("invalid write timeout ", *toWrite)
		return
	}

	_, _, err := net.SplitHostPort(*addr)
	if err != nil {
		println("invalid address ", *addr, " error", err.Error())
		return
	}
	now := time.Now()
	_, err = net.DialTimeout(snet, *addr, time.Duration(*to)*time.Millisecond)
	if err != nil {
		println("Connection error ", err.Error())
		return
	}
	el := time.Since(now)
	fmt.Printf("Connect duration %v ms\n", el.Seconds()*1000.0)
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)
	// var myWaitGroup sync.WaitGroup
	// myWaitGroup.Add(2)
	// go func() {
	// 	reader := bufio.NewReader(os.Stdin)
	// 	for {
	// 		select {
	// 		case <-c:
	// 			myWaitGroup.Done()
	// 			break
	// 		default:
	// 			println("enter:")
	// 			text, _ := reader.ReadString('\n')
	// 			conn.SetWriteDeadline(time.Now().Add(time.Duration(*toWrite) * time.Millisecond))
	// 			conn.Write([]byte(text))
	// 		}
	// 	}
	// }()

	// go func() {
	// 	var b [1024]byte
	// 	for {
	// 		select {
	// 		case <-c:
	// 			myWaitGroup.Done()
	// 			break
	// 		default:
	// 			n, err := conn.Read(b[:])
	// 			if err != nil {
	// 				println("Connection error ", err.Error())
	// 				return
	// 			}
	// 			print(string(b[:n]))
	// 		}
	// 	}
	// }()

	// myWaitGroup.Wait()

}

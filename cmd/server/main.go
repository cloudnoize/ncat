package main

import "net"

func main() {
	l, err := net.Listen("tcp", ":5432")
	if err != nil {
		println(err.Error())
		return
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			println(err.Error())
			continue
		}
		go func() {
			var b [1024]byte
			for {
				n, err := conn.Read(b[:])
				if err != nil {
					println(err.Error())
					return
				}
				println(b[:n])
				_, err = conn.Write(b[:n])
				if err != nil {
					println(err.Error())
					return
				}
			}

		}()
	}
}

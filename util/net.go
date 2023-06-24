package util

import (
	"io"
	"log"
	"net"
	"time"
)


func Relay(remoteConn net.Conn, conn net.Conn) {
	stop := make(chan bool)
	go func() {
		_, err := io.Copy(remoteConn, conn)
		if err != nil {
			log.Println(err)
		}
		remoteConn.SetReadDeadline(time.Now())

		stop <- true
	}()

	_, err := io.Copy(conn, remoteConn)
	if err != nil {
		log.Println(err)
	}
	conn.SetReadDeadline(time.Now())

	<-stop
}
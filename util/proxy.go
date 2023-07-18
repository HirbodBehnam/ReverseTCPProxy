package util

import (
	"errors"
	"io"
	"net"
	"sync"
)

// ProxyConnection will proxy two connections
func ProxyConnection(a1 net.Conn, a2 net.Conn) error {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	var err1, err2 error
	go func() {
		_, err1 = io.Copy(a1, a2)
		wg.Done()
	}()
	go func() {
		_, err2 = io.Copy(a2, a1)
		wg.Done()
	}()
	wg.Wait()
	return errors.Join(err1, err2)
}

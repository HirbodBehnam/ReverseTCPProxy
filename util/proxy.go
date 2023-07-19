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
	// This might look very sus but it's not! On FIN, only one side gets closed
	wg.Add(1)
	// We also need to ensure that one wg.Done gets called
	wgDoneCaller := new(sync.Once)
	var err1, err2 error
	go func() {
		_, err1 = io.Copy(a1, a2)
		wgDoneCaller.Do(func() {
			wg.Done()
		})
	}()
	go func() {
		_, err2 = io.Copy(a2, a1)
		wgDoneCaller.Do(func() {
			wg.Done()
		})
	}()
	wg.Wait()
	return errors.Join(err1, err2)
}

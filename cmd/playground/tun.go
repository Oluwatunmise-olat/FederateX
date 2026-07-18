package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

const (
	// Linux ioctl command to configure a TUN device
	tunSetIff = 0x400454ca

	iffTun  = 0x0001 // Create a TUN device
	iffNoPi = 0x1000 // Don't add packet info header just give raw IP packets
)

type interfaceRequest struct {
	name  [16]byte
	flags uint16
	_     [22]byte
}

func CreateTUN(name string) (*os.File, error) {
	tun, err := os.OpenFile("/dev/net/tun", os.O_RDWR, 0)

	if err != nil {
		return nil, fmt.Errorf("open /dev/net/tun: %w", err)
	}

	var req interfaceRequest

	copy(req.name[:], name)

	req.flags = iffTun | iffNoPi

	_, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		tun.Fd(),
		uintptr(tunSetIff),
		uintptr(unsafe.Pointer(&req)),
	)

	if errno != 0 {
		tun.Close()
		
		return nil, fmt.Errorf("ioctl TUNSETIFF: %v", errno)
	}

	return tun, nil
}

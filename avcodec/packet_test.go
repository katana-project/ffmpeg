package avcodec

import "testing"

func TestNewPacket(t *testing.T) {
	p := &Packet{}
	p.Alloc()

	defer p.Free()
}

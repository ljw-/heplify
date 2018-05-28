package decoder

import (
	"testing"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/negbie/heplify/config"
)

var rawPacket = []byte{0x0, 0xa, 0xa0, 0x0, 0xbe, 0xa8, 0x0, 0x26, 0x52, 0xe, 0xd3, 0x41, 0x8, 0x0, 0x45, 0x0, 0x2, 0xbd, 0xa1, 0xc3, 0x0, 0x0, 0x3e, 0x11, 0x69, 0x26, 0xc0, 0xa8, 0xf7, 0xfa, 0xc0, 0xa8, 0xf5, 0xfa, 0x13, 0xc4, 0x13, 0xc4, 0x2, 0xa9, 0x0, 0x0, 0x53, 0x49, 0x50, 0x2f, 0x32, 0x2e, 0x30, 0x20, 0x32, 0x30, 0x30, 0x20, 0x4f, 0x4b, 0xd, 0xa, 0x43, 0x61, 0x6c, 0x6c, 0x2d, 0x49, 0x44, 0x3a, 0x20, 0x42, 0x43, 0x30, 0x39, 0x39, 0x38, 0x38, 0x34, 0x40, 0x36, 0x64, 0x66, 0x63, 0x66, 0x66, 0x65, 0x38, 0xd, 0xa, 0x43, 0x53, 0x65, 0x71, 0x3a, 0x20, 0x32, 0x31, 0x35, 0x38, 0x33, 0x34, 0x34, 0x38, 0x39, 0x20, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0xd, 0xa, 0x46, 0x72, 0x6f, 0x6d, 0x3a, 0x20, 0x3c, 0x73, 0x69, 0x70, 0x3a, 0x31, 0x39, 0x32, 0x2e, 0x31, 0x36, 0x38, 0x2e, 0x31, 0x31, 0x31, 0x2e, 0x31, 0x31, 0x31, 0x3a, 0x35, 0x30, 0x36, 0x30, 0x3e, 0x3b, 0x74, 0x61, 0x67, 0x3d, 0x36, 0x64, 0x66, 0x63, 0x66, 0x66, 0x65, 0x38, 0x2b, 0x31, 0x2b, 0x62, 0x30, 0x61, 0x39, 0x30, 0x30, 0x30, 0x33, 0x2b, 0x63, 0x39, 0x65, 0x66, 0x63, 0x32, 0x30, 0x62, 0xd, 0xa, 0x54, 0x6f, 0x3a, 0x20, 0x3c, 0x73, 0x69, 0x70, 0x3a, 0x31, 0x39, 0x32, 0x2e, 0x31, 0x36, 0x38, 0x2e, 0x31, 0x31, 0x31, 0x2e, 0x31, 0x31, 0x31, 0x3a, 0x35, 0x30, 0x36, 0x30, 0x3b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x3d, 0x75, 0x64, 0x70, 0x3e, 0x3b, 0x74, 0x61, 0x67, 0x3d, 0x31, 0x38, 0x30, 0x34, 0x61, 0x34, 0x37, 0x64, 0x2b, 0x31, 0x2b, 0x65, 0x31, 0x30, 0x35, 0x30, 0x34, 0x37, 0x30, 0x2b, 0x62, 0x31, 0x32, 0x38, 0x61, 0x35, 0x36, 0x39, 0xd, 0xa, 0x56, 0x69, 0x61, 0x3a, 0x20, 0x53, 0x49, 0x50, 0x2f, 0x32, 0x2e, 0x30, 0x2f, 0x55, 0x44, 0x50, 0x20, 0x31, 0x39, 0x32, 0x2e, 0x31, 0x36, 0x38, 0x2e, 0x31, 0x31, 0x31, 0x2e, 0x31, 0x31, 0x31, 0x3a, 0x35, 0x30, 0x36, 0x30, 0x3b, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x3d, 0x7a, 0x39, 0x68, 0x47, 0x34, 0x62, 0x4b, 0x2b, 0x32, 0x31, 0x66, 0x31, 0x31, 0x33, 0x65, 0x37, 0x65, 0x33, 0x64, 0x30, 0x34, 0x63, 0x38, 0x34, 0x36, 0x31, 0x34, 0x38, 0x61, 0x39, 0x61, 0x64, 0x37, 0x36, 0x30, 0x37, 0x61, 0x65, 0x66, 0x61, 0x31, 0x2b, 0x36, 0x64, 0x66, 0x63, 0x66, 0x66, 0x65, 0x38, 0x2b, 0x31, 0xd, 0xa, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a, 0x20, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0xd, 0xa, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2d, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x3a, 0x20, 0x37, 0x38, 0xd, 0xa, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2d, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x20, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x64, 0x70, 0xd, 0xa, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x3a, 0x20, 0x31, 0x30, 0x30, 0x72, 0x65, 0x6c, 0x2c, 0x20, 0x74, 0x69, 0x6d, 0x65, 0x72, 0xd, 0xa, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x2d, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x3a, 0x20, 0x65, 0x6e, 0xd, 0xa, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x2d, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0xd, 0xa, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x3a, 0x20, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x64, 0x70, 0x2c, 0x20, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x73, 0x75, 0x70, 0x2c, 0x20, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x2f, 0x6d, 0x69, 0x78, 0x65, 0x64, 0xd, 0xa, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x3a, 0x20, 0x49, 0x4e, 0x56, 0x49, 0x54, 0x45, 0x2c, 0x20, 0x41, 0x43, 0x4b, 0x2c, 0x20, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x2c, 0x20, 0x42, 0x59, 0x45, 0x2c, 0x20, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x2c, 0x20, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x2c, 0x20, 0x50, 0x52, 0x41, 0x43, 0x4b, 0x2c, 0x20, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x2c, 0x20, 0x49, 0x4e, 0x46, 0x4f, 0x2c, 0x20, 0x52, 0x45, 0x46, 0x45, 0x52, 0xd, 0xa, 0xd, 0xa, 0x76, 0x3d, 0x30, 0xd, 0xa, 0x6f, 0x3d, 0x2d, 0x20, 0x30, 0x20, 0x30, 0x20, 0x49, 0x4e, 0x20, 0x49, 0x50, 0x34, 0x20, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0xd, 0xa, 0x73, 0x3d, 0x2d, 0xd, 0xa, 0x63, 0x3d, 0x49, 0x4e, 0x20, 0x49, 0x50, 0x34, 0x20, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0xd, 0xa, 0x74, 0x3d, 0x30, 0x20, 0x30, 0xd, 0xa, 0x6d, 0x3d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x20, 0x30, 0x20, 0x52, 0x54, 0x50, 0x2f, 0x41, 0x56, 0x50, 0x20, 0x38}

func BenchmarkProcess(b *testing.B) {
	config.Cfg.Dedup = false
	//config.Cfg.DiscardMethod = "REGISTER"
	//config.Cfg.Mode = "SIPLOG"
	d := NewDecoder(layers.LinkTypeEthernet)
	ci := gopacket.CaptureInfo{Timestamp: time.Now(), CaptureLength: 715, Length: 715, InterfaceIndex: 4}

	for i := 0; i < b.N; i++ {
		val, _ := d.Process(rawPacket, &ci)
		_ = val
	}
}

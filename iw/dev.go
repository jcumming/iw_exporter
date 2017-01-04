package iw

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
)

// Device is a wireless device and its metadata.
type Device struct {
	Iface   string
	Channel int
	Width   int
	TxPower float64
}

var (
	reInt  = regexp.MustCompile(`^\s*Interface ([^\s]+)`)
	reChan = regexp.MustCompile(`^\s*channel ([0-9]+) .*width: ([0-9]+)`)
	reTx   = regexp.MustCompile(`^\s*txpower ([0-9.]+)`)
)

// Devices reads "iw dev" text to return filled in Device structs
func Devices(r io.Reader) (devs []*Device) {
	var iw *Device

	s := bufio.NewScanner(r)
	for s.Scan() {
		t := s.Text()

		if m := reInt.FindStringSubmatch(t); m != nil {
			iw = &Device{Iface: m[1]}
			devs = append(devs, iw)
		} else if iw == nil {
			continue
		} else if m := reChan.FindStringSubmatch(t); m != nil {
			if ch, err := strconv.Atoi(m[1]); err == nil {
				iw.Channel = ch
			}
			if w, err := strconv.Atoi(m[2]); err == nil {
				iw.Width = w
			}
		} else if m := reTx.FindStringSubmatch(t); m != nil {
			if p, err := strconv.ParseFloat(m[1], 64); err == nil {
				iw.TxPower = p
			}
		}
	}
	return devs
}

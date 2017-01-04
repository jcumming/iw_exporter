package iw

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestReadDevices(t *testing.T) {
	f, err := os.Open("testdata/dev")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	devs := Devices(f)

	assert.Equal(t, 3, len(devs), "len(devs)")

	assert.Equal(t, "broken0", devs[0].Iface)
	assert.Equal(t, "wlan1", devs[1].Iface)
	assert.Equal(t, "wlan0", devs[2].Iface)

	assert.Equal(t, 1, devs[1].Channel)
	assert.Equal(t, 36, devs[2].Channel)

	assert.Equal(t, 30.0, devs[1].TxPower)
	assert.Equal(t, 23.0, devs[2].TxPower)
}

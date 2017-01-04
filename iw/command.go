package iw

// This file is particularly ugly.

import (
	"os/exec"

	"github.com/prometheus/client_golang/prometheus"
)

// DeviceCmd returns devices from "iw dev".
func DeviceCmd() (devs []*Device, err error) {
	cmd := exec.Command("iw", "dev")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err = cmd.Start(); err != nil {
		return nil, err
	}

	devs = Devices(stdout)

	if err = cmd.Wait(); err != nil {
		return nil, err
	}

	return devs, nil
}

// StationDumpCmd sends station metrics to Prometheus a collector channel.
func StationDumpCmd(iface string, ch chan<- prometheus.Metric) error {
	cmd := exec.Command("iw", "dev", iface, "station", "dump")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}

	if err = ReadStation(iface, stdout, ch); err != nil {
		return err
	}

	if err = cmd.Wait(); err != nil {
		return err
	}

	return nil
}

package iw

import (
	"fmt"
	"io"
	"os"
	"sync"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
)

func getAllMetrics(t *testing.T, r io.Reader) []prometheus.Metric {
	var got []prometheus.Metric
	ch := make(chan prometheus.Metric)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for m := range ch {
			got = append(got, m)
		}
		wg.Done()
	}()
	if err := ReadStation("test", r, ch); err != nil {
		t.Fatal(err)
	}
	close(ch)
	wg.Wait()

	return got
}

func TestReadStation(t *testing.T) {
	f, err := os.Open("testdata/station_dump")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	got := getAllMetrics(t, f)

	// TODO: some sanity checking on the metrics we get back.
	// maybe lookForMetric(got, "iw_signal_strength_dbm", station, 73.0)
	var m dto.Metric
	if err = got[0].Write(&m); err != nil {
		t.Fatal(err)
	}
	fmt.Printf("example %+v %+v", got[0].Desc(), m)
}

func TestMalformedData(t *testing.T) {
	f, err := os.Open("testdata/station_dump_bad")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	// TODO: consider doing more than just ensuring no panics.
	getAllMetrics(t, f)
}

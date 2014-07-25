package GoReporter

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestBuildReporter(t *testing.T) {

	const testUri = "http://sensable.uri/test"

	sensable := Sensable{
		SensorId:   "test-sensor-id",
		Unit:       "°c",
		SensorType: "temperature",
		Location: Location{
			Latitude:  52.5,
			Longitude: 13.3,
		},
		Name: "my test sensor",
	}

	settings := Settings{
		AccessToken: "some-access-token",
		Private:     false,
	}

	reporter := sensable.BuildReporter(settings, testUri)

	assert.Equal(t, testUri, reporter.uri, "reporter uri should be the configured one")
	assert.Equal(t, settings, reporter.settings, "reporter settings should be the configured ones")
	assert.Equal(t, sensable, reporter.sensable, "reporter sensable should be the one the reporter was created from")
}

func TestReportSucceeds(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))
	defer ts.Close()

	sensable := Sensable{
		SensorId:   "test-sensor-id",
		Unit:       "°c",
		SensorType: "temperature",
		Location: Location{
			Latitude:  52.5,
			Longitude: 13.3,
		},
		Name: "my test sensor",
	}

	settings := Settings{
		AccessToken: "some-access-token",
		Private:     false,
	}

	reporter := sensable.BuildReporter(settings, ts.URL)

	sample := Sample{
		Value:     32.5,
		Timestamp: int64(time.Now().UnixNano() / 1e6),
		State:     "it's getting warmer",
	}

	success, err := reporter.Report(sample)

	assert.True(t, success, "the call to report should succeed")
	assert.Nil(t, err, "the call to report should not return an error")
}

func TestReportFails(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	}))
	defer ts.Close()

	sensable := Sensable{
		SensorId:   "test-sensor-id",
		Unit:       "°c",
		SensorType: "temperature",
		Location: Location{
			Latitude:  52.5,
			Longitude: 13.3,
		},
		Name: "my test sensor",
	}

	settings := Settings{
		AccessToken: "some-access-token",
		Private:     false,
	}

	reporter := sensable.BuildReporter(settings, ts.URL)

	sample := Sample{
		Value:     32.5,
		Timestamp: int64(time.Now().UnixNano() / 1e6),
		State:     "it's getting warmer",
	}

	success, err := reporter.Report(sample)

	assert.False(t, success, "the call to report should not succeed")
	assert.NotNil(t, err, "the call to report should not return an error")
}

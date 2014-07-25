/*
Package GoReporter contains data structures
and methods to post sensed data to http://sensable.io.
*/
package GoReporter

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

/*
Default URI where to reach the Sensable API.
*/
const SensableApiUri = "http://sensable.io/sensable"

/*
Sensable is an object that describe a sensor.
*/
type Sensable struct {
	SensorId   string   `json:"sensorid"`
	Unit       string   `json:"unit"`
	SensorType string   `json:"sensortype"`
	Location   Location `json:"location"`
	Name       string   `json:"name"`
}

func (sensable Sensable) BuildReporter(settings Settings, uri ...string) (reporter Reporter) {
	sensableUri := SensableApiUri
	if uri != nil {
		sensableUri = uri[0]
	}
	reporter = Reporter{
		uri:      sensableUri,
		settings: settings,
		sensable: sensable,
	}
	return
}

/*
Settings to be applied to this reporter.
*/
type Settings struct {
	AccessToken string `json:"accessToken"` //the accessToken that can be found in your Sensable account
	Private     bool   `json:"-"`           //indicates whether or not this Sensable should be kept private (currently ignored)
}

/*
Sample is a sample captured by the Sensable.
*/
type Sample struct {
	Value     float64 `json:"value"`
	Timestamp int64   `json:"timestamp"`
	State     string  `json:"state"`
}

/*
Location is the geolocation associated with this Sensable.
*/
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type payload struct {
	Sensable
	Settings
	Sample Sample `json:"sample"`
}

func (payload payload) upload(uri string) (bool, error) {
	b, err := json.Marshal(payload)
	if err != nil {
		return false, err
	}
	res, err := http.Post(uri, "application/json", bytes.NewReader(b))
	defer res.Body.Close()
	if err != nil {
		return false, err
	}
	if res.StatusCode == 200 {
		return true, nil
	}
	return false, errors.New("Server responded with " + res.Status)
}

/*
Reporter is the object that takes care of reporting new
samples for the given Sensable via the Sensable API.
*/
type Reporter struct {
	uri      string
	settings Settings
	sensable Sensable
}

func (reporter Reporter) Report(sample Sample) (bool, error) {
	payload := payload{
		Sensable: Sensable{
			reporter.sensable.SensorId,
			reporter.sensable.Unit,
			reporter.sensable.SensorType,
			reporter.sensable.Location,
			reporter.sensable.Name,
		},
		Settings: Settings{
			reporter.settings.AccessToken,
			reporter.settings.Private,
		},
		Sample: sample,
	}
	return payload.upload(reporter.uri)
}

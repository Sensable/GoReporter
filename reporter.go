package GoReporter

import (
        "encoding/json"
        )

const uri = "http://sensable.io/sensable"

type Sample struct {
    Data float64 `json:"data"`
    Time int64 `json:"time"`
}

type Sensable struct {
    SensorId string `json:"sensorid"`
    Unit string `json:"unit"`
    SensorType string `json:"sensortype"`
    Latitude float64 `json:"latitude,string"`
    Longitude float64 `json:"longitude,string"`
    Name string `json:"name"`
    Sample Sample `json:"sample"`
}

type payload struct {
    Sensable
    Options
}

type Options struct {
    AccessToken string `json:"accessToken"`
    Private bool `json:"-"`
}

func (payload payload) Upload(options Options) (string, error) {
    b, err := json.Marshal(payload)
    return string(b), err
}

func (sensable Sensable) BuildReporter(options Options) func(sample Sample) (string, error) {
    payload := payload {
        Sensable: Sensable {
            sensable.SensorId,
            sensable.Unit,
            sensable.SensorType,
            sensable.Latitude,
            sensable.Longitude,
            sensable.Name,
            sensable.Sample,
        },
        Options: Options {
            options.AccessToken,
            options.Private,
        },
    }
    return func (sample Sample) (string, error) {
        payload.Sample = sample
        return payload.Upload(options)
    }
}
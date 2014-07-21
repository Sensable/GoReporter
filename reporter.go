package GoReporter

import (
        "encoding/json";
        )

const SensableApiUri = "http://sensable.io/sensable"

type Api struct {
    Uri string
}

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
}

type payload struct {
    Sensable
    Settings
    Sample Sample `json:"sample"`
}

type Settings struct {
    AccessToken string `json:"accessToken"`
    Private bool `json:"-"`
}

func (payload payload) upload(settings Settings, api Api) (string, error) {
    b, err := json.Marshal(payload)
    return string(b), err
}

func (sensable Sensable) BuildReporter(settings Settings, api... Api) func(sample Sample) (string, error) {
    var apiConfiguration Api

    if api == nil {
        apiConfiguration = Api {
            Uri: SensableApiUri,
        }
    } else {
        apiConfiguration = api[0]
    }
    payload := payload {
        Sensable: Sensable {
            sensable.SensorId,
            sensable.Unit,
            sensable.SensorType,
            sensable.Latitude,
            sensable.Longitude,
            sensable.Name,
        },
        Settings: Settings {
            settings.AccessToken,
            settings.Private,
        },
    }
    return func (sample Sample) (string, error) {
        payload.Sample = sample
        return payload.upload(settings, apiConfiguration)
    }
}
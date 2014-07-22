package GoReporter

import (
        "encoding/json";
        "net/http";
        "bytes";
        "errors"
        )

const SensableApiUri = "http://sensable.io/sensable"

type Api struct {
    Uri string
}

type Sample struct {
    Data float64 `json:"data"`
    Time int64 `json:"time"`
    State string `json:"state"`
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

func (payload payload) upload(settings Settings, api Api) (bool, error) {
    b, err := json.Marshal(payload)
    if err != nil {
        return false, err
    }
    res, err := http.Post(api.Uri, "application/json", bytes.NewReader(b));
    defer res.Body.Close();
    if err != nil {
        return false, err
    }
    if res.StatusCode == 200 {
        return true, nil
    }
    return false, errors.New("Server responded with " + res.Status)
}

func (sensable Sensable) BuildReporter(settings Settings, api... Api) func(sample Sample) (bool, error) {
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
    return func (sample Sample) (bool, error) {
        payload.Sample = sample
        return payload.upload(settings, apiConfiguration)
    }
}
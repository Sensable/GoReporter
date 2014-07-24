package GoReporter

import (
        "encoding/json";
        "net/http";
        "bytes";
        "errors"
        )

const SensableApiUri = "http://sensable.io/sensable"

type Sample struct {
    Data float64 `json:"data"`
    Time int64 `json:"time"`
    State string `json:"state"`
}

type Sensable struct {
    SensorId string `json:"sensorid"`
    Unit string `json:"unit"`
    SensorType string `json:"sensortype"`
    Latitude float64 `json:"-"`
    Longitude float64 `json:"-"`
    Name string `json:"name"`
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
    res, err := http.Post(uri, "application/json", bytes.NewReader(b));
    defer res.Body.Close();
    if err != nil {
        return false, err
    }
    if res.StatusCode == 200 {
        return true, nil
    }
    return false, errors.New("Server responded with " + res.Status)
}

type Settings struct {
    AccessToken string `json:"accessToken"`
    Private bool `json:"-"`
}

type Reporter struct {
    uri string
    settings Settings
    sensable Sensable
}

func (reporter Reporter) Report(sample Sample) (bool, error) {
    payload := payload {
        Sensable: Sensable {
            reporter.sensable.SensorId,
            reporter.sensable.Unit,
            reporter.sensable.SensorType,
            reporter.sensable.Latitude,
            reporter.sensable.Longitude,
            reporter.sensable.Name,
        },
        Settings: Settings {
            reporter.settings.AccessToken,
            reporter.settings.Private,
        },
        Sample: sample,
    }
    return payload.upload(reporter.uri)
}

func (sensable Sensable) BuildReporter(settings Settings, uri... string) (reporter Reporter) {
    sensableUri := SensableApiUri
    if uri != nil {
        sensableUri = uri[0]
    }
    reporter = Reporter {
        uri: sensableUri,
        settings: settings,
        sensable: sensable,
    }
    return
}
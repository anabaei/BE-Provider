package main

import (
    "encoding/json"
    "errors"
    "net/http"
    "io"
)


type envelope map[string]any

func (app *application) writeJSON(w http.ResponseWriter, status int, data any) error {
    js, err := json.MarshalIndent(data, "", "\t")
    if err != nil {
        return err
    }
    js = append(js, '\n')
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write(js)
    return nil
}

//this protect our help service by
// setting maxbyte to 1MB
// disallowing unknown fields
// and ensuring that the request body only contains a single JSON value
// this is to prevent a user from sending a request with a large body that could be used to exhaust our server's memory



func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
    maxBytes := 1048576 // 1 MB
    r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

    dec := json.NewDecoder(r.Body)
    dec.DisallowUnknownFields()

    err := dec.Decode(dst)
    if err != nil {
        return err
    }

    // Check if there is extra data in the JSON body
    err = dec.Decode(&struct{}{})
    if err != io.EOF {
        return errors.New("body must only contain a single JSON value")
    }

    return nil
}

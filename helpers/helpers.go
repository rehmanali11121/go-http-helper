package helpers

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
)

// RespondWithError sends an HTTP error response with the given status code and message
func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
    http.Error(w, message, statusCode)
}

// RespondWithJSON sends an HTTP response with the given status code and JSON body
func RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(data)
}

// HTTPRequest sends an HTTP request with the given method, URL, headers, and body,
// and returns the response body and any error encountered
func HTTPRequest(method, url string, headers map[string]string, body interface{}) ([]byte, error) {
    requestBody, err := json.Marshal(body)
    if err != nil {
        return nil, err
    }

    req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
    if err != nil {
        return nil, err
    }

    // Set request headers
    for key, value := range headers {
        req.Header.Set(key, value)
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    responseBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    return responseBody, nil
}

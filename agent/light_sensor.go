package agent

import (
    "github.com/mozillazg/request"
    "net/url"
    "errors"
    "strconv"
)

type LightSensor struct {
    url2Request *url.URL
    request     *request.Request
}

var NOT_JSON = errors.New("not json")
var RESPONSE_BODY_ILLEGAL = errors.New("body illegal")

const ADD_THREAD = "/thresholds/"
func (ls *LightSensor) AddThreshold(value int) (isSuccessful bool, err error) {
    isSuccessful = false
    res, err := ls.request.Post(ls.url2Request.String() + ADD_THREAD + strconv.Itoa(value))
    if err != nil {
        return
    }
    resJson, jsonErr := res.Json()
    if jsonErr != nil {
        err = NOT_JSON
        return
    }
    isSuccessful, errNoField := resJson.Get("is_success").Bool()
    if errNoField != nil {
        err = RESPONSE_BODY_ILLEGAL
        return
    }
    return
}

const URL_GET_LIGHT_VALUE = "/light"
func (ls *LightSensor) getLightValue() (lightValue int, err error) {
    res, err := ls.request.Get(ls.url2Request.String() + URL_GET_LIGHT_VALUE)
    if err != nil {
        return
    }
    resJson, jsonErr := res.Json()
    if jsonErr != nil {
        err = NOT_JSON
        return
    }
    lightValue, errNoField := resJson.Get("value").Int()
    if errNoField != nil {
        err = RESPONSE_BODY_ILLEGAL
        return
    }
    return
}

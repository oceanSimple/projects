package handler

import "github.com/tidwall/gjson"

type addPlayerParams struct {
	Account string `json:"account"`
}

var addPlayer = func(data gjson.Result) []byte {
	return nil
}

package api

import "encoding/json"

type Error struct {
	Msg string
}

func errorToJson(error Error) string {
	e, _ := json.Marshal(&error)
	return string(e)
}

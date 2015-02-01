package impl

import (
	"../"
	rj "github.com/fkmhrk-go/rawjson"
)

func json(r service.Result) rj.RawJsonObject {
	j, _ := rj.ObjectFromString(r.Body())
	return j
}

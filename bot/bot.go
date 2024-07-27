package bot

import (
	"os"

	"github.com/tencent-connect/botgo/openapi"
)

type PaperAirplaneBot struct {
	openapi.OpenAPI
	hostUrl string
}

var hostUrlTemplate struct {
	User  string
	Group string
}

var PaperAirplane *PaperAirplaneBot

const (
	hostUser         = "https://api.sgroup.qq.com/v2/users/%s/messages"
	hostGroup        = "https://api.sgroup.qq.com/v2/groups/%s/messages"
	hostUserSandbox  = "https://sandbox.api.sgroup.qq.com/v2/users/%s/messages"
	hostGroupSandbox = "https://sandbox.api.sgroup.qq.com/v2/groups/%s/messages"
)

func BuildClient(api openapi.OpenAPI) {
	PaperAirplane = new(PaperAirplaneBot)
	PaperAirplane.OpenAPI = api

	hostUrlTemplate.User = hostUserSandbox
	hostUrlTemplate.Group = hostGroupSandbox

	if os.Getenv("AIRP_MODE") == "release" {
		hostUrlTemplate.User = hostUser
		hostUrlTemplate.Group = hostGroup
	}
}

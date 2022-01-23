package entity

import "mhdsg/pkg/base"

type Sg_server_user_relation struct {
	base.AbstractMongoEntendEntity
	Userid          int64 `json:"userid"`
	Server_identity int32 `json:"server_identity"`
}

package utils

import (
	"strings"

	pbCommon "github.com/0x726f6f6b6965/mono-service/protos/common/v1"
)

func GetRoleType(input string) pbCommon.RoleType {
	switch strings.ToUpper(input) {
	case "ADMIN":
		return pbCommon.RoleType_ROLE_TYPE_ADMIN
	case pbCommon.RoleType_ROLE_TYPE_ADMIN.String():
		return pbCommon.RoleType_ROLE_TYPE_ADMIN
	case "VIEWER":
		return pbCommon.RoleType_ROLE_TYPE_VIEWER
	case pbCommon.RoleType_ROLE_TYPE_VIEWER.String():
		return pbCommon.RoleType_ROLE_TYPE_VIEWER
	default:
		return pbCommon.RoleType_ROLE_TYPE_GUEST
	}
}

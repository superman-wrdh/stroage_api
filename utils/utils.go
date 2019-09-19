package utils

import (
	"github.com/google/uuid"
	"strings"
)

func GetUUID4() string {
	uuid, _ := uuid.NewUUID()
	uid := uuid.String()
	uid = strings.Replace(uid, "-", "", -1)
	return uid
}

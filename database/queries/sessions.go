package queries

import (
	"photon/database"
	"photon/model"
)

func StoreSessionId(user *model.Credential, sessionId string) error {
	rdb := database.GetCacheInstance()
	userId := user.Id.String()

	_, err := rdb.LPush(database.Ctx, sessionId, userId).Result()

	return err
}

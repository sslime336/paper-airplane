package spark

import (
	"encoding/json"
	"errors"
	"paper-airplane/db"
	"time"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

var sessionCache *cache.Cache

func initCache() {
	sessionCache = cache.New(8*time.Minute, 30*time.Minute)

	sessionCache.OnEvicted(func(openId string, i any) {
		sess := i.(*Session)
		mh, _ := json.Marshal(sess.Message)
		var msess gorm.Session
		if err := db.Sqlite.Where("openId = ?", openId).First(&msess).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			db.Sqlite.Create(&db.Session{
				OpenId:         openId,
				MessageHistory: mh,
			})
		} else {
			db.Sqlite.Model(&msess).Update("messageHistory", mh)
		}
	})
}

package spark

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/sslime336/paper-airplane/db"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var sessionCache *cache.Cache

func initCache() {
	sessionCache = cache.New(8*time.Minute, 30*time.Minute)

	sessionCache.OnEvicted(func(openId string, i any) {
		log.Debug("session cache out of time", zap.String("session.openId", openId))
		sess := i.(*Session)
		mh, _ := json.Marshal(sess.Message)
		var msess db.Session
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

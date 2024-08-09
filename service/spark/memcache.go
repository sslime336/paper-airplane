package spark

import (
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/sslime336/paper-airplane/dao"
	"github.com/sslime336/paper-airplane/db/orm"
	"gorm.io/gen/field"
)

var sessionCache *cache.Cache

func initSessionCache() {
	sessionCache = cache.New(2*time.Hour, 5*time.Hour)
	sessionCache.OnEvicted(func(openId string, data any) {
		session := data.(*Session)
		session.Conn.Close()
		if session.TokenOverflow() {
			session.ResetToken()
		}
		dao.Session.Where(dao.Session.OpenId.Eq(openId)).Assign(field.Attrs(&orm.Session{
			OpenId:          openId,
			JsonizedHistory: session.JsonizedMessage(),
			TotalTokens:     session.TotalTokens,
		})).FirstOrCreate()
	})
}

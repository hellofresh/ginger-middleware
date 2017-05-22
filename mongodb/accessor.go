package mongodb

import (
	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

// DatabaseAccessor represents a mongo database encapsulation
type DatabaseAccessor struct {
	*mgo.Session
}

// InitDB creates a new mongo db server
func InitDB(dsn string) (*DatabaseAccessor, error) {
	log.Debugf("Trying to connect to %s", dsn)
	session, err := mgo.Dial(dsn)

	if err == nil {
		log.Debug("Connected to mongodb")
		session.SetMode(mgo.Monotonic, true)
		return &DatabaseAccessor{session}, nil
	}

	return nil, err
}

// Set a session to a context
func (da *DatabaseAccessor) Set(c *gin.Context, session *mgo.Session) {
	db := da.DB("")
	c.Set("db", db)
	c.Set("mgoSession", session)
}

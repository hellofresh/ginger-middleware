package mongodb

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// Middleware represents a database connection
type Middleware struct {
	Dba *DatabaseAccessor
}

//Middleware is the gin middleware for the database, this is different from the others since it's a
//database handler
func (m *Middleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Debug("Starting Database middleware")

		reqSession := m.Dba.Clone()
		defer reqSession.Close()
		m.Dba.Set(c, reqSession)

		c.Next()
	}
}

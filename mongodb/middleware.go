package mongodb

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

//Middleware is the gin middleware for the database, this is different from the others since it's a
//database handler
func Middleware(accessor *DatabaseAccessor) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Debug("Starting Database middleware")

		reqSession := accessor.Clone()
		defer reqSession.Close()
		accessor.Set(c, reqSession)

		c.Next()
	}
}

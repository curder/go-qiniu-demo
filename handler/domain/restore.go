package domain

import (
	"github.com/curder/go-qiniu-demo/handler"
	"github.com/curder/go-qiniu-demo/pkg/errno"
	"github.com/curder/go-qiniu-demo/pkg/log"
	"github.com/curder/go-qiniu-demo/service/domain"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 恢复域名
func Restore(c *gin.Context) {
	var (
		domainID     int
		rowsAffected int64
		err          error
	)

	log.Info("domain restore function called.")

	domainID, err = strconv.Atoi(c.Param("id"))

	if rowsAffected, err = domain.DomainSvc.Restore(uint64(domainID)); err != nil && rowsAffected != 0 {
		log.Warnf("[domain] restore domain err, %v", err)
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}

	handler.SendResponse(c, nil, domainID)
}

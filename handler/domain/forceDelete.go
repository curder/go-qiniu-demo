package domain

import (
	"github.com/curder/go-qiniu-demo/handler"
	"github.com/curder/go-qiniu-demo/pkg/errno"
	"github.com/curder/go-qiniu-demo/pkg/log"
	"github.com/curder/go-qiniu-demo/service/domain"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 强制删除
func ForceDelete(c *gin.Context) {
	var (
		domainID     int
		rowsAffected int64
		err          error
	)

	log.Info("domain force delete function called.")

	domainID, err = strconv.Atoi(c.Param("id"))

	if rowsAffected, err = domain.DomainSvc.ForceDeleteDomain(uint64(domainID)); err != nil && rowsAffected != 0 {
		log.Warnf("[domain] force delete domain err, %v", err)
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}

	handler.SendResponse(c, nil, domainID)
}

package domain

import (
	"github.com/curder/go-qiniu-demo/handler"
	"github.com/curder/go-qiniu-demo/pkg/errno"
	"github.com/curder/go-qiniu-demo/pkg/log"
	"github.com/curder/go-qiniu-demo/service/domain"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 更新
func Update(c *gin.Context) {
	var (
		req       UpdateDomainRequest
		domainID  int
		domainMap map[string]interface{}
		err       error
	)

	log.Info("domain update function called.")

	domainID, _ = strconv.Atoi(c.Param("id"))

	if err = c.ShouldBind(&req); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	log.Infof("req %#v", req)

	domainMap = make(map[string]interface{}, 1)
	domainMap["bucket_id"] = req.BucketID
	domainMap["protocol"] = req.Protocol
	domainMap["hostname"] = req.Hostname
	domainMap["description"] = req.Description

	if err = domain.DomainSvc.UpdateDomain(uint64(domainID), domainMap); err != nil {
		log.Warnf("[domain] update domain err, %v", err)
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}

	handler.SendResponse(c, nil, domainID)
}

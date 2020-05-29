package domain

import (
	"github.com/curder/go-qiniu-demo/handler"
	"github.com/curder/go-qiniu-demo/model"
	"github.com/curder/go-qiniu-demo/pkg/errno"
	"github.com/curder/go-qiniu-demo/pkg/log"
	"github.com/curder/go-qiniu-demo/service/domain"
	"github.com/gin-gonic/gin"
)

// 域名列表
func Index(c *gin.Context) {
	log.Info("Get domain list function called.")

	var (
		domains []*model.DomainModel
		err     error
	)
	if domains, err = domain.DomainSvc.GetDomains(); err != nil {
		handler.SendResponse(c, errno.ErrDomainNotFound, nil)
		return
	}

	handler.SendResponse(c, nil, domains)
}

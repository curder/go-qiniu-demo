package domain

import (
	"github.com/curder/go-qiniu-demo/handler"
	"github.com/curder/go-qiniu-demo/model"
	"github.com/curder/go-qiniu-demo/pkg/errno"
	"github.com/curder/go-qiniu-demo/pkg/log"
	"github.com/curder/go-qiniu-demo/service/domain"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 存储桶详情
func Show(c *gin.Context) {
	var (
		domainID    int
		domainModel *model.DomainModel
		response    ShowDomainInfoResponse
		err         error
	)

	log.Info("domain show function called.")

	domainID, _ = strconv.Atoi(c.Param("id"))

	if domainModel, err = domain.DomainSvc.GetByID(uint64(domainID)); err != nil {
		log.Warnf("[domain] show domain err, %v", err)
		handler.SendResponse(c, errno.ErrDomainNotFound, nil)
		return
	}

	response = ShowDomainInfoResponse{
		ID:       domainModel.ID,
		BucketID: domainModel.BucketID,
		Protocol: domainModel.Protocol,
		Hostname: domainModel.Hostname,
		Description: domainModel.Description,
		CreatedAt:   domainModel.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   domainModel.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	handler.SendResponse(c, nil, response)
}

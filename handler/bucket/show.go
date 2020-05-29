package bucket

import (
	"github.com/curder/go-qiniu-demo/handler"
	"github.com/curder/go-qiniu-demo/model"
	"github.com/curder/go-qiniu-demo/pkg/errno"
	"github.com/curder/go-qiniu-demo/pkg/log"
	"github.com/curder/go-qiniu-demo/service/bucket"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 存储桶详情
func Show(c *gin.Context) {
	var (
		bucketID    int
		bucketModel *model.BucketModel
		response    ShowBucketInfoResponse
		err         error
	)

	log.Info("bucket show function called.")

	bucketID, _ = strconv.Atoi(c.Param("id"))

	if bucketModel, err = bucket.BucketSvc.GetBucketByID(uint64(bucketID)); err != nil {
		log.Warnf("[bucket] show bucket err, %v", err)
		handler.SendResponse(c, errno.ErrBucketNotFound, nil)
		return
	}

	response = ShowBucketInfoResponse{
		ID:          bucketModel.ID,
		AccountID:   bucketModel.AccountID,
		Name:        bucketModel.Name,
		Description: bucketModel.Description,
		CreatedAt:   bucketModel.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   bucketModel.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	handler.SendResponse(c, nil, response)
}

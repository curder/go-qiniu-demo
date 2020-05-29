package bucket

import (
	"github.com/curder/go-qiniu-demo/handler"
	"github.com/curder/go-qiniu-demo/pkg/errno"
	"github.com/curder/go-qiniu-demo/pkg/log"
	"github.com/curder/go-qiniu-demo/service/bucket"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 恢复存储桶
func Restore(c *gin.Context) {
	var (
		bucketID     int
		rowsAffected int64
		err          error
	)

	log.Info("bucket restore function called.")

	bucketID, err = strconv.Atoi(c.Param("id"))

	if rowsAffected, err = bucket.BucketSvc.Restore(uint64(bucketID)); err != nil && rowsAffected != 0 {
		log.Warnf("[bucket] restore bucket err, %v", err)
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}

	handler.SendResponse(c, nil, bucketID)
}

package bucket

import (
	"github.com/curder/go-qiniu-demo/handler"
	"github.com/curder/go-qiniu-demo/pkg/errno"
	"github.com/curder/go-qiniu-demo/pkg/log"
	"github.com/curder/go-qiniu-demo/service/bucket"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 删除存储桶
func Delete(c *gin.Context) {
	var (
		bucketID     int
		rowsAffected int64
		err          error
	)

	log.Info("account delete function called.")

	bucketID, _ = strconv.Atoi(c.Param("id"))

	if rowsAffected, err = bucket.BucketSvc.Delete(uint64(bucketID)); err != nil && rowsAffected != 0 {
		log.Warnf("[bucket] delete bucket err, %v", err)
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}

	handler.SendResponse(c, nil, bucketID)

}

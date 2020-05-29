package bucket

import (
	"github.com/curder/go-qiniu-demo/handler"
	"github.com/curder/go-qiniu-demo/pkg/errno"
	"github.com/curder/go-qiniu-demo/pkg/log"
	"github.com/curder/go-qiniu-demo/service/bucket"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 更新
func Update(c *gin.Context) {
	var (
		req UpdateBucketInfoRequest
		bucketID  int
		bucketMap map[string]interface{}
		err       error
	)

	log.Info("bucket update function called.")

	bucketID, _ = strconv.Atoi(c.Param("id"))

	if err = c.ShouldBind(&req); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	log.Infof("req %#v", req)

	bucketMap = make(map[string]interface{}, 1)
	bucketMap["name"] = req.Name
	bucketMap["description"] = req.Description

	if err = bucket.BucketSvc.UpdateBucket(uint64(bucketID), bucketMap); err != nil {
		log.Warnf("[bucket] update bucket err, %v", err)
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}

	handler.SendResponse(c, nil, bucketID)
}

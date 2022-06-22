package ctrl

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/scm/pkg/constx"
	"github.com/thanhpp/scm/pkg/ginutil"
)

const (
	URLFileName = "filename"
)

type FileCtrl struct {
}

func NewFileCtrl() *FileCtrl {
	return &FileCtrl{}
}

func (ctrl FileCtrl) ServeFile(c *gin.Context) {
	filename := c.Param(URLFileName)
	if len(filename) == 0 {
		ginutil.RespErr(c, http.StatusNotAcceptable, errors.New("empty filename"))
		return
	}

	c.File(fmt.Sprintf("%s/%s", constx.SaveFilePaths, filename))
}

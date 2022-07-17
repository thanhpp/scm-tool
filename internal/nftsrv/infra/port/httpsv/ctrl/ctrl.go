package ctrl

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/scm/internal/nftsrv/app"
	"github.com/thanhpp/scm/internal/nftsrv/infra/port/httpsv/dto"
	"github.com/thanhpp/scm/pkg/ginutil"
)

type NFTMinterCtrl struct {
	app *app.App
}

func NewNFTMinterCtrl(app *app.App) *NFTMinterCtrl {
	return &NFTMinterCtrl{
		app: app,
	}
}

func (ctrl NFTMinterCtrl) MintNFT(c *gin.Context) {
	req := new(dto.MintSeriNFTReq)

	if err := c.ShouldBind(req); err != nil {
		ginutil.RespErr(c, http.StatusNotAcceptable, err, ginutil.WithData(req))
		return
	}

	seriNFT, err := ctrl.app.MintSeriNFT(c, req.Seri, req.Metadata)
	if err != nil {
		ginutil.RespErr(c, http.StatusInternalServerError, err, ginutil.WithData(req))
		return
	}

	resp := new(dto.MintSeriNFTResp)
	resp.Set200OK()
	resp.SetData(seriNFT)

	c.JSON(http.StatusOK, resp)
}

func (ctrl NFTMinterCtrl) GetBySeri(c *gin.Context) {
	seri := c.Param("seri")

	seriNFT, err := ctrl.app.GetSeriNFTBySeri(c, seri)
	if err != nil {
		ginutil.RespErr(c, http.StatusInternalServerError, err, ginutil.WithData(seri))
		return
	}

	resp := new(dto.GetSeriNFTResp)
	resp.Set200OK()
	resp.SetData(seriNFT)

	c.JSON(http.StatusOK, resp)
}

func (ctrl NFTMinterCtrl) GetByTxHash(c *gin.Context) {
	txHash := c.Param("txHash")

	seriNFT, err := ctrl.app.GetSeriNFTByTxHash(c, txHash)
	if err != nil {
		ginutil.RespErr(c, http.StatusInternalServerError, err, ginutil.WithData(txHash))
		return
	}

	resp := new(dto.GetSeriNFTResp)
	resp.Set200OK()
	resp.SetData(seriNFT)

	c.JSON(http.StatusOK, resp)
}

func (ctrl NFTMinterCtrl) GetByTokenID(c *gin.Context) {
	tokenID := c.Param("tokenID")

	uint64TokenID, err := strconv.ParseInt(tokenID, 10, 64)
	if err != nil {
		ginutil.RespErr(c, http.StatusNotAcceptable, err, ginutil.WithData(tokenID))
		return
	}

	seriNFT, err := ctrl.app.GetSeriNFTByTokenID(c, uint64TokenID)
	if err != nil {
		ginutil.RespErr(c, http.StatusInternalServerError, err, ginutil.WithData(tokenID))
		return
	}

	resp := new(dto.GetSeriNFTResp)
	resp.Set200OK()
	resp.SetData(seriNFT)

	c.JSON(http.StatusOK, resp)
}

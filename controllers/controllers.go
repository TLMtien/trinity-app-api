package controllers

import (
	"net/http"
	"strconv"
	"trinity_app/models"
	"trinity_app/services"
	"trinity_app/utils"

	"github.com/gin-gonic/gin"
)

type Controllers interface {
	GenerateVoucher(ctx *gin.Context)
	Purchase(ctx *gin.Context)
	CheckEligibilityCampaign(ctx *gin.Context)
	CheckValidVoucher(ctx *gin.Context)
}

type controller struct {
	services services.Services
}

func NewControllers(services services.Services) Controllers {
	return &controller{services}
}

// 	Generate Voucher godoc
//	@Summary		Generate Voucher
//	@Description	Generate Voucher
//	@Tags			voucher
//	@Accept		json
//	@Produce		json
//	@Param		voucher		body		utils.RequestGenerateVoucher	true	"User Request Body"
//	@Success		200		{object}	utils.Response{}
//	@Failure		400		{object}	utils.ResponseFailureBadRequest
//	@Failure		500		{object}	utils.ResponseFailureServerError{}
//	@Router		/api/voucher	[post]
func (c *controller) GenerateVoucher(ctx *gin.Context) {
	var voucher = models.Voucher{}
	if err := ctx.ShouldBindJSON(&voucher); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponseJson("[ctl]"+err.Error()))
		return
	}

	v, err, code := c.services.GenerationVoucher(ctx, &voucher)
	if err != nil {
		ctx.JSON(code, utils.ErrorResponseJson(err.Error()))
		return
	}

	ctx.JSON(code, utils.SuccessResponseJson("Generate voucher success", v))
}

// 	Campaign godoc
//	@Summary		Check Eligibility Campaign
//	@Description	Check Eligibility Campaign
//	@Tags			campaign
//	@Accept		json
//	@Produce		json
// 	@Param        	id   		path      	int  true  "campaign id" default(1)
//	@Success		200		{object}	utils.Response{}
//	@Failure		400		{object}	utils.ResponseFailureBadRequest
//	@Failure		500		{object}	utils.ResponseFailureServerError{}
//	@Router		/api/campaign/{id}		[get]
func (c *controller) CheckEligibilityCampaign(ctx *gin.Context) {
	campaignID := ctx.Param("id")
	campaignIDConvert, err := strconv.ParseUint(campaignID, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponseJson("[ctl]"+err.Error()))
		return
	}

	p, err, code := c.services.CheckEligibilityCampaign(ctx, uint32(campaignIDConvert))

	if err != nil {
		ctx.JSON(code, utils.ErrorResponseJson(err.Error()))
		return
	}

	ctx.JSON(code, utils.SuccessResponseJson("campaign is valid", p))
}

// 	Voucher godoc
//	@Summary		Check Valid Voucher
//	@Description	Check Valid Voucher
//	@Tags			voucher
//	@Accept		json
//	@Produce		json
//	@Param		voucher_code		query		string	true	"voucher code"
//	@Success		200		{object}	utils.Response{}
//	@Failure		400		{object}	utils.ResponseFailureBadRequest
//	@Failure		500		{object}	utils.ResponseFailureServerError{}
//	@Router		/api/voucher 	[get]
func (c *controller) CheckValidVoucher(ctx *gin.Context) {
	voucherCode := ctx.Query("voucher_code")
	if voucherCode == "" {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponseJson("[ctl]voucher code is required"))
		return
	}
	v, err, code := c.services.CheckValidVoucher(ctx, voucherCode)
	if err != nil {
		ctx.JSON(code, utils.ErrorResponseJson(err.Error()))
		return
	}

	ctx.JSON(code, utils.SuccessResponseJson("voucher is valid", v))

}

// 	Purchase godoc
//	@Summary		Purchase
//	@Description	Purchase
//	@Tags			purchase
//	@Accept		json
//	@Produce		json
//	@Param		voucher		body		utils.RequestPurchase	true	"Purchae Request Body"
//	@Success		200		{object}	utils.Response{}
//	@Failure		400		{object}	utils.ResponseFailureBadRequest
//	@Failure		500		{object}	utils.ResponseFailureServerError{}
//	@Router		/api/purchase	[post]
func (c *controller) Purchase(ctx *gin.Context) {
	var purchase = models.Purchase{}
	if err := ctx.ShouldBindJSON(&purchase); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponseJson("[ctl]"+err.Error()))
		return
	}
	p, err, code := c.services.Purchase(ctx, &purchase)
	if err != nil {
		ctx.JSON(code, utils.ErrorResponseJson(err.Error()))
		return
	}

	ctx.JSON(code, utils.SuccessResponseJson("Create purchase success", p))
}

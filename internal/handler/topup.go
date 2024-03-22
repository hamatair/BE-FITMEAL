package handler

import (
	"errors"
	"intern-bcc/entity"
	"intern-bcc/model"
	"intern-bcc/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) TopUp(c *gin.Context) {
	var amount model.TopUp

	err := c.ShouldBindJSON(&amount)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
	}

	user, ok := c.Get("user")
	if !ok {
		response.Error(c, http.StatusNotFound, "failed to get data user", errors.New(""))
	}

	resp, err := h.Service.TopUpService.InitializeTopUp(model.TopUpReq{
		Amount: amount.Amount,
		UserId: user.(entity.User).ID,
	})
	if err != nil {
		response.Error(c, http.StatusBadGateway, "failed to initialize topup", err)
	}

	response.Success(c, http.StatusOK, "success to initialize", resp)
}

func (h *Handler) VerifyPayment(c *gin.Context) {
	var notificationPayload map[string]interface{}
	err := c.ShouldBindJSON(&notificationPayload)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
	}
	// err := json.NewDecoder(c.Request.Response.Body).Decode(&notificationPayload)
	// if err != nil {
	// 	response.Error(c, http.StatusBadRequest, "failed to bind input", err)
	// }

	user, ok := c.Get("user")
	if !ok {
		response.Error(c, http.StatusNotFound, "failed to get user data", errors.New(""))
	}

	err = h.Service.TopUpService.ConfirmedTopUp(user.(entity.User).ID.String(), notificationPayload)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to verify payment", err)
	}

	response.Success(c, http.StatusOK, "success to verify payment", nil)

}

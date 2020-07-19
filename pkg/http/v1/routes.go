package v1

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	sharedAPI "github.com/phev8/go_commons/pkg/api/shared"
	"github.com/phev8/go_service_api/pkg/utils"
	"google.golang.org/grpc/status"
)

func (h *HttpEndpoints) AddAPI(rg *gin.RouterGroup) {
	rg.POST("/service-a", h.getDataServiceA)
	rg.POST("/service-b", h.getDataServiceB)
}

func (h *HttpEndpoints) getDataServiceA(c *gin.Context) {
	var req sharedAPI.RequestObject
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.clients.ServiceA.GetDataFromB(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) getDataServiceB(c *gin.Context) {
	var req sharedAPI.RequestObject
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.clients.ServiceB.GetData(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

package health

import (
	"github.com/adf-code/beta-book-api/internal/delivery/response"
	"net/http"
)

// Health godoc
// @Summary      Health Check
// @Description  Health check for service
// @Tags         health
// @Success      200  {object}  response.APIResponse
// @Router       /healthz [get]
func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info().Msg("📥 Incoming health check request")
	response.Success(w, 200, "health", "healthCheck", "Success Health Check", nil)
}

package service

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// @Summary Delete car id
// @Description Delete information about a car by its registration number
// @Tags cars
// @Produce json
// @Param	regnum	path	string	true	"Car Registration Number"
// @Success 204 "Successful deletion, no content returned"
// @Failure 404 {object} map[string]string "Resource not found"
// @Router /cars/{regnum} [delete]
func (s *Service) delCars() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.log.Info("user called method delete")

		s.log.Info("remote object ID", "regnum=", mux.Vars(r)["regnum"])
		if err := s.store.Delete(mux.Vars(r)["regnum"]); err != nil {
			s.log.Error("delete failed:", "error", err)
			w.WriteHeader(http.StatusNotFound)
			errorResponse := map[string]string{
				"error":   "Resource not found",
				"message": "The resource with the specified ID does not exist.",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(errorResponse)
			return
		} else {
			s.log.Info("successful object deletion", "regnum=", mux.Vars(r)["id"])
			w.WriteHeader(http.StatusNoContent)
		}
	}
}

package service

import (
	"carcat/internal/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// @Summary Patch car by id
// @Description Update information about a car by its registration number
// @Tags cars
// @Accept json
// @Param	regnum	query	string	true	"Car Registration Number"
// @Param	jsonfile	body	models.Patch false "Patch data"
// @Success 200 "Patch request was successful"
// @Failure 400 "Bad Request - wrong JSON format or update request failed"
// @Router /cars/{regnum} [patch]
func (s *Service) patchCars() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.log.Info("user calls patch request")
		var patch models.Patch

		if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
			s.log.Error("wrong json file", "json=", r.Body)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		patch.RegNum = mux.Vars(r)["regnum"]
		if err := s.store.Update(patch); err != nil {
			s.log.Error("update request failed", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		s.log.Info("patch request was successful")
		w.WriteHeader(http.StatusOK)
	}
}

package service

import (
	"carcat/internal/models"
	"carcat/internal/storage"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

// @Summary Post car by id in external API
// @Description Post information about a car by its registration number by using external API
// @Tags cars
// @Accept json
// @Produce json
// @Param	regNums	body	ReqBody	true	"Request data"
// @Success 200 {object} SuccessResponse "All cars successfully added or some cars were added, but some registration numbers were invalid."
// @Failure 400 "Bad Request - All provided registration numbers were invalid."
// @Router /cars [post]
func (s *Service) postCars() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.log.Info("user called 'POST' request")
		var req ReqBody
		var invalidRegNums []string
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.log.Error("failed decode REGnums", "regnum=", r.Body)
			str := map[string]string{
				"error":   "Resource not found",
				"message": "The resource with the specified attributes do not exist.",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(str)
			return
		}
		var exapi bool
		for _, reg := range req.RegNums {
			car, err := getcarinfo(reg)
			if err != nil {
				var errApi *ExternalApi
				s.log.Error("Get external API:", "error", err)

				if errors.As(err, &errApi) {
					exapi = true
					break
				}
				invalidRegNums = append(invalidRegNums, reg)
				continue
			} else {
				if err := s.store.Create(car); err != nil {
					var errRegNumExists *storage.ErrRegNumExists

					if errors.As(err, &errRegNumExists) {
						s.log.Info("Element is already existes:", "error", err)
						continue
					}
					s.log.Error("Insert failed:", "error", err)
					invalidRegNums = append(invalidRegNums, car.RegNum)
				}
			}
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		var reque SuccessResponse
		if exapi {
			json.NewEncoder(w).Encode(SuccessResponse{
				Message: "External API does not work",
			})
			return
		}
		if len(invalidRegNums) != 0 {
			s.log.Debug("Invalid regNums", "regnums=", invalidRegNums)
			reque.BadRegNums = append(reque.BadRegNums, invalidRegNums...)
			reque.ErrorResponse = map[string]string{
				"error":   "Elements is already existes",
				"message": "List of already exsistes element.",
			}
			json.NewEncoder(w).Encode(reque)
			return
		}
		reque.Message = "All elements added"
		json.NewEncoder(w).Encode(reque)
		s.log.Info("successful 'POST' request")
	}
}

type ReqBody struct {
	RegNums []string `json:"regNums"`
}
type SuccessResponse struct {
	BadRegNums    []string          `json:"badRegNums,omitempty"`
	ErrorResponse map[string]string `json:"errorResponse,omitempty"`
	Message       string            `json:"message,omitempty"`
}

func getcarinfo(reg string) (models.Catalog, error) {
	var car models.Catalog
	url := fmt.Sprintf("%s/info?regNum=%s", os.Getenv("externalAPI"), reg)
	response, err := http.Get(url)
	if err != nil {
		return car, &ExternalApi{}
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusBadRequest {
		return car, fmt.Errorf("failed to fetch car info: status code %d", response.StatusCode)
	}
	if response.StatusCode == http.StatusInternalServerError {
		return car, &ExternalApi{}
	}
	if err := json.NewDecoder(response.Body).Decode(&car); err != nil {
		return car, err
	}

	return car, nil
}

type ExternalApi struct {
}

func (e *ExternalApi) Error() string {
	return "External API does not work"
}

package service

import (
	"carcat/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

type ErrNotValidData struct {
	str  string
	bnum bool
}

func (e *ErrNotValidData) Error() string {
	if e.bnum {
		return fmt.Sprintf("wrong number attribute: %s", e.str)
	}
	if e.str != "" {
		return fmt.Sprintf("wrong string attribute: %s", e.str)
	}
	return "unknown attribute in URL"
}

// @Summary GET car by id
// @Description GET information about a car by its registration number
// @Tags cars
// @Produce json
// @Param	regNum	query	string	false	"Registration Number"
// @Param	mark	query	string	false	"Mark of car"
// @Param	model	query	string	false	"model of car"
// @Param	year	query	int	false	"Year of car"
// @Param	lyear	query	int	false	"Lower limit of car definition"
// @Param	tyear	query	int	false	"Top limit of car definition"
// @Param	name	query	string	false	"Owner name"
// @Param	surname	query	string	false	"Owner surname"
// @Param	patronymic	query	string	false	"Owner patronymic"
// @Param	page	query	int	false	"Page "
// @Param	pagesize	query	int	false	"Number of elements per page"
// @Success 200 {object} SuccessGetResponse "Information of cars by using filter"
// @Failure 400 {object} map[string]string "Incorrect filtering attributes"
// @Router /cars [get]
func (s *Service) getCars() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.log.Info("user calls a GET request")
		fil, err := filterAtrib(r.URL.Query())
		if err != nil {
			handlerError(w)
			s.log.Info("user error wrong filter attribute GET cars/:", "error", err)
			return
		}
		filJSON, _ := json.Marshal(fil)
		s.log.Info("filter of GET request: ", "filter", filJSON)
		pages, err := paginat(r.URL.Query())
		if err != nil {
			handlerError(w)
			s.log.Info("user error wrong filter attribute of pagination:", "error", err)
			return
		}
		if pages.Pagesize == 0 {
			w.WriteHeader(http.StatusBadRequest)
			errorResponse := map[string]string{
				"error":   "Resource not found",
				"message": "The resource with the specified filter attributes do not exist.",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(errorResponse)
			return
		}
		cars, err := s.store.Read(fil)
		if err != nil {
			s.log.Error("reading from database: %w", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if len(cars) == 0 {
			json.NewEncoder(w).Encode(SuccessGetResponse{
				Message: "Cars with the specified attributes do not exist",
			})
			return
		}
		pages.Pages = len(cars) / pages.Pagesize
		if len(cars)%pages.Pagesize != 0 {
			pages.Pages += 1
		}
		s.log.Info("max pages of user request equal: ", "pages=", pages.Pages) //Возможно ошибка
		if pages.Pages < pages.CurPage || pages.CurPage <= 0 {
			handlerError(w)
			s.log.Info("User error: out of pages")
			return
		}
		s.log.Info("successful \"GET\" response")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		startIndex := pages.Pagesize * (pages.CurPage - 1)
		endIndex := startIndex + pages.Pagesize
		if endIndex > len(cars) {
			endIndex = len(cars)
		}
		json.NewEncoder(w).Encode(SuccessGetResponse{
			Cars: cars[startIndex:endIndex],
		})
	}
}

type SuccessGetResponse struct {
	Cars    []models.Catalog `json:"cars,omitempty"`
	Message string           `json:"message,omitempty"`
}

func handlerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	errorResponse := map[string]string{
		"error":   "Resource not found",
		"message": "The resource with the specified filter attribute does not exist.",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(errorResponse)
}
func paginat(query url.Values) (models.Pages, error) {
	var pages models.Pages
	var err error
	if query.Get("pagesize") != "" {
		pages.Pagesize, err = strconv.Atoi(query.Get("pagesize"))
		if err != nil {
			return models.Pages{}, err
		}
	} else {
		pages.Pagesize = 5
	}
	if query.Get("page") != "" {
		pages.CurPage, err = strconv.Atoi(query.Get("page"))
		if err != nil {
			return models.Pages{}, err
		}
	} else {
		pages.CurPage = 1
	}

	return pages, nil
}
func filterAtrib(query url.Values) (models.Filter, error) {
	var filter models.Filter
	var err error
	if query.Get("year") != "" {
		filter.Year, err = strconv.Atoi(query.Get("year"))
		if err != nil {
			return models.Filter{}, &ErrNotValidData{
				bnum: true,
				str:  query.Get("lyear"),
			}
		}
	}
	if query.Get("lyear") != "" {
		if filter.Year != 0 {
			return models.Filter{}, fmt.Errorf("attribute year can not exsist with lyear and tyear")
		}
		filter.Lyear, err = strconv.Atoi(query.Get("lyear"))
		if err != nil {
			return models.Filter{}, &ErrNotValidData{
				bnum: true,
				str:  query.Get("lyear"),
			}
		}
	}
	if query.Get("tyear") != "" {
		if filter.Year != 0 {
			return models.Filter{}, fmt.Errorf("attribute year can not exsist with lyear and tyear")
		}
		filter.Tyear, err = strconv.Atoi(query.Get("tyear"))
		if err != nil {
			return models.Filter{}, &ErrNotValidData{
				bnum: true,
				str:  query.Get("lyear"),
			}
		}
	}
	if filter.Tyear < 0 {
		return models.Filter{}, fmt.Errorf("attribute tyear less zero")
	}
	if filter.Lyear < 0 {
		return models.Filter{}, fmt.Errorf("attribute lyear less zero")
	}
	if filter.Tyear < filter.Lyear {
		return models.Filter{}, fmt.Errorf("attribute tyear less then lyear")
	}
	filter.Regnum = query.Get("regNum")
	filter.Mark = query.Get("mark")
	filter.Model = query.Get("model")

	filter.Name, err = checkVaildPerson(query.Get("name"))
	if err != nil {
		return models.Filter{}, err
	}
	filter.Surname, err = checkVaildPerson(query.Get("surname"))
	if err != nil {
		return models.Filter{}, err
	}
	filter.Patronymic, err = checkVaildPerson(query.Get("patronymic"))
	if err != nil {
		return models.Filter{}, err
	}

	return filter, nil
}
func checkVaildPerson(str string) (string, error) {
	if str == "" {
		return "", nil
	}
	if isValidPerson(str) {
		return str, nil
	}
	return "", &ErrNotValidData{
		str: str,
	}
}
func isValidPerson(str string) bool {
	re := regexp.MustCompile(`^[A-ZА-ЯЁ][a-zа-яё\s-]*$`)
	return re.MatchString(str)
}

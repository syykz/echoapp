package controller

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"echoapp/usecase"
)

type Sample interface {
	Handle()
}

type SampleController struct {
	e *echo.Echo
	su usecase.SampleUsecase
}

func NewSample(e *echo.Echo, su usecase.SampleUsecase) *SampleController {
	return &SampleController {
		e: e,
		su: su,
	}
}

func (s *SampleController) Handle() {
	s.e.GET("/", s.root)
	s.e.GET("/users/:id", s.getUser)
}

type SampleOutput struct {
	ID string `json:"id"`
	Context string `json:"context"`
}
type SampleOutputs struct {
	SampleOutputs []*SampleOutput `json:"sample_outputs"`
}

func (s *SampleController)root(c echo.Context) error {
	samples, err := s.su.FindAll()
	if err != nil {
		return err
	}

	output := SampleOutputs{ 		
		SampleOutputs: []*SampleOutput{},
	}

	for _, sample := range samples {
		output.SampleOutputs = append(output.SampleOutputs, &SampleOutput{
			ID: sample.ID,
			Context: sample.Name,
		})
	}

	return c.JSON(http.StatusOK, output)
}

func (s *SampleController) getUser(c echo.Context) error {
	id := c.Param("id")
	
	user, err := s.su.FindById(id)

	if err != nil {
		return err
	}

	output := SampleOutput{
		ID: user.ID,
		Context: user.Name,
	}

    return c.JSON(http.StatusOK, output)
}
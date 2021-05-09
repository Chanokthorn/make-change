package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"make-change/app/interface/internal"
	"make-change/app/interface/presenter"
	"make-change/app/usecase/interactor"
	"net/http"
)

type ChangeController struct {
	changeInteractor interactor.ChangeInteractor
	changePresenter  presenter.ChangePresenter
}

func NewChangeController(changeInteractor interactor.ChangeInteractor, changePresenter presenter.ChangePresenter) *ChangeController {
	return &ChangeController{changeInteractor: changeInteractor, changePresenter: changePresenter}
}

func (c *ChangeController) MakeChange(e echo.Context) error {
	payload := new(internal.MakeChangeRequest)
	if err := e.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	noteMap, err := c.changeInteractor.MakeChange(payload.Given, payload.Price)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("unable to calculate change: %v", err))
	}
	err = e.JSON(http.StatusOK, c.changePresenter.MakeChangeResponse(noteMap))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to parse change to JSON")
	}
	return nil
}

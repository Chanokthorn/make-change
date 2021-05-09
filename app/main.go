package main

import (
	"github.com/labstack/echo/v4"
	"make-change/app/domain"
	"make-change/app/interface/controller"
	"make-change/app/interface/presenter"
	"make-change/app/interface/repository"
	"make-change/app/usecase/interactor"
)

func main() {
	e := echo.New()
	noteRepository := repository.NewNoteRepository(domain.NoteMap{
		1000: 10,
		500:  15,
		100:  15,
		50:   20,
		20:   30,
		10:   20,
		5:    20,
		1:    20,
		0.25: 50,
	})
	changeInteractor := interactor.NewChangeInteractor(noteRepository)
	changePresenter := presenter.NewChangePresenter()
	changeController := controller.NewChangeController(changeInteractor, changePresenter)
	e.POST("/", changeController.MakeChange)
	e.Logger.Fatal(e.Start(":3000"))
}

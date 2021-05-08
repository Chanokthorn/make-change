package main

import (
	"github.com/labstack/echo/v4"
	"q-chang/app/domain"
	"q-chang/app/interface/controller"
	"q-chang/app/interface/presenter"
	"q-chang/app/interface/repository"
	"q-chang/app/usecase/interactor"
)

func main() {
	e := echo.New()
	noteRepository := repository.NewNoteRepository(domain.NoteMap{
		//1000: 10,
		//500:  0,
		//100:  0,
		//50:   0,
		//20:   0,
		//10:   20,
		//5:    10,
		//1:    5,
		5:   20,
		2.5: 10,
		0.5: 5,
		//0.25: 50,
		//1000: 10,
		//500: 15,
		//100: 15,
		//50: 20,
		//20: 30,
		//10: 20,
		//5: 20,
		//1: 20,
		//0.25: 50,
	})
	changeInteractor := interactor.NewChangeInteractor(noteRepository)
	changePresenter := presenter.NewChangePresenter()
	changeController := controller.NewChangeController(changeInteractor, changePresenter)
	e.POST("/", changeController.MakeChange)
	e.Logger.Fatal(e.Start(":3000"))
}

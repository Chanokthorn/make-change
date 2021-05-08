package presenter

import (
	"q-chang/app/domain"
	"q-chang/app/interface/internal"
)

type ChangePresenter interface {
	MakeChangeResponse(noteMap domain.NoteMap) (result internal.MakeChangeResponse)
}

type changePresenter struct{}

func NewChangePresenter() *changePresenter {
	return &changePresenter{}
}

func (c *changePresenter) MakeChangeResponse(noteMap domain.NoteMap) (result internal.MakeChangeResponse) {
	for k, v := range noteMap {
		result = append(result, *domain.NewNote(k, v))
	}
	return
}

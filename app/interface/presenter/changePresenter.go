package presenter

import (
	"make-change/app/domain"
	"make-change/app/interface/internal"
	"sort"
)

type ChangePresenter interface {
	MakeChangeResponse(noteMap domain.NoteMap) (result internal.MakeChangeResponse)
}

type changePresenter struct{}

func NewChangePresenter() *changePresenter {
	return &changePresenter{}
}

func (c *changePresenter) MakeChangeResponse(noteMap domain.NoteMap) (result internal.MakeChangeResponse) {
	var keys []float64
	for key := range noteMap {
		keys = append(keys, key)
	}
	sort.Float64s(keys)
	for _, key := range keys {
		result = append(result, *internal.NewNoteResponse(key, noteMap[key]))
	}
	return
}

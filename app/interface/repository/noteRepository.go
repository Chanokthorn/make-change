package repository

import (
	"fmt"
	"make-change/app/domain"
)

type NoteRepository interface {
	GetNoteValueToCountMap() domain.NoteMap
	ReduceNote(noteMap domain.NoteMap) error
}

type noteRepository struct {
	noteMap domain.NoteMap
}

func NewNoteRepository(noteMap domain.NoteMap) NoteRepository {
	return &noteRepository{noteMap: noteMap}
}

func (n *noteRepository) GetNoteValueToCountMap() domain.NoteMap {
	return n.noteMap.Copy()
}

func (n *noteRepository) ReduceNote(m domain.NoteMap) error {
	for k, v := range m {
		if _, ok := n.noteMap[k]; ok {
			n.noteMap[k] -= v
		} else {
			return fmt.Errorf("note %v not found", k)
		}
	}
	return nil
}

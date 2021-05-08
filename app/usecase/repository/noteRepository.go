package repository

import "q-chang/app/domain"

type NoteRepository interface {
	GetNoteValueToCountMap() domain.NoteMap
	ReduceNote(noteMap domain.NoteMap) error
}

package repository

import "make-change/app/domain"

type NoteRepository interface {
	GetNoteValueToCountMap() domain.NoteMap
	ReduceNote(noteMap domain.NoteMap) error
}

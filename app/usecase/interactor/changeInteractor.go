package interactor

import (
	"errors"
	"q-chang/app/domain"
	"q-chang/app/usecase/repository"
)

type ChangeInteractor interface {
	MakeChange(given, price float64) (domain.NoteMap, error)
}

type changeInteractor struct {
	noteRepository repository.NoteRepository
}

func NewChangeInteractor(noteRepository repository.NoteRepository) ChangeInteractor {
	return &changeInteractor{noteRepository: noteRepository}
}

type memo map[float64][]domain.NoteMap

func newMemo() memo {
	return make(map[float64][]domain.NoteMap)
}

func (m memo) get(amount float64, noteMap domain.NoteMap) (bool, domain.NoteMap) {
	if _, ok := m[amount]; !ok {
		return false, nil
	}
	for _, nm := range m[amount] {
		if nm.Feasible(noteMap) {
			return true, nm
		}
	}
	return false, nil
}

func (m memo) set(amount float64, noteMap domain.NoteMap) {
	m[amount] = append(m[amount], noteMap)
}

func (c *changeInteractor) coinChange(amount float64, noteMap, noteLimitMap domain.NoteMap, mem memo) (bool, domain.NoteMap) {
	if amount == 0 {
		return true, make(domain.NoteMap)
	}
	if ok, noteMapMemo := mem.get(amount, noteMap); ok {
		return true, noteMapMemo
	}
	for value, count := range noteMap {
		if value > amount {
			continue
		}
		if count+1 > noteLimitMap[value] {
			continue
		}
		newNoteMap := noteMap.Copy()
		newNoteMap[value] += 1
		ok, resultNoteMap := c.coinChange(amount-value, newNoteMap, noteLimitMap, mem)
		if !ok {
			continue
		}
		resultNoteMap[value] += 1
		mem.set(amount, resultNoteMap.Copy())
		return true, resultNoteMap
	}
	return false, nil
}

func (c *changeInteractor) calculateChange(amount float64) (domain.NoteMap, error) {
	noteMap := c.noteRepository.GetNoteValueToCountMap()
	mem := newMemo()
	if ok, result := c.coinChange(amount, noteMap.CopyZero(), noteMap, mem); ok {
		return result, nil
	}
	return nil, errors.New("infeasible")
}

func (c *changeInteractor) MakeChange(given, price float64) (domain.NoteMap, error) {
	amount := given - price
	return c.calculateChange(amount)
}

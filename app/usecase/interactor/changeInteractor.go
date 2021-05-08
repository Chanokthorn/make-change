package interactor

import (
	"errors"
	"math"
	"q-chang/app/domain"
	"q-chang/app/usecase/repository"
	"q-chang/app/utils"
	"sort"
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

const maxInt = math.MaxInt64 - 1

func (c *changeInteractor) getMultipliedVariables(notes []float64, amount float64) (notesMultiplied []int, amountMultiplied int) {
	for _, note := range notes {
		notesMultiplied = append(notesMultiplied, int(note*100)) //TODO
	}
	amountMultiplied = int(amount * 100) //TODO
	return
}

func (c *changeInteractor) runDP(notes []float64, limits []int, amount float64) (dp [][]int, usage [][]bool) {
	notesMultiplied, amountMultiplied := c.getMultipliedVariables(notes, amount)

	dp = utils.Create2DIntArray(len(notesMultiplied), amountMultiplied+1)
	usage = utils.Create2DBoolArray(len(notesMultiplied), amountMultiplied+1)

	for i := range notesMultiplied {
		for j := 0; j <= amountMultiplied; j++ {
			if j == 0 {
				dp[i][j] = 0
			} else {
				dp[i][j] = maxInt
			}
		}
	}

	for i, note := range notesMultiplied {
		useNoteCounts := make([]int, amountMultiplied+1)
		for val := 1; val <= amountMultiplied; val++ {
			var useNoteValid bool
			if val-note >= 0 && dp[i][val-note] != maxInt && useNoteCounts[val-note]+1 <= limits[i] {
				useNoteValid = true
			}
			if i == 0 {
				if useNoteValid {
					dp[i][val] = dp[i][val-note] + 1
					usage[i][val] = true
					useNoteCounts[val] = useNoteCounts[val-note] + 1
				}
				continue
			}
			usePreviousNoteResult := dp[i-1][val]
			if useNoteValid {
				useNoteResult := dp[i][val-note] + 1
				if usePreviousNoteResult > useNoteResult {
					dp[i][val] = useNoteResult
					usage[i][val] = true
					useNoteCounts[val] = useNoteCounts[val-note] + 1
					continue
				}
			} else {
				dp[i][val] = usePreviousNoteResult
			}
		}
	}
	return
}

func (c *changeInteractor) backtraceDP(notes []float64, amount float64, dp [][]int, usage [][]bool) (domain.NoteMap, error) {
	notesMultiplied, amountMultiplied := c.getMultipliedVariables(notes, amount)
	if dp[len(notesMultiplied)-1][amountMultiplied] == maxInt {
		return nil, errors.New("infeasible")
	}
	result := make(domain.NoteMap)
	noteIndex := len(notesMultiplied) - 1
	for {
		if amountMultiplied == 0 {
			return result, nil
		}
		if usage[noteIndex][amountMultiplied] {
			result[notes[noteIndex]] += 1
			amountMultiplied -= notesMultiplied[noteIndex]
		} else {
			noteIndex -= 1
		}
	}
}

func (c *changeInteractor) sortNotes(noteMap domain.NoteMap) (notes []float64, limits []int) {
	var keys []float64
	for key := range noteMap {
		keys = append(keys, key)
	}
	sort.Float64s(keys)
	for _, key := range keys {
		notes = append(notes, key)
		limits = append(limits, noteMap[key])
	}
	return
}

func (c *changeInteractor) MakeChange(given, price float64) (domain.NoteMap, error) {
	amount := given - price
	currentNoteMap := c.noteRepository.GetNoteValueToCountMap()
	notes, limits := c.sortNotes(currentNoteMap)
	dp, usage := c.runDP(notes, limits, amount)
	sol, err := c.backtraceDP(notes, amount, dp, usage)
	//spew.Dump(sol)
	//return nil, nil
	if err != nil {
		return nil, errors.New("infeasible")
	}
	//err = c.noteRepository.ReduceNote(sol)
	//if err != nil {
	//	return nil, fmt.Errorf("unable to reduce notes: %v", err)
	//}
	return sol, nil
}

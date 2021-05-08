package domain

type Note struct {
	Value  float64
	Amount int
}

func NewNote(value float64, amount int) *Note {
	return &Note{Value: value, Amount: amount}
}

type NoteList []Note

type NoteMap map[float64]int

func (nm *NoteMap) Add(nm2 NoteMap) {
	for k, v := range nm2 {
		(*nm)[k] += v
	}
}

func (nm *NoteMap) Subtract(nm2 NoteMap) {
	for k, v := range nm2 {
		(*nm)[k] -= v
	}
}

func (nm *NoteMap) Copy() NoteMap {
	result := make(NoteMap)
	for k, v := range *nm {
		result[k] = v
	}
	return result
}

func (nm *NoteMap) CopyZero() NoteMap {
	result := make(NoteMap)
	for k, _ := range *nm {
		result[k] = 0
	}
	return result
}

func (nm *NoteMap) Feasible(nm2 NoteMap) bool {
	for k, v := range nm2 {
		if (*nm)[k] < v {
			return false
		}
	}
	return true
}

func (nm NoteMap) GetSum() int {
	count := 0
	for _, v := range nm {
		count += v
	}
	return count
}

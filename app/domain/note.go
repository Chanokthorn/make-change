package domain

type NoteMap map[float64]int

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

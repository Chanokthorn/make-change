package internal

type NoteResponse struct {
	Note   float64 `json:"note"`
	Amount int     `json:"amount"`
}

func NewNoteResponse(value float64, amount int) *NoteResponse {
	return &NoteResponse{Note: value, Amount: amount}
}

type NoteList []NoteResponse

type MakeChangeResponse NoteList

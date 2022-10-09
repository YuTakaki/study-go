package src

type Todo struct {
	ID         int    `json:"id"`
	Todo       string `json:"todo"`
	IsComplete bool   `json:"isComplete"`
}

var TodoLists = []Todo{
	{
		ID:         1,
		Todo:       "eat",
		IsComplete: true,
	},
	{
		ID:         2,
		Todo:       "code",
		IsComplete: false,
	},
	{
		ID:         3,
		Todo:       "sleep",
		IsComplete: false,
	},
}
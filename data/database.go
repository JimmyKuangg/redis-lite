package data

func NewDatabase() *Database {
	return &Database{
		data: make(map[string]Entry),
	}
}

package data

type Database struct {
	data map[string]string
}

type Command struct {
	Name string
	Args []string
}
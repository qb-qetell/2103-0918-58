package err


// Type Error stands for Programmatic error.
type Error struct {
	id string
	descrptn string
	lowerLevelVrsion error
}


// ---- //
func Create (descrptn string) (Error) {
	return Error {"", descrptn, nil}
}

func CreateModel2 (id, descrptn string, lowerLevelVrsion error) (Error) {
	return Error {id, descrptn, lowerLevelVrsion}
}

func (t Error) GetId () (string) {
	return t.id
}

func (t Error) Error () (string) {
	return t.descrptn
}

func (t Error) Unwrap () (error) {
	return t.lowerLevelVrsion
}

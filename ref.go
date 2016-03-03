package go4git

func (r *Repository) References() ([]Reference, error) {
	return make([]Reference, 0), nil
}


type Reference struct {
	name string
	targetId string
}

func (r Reference) IsBranch() bool {
	return false
}

func (r Reference) IsTag() bool {
	return false
}

func (r Reference) IsRemote() bool {
	return false
}


func (r Reference) HasLog() bool {
	return false
}


func (r Reference) Type() string {
	return ""
}


func (r Reference) Name() string {
	return ""
}

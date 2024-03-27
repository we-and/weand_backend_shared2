package structs

type Message struct {
	Object     string
	TeamName   string
	TeamId     uint32
	Content    string
	Title      string
	ActionLink string
	RelatedId  uint32
}

func (j *Message) IsAnnouncement() bool {
	return j.Object == "ANNOUNCEMENT"
}
func (j *Message) IsPoll() bool {
	return j.Object == "POLL"
}
func (j *Message) CopyFrom(v *Message) {
	if v == nil {
		return
	}
	j.Object = (*v).Object
	j.TeamName = (*v).TeamName
	j.TeamId = (*v).TeamId
}

func (j *Message) IsFile() bool {
	return j.Object == "FILE"
}
func (j *Message) IsRSVP() bool {
	return j.Object == "RSVP"
}

func (j *Message) GetShortObject() string {
	if j.IsFile() {
		return "FIL"
	}
	if j.IsAnnouncement() {
		return "ANN"
	}
	if j.IsPoll() {
		return "POL"
	}
	if j.IsRSVP() {

		return "RSV"

	}

	return "UNK"
}

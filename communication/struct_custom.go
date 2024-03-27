package communication

type CustomerPersonMedium struct {
	PersonPid   uint32  `json:"person_pid"`
	ShortMedium string  `json:"shortmedium"`
	PhoneBase   *string `json:"phone_base"`
	Email       *string `json:"email"`
}

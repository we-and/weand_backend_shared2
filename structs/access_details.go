package structs

type Me struct {
	UserId   uint32
	PersonId uint32
	Rights   string
	UserKey  string
}

func (c *Me) HasUserKey() bool {
	return len(c.UserKey) > 0
}
func (c *Me) HasUserId() bool {
	return c.UserId != 0
}
func (c *Me) CheckPersonId(personId uint32) bool {
	return personId == c.PersonId
}
func (c *Me) CheckUserId(userId uint32) bool {
	return userId == c.UserId
}

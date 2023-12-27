package http

func (c *UserController) GetByID(id int64) (*User, error) {
	return c.usecase.GetByID(id)
}
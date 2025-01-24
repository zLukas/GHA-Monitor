package credentials

type Credentials struct{
	token string
	Repo string
	Owner string

}

func (c *Credentials)GetToken() string{
	return c.token
}


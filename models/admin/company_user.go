package admin

type User struct {
	OpenId string
}

func (User)TableName() string {
	return "company_user"
}
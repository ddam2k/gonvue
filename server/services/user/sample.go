package user

func SampleUserList() []Account {
	return []Account{
		{
			Id:       "mgkim",
			password: "1234",
			Name:     "김명길",
			Email:    "mgkim@mgkim.com",
		},
		{
			Id:       "bonwoo",
			password: "1234",
			Name:     "이본우",
			Email:    "bon@bon.com",
		},
	}
}

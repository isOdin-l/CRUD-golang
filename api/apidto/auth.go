package apidto

type SignUpAPI struct {
	Name     string
	Username string
	Password string
}

type SignInAPI struct {
	Username string
	Password string
}

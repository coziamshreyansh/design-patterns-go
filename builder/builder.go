package builder

type User struct {
	name   string
	email  string
	mobile string
}

type UserBuilder struct {
	name   string
	email  string
	mobile string
}

func CreateUserBuilderInstance() *UserBuilder {
	return &UserBuilder{}
}

func (b *UserBuilder) WithName(name string) *UserBuilder {
	b.name = name
	return b
}

func (b *UserBuilder) WithEmail(email string) *UserBuilder {
	b.email = email
	return b
}

func (b *UserBuilder) WithMobile(mobile string) *UserBuilder {
	b.mobile = mobile
	return b
}

func (b *UserBuilder) Build() User {
	return User{
		name:   b.name,
		email:  b.email,
		mobile: b.mobile,
	}
}

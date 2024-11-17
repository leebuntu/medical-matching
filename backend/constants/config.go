package constants

var DatabaseNames = map[string]string{
	UserDB: "users.db",
	// TODO: Add other databases
}

const (
	DBPath = "../db/"
	UserDB = "UserDB"
)

const (
	InternalServerError    = "내부 서버 오류"
	DuplicateUser          = "이미 존재하는 계정"
	BadRequest             = "잘못된 요청"
	WrongAccountOrPassword = "계정 또는 비밀번호가 틀렸습니다."
	WelcomeRegister        = "환영합니다!"
)

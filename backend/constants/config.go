package constants

var DatabaseNames = map[string]string{
	UserDB:     "users.db",
	HospitalDB: "hospitals.db",
	ReviewDB:   "reviews.db",
	// TODO: Add other databases
}

const (
	DBPath     = "../db/"
	UserDB     = "UserDB"
	HospitalDB = "HospitalDB"
	ReviewDB   = "ReviewDB"
)

const (
	Unauthorized               = "인증되지 않은 요청"
	InternalServerError        = "내부 서버 오류"
	DuplicateUser              = "이미 존재하는 계정"
	BadRequest                 = "잘못된 요청"
	WrongAccountOrPassword     = "계정 또는 비밀번호가 틀렸습니다."
	WelcomeRegister            = "환영합니다!"
	UpdateProfileSuccess       = "프로필 업데이트 완료"
	AddPaymentMethodSuccess    = "결제 수단 추가 완료"
	DeletePaymentMethodSuccess = "결제 수단 삭제 완료"
)

package constants

var DatabaseNames = map[string]string{
	UserDB:     "users.db",
	HospitalDB: "hospitals.db",
	ReviewDB:   "reviews.db",
	// TODO: Add other databases
}

const (
	DBPath           = "../db/"
	TestDataPath     = DBPath + "testdata/"
	UserDB           = "UserDB"
	HospitalDB       = "HospitalDB"
	ReviewDB         = "ReviewDB"
	HospitalTestData = "hospitals.csv"
	SymptomTestData  = "symptoms.csv"
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
	NotFound                   = "찾을 수 없는 리소스"
	EndMatchingSuccess         = "매칭 종료"
)

const (
	Waiting = iota + 1
	Distance
	Review
	HaveParkingLot
	LeastWalk
)

var Weights = map[int]float64{
	1: 3.0,
	2: 2.0,
	3: 1.5,
	4: 1.0,
}

const (
	TotalPriority = 5
)

const (
	PerWatingPersonScore = -5
	HaveParkingLotScore  = 30
	PerWalkMinuteScore   = -5
	PerDrivingTimeScore  = -1.66
)

const (
	BeforeMatching = iota
	StartMatching
	WhileMatching
	MatchingCompleted
	MatchingFailed
	MatchingEnded
)

const (
	ReviewPerPage = 10
	RecordPerPage = 5
)

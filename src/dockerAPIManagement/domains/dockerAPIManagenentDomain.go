package domains

// business logic
type UseCase interface {
	CreateService()
	UpdateService()
	DeleteService()
}

// อะไรเชื่อมต่อกับ DB
type Repository interface {
}

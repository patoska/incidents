package models

// var DB *gorm.DB
var (
	DB     = make(map[uint32]*Incident)
	LastID = uint32(0)
)

func InitDB() {
	
}
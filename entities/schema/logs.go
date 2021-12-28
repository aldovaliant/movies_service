package schema

type Logs struct {
	Base
	Name string `gorm:"not null"`
}

func (Logs) TableName() string {
	return "logs"
}

func (Logs) Pk() string {
	return "id"
}

func (f Logs) Ref() string {
	return f.TableName() + "(" + f.Pk() + ")"
}

func (f Logs) AddForeignKeys() {
}

func (f Logs) InsertDefaults() {
}

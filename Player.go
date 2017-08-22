package baseball

type Player struct {
	PlayerID     string    `gorm:"column:playerID;primary_key"`
	NameFirst    string    `gorm:"column:nameFirst"`
	NameLast     string    `gorm:"column:nameLast"`
	ManagerID    string    `gorm:"column:managerID"`
	HofID        string    `gorm:"column:hofID"`
	BirthYear    *int      `gorm:"column:birthYear"`
	BirthMonth   *int      `gorm:"column:birthMonth"`
	BirthDay     *int      `gorm:"column:birthDay"`
	BirthCountry string    `gorm:"column:birthCountry"`
	BirthState   string    `gorm:"column:birthState"`
	BirthCity    string    `gorm:"column:birthCity"`
	DeathYear    string    `gorm:"column:deathYear"`
	DeathMonth   string    `gorm:"column:deathMonth"`
	DeathDay     string    `gorm:"column:deathDay"`
	DeathCountry string    `gorm:"column:deathCountry"`
	DeathState   string    `gorm:"column:deathState"`
	DeathCity    string    `gorm:"column:deathCity"`
	NameNote     string    `gorm:"column:nameNote"`
	NameGiven    string    `gorm:"column:nameGiven"`
	NameNick     string    `gorm:"column:nameNick"`
	Weight       *float32  `gorm:"column:weight"`
	Height       *float32  `gorm:"column:height"`
	Bats         string    `gorm:"column:bats"`
	Throws       string    `gorm:"column:throws"`
	Debut        string    `gorm:"column:debut"`
	FinalGame    string    `gorm:"column:finalGame"`
	College      string    `gorm:"column:college"`
	Lahman40ID   string    `gorm:"column:lahman40ID"`
	Lahman45ID   string    `gorm:"column:lahman45ID"`
	RetroID      string    `gorm:"column:retroID"`
	HoltzID      string    `gorm:"column:holtzID"`
	BbrefID      string    `gorm:"column:bbrefID"`
	BattingData  []Batting `gorm:"ForeignKey:playerID"json:"batting_data,omitempty"`
	URL          string    `gorm:"-"`
	Average      float32   `gorm:"-"`
}

func (Player) TableName() string {
	return "Master"
}

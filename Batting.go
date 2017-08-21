package main

type Batting struct {
	PlayerID   string `gorm:"column:playerID;primary_key"`
	YearID     string `gorm:"column:yearID;primary_key"`
	Stint      string `gorm:"column:stint;primary_key"`
	TeamID     string `gorm:"column:teamID"`
	LgID       string `gorm:"column:lgID"`
	G          int    `gorm:"column:G"`
	G_batting  int    `gorm:"column:G_batting"`
	AB         int    `gorm:"column:AB"`
	R          int    `gorm:"column:R"`
	H          int    `gorm:"column:H"`
	SecondBase int    `gorm:"column:2B"`
	ThirdBase  int    `gorm:"column:3B"`
	HR         int    `gorm:"column:HR"`
	RBI        int    `gorm:"column:RBI"`
	SB         int    `gorm:"column:SB"`
	CS         int    `gorm:"column:CS"`
	BB         int    `gorm:"column:BB"`
	SO         int    `gorm:"column:SO"`
	IBB        int    `gorm:"column:IBB"`
	HBP        int    `gorm:"column:HBP"`
	SH         int    `gorm:"column:SH"`
	SF         int    `gorm:"column:SF"`
	GIDP       int    `gorm:"column:GIDP"`
	G_old      int    `gorm:"column:G_old"`
}

func (Batting) TableName() string {
	return "Batting"
}

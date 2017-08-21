package main

import "fmt"
import "os"
import "github.com/jinzhu/gorm"

// import "encoding/json"
import "strings"
import _ "github.com/go-sql-driver/mysql"

// import "strconv"

import "github.com/gin-gonic/gin"

const HOST = "http://localhost:3000"

func Database() *gorm.DB {
	db, err := gorm.Open("mysql", "robertwilkinson:@/baseball?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	return db

}

type Batting struct {
	PlayerID  string `gorm:"column:playerID;primary_key"`
	YearID    string `gorm:"column:yearID;primary_key"`
	Stint     string `gorm:"column:stint;primary_key"`
	TeamID    string `gorm:"column:teamID"`
	LgID      string `gorm:"column:lgID"`
	G         string `gorm:"column:G"`
	G_batting string `gorm:"column:G_batting"`
	AB        string `gorm:"column:AB"`
	R         string `gorm:"column:R"`
	Player    Player `gorm:"ForeignKey:playerID"json:"name,omitempty"`
}

type Player struct {
	NameFirst    string    `gorm:"column:nameFirst"`
	NameLast     string    `gorm:"column:nameLast"`
	LahmanID     *uint     `gorm:"column:lahmanID",primary_key`
	PlayerID     string    `gorm:"column:playerID"`
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
}

func (Player) TableName() string {
	return "master"
}
func (Batting) TableName() string {
	return "Batting"
}

func CreateSearchQuery(c *gin.Context) string {
	err := c.Request.ParseForm()

	if err != nil {
		fmt.Println(err)
	}
	form := c.Request.Form
	acceptable_keys := map[string]bool{"first_name": true, "last_name": true}
	key_conversion := map[string]string{"first_name": "nameFirst", "last_name": "nameLast"}
	query_base := ""

	counter := 0
	for key, value := range form {
		if counter > 0 {
			query_base = query_base + " AND "
		}
		counter = counter + 1
		fmt.Println(key)
		fmt.Println(value)
		if acceptable_keys[key] {
			converted_key := key_conversion[key]
			// check if the value is an array to do IN
			if len(value) > 1 {
				// its an array
				query_base = query_base + fmt.Sprintf(" %v in (%v) ", converted_key, strings.Join(value, "', '"))

			} else {
				query_base = query_base + fmt.Sprintf(" %v = '%v' ", converted_key, value[0])
			}
		}
	}
	return query_base
	//    fmt.printLn(string(c.Request.Form))
}

func FindPlayer(c *gin.Context) {
	db := Database()
	// var players []Player
	query := CreateSearchQuery(c)
	fmt.Println(query)

	//	if err := db.Where(query).Find(&players).Related(&batting, "playerID"); err != nil {
	//		fmt.Println(err)
	//	}

	// 	var player Player
	// 	var batting Batting
	// 	db.Debug().First(&batting)
	// 	d := db.Debug().Model(&batting).Related(&player, "Player")
	// 	fmt.Println(d)
	// 	fmt.Println(player)
	// 	batting.Player = player
	var player Player
	var batting []Batting
	db.Debug().First(&player)
	d := db.Debug().Model(&player).Related(&batting, "playerID")
	fmt.Println(d)
	fmt.Println(player)
	player.BattingData = batting

	c.JSON(200, player)

}

func createUrls(players []Player) []Player {
	for i := range players {
		players[i].URL = fmt.Sprint(HOST, "/", players[i].LahmanID)
	}
	return players
}

// func GetPlayer(c *gin.Context) {
// 	db := Database()
// 	fmt.Println(c.Param("id"))
// 	var (
// 		nameFirst, nameLast                                string
// 		games, games_batting, at_bats, runs, hits, doubles int
// 	)
// 	batters := make([]gin.H, 0)
// 	var current_player Player
// 	//querynameFirst := c.Request.Form["firstname"]
// 	//querynameLast := c.Request.Form["lastname"]
// 	queryID := c.Param("id")
//
// 	rows, err := db.Raw("select master.nameFirst, master.nameLast, batting.yearID, batting.teamID, batting.lgID, batting.G, batting.G_batting, batting.AB, batting.R, batting.H, batting.2B  from Batting  JOIN master ON master.playerID=batting.playerID WHERE master.lahmanID=?;", queryID).Rows()
//
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		rows.Scan(&current_player.FirstName, &current_player.LastName, &current_player.YearID, &current_player.TeamID, &current_player.LgID, &current_player.Games, &current_player.GamesBatting, &current_player.AtBats, &current_player.Runs, &current_player.Hits, &current_player.Doubles)
// 		name := gin.H{"firstname": nameFirst, "lastname": nameLast}
// 		batting := gin.H{"games": games, "games_batting": games_batting, "at_bats": at_bats, "runs": runs, "hits": hits, "doubles": doubles}
// 		s := gin.H{"batting": batting, "name": name}
//
// 		//			player.firstname = nameFirst
// 		//			player.lastname = nameLast
// 		//			player.weight = weight
// 		//			player.height = height
// 		//			player.year_born = birthYear
// 		//			output += nameFirst + ", " + nameLast + "Weight: " + strconv.Itoa(weight) + " Height: " + strconv.Itoa(height) + " Born: " + strconv.Itoa(birthYear) + "\n"
// 		batters = append(batters, s)
// 	}
// 	//		c.JSON(200, gin.H{"firstname": player.firstname, "lastname": player.lastname, "weight": player.weight, "height": player.height, "year_born": player.year_born})
// 	c.JSON(200, batters)
// }
func main() {
	db := Database()
	db.AutoMigrate(&Player{})
	db.AutoMigrate(&Batting{})

	r := gin.Default()
	//	r.GET("/players/:id", GetPlayer)
	r.GET("/search", FindPlayer)
	r.Run(":" + os.Getenv("PORT"))

}

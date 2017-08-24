package baseball

import "fmt"
import "github.com/jinzhu/gorm"

// import "encoding/json"
import "os"
import "strings"
import "log"
import _ "github.com/go-sql-driver/mysql"

// import "strconv"

import "github.com/gin-gonic/gin"

func Database() *gorm.DB {
	var db_string string
	if os.Getenv("ENV") == "production" {
		connectionName := os.Getenv("CLOUDSQL_CONNECTION_NAME")
		user := os.Getenv("CLOUDSQL_USER")
		password := os.Getenv("CLOUDSQL_PASSWORD")
		db_string = fmt.Sprintf("%s:%s@cloudsql(%s)/lahman2016", user, password, connectionName)
	} else {
		db_string = "root:@/lahman2016?charset=utf8&parseTime=True&loc=Local"
	}
	log.Printf("db_string: %v", db_string)
	db, err := gorm.Open("mysql", db_string)
	if err != nil {
		fmt.Println(err)
	}
	return db

}

// func (player *Player) AfterFind(scope *gorm.Scope) {
// 	fmt.Println("after find")
// 	// for _, field := range scope.Fields() {
// 	// 	fmt.Println(field.Field)
// 	// }
// 	for _, field := range scope.GetStructFields() {
// 		if field.DBName == "playerID" {
// 			fmt.Printf("%+v\n", field)
// 		}
// 	}

// }
// func (batting *Batting) AfterFind(scope *gorm.Scope) {
// 	for _, field := range scope.Fields() {
// 		fmt.Println(field.Field)
// 	}
// 	for _, field := range scope.GetStructFields() {
// 		if field.DBName == "playerID" {
// 			fmt.Printf("%+v\n", field)
// 		}
// 	}

// }

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
	var players []Player
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
	db.Debug().Where(query).Find(&players)

	c.JSON(200, createUrls(players))

}

func createUrls(players []Player) []Player {

	HOST := os.Getenv("HTTP_ORIGIN")
	for i := range players {
		players[i].URL = fmt.Sprint(HOST, "/players/", players[i].PlayerID)
	}
	return players
}

func GetPlayer(c *gin.Context) {
	db := Database()
	fmt.Println(c.Param("id"))
	var player Player
	var batting []Batting
	//querynameFirst := c.Request.Form["firstname"]
	//querynameLast := c.Request.Form["lastname"]
	fmt.Println(db.HasTable(&player))
	queryID := c.Param("id")
	fmt.Println(queryID)

	db.Debug().Where("playerID = ?", queryID).Find(&player)
	db.Debug().Model(&player).Related(&batting, "playerID")
	player.BattingData = batting
	hits := 0
	at_bats := 0
	for _, battingData := range batting {
		at_bats = at_bats + battingData.AB
		hits = hits + battingData.H
	}
	avg := float32(hits) / float32(at_bats)
	player.Average = avg

	c.JSON(200, player)
}
func Test(c *gin.Context) {
	db := Database()
	var one, two, three, four string
	db.Debug().Table("Master").Select("playerID").Row().Scan(&one, &two, &three, &four)
	c.String(200, one)
}

func Router() *gin.Engine {
	db := Database()
	db.AutoMigrate(&Player{})
	db.AutoMigrate(&Batting{})

	r := gin.New()
	r.GET("/", Test)
	r.GET("/players/:id", GetPlayer)
	r.GET("/search", FindPlayer)
	return r
}

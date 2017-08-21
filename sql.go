package main

import "database/sql"
import "fmt"
import "os"

// import "encoding/json"
import "strings"
import _ "github.com/go-sql-driver/mysql"

// import "strconv"

import "github.com/gin-gonic/gin"

func Database() *sql.DB {
	db, err := sql.Open("mysql", "robertwilkinson:@/baseball")
	if err != nil {
		fmt.Println(err)
	}
	return db

}

type Player struct {
	LahmanID     string `json:"ID"`
	FirstName    string `json:"nameFirst"`
	LastName     string `json:"nameLast"`
	Weight       string `json: "weight" `
	Height       string `json: "height"`
	BirthYear    string `json: "birthYear"`
	Url          string `json: "url"`
	TeamID       int
	YearID       int
	LgID         int
	Games        int
	GamesBatting int
	AtBats       int
	Runs         int
	Hits         int
	Doubles      int
}

func CreateSearchQuery(c *gin.Context) string {
	err := c.Request.ParseForm()

	if err != nil {
		fmt.Println(err)
	}
	form := c.Request.Form
	acceptable_keys := map[string]bool{"first_name": true, "last_name": true}
	key_conversion := map[string]string{"first_name": "nameFirst", "last_name": "nameLast"}
	query_base := "select lahmanID, nameFirst, nameLast, weight, height, birthYear  from master where "

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

	query := CreateSearchQuery(c)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}

	players := make([]Player, 0)
	defer rows.Close()
	for rows.Next() {
		var current_player Player
		rows.Scan(&current_player.LahmanID, &current_player.FirstName, &current_player.LastName, &current_player.Weight, &current_player.Height, &current_player.BirthYear)
		url := fmt.Sprintf("%v/players/%v", "http://localhost:3000", current_player.LahmanID)
		current_player.Url = url
		players = append(players, current_player)
	}
	//		c.JSON(200, gin.H{"firstname": player.firstname, "lastname": player.lastname, "weight": player.weight, "height": player.height, "year_born": player.year_born})
	c.JSON(200, players)

}

func GetPlayer(c *gin.Context) {
	db := Database()
	fmt.Println(c.Param("id"))
	var (
		nameFirst, nameLast                                string
		games, games_batting, at_bats, runs, hits, doubles int
	)
	batters := make([]gin.H, 0)
	var current_player Player
	//querynameFirst := c.Request.Form["firstname"]
	//querynameLast := c.Request.Form["lastname"]
	queryID := c.Param("id")

	rows, err := db.Query("select master.nameFirst, master.nameLast, batting.yearID, batting.teamID, batting.lgID, batting.G, batting.G_batting, batting.AB, batting.R, batting.H, batting.2B  from Batting  JOIN master ON master.playerID=batting.playerID WHERE master.lahmanID=?;", queryID)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&current_player.FirstName, &current_player.LastName, &current_player.YearID, &current_player.TeamID, &current_player.LgID, &current_player.Games, &current_player.GamesBatting, &current_player.AtBats, &current_player.Runs, &current_player.Hits, &current_player.Doubles)
		name := gin.H{"firstname": nameFirst, "lastname": nameLast}
		batting := gin.H{"games": games, "games_batting": games_batting, "at_bats": at_bats, "runs": runs, "hits": hits, "doubles": doubles}
		s := gin.H{"batting": batting, "name": name}

		//			player.firstname = nameFirst
		//			player.lastname = nameLast
		//			player.weight = weight
		//			player.height = height
		//			player.year_born = birthYear
		//			output += nameFirst + ", " + nameLast + "Weight: " + strconv.Itoa(weight) + " Height: " + strconv.Itoa(height) + " Born: " + strconv.Itoa(birthYear) + "\n"
		batters = append(batters, s)
	}
	//		c.JSON(200, gin.H{"firstname": player.firstname, "lastname": player.lastname, "weight": player.weight, "height": player.height, "year_born": player.year_born})
	c.JSON(200, batters)
}
func main() {
	r := gin.Default()
	r.GET("/players/:id", GetPlayer)
	r.GET("/search", FindPlayer)
	r.Run(":" + os.Getenv("PORT"))

}

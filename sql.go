package main

import "database/sql"
import "fmt"
import "os"
import "strings"
import _ "github.com/go-sql-driver/mysql"

// import "strconv"

import "github.com/gin-gonic/gin"

func main() {
	//	type player struct {
	//		firstname string `json:"firstname"`
	//		lastname  string `json:"lastname"`
	//		weight    int    `json:"weight"`
	//		height    int    `json:"height"`
	//		year_born int    `json:"year_born"` //	}
	var (
		output                              string
		lahmanID  int
	)
	db, err := sql.Open("mysql", "robertwilkinson:@/baseball")
	if err != nil {
		fmt.Println(err)
	}
	r := gin.Default()
	r.GET("/player", func(c *gin.Context) {
		var (
			nameFirst, nameLast                                                      string
			teamID, yearID, lgID, games, games_batting, at_bats, runs, hits, doubles int
		)
		output = ""
		batters := make([]gin.H, 0)
		//querynameFirst := c.Request.Form["firstname"]
		//querynameLast := c.Request.Form["lastname"]
		c.Request.ParseForm()
		queryID := c.Request.Form.Get("id")
		//		querynameFirst := c.Params.ByName("querynameFirst")
		//		querynameLast := c.Params.ByName("querynameLast")

		rows, err := db.Query("select master.nameFirst, master.nameLast, batting.yearID, batting.teamID, batting.lgID, batting.G, batting.G_batting, batting.AB, batting.R, batting.H, batting.2B  from Batting  JOIN master ON master.playerID=batting.playerID WHERE master.lahmanID=?;", queryID)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&nameFirst, &nameLast, &yearID, &teamID, &lgID, &games, &games_batting, &at_bats, &runs, &hits, &doubles)
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
	})
	r.GET("/search", func(c *gin.Context) {
    err := c.Request.ParseForm()
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
      if(acceptable_keys[key]) {
        converted_key := key_conversion[key]
        // check if the value is an array to do IN
        if(len(value) > 1) {
          // its an array
          query_base = query_base + fmt.Sprintf(" %v in (%v) ", converted_key, strings.Join(value, "', '"))

        } else {
          query_base = query_base + fmt.Sprintf(" %v = '%v' ",converted_key, value[0])
        }
      }
    }
//    fmt.printLn(string(c.Request.Form))
    fmt.Println(query_base)

		rows, err := db.Query(query_base)
		if err != nil {
			fmt.Println(err)
		}

      type player struct {
        lahmanID string `json:"ID"`
        nameFirst string `json:"nameFirst"`
        nameLast string `json:"nameLast"`
        weight string
        height string
        birthYear string
        url string
      }
		players := make([]player, 0)
		defer rows.Close()
		for rows.Next() {
      var current_player player
			rows.Scan(&current_player.lahmanID, &current_player.nameFirst, &current_player.nameLast, &current_player.weight, &current_player.height, &current_player.birthYear)
      url := fmt.Sprintf("%v/players/%v", "localhost:3000", lahmanID)
      current_player.url = url
			//			player.firstname = nameFirst
			//			player.lastname = nameLast
			//			player.weight = weight
			//			player.height = height
			//			player.year_born = birthYear
			players = append(players, current_player)
			//			output += nameFirst + ", " + nameLast + "Weight: " + strconv.Itoa(weight) + " Height: " + strconv.Itoa(height) + " Born: " + strconv.Itoa(birthYear) + "\n"
		}
		//		c.JSON(200, gin.H{"firstname": player.firstname, "lastname": player.lastname, "weight": player.weight, "height": player.height, "year_born": player.year_born})
    fmt.Println(players[0])
		c.JSON(200, players[0])
	})
  r.Run(":" + os.Getenv("PORT"))

}

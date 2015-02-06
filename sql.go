package main

import "database/sql"
import "fmt"
import _ "github.com/go-sql-driver/mysql"

// import "strconv"

import "github.com/gin-gonic/gin"

func main() {
	type player struct {
		firstname string `json:"firstname"`
		lastname  string `json:"lastname"`
		weight    int    `json:"weight"`
		height    int    `json:"height"`
		year_born int    `json:"year_born"`
	}
	var (
		nameFirst                           string
		nameLast                            string
		output                              string
		lahmanID, weight, height, birthYear int
	)
	db, err := sql.Open("mysql", "root:@/baseball")
	if err != nil {
		fmt.Println(err)
	}
	r := gin.Default()
	r.GET("/player", func(c *gin.Context) {
		output = ""
		//querynameFirst := c.Request.Form["firstname"]
		//querynameLast := c.Request.Form["lastname"]
		c.Request.ParseForm()
		querynameFirst := c.Request.Form.Get("firstname")
		querynameLast := c.Request.Form.Get("lastname")
		//		querynameFirst := c.Params.ByName("querynameFirst")
		//		querynameLast := c.Params.ByName("querynameLast")

		rows, err := db.Query("select lahmanID, nameFirst, nameLast, weight, height, birthYear  from master where  nameFirst= ? AND nameLast= ? ", querynameFirst, querynameLast)
		if err != nil {
			fmt.Println(err)
		}
		players := make([]gin.H, 0)
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&lahmanID, &nameFirst, &nameLast, &weight, &height, &birthYear)
			s := gin.H{"ID": lahmanID, "firstname": nameFirst, "lastname": nameLast, "weight": weight, "height": height, "birthyear": birthYear}

			//			player.firstname = nameFirst
			//			player.lastname = nameLast
			//			player.weight = weight
			//			player.height = height
			//			player.year_born = birthYear
			players = append(players, s)
			//			output += nameFirst + ", " + nameLast + "Weight: " + strconv.Itoa(weight) + " Height: " + strconv.Itoa(height) + " Born: " + strconv.Itoa(birthYear) + "\n"
		}
		//		c.JSON(200, gin.H{"firstname": player.firstname, "lastname": player.lastname, "weight": player.weight, "height": player.height, "year_born": player.year_born})
		c.JSON(200, players)
	})
	r.Run(":8080")

}

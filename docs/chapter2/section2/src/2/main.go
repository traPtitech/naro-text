package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type City struct {
	ID          int    `json:"id,omitempty"  db:"ID"`
	Name        string `json:"name,omitempty"  db:"Name"`
	CountryCode string `json:"countryCode,omitempty"  db:"CountryCode"`
	District    string `json:"district,omitempty"  db:"District"`
	Population  int    `json:"population,omitempty"  db:"Population"`
}

type Me struct {//[!code ++]
	Username string `json:"username,omitempty"  db:"username"`//[!code ++]
}//[!code ++]

var (
	db *sqlx.DB
)

func main() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}

	conf := mysql.Config{
		User:      os.Getenv("DB_USERNAME"),
		Passwd:    os.Getenv("DB_PASSWORD"),
		Net:       "tcp",
		Addr:      os.Getenv("DB_HOSTNAME") + ":" + os.Getenv("DB_PORT"),
		DBName:    os.Getenv("DB_DATABASE"),
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	_db, err := sqlx.Open("mysql", conf.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("conntected")
	db = _db
	
	e := echo.New()

	e.GET("/cities/:cityName", getCityInfoHandler)
	e.POST("/cities", postCityHandler) 

	withLogin := e.Group("")//[!code ++]
		withLogin.Use(checkLogin)//[!code ++]
		withLogin.GET("/cities/:cityName", getCityInfoHandler)//[!code ++]
		withLogin.GET("/whoami", getWhoAmIHandler)//[!code ++]

	e.Start(":8080")
}


func getCityInfoHandler(c echo.Context) error {
	cityName := c.Param("cityName")
	fmt.Println(cityName)

	var city City
	if err := db.Get(&city, "SELECT * FROM city WHERE Name=?", cityName); errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("No such city Name = %s", cityName))
	} else if err != nil {
		log.Fatalf("failed to get city: %s", err)
	}

	return c.JSON(http.StatusOK, city)
}


func postCityHandler(c echo.Context) error { 
	var city City        
	err := c.Bind(&city) 
	if err != nil {      
		return echo.NewHTTPError(http.StatusBadRequest, "bad request body") 
	} 
	
	result, err := db.Exec("INSERT INTO city (Name, CountryCode, District, Population) VALUES (?, ?, ?, ?)", city.Name, city.CountryCode, city.District, city.Population) 
	if err != nil {                                                                                                                                                       
		log.Fatalf("failed to insert city data: %s", err) 
	} 
	
	id, _ := result.LastInsertId() 
	city.ID = int(id)              
	
	return c.JSON(http.StatusCreated, city) 
} 

func checkLogin(next echo.HandlerFunc) echo.HandlerFunc {//[!code ++]
	return func(c echo.Context) error {//[!code ++]
		sess, err := session.Get("sessions", c)//[!code ++]
		if err != nil {//[!code ++]
			fmt.Println(err)//[!code ++]
			return c.String(http.StatusInternalServerError, "something wrong in getting session")//[!code ++]
		}//[!code ++]

		if sess.Values["userName"] == nil {//[!code ++]
			return c.String(http.StatusForbidden, "please login")//[!code ++]
		}//[!code ++]
		c.Set("userName", sess.Values["userName"].(string))//[!code ++]

		return next(c)//[!code ++]
	}//[!code ++]
}//[!code ++]

func getWhoAmIHandler(c echo.Context) error {//[!code ++]
	sess, _ := session.Get("sessions", c)//[!code ++]
  //[!code ++]
	return c.JSON(http.StatusOK, Me{//[!code ++]
		Username: sess.Values["userName"].(string),//[!code ++]
	})//[!code ++]
}//[!code ++]

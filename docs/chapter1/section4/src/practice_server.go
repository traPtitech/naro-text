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
	fmt.Println("connected")
	db = _db
	//#region echo
	e := echo.New()

	e.GET("/cities/:cityName", getCityInfoHandler)
	e.POST("/cities", postCityHandler) //[!code ++]

	e.Start(":8080")
}

// #endregion echo
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

// #region func
func postCityHandler(c echo.Context) error { //[!code ++]
	var city City        //[!code ++]
	err := c.Bind(&city) //[!code ++]
	if err != nil {      //[!code ++]
		return echo.NewHTTPError(http.StatusBadRequest, "bad request body") //[!code ++]
	} //[!code ++]
	//[!code ++]
	result, err := db.Exec("INSERT INTO city (Name, CountryCode, District, Population) VALUES (?, ?, ?, ?)", city.Name, city.CountryCode, city.District, city.Population) //[!code ++]
	if err != nil {                                                                                                                                                       //[!code ++]
		log.Printf("DB Error: %s", err)                                                   //[!code ++]
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error") //[!code ++]
	} //[!code ++]
	//[!code ++]
	id, _ := result.LastInsertId() //[!code ++]
	city.ID = int(id)              //[!code ++]
	//[!code ++]
	return c.JSON(http.StatusCreated, city) //[!code ++]
} //[!code ++]
//#endregion func

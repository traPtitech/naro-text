package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Asia%%2FTokyo&charset=utf8mb4",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	_db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("conntected")
	db = _db
//#region echo
	e := echo.New()

	e.GET("/cities/:cityName", getCityInfoHandler)
	e.POST("/cities", postCityHandler) //[!code ++]

	e.Start(":3000")
}
//#endregion echo
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
//#region func
func postCityHandler(c echo.Context) error { //[!code ++]
	var city City //[!code ++]
	err := c.Bind(&city) //[!code ++]
	if err != nil { //[!code ++]
		return echo.NewHTTPError(http.StatusBadRequest, "bad request body") //[!code ++]
	} //[!code ++]
 //[!code ++]
	result, err := db.Exec("INSERT INTO city (Name, CountryCode, District, Population) VALUES (?, ?, ?, ?)", city.Name, city.CountryCode, city.District, city.Population) //[!code ++]
	if err != nil { //[!code ++]
		log.Fatalf("failed to insert city data: %s", err) //[!code ++]
	} //[!code ++]
 //[!code ++]
	id, _ := result.LastInsertId() //[!code ++]
	city.ID = int(id) //[!code ++]
 //[!code ++]
	return c.JSON(http.StatusCreated, city) //[!code ++]
} //[!code ++]
//#endregion func
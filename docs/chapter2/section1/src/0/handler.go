package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type City struct {
	ID          int            `json:"id,omitempty"  db:"ID"`
	Name        sql.NullString `json:"name,omitempty"  db:"Name"`
	CountryCode sql.NullString `json:"countryCode,omitempty"  db:"CountryCode"`
	District    sql.NullString `json:"district,omitempty"  db:"District"`
	Population  sql.NullInt64  `json:"population,omitempty"  db:"Population"`
}

func getCityInfoHandler(c echo.Context) error {
	cityName := c.Param("cityName")
	fmt.Println(cityName)

	city := City{}
	db.Get(&city, "SELECT * FROM city WHERE Name=?", cityName)
	if !city.Name.Valid {
		return c.NoContent(http.StatusNotFound)
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

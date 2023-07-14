package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/srinathgs/mysqlstore"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"
)

type City struct {
	ID          int            `json:"id,omitempty"  db:"ID"`
	Name        sql.NullString `json:"name,omitempty"  db:"Name"`
	CountryCode sql.NullString `json:"countryCode,omitempty"  db:"CountryCode"`
	District    sql.NullString `json:"district,omitempty"  db:"District"`
	Population  sql.NullInt64  `json:"population,omitempty"  db:"Population"`
}

var (
	db   *sqlx.DB
	salt = ""
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

	salt = os.Getenv("HASH_SALT")

	_db, err := sqlx.Open("mysql", conf.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	db = _db

	// #region setup_table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (Username VARCHAR(255) PRIMARY KEY, HashedPass VARCHAR(255))")

	if err != nil {
		log.Fatal(err)
	}
	// #endregion setup_table

	// #region setup_session
	store, err := mysqlstore.NewMySQLStoreFromConnection(db.DB, "sessions", "/", 60*60*24*14, []byte("secret-token"))
	if err != nil {
		log.Fatal(err)
	}
	// #endregion setup_session

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(session.Middleware(store))

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.POST("/login", loginHandler)
	e.POST("/signup", signUpHandler)

	withAuth := e.Group("")
	withAuth.Use(userAuthMiddleware)
	withAuth.GET("/cities/:cityName", getCityInfoHandler)
	withAuth.POST("/cities", postCityHandler)
	withAuth.GET("/whoami", getWhoAmIHandler)

	e.Start(":8080")
}

type LoginRequestBody struct {
	Username string `json:"username,omitempty" form:"username"`
	Password string `json:"password,omitempty" form:"password"`
}

type User struct {
	Username   string `json:"username,omitempty"  db:"Username"`
	HashedPass string `json:"-"  db:"HashedPass"`
}

type Me struct {
	Username string `json:"username,omitempty"  db:"username"`
}

func signUpHandler(c echo.Context) error {
	// #region request
	req := LoginRequestBody{}
	c.Bind(&req)
	// #endregion request
	// #region valid

	if req.Password == "" || req.Username == "" {
		return c.String(http.StatusBadRequest, "Username or Password is empty")
	}
	// #endregion valid
	// #region hash
	pw := req.Password + salt

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	// #endregion hash

	// ユーザーの存在チェック
	var count int

	// #region check_user
	err = db.Get(&count, "SELECT COUNT(*) FROM users WHERE Username=?", req.Username)
	if err != nil {
		log.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	if count > 0 {
		return c.String(http.StatusConflict, "Username is already used")
	}
	// #endregion check_user
	// #region add_user
	_, err = db.Exec("INSERT INTO users (Username, HashedPass) VALUES (?, ?)", req.Username, hashedPass)
	if err != nil {
		log.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusCreated)
	// #endregion add_user
}

func loginHandler(c echo.Context) error {
	// #region post_req
	req := LoginRequestBody{}
	c.Bind(&req)
	// #endregion post_req
	user := User{}
	err := db.Get(&user, "SELECT * FROM users WHERE username=?", req.Username)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %v", err))
	}
	// #endregion post_req
	// #region post_hash
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPass), []byte(req.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return c.NoContent(http.StatusForbidden)
		} else {
			return c.NoContent(http.StatusInternalServerError)
		}
	}
	// #endregion post_hash
	// #region add_session
	sess, err := session.Get("sessions", c)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "something wrong in getting session")
	}
	sess.Values["userName"] = req.Username
	sess.Save(c.Request(), c.Response())

	return c.NoContent(http.StatusOK)
	// #endregion add_session
}

func userAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	// #region get_session
	return func(c echo.Context) error {
		sess, err := session.Get("sessions", c)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "something wrong in getting session")
		}
		// #endregion get_session
		// #region check_session
		if sess.Values["userName"] == nil {
			return c.String(http.StatusUnauthorized, "please login")
		}
		c.Set("userName", sess.Values["userName"].(string))

		return next(c)
	}
	// #endregion check_session
}

func getCityInfoHandler(c echo.Context) error {
	cityName := c.Param("cityName")

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

// #region whoami
func getWhoAmIHandler(c echo.Context) error {
	sess, _ := session.Get("sessions", c)

	return c.JSON(http.StatusOK, Me{
		Username: sess.Values["userName"].(string),
	})
}

// #endregion whoami

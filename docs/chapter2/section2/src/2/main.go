e := echo.New()
e.Use(middleware.Logger())
e.Use(session.Middleware(store))

e.POST("/login", loginHandler)
e.POST("/signup", signUpHandler)
e.GET("/ping", func (c echo.Context) error { return c.String(http.StatusOK,"pong")}) //[!code ++]

withAuth := e.Group("")
withAuth.Use(userAuthMiddleware)
withAuth.GET("/cities/:cityName", getCityInfoHandler)
withAuth.POST("/cities", postCityHandler)
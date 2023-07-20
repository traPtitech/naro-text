e.Use(session.Middleware(store))
e.Use(middleware.CORSWithConfig(middleware.CORSConfig{//[!code ++]
	AllowOrigins: []string{"http://localhost:5173"},//[!code ++]
	AllowMethods: []string{http.MethodGet, http.MethodPost},//[!code ++]
}))//[!code ++]

e.POST("/login", loginHandler)
e.POST("/signup", signUpHandler)
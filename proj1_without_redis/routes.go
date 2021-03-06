// routes.go

package main

import (
	"github.com/gin-gonic/gin"
)

func initializeRoutes() {

	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(setUserStatus())

	// Handle the index route
	router.GET("/", showIndexPage)

	// Group user related routes together
	userRoutes := router.Group("/")
	{
		// Handle the GET requests at /signin
		// Show the login page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("signin", ensureNotLoggedIn(), showLoginPage)

		// Handle POST requests at /profile
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("profile", ensureNotLoggedIn(), performLogin)

		// Handle GET requests at /logout
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("logout", ensureLoggedIn(), logout)

    // Handle GET request at /profile
		userRoutes.GET("profile", viewProfile)

		// Handle the GET requests at /signup
		// Show the registration page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("signup", ensureNotLoggedIn(), showRegistrationPage)

		// Handle POST requests at /signup
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("signup", ensureNotLoggedIn(), register)
	}
}

func showIndexPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Home Page",
	}, "index.html")
}

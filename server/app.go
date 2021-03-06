package main

import (
	"log"

	"github.com/alfuhigi/gofiber-blog-example/middleware"

	"github.com/alfuhigi/gofiber-blog-example/db"
	"github.com/alfuhigi/gofiber-blog-example/handlers"
	"github.com/alfuhigi/gofiber-blog-example/providers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
)

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Pass      string `json:"pass"`
	FirstName string `josn:"firstname"`
	LastName  string `josn:"lastname"`
	Email     string `json:"email"`
	IsActive  bool   `json:"-"`
}
type Post struct {
	*User
	Title   string `json:"title"`
	Date    func() string
	Content string `json:"content"`
}

var posts = []Post{}

var context = fiber.Map{"Title": "Blog"}
var title = "عنوان الموقع"

func main() {

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Static("/static", "public")

	SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

func SetupRoutes(a *fiber.App) {
	entity := db.NewEntity(providers.Connect())
	hn := handlers.NewHandler(entity)
	app := a.Group("/")
	app.Get("/login", hn.Login)
	app.Post("/login", hn.TryLogin)
	app.Get("/createpage", middleware.Protected(), hn.CreateNewPage)
	app.Get("/robots.txt", hn.Robots)
	app.Get("/contact", hn.GetContact)
	app.Get("/contact", hn.PostContact)
	app.Get("/about", hn.About)

	// app.Get("/login", hn.Login)
	app.Get("/logout", hn.Logout)
	app.Get("/register", hn.Register)
	app.Get("/:slug", hn.PageBySlug)
	app.Get("/", middleware.Protected(), hn.Index)

}

package main

import (
	"log"

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

var posts = []Post{
	{Title: "first post", Date: func() string { return "test" }, Content: "this is content"},
	{Title: "second post", Date: func() string { return "test" }, Content: "this is content"},
	{Title: "third post", Date: func() string { return "test" }, Content: "this is content"},
	{Title: "forth post", Date: func() string { return "test" }, Content: "this is content"},
	{Title: "forth post", Date: func() string { return "test" }, Content: "this is content"},
	{Title: "forth post", Date: func() string { return "test" }, Content: "this is content"},
	{Title: "forth post", Date: func() string { return "test" }, Content: "this is content"},
	{Title: "forth post", Date: func() string { return "test" }, Content: "this is content"},
	{Title: "forth post", Date: func() string { return "test" }, Content: "this is content"},
	{Title: "forth post", Date: func() string { return "test" }, Content: "this is content"},
	{Title: "forth post", Date: func() string { return "test" }, Content: "this is content"},
	{Title: "forth post", Date: func() string { return "test" }, Content: "this is content"},
	{Title: "forth post", Date: func() string { return "test" }, Content: "this is content"},
	{Title: "forth post", Date: func() string { return "test" }, Content: "this is content"},
	{Title: "forth post", Date: func() string { return "test" }, Content: "this is content"},
	{Title: "forth post", Date: func() string { return "test" }, Content: "this is content"},
	{Title: "forth post", Date: func() string { return "test" }, Content: "this is content"},
	{Title: "forth post", Date: func() string { return "test" }, Content: "this is content"},
}

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
	app.Get("/robots.txt", func(ctx *fiber.Ctx) error {
		return ctx.Render("robots", nil)
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{"Title": title, "HomeActive": true, "posts": posts}, "layout")
	})

	app.Get("/contact", func(ctx *fiber.Ctx) error {
		return ctx.Render("contact", fiber.Map{"Title": title, "ContactActive": true}, "layout")
	})

	app.Get("/about", func(ctx *fiber.Ctx) error {
		return ctx.Render("about", fiber.Map{"Title": title, "AboutActive": true}, "layout")
	})

	app.Get("/login", func(ctx *fiber.Ctx) error {
		return ctx.Render("login", fiber.Map{"Title": title, "LoginActive": true}, "layout")
	})
	app.Get("/logout", func(ctx *fiber.Ctx) error {
		return nil
	})

	app.Get("/register", func(ctx *fiber.Ctx) error {
		return ctx.Render("register", fiber.Map{"Title": title, "RegisterActive": true}, "layout")
	})

	log.Fatal(app.Listen(":3000"))
}

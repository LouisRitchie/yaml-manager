package main

import (
	"encoding/json"
	"os"
    "io"
    "net/http"
    "fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
	"github.com/mewben/config-echo"
	//_ "github.com/lib/pq"
	//"github.com/mewben/db-go-env"
)

// Initialize Port and DB Connection config
func init() {
	type Config struct {
		SERVERPORT string
		//DB   db.Database
	}

	configFile, err := os.Open("env.json")
	if err != nil {
		panic(err)
	}

	var devConfig Config
	jsonParser := json.NewDecoder(configFile)

	if err = jsonParser.Decode(&devConfig); err != nil {
		panic(err)
	}

	// setup postgres db connection
	//db.Setup(devConfig.DB)

	// setup port
	// This sets the global Port string
	// If you set an environment variable DATABASE_URL,
	// it sets Mode = "prod" and uses the env Port instead
	config.Setup(devConfig.SERVERPORT)
}


func upload(c echo.Context) error {
  // Read form fields
  name := c.FormValue("name")
  email := c.FormValue("email")

  //------------
  // Read files
  //------------

  // form
  file, err := c.FormFile("file")
  if err != nil {
    return err
  }

  // Source
  src, err := file.Open()
  if err != nil {
    return err
  }
  defer src.Close()

  // Destination
  dst, err := os.Create(file.Filename)
  if err != nil {
    return err
  }
  defer dst.Close()

  // Copy
  if _, err = io.Copy(dst, src); err != nil {
    return err
  }

  return c.HTML(http.StatusOK, fmt.Sprintf("<p>Successfully uploaded your file with fields name=%s and email=%s.</p>", name, email))
}

func main() {
	app := echo.New()

	app.Use(middleware.Recover())
	app.Use(middleware.Gzip())
	app.Use(middleware.Secure())
	app.Use(middleware.BodyLimit("100K"))

	if config.Mode == "dev" {
		// Enable Debug
		app.Use(middleware.Logger())
		app.SetDebug(true)
	}

	app.File("/*", "public/index.html")
	app.Static("/assets", "public/assets")

    app.POST("/upload", upload)

    fmt.Sprintf("why are things so fucked?")

	app.Run(fasthttp.New(config.Port))
}

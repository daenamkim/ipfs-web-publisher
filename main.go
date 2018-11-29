package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/go-xorm/xorm"

	_ "github.com/lib/pq"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

type Publish struct {
	Id         int64
	Name       string    `xorm:varchar not null`
	FileName   string    `xorm:varchar not null`
	GatewayUrl string    `xorm:varchar not null`
	CreatedAt  time.Time `xorm:created`
	UpdatedAt  time.Time `xorm:updated`
}

func IPFSWebPublisher() *iris.Application {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.Default())

	app.Logger().SetLevel("debug")

	return app
}

func main() {
	app := IPFSWebPublisher()

	orm, err := setupDatabase()
	if err != nil {
		// app.Logger().Fatalf("orm failed to initialized publish table: %v", err)
		log.Fatal(err)
	}

	publishApi := app.Party("/publish")
	publishApi.Get("/", func(ctx iris.Context) {
		// TODO: how to get list.
	})
	publishApi.Options("/", func(ctx iris.Context) {
		// for CORS
	})
	publishApi.Get("/{id:int}", func(ctx iris.Context) {
		id, err := ctx.Params().GetInt64("id")
		if err != nil {
			log.Fatal(err)
			return
		}

		publish := Publish{Id: id}
		_, err = orm.Get(&publish)
		if err != nil {
			log.Fatal(err)
			return
		}

		publishJson, err := json.Marshal(publish)
		if err != nil {
			log.Fatal(err)
			return
		}

		ctx.Write(publishJson)
	})

	publishApi.Post("/", func(ctx iris.Context) {
		appName := ctx.PostValue("app-name")

		publishAdd := &Publish{
			Name:       appName,
			FileName:   "",
			GatewayUrl: "",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		// Remove previous uploads.
		// TODO: need to keep for a while in random file name.
		// FIXME: cannot remove files and directories why?
		out, err := exec.Command("rm", "-rf", "uploads/*").Output()
		fmt.Printf("RM command %s", out)

		// Upload a file.
		_, err = ctx.UploadFormFiles("./uploads", beforeSave)
		if err != nil {
			log.Fatal(err)
			return
		}

		// TODO: auto detect a directory name.
		// Unzip.
		out, err = exec.Command("unzip", "uploads/upload.zip", "-d", "uploads/").Output()
		fmt.Printf("Unzip %s", out)

		// IPFS run.
		out, err = exec.Command("ipfs", "add", "-r", "uploads/upload").Output()
		lines := strings.Split(string(out), "\n")
		// for index, line := range lines {
		// 	fmt.Printf("%d, %q\n", index, line)
		// }
		// "added QmbFMke1KXqnYyBBWxB74N4c5SBnJMVAiMNRcGu6x1AwQH sample/img/test.png"
		// "added QmbFMke1KXqnYyBBWxB74N4c5SBnJMVAiMNRcGu6x1AwQH sample/index.html"
		// "added QmWami7sJTPgZCzZKU8T6N13XQoMaktE5jKRAmVLWPqj4w sample/img"
		// "added QmdUnBjZgps3brQgWSXHs9wLbicJwXaqhRJEVgtGjKFDqE sample"
		// ""
		publishAdd.GatewayUrl = strings.Split(lines[len(lines)-2], " ")[1]

		// TODO: how to get latest one after insert.
		_, err = orm.Insert(publishAdd)
		if err != nil {
			log.Fatal(err)
			return
		}

		// publishAdded := Publish{Id: ?}
		// _, err = orm.Get(&publishAdded)
		// if err != nil {
		// 	log.Fatal(err)
		// 	return
		// }

		// publishAddedJson, err := json.Marshal(publishAdded)
		// if err != nil {
		// 	log.Fatal(err)
		// 	return
		// }

		// ctx.Write(publishAddedJson)
	})

	publishApi.Patch("/{id:int}", func(ctx iris.Context) {
		result, _ := ctx.Params().GetInt64("id")
		appName := ctx.PostValue("app-name")
		condition := Publish{Id: result}
		data := &Publish{Name: appName}
		orm.Update(data, condition)
		ctx.UploadFormFiles("./uploads", beforeSave)
		ctx.Writef("Patch!! %#v, %#v", result, appName)
	})

	publishApi.Delete("/{id:int}", func(ctx iris.Context) {
		id, err := ctx.Params().GetInt64("id")
		if err != nil {
			log.Fatal(err)
			return
		}

		publish := Publish{Id: id}
		_, err = orm.Delete(publish)
		if err != nil {
			log.Fatal(err)
			return
		}
	})

	app.Get("/ping", func(ctx iris.Context) {
		ctx.Writef("pong")
	})

	app.StaticWeb("/", "./public")
	app.StaticWeb("/js", "./public/js")
	app.StaticWeb("/css", "./public/css")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	app.Run(iris.Addr(":" + port))
}

func beforeSave(ctx iris.Context, file *multipart.FileHeader) {
	// TODO:
	// ip := ctx.RemoteAddr()
	// ip = strings.Replace(ip, ".", "-", -1)
	// ip = strings.Replace(ip, ":", "-", -1)
	// now := time.Now().Unix()
	// file.Filename = ip + "-" + strconv.FormatInt(now, 10) + "-" + file.Filename
	file.Filename = "upload.zip"
}

func setupDatabase() (*xorm.Engine, error) {
	// Raw SQL tutorial.
	// connStr := "user=daenamkim dbname=ipfs sslmode=disable"
	// For access to database's URL.
	// connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	//  return
	// }
	//
	// var publishId int
	// err = db.QueryRow(`INSERT INTO publish (name, local_url, gateway_url) VALUES ($1, $2, $3)`, "Test Name", "Test Local URL", "Test Gateway URL").Scan(&publishId)
	// fmt.Printf("Created ID: %s", publishId)
	// db.Query(`DELETE FROM publish WHERE id = $1`, 3)
	// db.Query(`UPDATE publish SET local_url = $1, updated_at = $3 WHERE id = $2`, "NEW ONE!!!", 5, time.Now())
	// rows, err := db.Query("SELECT * FROM publish")
	// if err != nil {
	//	return
	// }

	// ORM initialization for database.
	connStr := "user=daenamkim dbname=ipfs sslmode=disable"
	orm, err := xorm.NewEngine("postgres", connStr)
	if err != nil {
		return orm, err
	}

	iris.RegisterOnInterrupt(func() {
		orm.Close()
	})

	// Create and sync a table.
	// TODO: When just using this sync2 first time wintout creating a table,
	// I don't know how to add CURRENT_TIMESTAMP and there is no actual time zone value such as "+09".
	// So I just created a table manually then used this. Later on I need to look into the ORM library.
	err = orm.Sync2(new(Publish))
	return orm, err
}

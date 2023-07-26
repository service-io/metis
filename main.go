package main

import (
	"fmt"
	"github.com/json-iterator/go/extra"
	"metis/config"
	"metis/router"
	"metis/util/logger"
	"strings"
	"time"
)

// var db *sql.DB

// type Album struct {
//	ID     int64
//	Title  string
//	Artist string
//	Price  float32
// }

func lowerCamelCase(f string) string {
	return strings.ToLower(f[:1]) + f[1:]
}

func main() {
	useLogger := logger.CommonLogger()
	tomlConfig := config.TomlConfig()
	baseRouter := router.BaseRouter()

	extra.RegisterFuzzyDecoders()
	extra.RegisterTimeAsInt64Codec(time.Millisecond)
	extra.SetNamingStrategy(lowerCamelCase)

	// docs.SwaggerInfo.BasePath = "/"
	// baseRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	useLogger.Info(fmt.Sprintf("config -> %v", tomlConfig))
	useLogger.Info("hello tabuyos, I'm born.")

	err := baseRouter.Run("127.0.0.1:8080")
	if err != nil {
		return
	}
}

// func main() {
//
//	fmt.Println("hello tabuyos, I'm born.")
//
//  var mutex sync.Mutex
//  mutex.Lock()
//
//	cfg := mysql.Config{
//		User:   "root",
//		Passwd: "root",
//		Net:    "tcp",
//		Addr:   "localhost:3307",
//		DBName: "metis",
//	}
//
//	var err error
//	db, err = sql.Open("mysql", cfg.FormatDSN())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	pingErr := db.Ping()
//	if pingErr != nil {
//		log.Fatal(pingErr)
//	}
//	fmt.Println("Connected!")
//
//	albums, err := albumsByArtist("John Coltrane")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Albums found: %v\n", albums)
//
//	alb, err := albumByID(2)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Album found: %v\n", alb)
//
//	albID, err := addAlbum(Album{
//		Title:  "The Modern Sound of Betty Carter",
//		Artist: "Betty Carter",
//		Price:  49.99,
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("ID of added album: %v\n", albID)
//
//  mutex.Unlock()
//
// }
//
// // albumsByArtist queries for albums that have the specified artist name.
// func albumsByArtist(name string) ([]Album, error) {
//	// An albums slice to hold data from returned rows.
//	var albums []Album
//
//	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
//	if err != nil {
//		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
//	}
//	defer func(rows *sql.Rows) {
//    err := rows.Close()
//    if err != nil {
//
//    }
//  }(rows)
//	// Loop through rows, using Scan to assign column data to struct fields.
//	for rows.Next() {
//		var alb Album
//		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
//			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
//		}
//		albums = append(albums, alb)
//	}
//	if err := rows.Err(); err != nil {
//		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
//	}
//	return albums, nil
// }
//
// // albumByID queries for the album with the specified ID.
// func albumByID(id int64) (Album, error) {
//	// An album to hold data from the returned row.
//	var alb Album
//
//	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
//	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
//		if err == sql.ErrNoRows {
//			return alb, fmt.Errorf("albumsById %d: no such album", id)
//		}
//		return alb, fmt.Errorf("albumsById %d: %v", id, err)
//	}
//	return alb, nil
// }
//
// // addAlbum adds the specified album to the database,
// // returning the album ID of the new entry
// func addAlbum(alb Album) (int64, error) {
//	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
//	if err != nil {
//		return 0, fmt.Errorf("addAlbum: %v", err)
//	}
//	id, err := result.LastInsertId()
//	if err != nil {
//		return 0, fmt.Errorf("addAlbum: %v", err)
//	}
//	return id, nil
// }

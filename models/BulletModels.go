package models

import (
	"database/sql"

	"github.com/anandagireesh/Evermos-Backend-Engineer-Assessment-Q-1/database"
)

type Bullets struct {
	Bullets string `json:"bullet"`
}

func AddBullet(data Bullets) string {

	var message string

	row := database.Db.QueryRow("SELECT bullet_name FROM bullets where bullet_name = '" + data.Bullets + "'")
	bullets := Bullets{}
	err := row.Scan(&bullets.Bullets)
	if err != nil {
		if err == sql.ErrNoRows {

			insert, err := database.Db.Prepare("INSERT INTO bullets(bullet_name) VALUES ( ? )")

			//if there is an error inserting, handle it
			if err != nil {
				panic(err.Error())
			}
			insert.Exec(data.Bullets)

			message = "Bullets successfully registered"

		} else {
			panic(err)
		}
	} else {
		message = "Bullet numbered already registered"
	}

	return message
}

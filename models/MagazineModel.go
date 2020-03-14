package models

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/anandagireesh/Evermos-Backend-Engineer-Assessment-Q-1/database"
)

type Magazine struct {
	Magazine string `json:"magazine"`
}
type FillMagazine struct {
	Magazine string `json:"magazine"`
	Bullets  string `json:"bullet"`
}

// register bullets to database model

func RegisterMagazine(data Magazine) string {

	var message string

	row := database.Db.QueryRow("SELECT magazine FROM magazine where magazine = '" + data.Magazine + "'")
	magazine := Magazine{}
	err := row.Scan(&magazine.Magazine)
	if err != nil {
		if err == sql.ErrNoRows {

			insert, err := database.Db.Prepare("INSERT INTO magazine(magazine) VALUES ( ? )")

			//if there is an error inserting, handle it
			if err != nil {
				panic(err.Error())
			}
			insert.Exec(data.Magazine)

			message = "Magazine successfully registered"

		} else {
			panic(err)
		}
	} else {
		message = "Magazine number already registered"
	}

	return message
}

func FillMAgazine(data FillMagazine) string {

	var message string

	row := database.Db.QueryRow("SELECT bullet FROM fill_magazine where bullet = '" + data.Bullets + "' AND status = 1")
	magazine := FillMagazine{}
	err := row.Scan(&magazine.Bullets)
	if err != nil {
		if err == sql.ErrNoRows {

			var count int
			stmt, err := database.Db.Prepare("SELECT COUNT(*) as count FROM fill_magazine where magazine = '" + data.Magazine + "' AND status = 1")
			if err != nil {
				log.Fatal(err)
			}
			err = stmt.QueryRow().Scan(&count)
			if err != nil {
				log.Fatal(err)
			}
			if count < 2 {

				log.Println(data.Magazine)

				insert, err := database.Db.Prepare("INSERT INTO fill_magazine(magazine,bullet,status) VALUES ( ?,?,? )")

				//if there is an error inserting, handle it
				if err != nil {
					panic(err.Error())
				}
				insert.Exec(data.Magazine, data.Bullets, 1)

				message = "Magazine filled successfully registered"

			} else {
				message = "Magazine already filled"
			}

		}

	} else {
		panic(err)
	}

	return message
}

func VerifyMagazine(data Magazine) string {

	var message string

	var count int
	stmt, err := database.Db.Prepare("SELECT COUNT(*) as count FROM fill_magazine where magazine = '" + data.Magazine + "' AND status = 1")
	if err != nil {
		log.Fatal(err)
	}
	err = stmt.QueryRow().Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count < 3 && count > 0 {

		ct := strconv.Itoa(count)

		message = "Magazine Verified. Total Bullets : " + ct

	} else if count == 0 {
		message = "Magazine not verified"
	}
	log.Println(count)
	return message
}

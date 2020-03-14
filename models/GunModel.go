package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/anandagireesh/Evermos-Backend-Engineer-Assessment-Q-1/database"
)

type Gun struct {
	Gun string `json:"gun"`
}
type LoadGun struct {
	Magazine string `json:"magazine"`
	Gun      string `json:"gun"`
}

// register bullets to database model

func RegisterGun(data Gun) string {

	var message string

	row := database.Db.QueryRow("SELECT gun FROM gun where gun = '" + data.Gun + "'")
	Gun := Gun{}
	err := row.Scan(&Gun.Gun)
	if err != nil {
		if err == sql.ErrNoRows {

			insert, err := database.Db.Prepare("INSERT INTO gun(gun) VALUES ( ? )")

			//if there is an error inserting, handle it
			if err != nil {
				panic(err.Error())
			}
			insert.Exec(data.Gun)

			message = "Gun successfully registered"

		} else {
			panic(err)
		}
	} else {
		message = "Gun already registered"
	}

	return message
}

func LoadMagazine(data LoadGun) string {

	var message string

	row := database.Db.QueryRow("SELECT gun FROM fill_gun where magazine = '" + data.Magazine + "' AND status = 1")
	loadGun := LoadGun{}
	err := row.Scan(&loadGun.Gun)
	if err != nil {
		if err == sql.ErrNoRows {

			fmt.Println("test one")

			var count int
			stmt, err := database.Db.Prepare("SELECT COUNT(*) as count FROM fill_magazine where magazine = '" + data.Magazine + "' AND status = 1")
			if err != nil {
				log.Fatal(err)
			}
			err = stmt.QueryRow().Scan(&count)
			if err != nil {
				log.Fatal(err)
			}
			if count == 2 {

				fmt.Println("test two")

				var count int
				stmt1, err := database.Db.Prepare("SELECT COUNT(*) as count FROM fill_gun where magazine = '" + data.Magazine + "' AND status = 1")
				if err != nil {
					log.Fatal(err)
				}
				err = stmt1.QueryRow().Scan(&count)
				if err != nil {
					log.Fatal(err)
				}
				if count < 2 {

					fmt.Println("test three")

					log.Println(data.Magazine)

					insert, err := database.Db.Prepare("INSERT INTO fill_gun(magazine,gun,status) VALUES ( ?,?,? )")

					//if there is an error inserting, handle it
					if err != nil {
						panic(err.Error())
					}
					insert.Exec(data.Magazine, data.Gun, 1)

					message = "Magazine loaded successfully registered"

				} else {
					message = "Cannot Add magazine"
				}

			}

		}

	} else {
		panic(err)
	}

	return message
}

func FireGun(data Gun) string {

	var message, magazine, bullet string

	row := database.Db.QueryRow("SELECT magazine FROM fill_gun where gun = '" + data.Gun + "' AND status = 1")
	//loadGun := Gun{}
	err := row.Scan(&magazine)
	if err != nil {
		if err == sql.ErrNoRows {

		}

	} else {

		row := database.Db.QueryRow("SELECT bullet FROM fill_magazine where magazine = '" + magazine + "' AND status = 1")
		//loadGun := Gun{}
		err := row.Scan(&bullet)
		if err != nil {
			if err == sql.ErrNoRows {

				insForm, err := database.Db.Prepare("UPDATE fill_gun SET status=? WHERE magazine=?")
				if err != nil {
					panic(err.Error())
				}
				insForm.Exec(0, magazine)

				message = "Gun Cannot Fired changing magazine try again"

			}

		} else {

			fmt.Println(bullet)
			insForm, err := database.Db.Prepare("UPDATE fill_magazine SET status=? WHERE bullet=?")
			if err != nil {
				panic(err.Error())
			}
			insForm.Exec(0, bullet)

			message = "fired successfully"

		}
	}

	return message
}

package dal

import (
	"api/golang-song-api/confi"
	"api/golang-song-api/data_model"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)


func UserAuthenticate( userName string, password string) data_model.User {

	var u data_model.User 
	found := false

	db, err := sql.Open(confi.DataBaseType, confi.DataBasePath)
	checkErr(err)

	query := fmt.Sprintf("SELECT UserName, FirstName, LastName, Email, IsAdmin FROM User WHERE lower(UserName) = lower('%s') AND Password = '%s'", userName, password);

	rows, err := db.Query(query)
	
	// Let the server send the error to the user
	if err != nil {
		checkErrNice(err)
		return u
	}	

	for rows.Next() {
		err = rows.Scan(&u.UserName, &u.FirstName, &u.LastName, &u.Email, &u.IsAdmin)

		checkErrDb(err, db)

		//token := helpers.GenerateSecureToken()
		//tokenExpires := helpers.GetNewTokenExpirationDate()
		//u.Token = token
		//u.TokenExpires = tokenExpires

		found = true
		break
	}

	rows.Close()
	db.Close()
	
	if (!found) {
		println("Authentication for user " + userName + " has failed.")
	} else {
		//UpdateTokenForUser(u)
	}

	return u	
}

func getUserAdmin(userName string) bool {
	var isAdmin bool
	db, err := sql.Open(confi.DataBaseType, confi.DataBasePath)
	checkErr(err)

	rows, err := db.Query(fmt.Sprintf("SELECT isAdmin FROM User WHERE UserName ='%s'", userName))
	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&isAdmin)
		checkErrDb(err, db)
	}

	rows.Close()
	db.Close()

	return isAdmin
}

func CheckUserExists(userName string) bool {
	db, err := sql.Open(confi.DataBaseType, confi.DataBasePath)
	checkErr(err)

	rows, err := db.Query(fmt.Sprintf("SELECT userName FROM User WHERE lower(UserName) = lower('%s')",userName))
	checkErrDb(err, db)

	found := false

	for rows.Next() {
		checkErrDb(err, db)
		found = true
	}

	rows.Close()
	db.Close()

	return found
}

func AddUser(u data_model.User) {
	db, err := sql.Open(confi.DataBaseType, confi.DataBasePath)
	checkErr(err)

	query := fmt.Sprintf("INSERT INTO User (UserName, Password, FirstName, LastName, Email, IsAdmin) Values ('%s', '%s', '%s', '%s', '%s', 0)",
		u.UserName, u.Password, u.FirstName, u.LastName, u.Email)

		println(query)

	_, err = db.Exec(query)
	checkErr(err)

	db.Close()
}

/*func UpdateTokenForUser(u data_model.User) {
	db, err := sql.Open(confi.DataBaseType, confi.DataBasePath)
	checkErr(err)

	query := fmt.Sprintf("UPDATE User SET Token = '%s', TokenExpires = '%s' WHERE  UserName = '%s'", u.Token, u.TokenExpires.Format("01-02-2006 15:04:05"), u.UserName )
	_, err = db.Exec(query)
	db.Close()
	checkErrNice(err)		
}*/

/*func CheckToken(userName string, token string) bool {
	var tokenExpires sql.NullTime
	foundValid := false
	needsRefresh := false
	
	db, err := sql.Open(confi.DataBaseType, confi.DataBasePath)
	checkErr(err)

	query := fmt.Sprintf("SELECT TokenExpires FROM User WHERE UserName = '%s' AND Token = '%s'", userName, token);
	println(query)
	rows, err := db.Query(query)

	for rows.Next() {
		err = rows.Scan(&tokenExpires)
		checkErrDb(err, db)

		if !tokenExpires.Valid && tokenExpires.Time.After(time.Now()) {
			foundValid = true
		}

		refreshTime := time.Now().Add(time.Minute * (-30))

		// If there has been more than 30 minutes since last authentication let's refresh the tokenExpire time
		if tokenExpires.Time.After(refreshTime) {
			needsRefresh = true
		}

	}

	rows.Close()
	db.Close()

	if needsRefresh {
		var u data_model.User
		u.UserName = userName
		u.Token = token
		u.TokenExpires = helpers.GetNewTokenExpirationDate()
		UpdateTokenForUser(u)
	}

	return foundValid
}*/




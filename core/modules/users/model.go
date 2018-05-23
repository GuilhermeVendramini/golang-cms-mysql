package users

import (
	"github.com/GuilhermeVendramini/golang-cms-mysql/config"
)

// Create a new User
func Create(user User) (User, error) {
	stmtIns, err := config.DB.Prepare("INSERT INTO users (Name, Email, Password, Admin) VALUES (?, ?, ?, ?)")
	defer stmtIns.Close()
	if err != nil {
		panic(err.Error())
	}
	stmtIns.Exec(user.Name, user.Email, user.Password, user.Admin)
	return user, nil
}

// Update user
func Update(user User, ID string) (User, error) {
	stmtUp, err := config.DB.Prepare("UPDATE users set Name = ?, Email = ?, Password = ?, Admin = ? WHERE Id = ?")
	defer stmtUp.Close()
	if err != nil {
		panic(err.Error())
	}
	stmtUp.Exec(user.Name, user.Email, user.Password, user.Admin, ID)
	return user, nil
}

// Remove user
func Remove(ID string) error {
	stmtDel, err := config.DB.Prepare("delete from users where id = ?")
	defer stmtDel.Close()
	if err != nil {
		panic(err.Error())
	}
	stmtDel.Exec(ID)
	return nil
}

// GetbyID return one user by ID
func GetbyID(ID string) (User, error) {
	user := User{}
	rows, err := config.DB.Query("SELECT Id, Name, Email, Password, Admin FROM users WHERE Id = ?", ID)
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Admin)
	}
	return user, err
}

// GetbyEmail return one user by email
func GetbyEmail(Email string) (User, error) {
	user := User{}
	rows, err := config.DB.Query("SELECT Id, Name, Email, Password, Admin FROM users WHERE Email = ?", Email)
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Admin)
	}
	return user, err
}

// GetAll return all users
func GetAll() ([]User, error) {
	user := User{}
	users := []User{}
	rows, _ := config.DB.Query("SELECT Id, Name, Email, Password, Admin FROM users")
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Admin)
		users = append(users, user)
	}
	return users, nil
}

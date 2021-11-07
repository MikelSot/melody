package persistence

import (
	"database/sql"
	"fmt"
	"github.com/MikelSot/melody/domain"
	"log"
)

const (
	create_user =`INSERT INTO users(name, last_name, email, password ) VALUES($1, $2, $3, $4) RETURNING id`
	update_user =`UPDATE users SET name = $1, last_name = $2, email = $3, nick_name = $4, description = $5 WHERE id=$6` // modifica aca es nikname NO nike_name (en la tabla pincipal igual)
	update_nikname_field = `UPDATE users SET nick_name=$1 WHERE id=$2` // a qui igual lo cambias despues
	update_picture_field = `UPDATE users SET picture=$1 WHERE id=$2`
	update_online_field = `UPDATE users SET online=$1 WHERE id=$2`
)

type user struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *user {
	return &user{db}
}

func (u user) Create(d *domain.CreateUser) error {
	stmt, err := u.db.Prepare(create_user)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		d.Name,
		d.LastName,
		d.Email,
		d.Password,
	).Scan(&d.ID)
	if err != nil {
		return err
	}
	log.Println("CREADO")
	return  nil
}

func (u user) Update(up *domain.UpdateUser) error {
	stmt, err := u.db.Prepare(update_user)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		up.Name,
		up.LastName,
		up.Email,
		up.Nickname,
		up.Description,
		up.ID,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no existe un usuario con id: %d", up.ID)
	}

	log.Println("ACTUALIZADO")
	return  nil
}

func (u user) UpdateNicknameField(nick *domain.FieldString) error {
	stmt, err := u.db.Prepare(update_nikname_field)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(nick.Field,nick.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no existe un usuario con nickname: %v", nick.Field)
	}
	log.Println("ACTUALIZADO NICKNAME")
	return nil
}

func (u user) UpdatePictureField(pic *domain.FieldString) error {
	stmt, err := u.db.Prepare(update_picture_field)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(pic.Field,pic.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no se puedo actualizar la imagen del usuario")
	}
	log.Println("ACTUALIZADO PICTURE")
	return nil
}

func (u user) UpdateOnlineField(state *domain.State) error {
	stmt, err := u.db.Prepare(update_online_field)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(state.Field,state.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no se puedo actualizar el campo online del usuario: %v", state.ID)
	}
	log.Println("ACTUALIZADO ONLINE")
	return nil
}
package dao

import "github.com/MikelSot/melody/domain"

func (m message) GetById(id uint) (interface{}, error) {
	ms := domain.Message{}
	db.First(&ms, id)
	return ms, nil
}


// TODO: obtener las conversaciones por CHAT_ID y ordenamos de manera ascemdente [x]
// TODO: obtener los mensajes que sean de tipo archiv, etc y lo agrupo por estos (archiv, etc), OJO puede ser solo una consulta o varias (en el frond lo podemos serrar por el id del tipo) [x]
// TODO: obtener todos los mensajes destcados  [x]




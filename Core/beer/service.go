package beer

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type UseCase interface {
	GetAll() ([]*Beer, error)
	Get(Id int64) (*Beer, error)
	Store(b *Beer) error
	Update(b *Beer) error
	Remove(Id int64) error
}

// a struct Service agora tem uma conexão com o banco de dados dentro dela
type Service struct {
	DB *sql.DB
}

// esta função retorna um ponteiro em memória para uma estrutura
// a função agora recebe uma conexão com o banco de dados
func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

// vamos implementar as funções na próxima etapa
func (s *Service) GetAll() ([]*Beer, error) {
	//result é um slice de ponteiros do tipo Beer
	var result []*Beer

	//vamos sempre usar a conexão que está dentro do Service
	rows, err := s.DB.Query("select Id, name, type, style from beer")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b Beer
		err = rows.Scan(&b.Id, &b.Name, &b.Type, &b.Style)
		if err != nil {
			return nil, err
		}

		result = append(result, &b)
	}
	return result, nil
}

func (s *Service) Get(Id int64) (*Beer, error) {
	//b é um tipo Beer
	var b Beer

	//o comando Prepare verifica se a consulta está válIda
	stmt, err := s.DB.Prepare("select Id, name, type, style from beer where Id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(Id).Scan(&b.Id, &b.Name, &b.Type, &b.Style)
	if err != nil {
		return nil, err
	}

	//deve retornar a posição da memória de b
	return &b, nil
}

func (s *Service) Store(b *Beer) error {
	//iniciamos uma transação
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("insert into beer(Id, name, type, style) values (?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	//o comando Exec retorna um Result, mas não temos interesse nele, por isso podemos ignorá-lo com o _
	_, err = stmt.Exec(b.Id, b.Name, b.Type, b.Style)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *Service) Update(b *Beer) error {
	if b.Id == 0 {
		//podemos também retornar um erro de aplicação
		//que criamos para definir uma condição de erro, como um possível update sem Where
		return fmt.Errorf("invalId Id")
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("update beer set name=?, type=?, style=? where Id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	//o comando Exec retorna um Result, mas não temos interesse nele, por isso podemos ignorá-lo com o _
	_, err = stmt.Exec(b.Name, b.Type, b.Style, b.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *Service) Remove(Id int64) error {
	if Id == 0 {
		//podemos também retornar um erro de aplicação
		//que criamos para definir uma condição de erro, como um possível update sem Where
		return fmt.Errorf("invalId Id")
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	//o comando Exec retorna um Result, mas não temos interesse nele, por isso podemos ignorá-lo com o _
	_, err = tx.Exec("delete from beer where Id=?", Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

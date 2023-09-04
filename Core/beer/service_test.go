package beer

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestBeerService(t *testing.T) {
	// Conectar a um banco de dados SQLite de teste (usando um banco de dados em memória para testes)
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Erro ao abrir o banco de dados de teste: %v", err)
	}
	defer db.Close()

	// Criar a tabela de teste (certifique-se de que sua implementação crie a tabela corretamente)
	_, err = db.Exec(`
		CREATE TABLE beer (
			Id INTEGER PRIMARY KEY,
			name TEXT,
			type INTEGER,
			style INTEGER
		)
	`)
	if err != nil {
		t.Fatalf("Erro ao criar a tabela de teste: %v", err)
	}

	// Criar uma instância do serviço de cerveja
	service := NewService(db)

	t.Run("Test Store, Get, Update, Remove", func(t *testing.T) {
		// Criar uma nova cerveja para armazenamento
		newBeer := &Beer{
			Name:  "NewBeer",
			Type:  TypeAle,
			Style: StyleAmber,
		}

		// Testar a função Store
		err := service.Store(newBeer)
		if err != nil {
			t.Fatalf("Erro ao armazenar nova cerveja: %v", err)
		}

		// Testar a função Get
		retrievedBeer, err := service.Get(1)
		if err != nil {
			t.Fatalf("Erro ao buscar cerveja por ID: %v", err)
		}

		if retrievedBeer.Name != newBeer.Name {
			t.Fatalf("Nome da cerveja incorreto. Esperado '%s', mas obteve '%s'", newBeer.Name, retrievedBeer.Name)
		}

		// Atualizar a cerveja recuperada
		retrievedBeer.Name = "UpdatedBeer"
		err = service.Update(retrievedBeer)
		if err != nil {
			t.Fatalf("Erro ao atualizar a cerveja: %v", err)
		}

		// Verifique se a cerveja foi atualizada corretamente
		updatedBeer, err := service.Get(1)
		if err != nil {
			t.Fatalf("Erro ao buscar cerveja atualizada por ID: %v", err)
		}

		if updatedBeer.Name != "UpdatedBeer" {
			t.Fatalf("Nome da cerveja atualizado incorretamente. Esperado 'UpdatedBeer', mas obteve '%s'", updatedBeer.Name)
		}

		// Testar a função Remove
		err = service.Remove(1)
		if err != nil {
			t.Fatalf("Erro ao remover cerveja por ID: %v", err)
		}

		// Verifique se a cerveja foi removida corretamente
		_, err = service.Get(1)
		if err == nil {
			t.Fatalf("Erro inesperado: a cerveja não foi removida")
		}
	})
}

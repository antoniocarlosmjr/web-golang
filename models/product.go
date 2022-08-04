package models

import (
	"github.com/antoniocarlosmjr/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func SearchAllProducts() []Produto {
	db := db.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	insereDadosBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insereDadosBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func Delete(idProduto string) {
	db := db.ConectaComBancoDeDados()
	deleteProduto, err := db.Prepare("delete from produtos where id $1")
	if err != nil {
		panic(err.Error())
	}

	deleteProduto.Exec(idProduto)
	defer db.Close()
}

func SearchProduct(id string) Produto {
	db := db.ConectaComBancoDeDados()
	produtoFromBanco, err := db.Query("select * from produtos where id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoObject := Produto{}

	for produtoFromBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoFromBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produtoObject.Id = id
		produtoObject.Nome = nome
		produtoObject.Descricao = descricao
		produtoObject.Preco = preco
		produtoObject.Quantidade = quantidade
	}
	defer db.Close()
	return produtoObject
}

func Update(id, quantidade int, nome, descricao string, preco float64) {
	db := db.ConectaComBancoDeDados()
	atualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, quantidade=$3, preco=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	atualizaProduto.Exec(nome, descricao, quantidade, preco, id)
	defer db.Close()
}

package classes

import (
	"PLP_Backend/database"
	"fmt"
	"log"
)

// Struct de Crimes para o rows
type Crimes struct {
	IDCrime         int    `json:"id_crime"`
	NomeHeroi       string `json:"nome_heroi"`
	NomeCrime       string `json:"nome_crime"`
	Severidade      string `json:"severidade"`
	DataCrime       string `json:"data_crime"`
	DescricaoEvento string `json:"descricao_evento"`
}

// Método para consultar crimes por herói e por severidade
func ConsultaCrimesPorHeroiESeveridade(nomeHeroi string, severidadeMinima int, severidadeMaxima int) ([]Crimes, error) {
	db := database.ConectaDB()
	defer db.Close() // Garantir que o banco de dados seja fechado após o uso

	// Consulta para buscar crimes com base no nome do herói e na severidade
	query := `
		SELECT 
			c.nome_crime, c.severidade, hc.data_crime, hc.descricao_evento, h.nome_heroi
		FROM 
			crimes c
		JOIN 
			herois_crimes hc ON c.id_crime = hc.id_crime
		JOIN 
			herois h ON hc.id_heroi = h.id_heroi
		WHERE 
			h.nome_heroi = $1 
		
		AND 
			c.severidade BETWEEN $2 AND $3
		AND 
			hc.esconder = false;
	`

	// Executa a consulta
	rows, err := db.Query(query, nomeHeroi, severidadeMinima, severidadeMaxima)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close() // Garantir que o resultado seja fechado após o uso

	// Cria uma slice para armazenar os crimes
	var crimes []Crimes

	// Itera sobre os resultados da consulta
	for rows.Next() {
		var crime Crimes
		err := rows.Scan(
			&crime.NomeCrime,
			&crime.Severidade,
			&crime.DataCrime,
			&crime.DescricaoEvento,
			//&esconder,        // Agora você armazena o valor de "esconder" em uma variável bool
			&crime.NomeHeroi, // Nome do herói
		)
		if err != nil {
			log.Fatal(err)
		}
		crimes = append(crimes, crime)

	}

	// Verifica se não encontrou nenhum crime
	if len(crimes) == 0 {
		return nil, fmt.Errorf("nenhum crime encontrado para o herói %s com severidade entre %d e %d", nomeHeroi, severidadeMinima, severidadeMaxima)
	}

	return crimes, nil
}

// Função para Consultar os Crimes por Herói
func ConsultaCrimesPorHeroi(nomeHeroi string) ([]Crimes, error) {
	db := database.ConectaDB()
	defer db.Close() // Garantir que o banco de dados seja fechado após o uso

	// Consulta para buscar crimes com base no nome do herói
	query := `
		SELECT 
			c.nome_crime, c.severidade, hc.data_crime, hc.descricao_evento, h.nome_heroi
		FROM 
			crimes c
		JOIN 
			herois_crimes hc ON c.id_crime = hc.id_crime
		JOIN 
			herois h ON hc.id_heroi = h.id_heroi
		WHERE 
			h.nome_heroi = $1
		AND 
			hc.esconder = false;
	`

	// Executa a consulta
	rows, err := db.Query(query, nomeHeroi)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close() // Garantir que o resultado seja fechado após o uso

	// Cria uma slice para armazenar os crimes
	var crimes []Crimes

	// Itera sobre os resultados da consulta
	for rows.Next() {
		var crime Crimes
		err := rows.Scan(
			&crime.NomeCrime,
			&crime.Severidade,
			&crime.DataCrime,
			&crime.DescricaoEvento,
			&crime.NomeHeroi,
		)
		if err != nil {
			log.Fatal(err)
		}
		crimes = append(crimes, crime)
	}

	// Verifica se não encontrou nenhum crime
	if len(crimes) == 0 {
		return nil, fmt.Errorf("nenhum crime encontrado para o herói %s", nomeHeroi)
	}

	return crimes, nil
}

// Função para Consultar os Crimes por Severidade
func ConsultaCrimesPorSeveridade(severidadeMinima int, severidadeMaxima int) ([]Crimes, error) {
	db := database.ConectaDB()
	defer db.Close() // Garantir que o banco de dados seja fechado após o uso

	// Consulta para buscar crimes com base na severidade
	query := `
		SELECT 
			c.id_crime, c.nome_crime, c.severidade, hc.data_crime, hc.descricao_evento, h.nome_heroi
		FROM 
			crimes c
		JOIN 
			herois_crimes hc ON c.id_crime = hc.id_crime
		JOIN 
			herois h ON hc.id_heroi = h.id_heroi
		WHERE 
			c.severidade BETWEEN $1 AND $2
		AND 
			hc.esconder = false;
	`

	// Executa a consulta
	rows, err := db.Query(query, severidadeMinima, severidadeMaxima)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close() // Garantir que o resultado seja fechado após o uso

	// Cria uma slice para armazenar os crimes
	var crimes []Crimes

	// Itera sobre os resultados da consulta
	for rows.Next() {
		var crime Crimes
		err := rows.Scan(
			&crime.IDCrime,
			&crime.NomeCrime,
			&crime.Severidade,
			&crime.DataCrime,
			&crime.DescricaoEvento,
			&crime.NomeHeroi,
		)
		if err != nil {
			log.Fatal(err)
		}
		crimes = append(crimes, crime)
	}

	// Verifica se não encontrou nenhum crime
	if len(crimes) == 0 {
		return nil, fmt.Errorf("nenhum crime encontrado com severidade entre %d e %d", severidadeMinima, severidadeMaxima)
	}

	return crimes, nil
}

func DeleteCrime(id_crime int, nome_heroi string) error {
	db := database.ConectaDB()
	defer db.Close() // Garantir que o banco de dados seja fechado após o uso

	// Consulta para deletar crimes com base no id do herói e do crime
	query := `
        UPDATE herois_crimes hc
        SET esconder = true
        FROM herois h
        WHERE hc.id_crime = $1 
        AND h.nome_heroi = $2
        AND hc.id_heroi = h.id_heroi;`

	// Executa a consulta
	_, err := db.Exec(query, id_crime, nome_heroi)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func AtualizarCrime(idCrime string, crime Crimes) error {
	db := database.ConectaDB()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Atualizar dados básicos da missão
	query :=
		`UPDATE herois_crimes
        SET id_crime = $1, descricao_evento = $2, data_crime = $3, id_heroi = $4
        WHERE id_ocorrencia = $5`

	_, err = tx.Exec(query,
		crime.NomeCrime,
		crime.DescricaoEvento,
		crime.DataCrime,
		crime.NomeHeroi,
		idCrime,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

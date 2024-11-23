package classes

import (
	"PLP_Backend/database"
	"fmt"
	"log"
)

// Struct de Missoes para o rows
type Missoes struct {
	IdMissao        string `json:"id_missao"` // Adicionar este campo
	NomeHeroi       string `json:"nome_heroi"`
	NomeMissao      string `json:"nome_missao"`
	DescricaoMissao string `json:"descricao"`
	NivelMissao     string `json:"nivel_dificuldade"`
	Resultado       string `json:"resultado"`
	Recompensa      string `json:"recompensa"`
}

func ConsultaMissoesPorHeroi(nomeHeroi string) ([]Missoes, error) {
	db := database.ConectaDB()
	defer db.Close()

	// Query atualizada para incluir todos os heróis da missão
	query :=
		`SELECT
            m.nome_missao, 
            m.descricao, 
            m.nivel_dificuldade, 
            m.resultado, 
            m.recompensa, 
            h.nome_heroi
        FROM
            missoes m
        JOIN
            herois_missoes hm ON m.id_missao = hm.id_missao
        JOIN
            herois h ON hm.id_heroi = h.id_heroi
        WHERE
            m.id_missao IN (
                SELECT DISTINCT hm.id_missao 
                FROM herois_missoes hm
                JOIN herois h ON hm.id_heroi = h.id_heroi
                WHERE h.nome_heroi = $1
            )
        AND m.esconder_missao = false
        ORDER BY m.nivel_dificuldade ASC;`

	rows, err := db.Query(query, nomeHeroi)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	//Itera sobre o resultado das consultas
	var missoes []Missoes
	for rows.Next() {
		var missao Missoes
		err := rows.Scan(
			&missao.NomeMissao,
			&missao.DescricaoMissao,
			&missao.NivelMissao,
			&missao.Resultado,
			&missao.Recompensa,
			&missao.NomeHeroi,
		)
		if err != nil {
			log.Fatal(err)
		}
		missoes = append(missoes, missao)
	}

	if len(missoes) == 0 {
		return nil, fmt.Errorf("nenhuma missão encontrada para o herói %s", nomeHeroi)
	}
	return missoes, nil
}

func ListarTodasMissoes() ([]Missoes, error) {
	db := database.ConectaDB()
	defer db.Close()
	query :=
		`SELECT DISTINCT
            m.id_missao,
            m.nome_missao, 
            m.descricao, 
            m.nivel_dificuldade, 
            m.resultado, 
            m.recompensa, 
            string_agg(h.nome_heroi, ', ') as herois
        FROM
            missoes m
        LEFT JOIN
            herois_missoes hm ON m.id_missao = hm.id_missao
        LEFT JOIN
            herois h ON hm.id_heroi = h.id_heroi
        WHERE
            m.esconder_missao = false
        GROUP BY
            m.id_missao, m.nome_missao, m.descricao, m.nivel_dificuldade, m.resultado, m.recompensa
        ORDER BY m.nivel_dificuldade ASC;`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var missoes []Missoes
	for rows.Next() {
		var missao Missoes
		err := rows.Scan(
			&missao.IdMissao,
			&missao.NomeMissao,
			&missao.DescricaoMissao,
			&missao.NivelMissao,
			&missao.Resultado,
			&missao.Recompensa,
			&missao.NomeHeroi,
		)
		if err != nil {
			return nil, err
		}
		missoes = append(missoes, missao)
	}

	return missoes, nil
}

func ConsultaMissaoPorId(idMissao string) (Missoes, error) {
	db := database.ConectaDB()
	defer db.Close()
	query := `
        SELECT 
            m.id_missao,
            m.nome_missao,
            m.descricao,
            m.nivel_dificuldade,
            m.resultado,
            m.recompensa,
            string_agg(h.nome_heroi, ', ') as herois
        FROM missoes m
        LEFT JOIN herois_missoes hm ON m.id_missao = hm.id_missao
        LEFT JOIN herois h ON hm.id_heroi = h.id_heroi
        WHERE m.id_missao = $1
        AND m.esconder_missao = false
        GROUP BY m.id_missao, m.nome_missao, m.descricao, m.nivel_dificuldade, m.resultado, m.recompensa
    `

	var missao Missoes
	err := db.QueryRow(query, idMissao).Scan(
		&missao.IdMissao,
		&missao.NomeMissao,
		&missao.DescricaoMissao,
		&missao.NivelMissao,
		&missao.Resultado,
		&missao.Recompensa,
		&missao.NomeHeroi,
	)

	if err != nil {
		return Missoes{}, err
	}

	return missao, nil
}

func AtualizarMissao(idMissao string, missao Missoes) error {
	db := database.ConectaDB()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Atualizar dados básicos da missão
	query := `
        UPDATE missoes 
        SET nome_missao = $1, descricao = $2, nivel_dificuldade = $3, 
            resultado = $4, recompensa = $5
        WHERE id_missao = $6
    `

	_, err = tx.Exec(query,
		missao.NomeMissao,
		missao.DescricaoMissao,
		missao.NivelMissao,
		missao.Resultado,
		missao.Recompensa,
		idMissao,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func DeletarMissao(idMissao string) error {
	db := database.ConectaDB()
	defer db.Close()

	query :=
		`UPDATE missoes 
        SET esconder_missao = true 
        WHERE id_missao = $1`

	result, err := db.Exec(query, idMissao)
	if err != nil {
		return fmt.Errorf("erro ao deletar missão: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar linhas afetadas: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("nenhuma missão encontrada com o ID %s", idMissao)
	}

	return nil
}

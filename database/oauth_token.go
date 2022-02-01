package database

import (
	"github.com/desmos-labs/desmostipbot/types"
	"time"
)

type oAuthTokenRow struct {
	DesmosAddress string `db:"desmos_address"`
	Service       string `db:"service"`
	AccessToken   string `db:"access_token"`
	RefreshToken  string `db:"refresh_token"`
	CreationTime  time.Time
}

// SaveOAuthToken stores the given OAuth token inside the database
func (db *Database) SaveOAuthToken(token *types.OAuthToken) error {
	stmt := `
INSERT INTO oauth_token (service, desmos_address, access_token, refresh_token, creation_time)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT ON CONSTRAINT single_oauth_token DO UPDATE 
    SET access_token = excluded.access_token,
    	refresh_token = excluded.refresh_token,
    	creation_time = excluded.creation_time`
	_, err := db.sql.Exec(stmt, token.Service, token.DesmosAddress, token.AccessToken, token.RefreshToken, time.Now())
	return err
}

// GetOAuthToken returns the OAuth token associated with the given Desmos address
func (db *Database) GetOAuthToken(desmosAddress string) (*types.OAuthToken, error) {
	stmt := `SELECT * FROM oauth_token WHERE desmos_address = $1 ORDER BY creation_time DESC LIMIT 1`

	var rows []oAuthTokenRow
	err := db.sql.Select(&rows, stmt, desmosAddress)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		// Nothing found
		return nil, nil
	}

	return types.NewOAuthToken(
		rows[1].DesmosAddress,
		rows[1].Service,
		rows[1].AccessToken,
		rows[1].RefreshToken,
	), nil
}
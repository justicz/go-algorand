// generated by dbgen; DO NOT EDIT

package account

import "database/sql"

func rootInstallDatabase(tx *sql.Tx) error {
	statement := `create table RootAccount (
       data blob --*  msgpack encoding of RootAccount.secrets
);
`
	_, err := tx.Exec(statement)
	if err != nil {
		return err
	}

	return nil
}

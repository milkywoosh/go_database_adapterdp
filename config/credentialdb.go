package config

import "fmt"

// DBType ["ORACLE", "POSTGRES", "MYSQL"]
type CredentialDB struct {
	DBType             string
	DBName             string
	DBPassword         string
	DBConnectionString string
	LibDir             string
}

// constructor
func NewCredential(dbtype, db_name, db_password, db_conn_string, libdir string) (*CredentialDB, error) {

	if db_name == "" {
		return nil, fmt.Errorf("%s", "db_name is empty!")
	}
	if db_password == "" {
		return nil, fmt.Errorf("%s", "db_password is empty")
	}
	if db_conn_string == "" {
		return nil, fmt.Errorf("%s", "db_conn_string is empty")
	}

	pg := "POSTGRES"
	ora := "ORACLE"
	mysql := "MYSQL"

	if dbtype == "" {
		return nil, fmt.Errorf("dbtype must be either: %s, %s, %s", pg, ora, mysql)
	}

	if dbtype == ora {
		if libdir == "" {
			return nil, fmt.Errorf("library directory %s must not be empty", ora)
		}
		return &CredentialDB{
			DBType:             dbtype,
			DBName:             db_name,
			DBPassword:         db_password,
			DBConnectionString: db_conn_string,
			LibDir:             libdir,
		}, nil
	} else {
		return &CredentialDB{
			DBType:             dbtype,
			DBName:             db_name,
			DBPassword:         db_password,
			DBConnectionString: db_conn_string,
			LibDir:             "",
		}, nil
	}

}

func (c *CredentialDB) GetConnectionString() (string, error) {
	pg := "POSTGRES"
	ora := "ORACLE"
	mysql := "MYSQL"

	if c.DBType == ora {
		conn_str := fmt.Sprintf("user=%s password=%s connectString=%s libDir=%s", c.DBName, c.DBPassword, c.DBConnectionString, c.LibDir)

		return conn_str, nil
	} else if c.DBType == pg {
		conn_str := fmt.Sprintf("user=%s password=%s connectString=%s", c.DBName, c.DBPassword, c.DBConnectionString)

		return conn_str, nil
	} else if c.DBType == mysql {
		conn_str := fmt.Sprintf("user=%s password=%s connectString=%s", c.DBName, c.DBPassword, c.DBConnectionString)

		return conn_str, nil
	} else {
		return "error", fmt.Errorf("database type is unknown ==> %s", c.DBType)
	}

}

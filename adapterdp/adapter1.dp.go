package adapterdp

import "fmt"

// NOTE try to integrate postgres and oracle here
type OracleDB struct{}

func (o *OracleDB) ConnectOra() {
	fmt.Printf("%s", "connect to Oracle")
}

type PostgresDB struct{}

func (pg *PostgresDB) ConnectPg() {
	fmt.Printf("%s", "connect to Postgres")
}

type ConnectionProcessor interface {
	Connecting()
}

type InitConnection struct {
	objectConnectionAdapter ConnectionProcessor
}

func NewInitConnection(objectConnectionAdapter ConnectionProcessor) InitConnection {
	return InitConnection{
		objectConnectionAdapter,
	}
}

func (ic *InitConnection) Connection() {
	ic.objectConnectionAdapter.Connecting()
}

type ConnectionAdapter struct {
	AdapteeDatabase interface{}
}

// satisfy ConnectionProcesssor interface
func (ca *ConnectionAdapter) Connecting() {
	switch dbType := ca.AdapteeDatabase.(type) {
	case *OracleDB:
		dbType.ConnectOra()
	case *PostgresDB:
		dbType.ConnectPg()
	default:
		fmt.Println("Unsupported database")
	}

}

func NewConnAdapter(adapteeDB interface{}) ConnectionAdapter {
	return ConnectionAdapter{
		adapteeDB,
	}
}

package entities

type DBconf struct {
	DriverName string
	Ip         string
	Port       string
	Account    string
	Password   string
	DB         string
	MaxIdle    int
	MaxConn    int
	EnCoding   string
}

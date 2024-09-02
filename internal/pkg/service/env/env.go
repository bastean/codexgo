package env

func Init() {
	SMTP()
	Broker()
	Database()
	JWT()
	Server()
}

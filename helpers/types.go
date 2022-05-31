package helpers

type Map map[string]interface{}

type ContextValues struct {
	UserID      string
	AccessUUID  string
	Role        string
	Permissions string
	UserDetails string
}

type EnviormentVariables struct {
	MYSQL_USERNAME     string
	MYSQL_PASSWORD     string
	MYSQL_HOST_ADDR    string
	MYSQL_HOST_PORT    string
	DBNAME             string
	JWT_ACCESS_SECRET  string
	JWT_REFRESH_SECRET string
	REDIS_DSN          string
	REDIS_PASSWORD     string
	APP_ENV            string
}

type RolesList struct {
	Admin string
	User  string
}

package helpers

var CtxValues = ContextValues{
	UserID:      "userID",
	AccessUUID:  "accessUUID",
	Role:        "role",
	Permissions: "permissions",
	UserDetails: "userDetails",
}

var EnvKeys = EnviormentVariables{
	MYSQL_USERNAME:     "MYSQL_USERNAME",
	MYSQL_PASSWORD:     "MYSQL_PASSWORD",
	DBNAME:             "DBNAME",
	MYSQL_HOST_ADDR:    "MYSQL_HOST_ADDR",
	MYSQL_HOST_PORT:    "MYSQL_HOST_PORT",
	JWT_ACCESS_SECRET:  "JWT_ACCESS_SECRET",
	JWT_REFRESH_SECRET: "JWT_REFRESH_SECRET",
	REDIS_DSN:          "REDIS_DSN",
	REDIS_PASSWORD:     "REDIS_PASSWORD",
	APP_ENV:            "APP_ENV",
}

var UserRoles = RolesList{
	Admin: "admin",
	User:  "user",
}

const (
	CreatedMessage          = "Successfully Inserted!"
	Unauthorized            = "You are not authorized!"
	RequiredInvitationToken = "Invitation Token is required!"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

package service

const config = "config.json"

var (
	Config GlobalConfig
)

type GlobalConfig struct {
	DatabaseFile string `json:"database_file"`
	AdminToken   string `json:"admin_token"`
}

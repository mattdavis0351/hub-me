package auth

import (
	"fmt"

	"github.com/spf13/viper"
)

func getEnvVariable(k string) string {
	viper.SetConfigFile(".env")
	viper.AddConfigPath("../")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error while reading config file: %s", err)
	}

	value, ok := viper.Get(k).(string)

	if !ok {
		fmt.Println("Invalid type assertion")
	}

	return value
}

// func setEnvVariable(t *oauth2.Token) {

// 	os.Setenv("HUB_ME_ACCESS_TOKEN", t.AccessToken)
// 	os.Setenv("HUB_ME_REFRESH_TOKEN", t.RefreshToken)
// 	os.Setenv("HUB_ME_TOKEN_TYPE", t.TokenType)

// }

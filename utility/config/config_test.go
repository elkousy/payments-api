package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_InitConfig_Toml(t *testing.T) {
	//Arrange
	os.Setenv("ENVIRONMENT", "dev")
	//Act
	InitConfig()
	//Assert
	assert.NotEmpty(t, AppName, "AppName")
	assert.NotEmpty(t, AppPort, "AppPort")
	assert.NotEmpty(t, OpsPort, "OpsPort")
	assert.NotEmpty(t, DebugPort, "DebugPort")
	assert.NotEmpty(t, DBHost, "DBHost")
	assert.NotEmpty(t, DBPort, "DBPort")
	assert.NotEmpty(t, DBName, "DBName")
	assert.NotEmpty(t, DBUser, "DBUser")
	assert.NotEmpty(t, DBPassword, "DBPassword")
	assert.NotEmpty(t, DBTimeout, "DBTimeout")
}

func Test_InitConfig_EnvVar(t *testing.T) {
	//Arrange
	os.Setenv("ENVIRONMENT", "docker")
	os.Setenv("APP_PORT", "8080")
	os.Setenv("OPS_PORT", "8081")
	os.Setenv("DB_USER", "raouf")
	os.Setenv("DB_PASSWORD", "raouf")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_NAME", "postgres")
	os.Setenv("DB_TIMEOUT", "5")
	//Act
	InitConfig()
	//Assert
	assert.Equal(t, AppPort, 8080)
	assert.Equal(t, OpsPort, 8081)
	assert.Equal(t, DBHost, "localhost")
	assert.Equal(t, DBPort, 5432)
	assert.Equal(t, DBName, "postgres")
	assert.Equal(t, DBUser, "raouf")
	assert.Equal(t, DBPassword, "raouf")
	assert.Equal(t, DBTimeout, 5)
}

// func TestNewConfig(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		setEnv  func()
// 		want    *Config
// 		wantErr bool
// 	}{
// 		{
// 			name: "new config ok",
// 			setEnv: func() {
// 				os.Setenv("ENVIRONMENT", "dev")
// 				viper.Set("APP_PORT", 9000)
// 				viper.Set("LEGACY_API_HOST", "legacy")
// 				viper.Set("LEGACY_API_TIMEOUT_SEC", 30)
// 				viper.Set("MAX_PICTURES_AT_INSERT", 1)
// 				viper.Set("PARTNERSHIP_WEB_USER_FR", "user_fr")
// 				viper.Set("PARTNERSHIP_WEB_PASSWORD_FR", "password_fr")
// 				viper.Set("PARTNERSHIP_WEB_USER_DE", "user_de")
// 				viper.Set("PARTNERSHIP_WEB_PASSWORD_DE", "password_de")
// 				viper.Set("PARTNERSHIP_IOS_USER_FR", "user_fr")
// 				viper.Set("PARTNERSHIP_IOS_PASSWORD_FR", "password_fr")
// 				viper.Set("PARTNERSHIP_IOS_USER_DE", "user_de")
// 				viper.Set("PARTNERSHIP_IOS_PASSWORD_DE", "password_de")
// 				viper.Set("PARTNERSHIP_ANDROID_USER_FR", "user_fr")
// 				viper.Set("PARTNERSHIP_ANDROID_PASSWORD_FR", "password_fr")
// 				viper.Set("PARTNERSHIP_ANDROID_USER_DE", "user_de")
// 				viper.Set("PARTNERSHIP_ANDROID_PASSWORD_DE", "password_de")
// 				viper.Set("ASSETS_DOMAIN_API_HOST", "assets")
// 				viper.Set("URL_BUILDER_API_HOST", "url")
// 				viper.Set("DELIVERY_API_HOST", "url")
// 				viper.Set("LEGACY_DB_NAME", "dbnew")
// 				viper.Set("LEGACY_DB_HOST", "db_host")
// 				viper.Set("LEGACY_DB_PORT", 7777)
// 				viper.Set("LEGACY_DB_FAILOVER_HOST", "db_failover_host")
// 				viper.Set("LEGACY_DB_FAILOVER_PORT", 7777)
// 				viper.Set("LEGACY_DB_USER", "db_user")
// 				viper.Set("LEGACY_DB_PASSWORD", "db_password")
// 				viper.Set("LEGACY_DB_TIMEOUT", 5)
// 				viper.Set("SUGGESTED_START_TIME", "18:00")
// 				viper.Set("SUGGESTED_END_TIME", "22:00")
// 				viper.Set("SUGGESTED_INTERVAL_MINUTES", 5)
// 				viper.Set("IS_DB_LOG", false)
// 				viper.Set("CATEGORIES_API_HOST", "url")
// 				viper.Set("USERS_DOMAIN_API_HOST", "url")
// 				viper.Set("PAYMENT_API_HOST", "url")
// 				viper.Set("PROMOTIONS_API_HOST", "url")
// 				viper.Set("OFFERS_API_HOST", "url")
// 				viper.Set("ARTICLES_API_HOST", "url")
// 				viper.Set("ARTICLE_LOGS_API_HOST", "url")

// 				viper.Set("DB_HOST", "pq_db_host")
// 				viper.Set("DB_PORT", "pq_db_port")
// 				viper.Set("DB_NAME", "pq_db_name")
// 				viper.Set("DB_USER", "pq_db_user")
// 				viper.Set("DB_PASSWORD", "pq_db_pwd")
// 				viper.Set("DB_TIMEOUT", 5)
// 				viper.Set("CACHE_DURATION_MINUTES", 30)
// 				viper.Set("VALID_LISTING_API_IDS", "1,2,3")

// 				viper.Set("AUTH0_AUDIENCE", "audience")
// 				viper.Set("AUTH0_ISSUER", "issuer")
// 				viper.Set("AUTH0_CERTIFICATE", "certificate")

// 			},
// 			want: &Config{
// 				AppPort:             9000,
// 				LegacyAPIHost:       "legacy",
// 				LegacyAPITimeoutSec: 30,
// 				MaxPicturesAtInsert: 1,
// 				WebPartnership: Partnership{
// 					UserFR:     "user_fr",
// 					PasswordFR: "password_fr",
// 					UserDE:     "user_de",
// 					PasswordDE: "password_de",
// 				},
// 				IosPartnership: Partnership{
// 					UserFR:     "user_fr",
// 					PasswordFR: "password_fr",
// 					UserDE:     "user_de",
// 					PasswordDE: "password_de",
// 				},
// 				AndroidPartnership: Partnership{
// 					UserFR:     "user_fr",
// 					PasswordFR: "password_fr",
// 					UserDE:     "user_de",
// 					PasswordDE: "password_de",
// 				},
// 				AssetsDomainAPIHost:          "assets",
// 				URLBuilderAPIHost:            "url",
// 				UsersAPIHost:                 "url",
// 				LegacyDbName:                 "dbnew",
// 				LegacyDbHost:                 "db_host",
// 				LegacyDbPort:                 7777,
// 				LegacyDbFailoverHost:         "db_failover_host",
// 				LegacyDbFailoverPort:         7777,
// 				LegacyDbUser:                 "db_user",
// 				LegacyDbPassword:             "db_password",
// 				LegacyDbTimeout:              5,
// 				StartTime:                    "18:00",
// 				EndTime:                      "22:00",
// 				IntervalMinutes:              5,
// 				IsDbLog:                      false,
// 				DeliveryAPIHost:              "url",
// 				CategoriesAPIHost:            "url",
// 				DbHost:                       "pq_db_host",
// 				DbPort:                       "pq_db_port",
// 				DbName:                       "pq_db_name",
// 				DbUser:                       "pq_db_user",
// 				DbPassword:                   "pq_db_pwd",
// 				DbTimeout:                    5,
// 				CacheDurationMinutes:         30,
// 				PaymentAPIHost:               "url",
// 				PromotionsAPIHost:            "url",
// 				OffersAPIHost:                "url",
// 				ArticlesAPIHost:              "url",
// 				ArticleLogsAPIHost:           "url",
// 				ValidListingAPIIDs:           "1,2,3",
// 				Auth0Audience:                "audience",
// 				Auth0Issuer:                  "issuer",
// 				Auth0Certificate:             "certificate",
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.setEnv()
// 			got, err := NewConfig()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("NewConfig() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

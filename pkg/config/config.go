package config

import (
	"github.com/spf13/viper"
)

type DataBase struct {
	DBUser     string `mapstructure:"DBUSER"`
	DBName     string `mapstructure:"DBNAME"`
	DBPassword string `mapstructure:"DBPASSWORD"`
	DBHost     string `mapstructure:"DBHOST"`
	DBPort     string `mapstructure:"DBPORT"`
}

type Token struct {
	AdminSecurityKey  string `mapstructure:"ADMIN_TOKENKEY"`
	SellerSecurityKey string `mapstructure:"SELLER_TOKENKEY"`
	UserSecurityKey   string `mapstructure:"USER_TOKENKEY"`
	TemperveryKey     string `mapstructure:"TEMPERVERY_TOKENKEY"`
}

type OTP struct {
	AccountSid string `mapstructure:"Account_SID"`
	AuthToken  string `mapstructure:"Auth_Token"`
	ServiceSid string `mapstructure:"Service_SID"`
}

type S3Bucket struct {
	AccessKeyID     string `mapstructure:"AccessKeyID"`
	AccessKeySecret string `mapstructure:"AccessKeySecret"`
	Region          string `mapstructure:"Region"`
	BucketName      string `mapstructure:"BucketName"`
}

type Razopay struct {
	RazopayKey    string `mapstructure:"RAZOPAYKEY"`
	RazopaySecret string `mapstructure:"PAZOPAYSECRET"`
}

type Config struct {
	DB      DataBase
	Token   Token
	Otp     OTP
	S3aws   S3Bucket
	Razopay Razopay
}

func LoadConfig() (*Config, error) {
	var (
		db      DataBase
		token   Token
		otp     OTP
		s3      S3Bucket
		razopay Razopay
	)

	// viper.SetConfigType(".env")
	// viper.SetConfigName(".env")
	// viper.AddConfigPath("/home/vajid/Brocamp/Mobile-mart/")
	// viper.SetConfigFile("/home/ubuntu/Mobile-Mart/.env")
	viper.SetConfigFile("./.env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&db)
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&token)
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&otp)
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&s3)
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&razopay)
	if err != nil {
		return nil, err
	}

	config := Config{DB: db, Token: token, Otp: otp, S3aws: s3, Razopay: razopay}
	return &config, nil
}

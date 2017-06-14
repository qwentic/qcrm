package config

import (
	"log"
	"math/big"
	"time"

	"github.com/UnnoTed/hide"
	"github.com/spf13/viper"
)

var (
	DBUser  string
	DBPass  string
	DBName  string
	DBParam string
	//EncryptionKey     []byte
	JWTExpirationTime time.Duration
	//JWTSecret         []byte

	//SendInBlue           string
	Port string
)

func init() {
	if err := hide.Default.SetInt32(new(big.Int).SetInt64(1580030173)); err != nil {
		panic(err)
	}

	//if err := hide.Default.SetInt64(new(big.Int).SetInt64(8230452606740808761)); err != nil {
	if err := hide.Default.SetInt64(new(big.Int).SetInt64(1580030173)); err != nil {
		panic(err)
	}

	if err := hide.Default.SetUint32(new(big.Int).SetUint64(1500450271)); err != nil {
		panic(err)
	}

	//if err := hide.Default.SetUint64(new(big.Int).SetUint64(12764787846358441471)); err != nil {
	if err := hide.Default.SetUint64(new(big.Int).SetUint64(1500450271)); err != nil {
		panic(err)
	}

	//if err := hide.Default.SetXor(new(big.Int).SetUint64(3469983624777167712)); err != nil {
	if err := hide.Default.SetXor(new(big.Int).SetUint64(1580030173)); err != nil {
		panic(err)
	}

	// setup database connection from config file
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("No configuration file loaded - using defaults")
	}

	// sets the path to the running path
	log.Println("Loaded config from path:", viper.ConfigFileUsed())
	//viper.SetDefault("port", "8000")

	// database
	DBUser = viper.GetString("dbuser")
	DBPass = viper.GetString("dbpass")
	DBName = viper.GetString("dbname")
	DBParam = viper.GetString("dbparam")
	// encryption / secret
	/*	EncryptionKey = []byte(viper.GetString("encryption_key"))
		JWTSecret = []byte(viper.GetString("jwt_secret"))
	*/
	// jwt
	jet := viper.GetInt("jwt_expiration_time")
	JWTExpirationTime = time.Duration(jet) * time.Hour

	// server
	Port = viper.GetString("port")

}

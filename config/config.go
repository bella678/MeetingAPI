package config
import(
    "log"
    "github.com/BruntSushi/toml"
)

//Represents database server and Credentials
type Config struct{
    Server string
    Database string
}

//Read and parse the configuration file
func(c *Config)Read(){
    if _,err:=toml.DecodeFile("config.toml",&c);err !=nil{
        log.Fatal(err)
    }
}
package main

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
)

type config struct {
	Server struct {
		Listen string
	}
	Database struct {
		Host   string
		Port   string
		User   string
		Pass   string
		Schema string
	}
	App struct {
		Mode string
		Key  string
	}
	Email struct {
		Email string
		Pass  string
	}
}

const configFile = "./config.toml"

func initConfig() (*config, error) {

	//Abrir y leer archivo de configuracion para conexiones
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return nil, errors.New("No se encontro el archivo de configuraciones")
	} else if err != nil {
		return nil, err
	}

	conf := &config{}
	if _, err := toml.DecodeFile(configFile, conf); err != nil {
		return nil, err
	}

	return conf, nil
}

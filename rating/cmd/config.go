package main

type apiConfig struct {
	Port int `yaml:"port"`
}

type config struct {
	API apiConfig `yaml:"api"`
}

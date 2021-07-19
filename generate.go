package main

import(
	"cryptogen"
)

func generate_certs(outputDir string) error{
	cryptogen.OutputDir = outputDir
	if err := cryptogen.Generate();err != nil{
		return err
	}
	return nil
}

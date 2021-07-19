package main

import(
	"cryptogen"
)

func generate_certs(outputDir) error{
	cryptogen.OutputDir = outputDir
	if err := cryptogen.Generate();err != nil{
		return err
	}
	return nil
}

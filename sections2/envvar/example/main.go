package main

import (
	"os"
	"io/ioutil"
	"bytes"
	"github.com/ifirlana/go-solution/sections2/envvar"
	"fmt"
)

type Config struct {
	Version string `json:"version" required:"true"`
	IsSafe string `json:"is_safe" default:"true"`
	Secret string `json:"secret"`
} 

func main()  {
	var err error

	tf, err := ioutil.TempFile("","tmp")
	if err != nil {
		panic(err)
	}

	defer tf.Close()
	defer os.Remove(tf.Name())

	secrets := `{
	"secret": "so secret"
	}`

	if _, err = tf.Write(bytes.NewBufferString(secrets).Bytes());err != nil {
		panic(err)
	}

	if err = os.Setenv("EXAMPLE_VERSION","1.0.0");err != nil {
		panic(err)
	}

	if err = os.Setenv("EXAMPLE_ISSAFE","false");err != nil {
		panic(err)
	}

	c := Config{}
	if err = envvar.LoadConfig(tf.Name(), "EXAMPLE", &c); err != nil {
		panic(err)
	}

	fmt.Println("secret file contains =", secrets)
	fmt.Println("EXAMPLE_VERSIONS = ", os.Getenv("EXAMPLE_VERSION"))
	fmt.Println("EXAMPLE_ISSAFE = ", os.Getenv("EXAMPLE_ISSAFE"))
	fmt.Printf("Final Config: %#v\n", c)
}
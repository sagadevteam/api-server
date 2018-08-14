package yaml_test

import (
	"fmt"
	"log"

<<<<<<< HEAD
	"gopkg.in/yaml.v2"
=======
        "gopkg.in/yaml.v2"
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
)

// An example showing how to unmarshal embedded
// structs from YAML.

type StructA struct {
	A string `yaml:"a"`
}

type StructB struct {
	// Embedded structs are not treated as embedded in YAML by default. To do that,
	// add the ",inline" annotation below
<<<<<<< HEAD
	StructA `yaml:",inline"`
	B       string `yaml:"b"`
=======
	StructA   `yaml:",inline"`
	B string `yaml:"b"`
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
}

var data = `
a: a string from struct A
b: a string from struct B
`

func ExampleUnmarshal_embedded() {
	var b StructB

	err := yaml.Unmarshal([]byte(data), &b)
	if err != nil {
<<<<<<< HEAD
		log.Fatalf("cannot unmarshal data: %v", err)
	}
	fmt.Println(b.A)
	fmt.Println(b.B)
	// Output:
	// a string from struct A
	// a string from struct B
=======
		log.Fatal("cannot unmarshal data: %v", err)
	}
        fmt.Println(b.A)
        fmt.Println(b.B)
        // Output:
        // a string from struct A
        // a string from struct B
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
}

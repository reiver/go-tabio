package tabio_test


import (
	"github.com/reiver/go-tabio"

	"fmt"
	"io"
	"log"
)


var readCloser io.ReadCloser


func ExampleRecordReader_Fields() {
	rows := tabio.NewRecordReader(readCloser)
	defer rows.Close()

	for rows.Next() {

		fields, err := rows.Fields()
		if nil != err {
			log.Fatal(err)
		}

		fmt.Printf("fields = %#v\n", fields)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}


func ExampleRecordReader_Scan() {
	rows := tabio.NewRecordReader(readCloser)
	defer rows.Close()

	for rows.Next() {

		var name string
		if err := rows.Scan(&name); nil != err {
			log.Fatal(err)
		}

		fmt.Printf("name = %q\n", name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}


func ExampleRecordReader_Unmarshal() {
	rows := tabio.NewRecordReader(readCloser)
	defer rows.Close()

	for rows.Next() {

		datum := struct{
			Name string
			Age  int
		}{}

		if err := rows.Unmarshal(&datum); nil != err {
			log.Fatal(err)
		}

		fmt.Printf("datum = %#v\n", datum)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}


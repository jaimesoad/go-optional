package opt

import "log"

func Log(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

func Fatal(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Panic(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func Ok(err error) bool {
	return err == nil
}

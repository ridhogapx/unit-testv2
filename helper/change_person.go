package helper

type Person struct {
	Name string
	Nick string
	Age  int8
}

func ChangePerson(person *Person) *Person {
	*person = Person{
		Name: "John",
		Nick: "Doe",
		Age:  40,
	}

	return person
}

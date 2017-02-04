package animals

type Dog struct {
	Name    string
	Hobbies []string
}

func (d *Dog) Bark() string {
	return "Woof!"
}

func (k *Dog) SetName(name string) {
	k.Name = name
}

func (k *Dog) GetName() string {
	return k.Name
}

func (k *Dog) GetHobbies() []string {
	return k.Hobbies
}

func (k *Dog) SetHobbies(hobbies []string) {
	k.Hobbies = hobbies
}

func (k *Dog) AddHobby(hobby string) {
	k.Hobbies = append(k.Hobbies, hobby)
}

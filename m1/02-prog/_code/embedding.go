package embed

type animal struct{ name string }

func (a animal) Name() string { return "my name is: " + a.name }

type monkey struct{ animal } // monkey embeds 'animal' // HLxxx

func (m monkey) play() string { return m.Name() } // can re-use animal.Name()

package model

type Character struct {
	/* About: CharacterSeq
	In the api specification, the response name is described as userCharacterID,
	but the id is a sequence number(CharacterSeq) and is difficult to understand,
	so the field name is left as a sequence.
	*/
	CharacterSeq int
	Id           int
	Name         string
}

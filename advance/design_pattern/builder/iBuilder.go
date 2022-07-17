package main

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func getBuilder(buiderType string) iBuilder {
	if buiderType == "normal" {
		return &normalBuilder{}
	}

	if buiderType == "igloo" {
		return &iglooBuilder{}
	}
	return nil
}

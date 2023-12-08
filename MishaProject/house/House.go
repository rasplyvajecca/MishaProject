package house

import (
	"MishaProject/MishaProject/children"
	"MishaProject/MishaProject/decorations"
	"MishaProject/MishaProject/family"
	"MishaProject/MishaProject/furniture"
	"MishaProject/MishaProject/rooms"
	"fmt"
)

type House struct {
	Size        float64
	City        string
	HouseFamily []family.Family

	HouseFurniture   []furniture.Furniture
	HouseChildren    []children.Children
	HouseRooms       []rooms.Rooms
	HouseDecorations []decorations.Decorations
}

func NewHouse() House {
	return House{
		Size:             49.999,
		City:             "Москва",
		HouseFamily:      family.NewFamily(),
		HouseChildren:    children.NewChildren(),
		HouseRooms:       rooms.NewRooms(),
		HouseFurniture:   furniture.NewFurniture(),
		HouseDecorations: decorations.NewDecorations(),
	}
}

func NewObjects(house House) {
	fmt.Printf("Новый дом:\n")
	fmt.Printf("Площадь: %.3f кв.м\n", house.Size)
	fmt.Printf("Город: %s\n", house.City)

	fmt.Println("Новая семья:")
	for _, familyObject := range house.HouseFamily {
		fmt.Printf("- %s, %d\n", familyObject.Member, familyObject.Age)
	}

	fmt.Println("Новые дети:")
	for _, childrenObject := range house.HouseChildren {
		fmt.Printf("- %s, %d\n", childrenObject.Name, childrenObject.Age)
	}

	fmt.Println("Новые комнаты:")
	for _, roomsObject := range house.HouseRooms {
		fmt.Printf("- %s: %.3f кв.м, кол-во: %d\n", roomsObject.Room, roomsObject.Size, roomsObject.Count)
	}

	fmt.Println("Новая мебель:")
	for _, furnitureObject := range house.HouseFurniture {
		fmt.Printf("- %s: %.3f кв.м, кол-во: %d,\n", furnitureObject.Object, furnitureObject.Size, furnitureObject.Count)
	}

	fmt.Println("Новый декор:")
	for _, decorationsObject := range house.HouseDecorations {
		fmt.Printf("- %s: цвет: %s, кол-во: %d,\n", decorationsObject.Object, decorationsObject.Color, decorationsObject.Count)
	}
}

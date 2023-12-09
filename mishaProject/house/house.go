package house

import (
	"MishaProject/mishaProject/children"
	"MishaProject/mishaProject/decorations"
	"MishaProject/mishaProject/family"
	"MishaProject/mishaProject/furniture"
	"MishaProject/mishaProject/rooms"
	"fmt"
)

type House struct {
	Size             float64
	City             string
	HouseFamily      []family.Family
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

func (house House) PrintHouseInfo() {
	fmt.Println("Новый дом:")
	fmt.Printf("Площадь: %.3f кв.м\n", house.Size)
	fmt.Printf("Город: %s\n", house.City)
}

func (house House) PrintFamilyInfo() {
	fmt.Println("Новая семья:")
	for _, familyObject := range house.HouseFamily {
		fmt.Printf("- %s, %d\n", familyObject.Member, familyObject.Age)
	}
}

func (house House) PrintChildrenInfo() {
	fmt.Println("Новые дети:")
	for _, childrenObject := range house.HouseChildren {
		fmt.Printf("- %s, %d\n", childrenObject.Name, childrenObject.Age)
	}
}

func (house House) PrintRoomsInfo() {
	fmt.Println("Новые комнаты:")
	for _, roomsObject := range house.HouseRooms {
		fmt.Printf("- %s: %.3f кв.м, кол-во: %d\n", roomsObject.Room, roomsObject.Size, roomsObject.Count)
	}
}

func (house House) PrintFurnitureInfo() {
	fmt.Println("Новая мебель:")
	for _, furnitureObject := range house.HouseFurniture {
		fmt.Printf("- %s: %.3f кв.м, кол-во: %d,\n", furnitureObject.Object, furnitureObject.Size, furnitureObject.Count)
	}
}

func (house House) PrintDecorationsInfo() {
	fmt.Println("Новый декор:")
	for _, decorationsObject := range house.HouseDecorations {
		fmt.Printf("- %s: цвет: %s, кол-во: %d,\n", decorationsObject.Object, decorationsObject.Color, decorationsObject.Count)
	}
}

func NewObjects(house House) {
	house.PrintHouseInfo()
	house.PrintFamilyInfo()
	house.PrintChildrenInfo()
	house.PrintRoomsInfo()
	house.PrintFurnitureInfo()
	house.PrintDecorationsInfo()
}

package main

import (
	"encoding/binary"
	"fmt"
	"github.com/xStrixU/godataserializer"
)

type Player struct {
	Name string
	Level int32
	Exp int32
	Money float64
	Flying bool
}

func (p *Player) Read(serializer *godataserializer.DataSerializer) {
	p.Name = serializer.ReadString()
	p.Level = serializer.ReadInt32()
	p.Exp = serializer.ReadInt32()
	p.Money = serializer.ReadFloat64()
	p.Flying = serializer.ReadBool()
}

func (p *Player) Write(serializer *godataserializer.DataSerializer) {
	serializer.WriteString(p.Name)
	serializer.WriteInt32(p.Level)
	serializer.WriteInt32(p.Exp)
	serializer.WriteFloat64(p.Money)
	serializer.WriteBool(p.Flying)
}

func main() {
	player := Player{
		Name: "xStrixU",
		Level: 52,
		Exp: 1064,
		Money: 781.22,
		Flying: true,
	}

	serializer := godataserializer.NewDataSerializer(binary.LittleEndian)
	player.Write(serializer)

	serializedPlayer := serializer.Bytes() // []byte

	deserializedPlayer := Player{}
	deserializedPlayer.Read(godataserializer.NewDataSerializer(binary.LittleEndian, serializedPlayer))

	fmt.Println(player.Name == deserializedPlayer.Name) // true
	fmt.Println(player.Level == deserializedPlayer.Level) // true
	fmt.Println(player.Exp == deserializedPlayer.Exp) // true
	fmt.Println(player.Money == deserializedPlayer.Money) // true
	fmt.Println(player.Flying == deserializedPlayer.Flying) // true
}
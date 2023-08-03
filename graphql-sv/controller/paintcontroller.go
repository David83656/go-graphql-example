package controller

import (
	"log"

	"github.com/David83656/go-dbgraphql/graphql-sv/database"
	"github.com/David83656/go-dbgraphql/graphql-sv/graph/model"
)

const (
	DB  = "test"
	COL = "test"
)

type ControllerPaintPinturas interface {
	SaveP(paint *model.Producto)
	FindAllP() []*model.Producto
}

type ControllerPaintClientes interface {
	SaveC(client *model.Cliente)
	FindAllC() []*model.Cliente
}

func SaveC(document *model.Cliente) interface{} {
	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer database.Close(client, ctx, cancel)
	cursor, err := database.SaveOne(client, ctx, DB, COL, document)
	if err != nil {
		panic(err)
	}
	return cursor.InsertedID
}

func FindAllC() []*model.Cliente {
	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer database.Close(client, ctx, cancel)

	cursor, err := database.Query(client, ctx, DB, COL)
	if err != nil {
		panic(err)
	}

	var results []*model.Cliente
	for cursor.Next(ctx) {
		var v *model.Cliente
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, v)
	}
	return results
}

func SaveP(document *model.Producto) interface{} {
	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer database.Close(client, ctx, cancel)
	cursor, err := database.SaveOne(client, ctx, DB, COL, document)
	if err != nil {
		panic(err)
	}
	return cursor.InsertedID
}

func FindAllP() []*model.Producto {
	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer database.Close(client, ctx, cancel)

	cursor, err := database.Query(client, ctx, DB, COL)
	if err != nil {
		panic(err)
	}

	var results []*model.Producto
	for cursor.Next(ctx) {
		var v *model.Producto
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, v)
	}
	return results
}

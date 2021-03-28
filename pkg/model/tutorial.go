package model

import (
	"github.com/graphql-go/graphql"
	_ "github.com/mattn/go-sqlite3"
	"go-graphql-tutorial/config"
	_ "go-graphql-tutorial/config"
	_ "go-graphql-tutorial/pkg/mysql"
	mysqlHelper "go-graphql-tutorial/pkg/mysql"
	"time"
)

type Tutorial struct {
	//gorm.Model
	//ID       int
	Id        int64     `gorm:"primary_key;AUTO_INCREMENT" json:"id"` // 主键
	Title     string    `gorm:"title" json:"title"`
	UpdatedAt time.Time `gorm:"updated_at,type:timestamp" json:"updated_at"`
	CreatedAt time.Time `gorm:"created_at,type:timestamp" json:"created_at"`
	//Author   Author
	//Comments []Comment
}

var tutorialType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Tutorial",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"updated_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			//"author": &graphql.Field{
			//	Type: authorType,
			//},
			//"comments": &graphql.Field{
			//	Type: graphql.NewList(commentType),
			//},
		},
	},
)

func init() {
	db := mysqlHelper.NewGormDB(&config.ReadEnv().MysqlURL)
	//db, err := gorm.Open(&config.ReadEnv().MysqlURL)
	//if err != nil {
	//	log.Fatal(err)
	//}
	defer db.Close()
	db.AutoMigrate(&Tutorial{})
	//db.AutoMigrate(&Comment{})
	//db.AutoMigrate(&Author{})
}

//func SingleTutorialSchema() *graphql.Field {
//	return &graphql.Field{
//		Type:        tutorialType,
//		Description: "Get Tutorial By ID",
//		Args: graphql.FieldConfigArgument{
//			"id": &graphql.ArgumentConfig{
//				Type: graphql.Int,
//			},
//		},
//		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
//			var tutorial Tutorial
//			//db, _ := gorm.Open("sqlite3", "tutorials.db")
//			db := mysqlHelper.NewGormDB(&config.ReadEnv().MysqlURL)
//			db.First(&tutorial, params.Args["id"].(int))
//			return tutorial, nil
//		},
//	}
//}

func ListTutorialSchema() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(tutorialType),
		Description: "Get Tutorial List",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var tutorials []Tutorial
			//db, _ := gorm.Open("sqlite3", "tutorials.db")
			db := mysqlHelper.NewGormDB(&config.ReadEnv().MysqlURL)

			db.Where("id>=?", 5).Find(&tutorials)
			return tutorials, nil
		},
	}
}

//func CreateTutorialMutation() *graphql.Field {
//	return &graphql.Field{
//		Type:        tutorialType,
//		Description: "Create a new Tutorial",
//		Args: graphql.FieldConfigArgument{
//			"id": &graphql.ArgumentConfig{
//				Type: graphql.NewNonNull(graphql.Int),
//			},
//			"title": &graphql.ArgumentConfig{
//				Type: graphql.NewNonNull(graphql.String),
//			},
//		},
//		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
//			//tutorial := Tutorial{ID: params.Args["id"].(int), Title: params.Args["title"].(string)}
//			tutorial := Tutorial{ Title: params.Args["title"].(string)}
//
//			//db, _ := gorm.Open("sqlite3", "tutorials.db")
//			db := mysqlHelper.NewGormDB(&config.ReadEnv().MysqlURL)
//			db.Save(&tutorial)
//			//var tutorials Tutorial
//			//
//			//id := params.Args["id"].(int)
//			//db.Where("id=?",id).Find(&tutorials)
//			return tutorial, nil
//		},
//	}
//}

func QueryTutorialMutationByID() *graphql.Field {
	return &graphql.Field{
		Type:        tutorialType,
		Description: "Create a new Tutorial",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"title": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			//tutorial := Tutorial{ID: params.Args["id"].(int), Title: params.Args["title"].(string)}
			//tutorial := Tutorial{ Title: params.Args["title"].(string)}

			//db, _ := gorm.Open("sqlite3", "tutorials.db")
			db := mysqlHelper.NewGormDB(&config.ReadEnv().MysqlURL)
			//db.Save(&tutorial)
			var tutorials Tutorial

			id := params.Args["id"].(int)
			db.Where("id=?", id).Find(&tutorials)
			return tutorials, nil
		},
	}
}

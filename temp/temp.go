package temp

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

//type Podcast struct {
//	ID     primitive.ObjectID `bson:"_id,omitempty"`
//	Name   string             `bson:"name,omitempty"`
//	Author string             `bson:"author,omitempty"`
//	Tags   []string           `bson:"tags,omitempty"`
//}

type Episode struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Podcast     primitive.ObjectID `bson:"podcast,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
}

func smain() {

	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	user := viper.GetString("database.user")
	password := viper.GetString("database.password")

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://" + user + ":" + password + "@cluster0.rkyph.mongodb.net/cluster0?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatalf("%s", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("%s", err)
	}

	testDB := client.Database("cluster0")
	podcastsCollection := testDB.Collection("podcasts")
	episodesCollection := testDB.Collection("episodes")

	podcastsResult, err := podcastsCollection.InsertOne(ctx, bson.D{
		{"title", "The Polyglot Developer Podcast"},
		{"author", "Chamod Perera"},
		{"tags", bson.A{"development", "programming", "coding"}},
	})
	if err != nil {
		log.Fatalf("Unable to Add %s", err)
	}
	fmt.Println(podcastsResult.InsertedID)

	episodesResult, err := episodesCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{"podcast", podcastsResult.InsertedID},
			{"title", "Episode #1"},
			{"description", "This is the first episode"},
			{"duration", 25},
		},
		bson.D{
			{"podcast", podcastsResult.InsertedID},
			{"title", "Episode #2"},
			{"description", "This is the second episode"},
			{"duration", 35},
		},
	})
	if err != nil {
		log.Fatalf("Unable to Add %s", err)
	}
	fmt.Println(episodesResult.InsertedIDs)
}

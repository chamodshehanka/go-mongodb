package temp

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func gmain() {
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

	cursor, err := episodesCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalf("Unable to fetch Data %s", err)
	}

	var episodes []bson.M
	//if err = cursor.All(ctx, &episodes); err != nil {
	//	log.Fatalf("%s", err)
	//}
	//for _, episode := range episodes {
	//	fmt.Println(episode["title"])
	//}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var episode bson.M
		if err = cursor.Decode(&episodes); err != nil {
			log.Fatalf("%s", err)
		}

		fmt.Println(episode)
	}

	var podcast bson.M
	if err = podcastsCollection.FindOne(ctx, bson.M{}).Decode(&podcast); err != nil {
		log.Fatalf("%s",err)
	}
	fmt.Println(podcast)

	filterCursor,err := episodesCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalf("%s", err)
	}
	var episodesFiltered []bson.M
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(episodesFiltered)

	opts := options.Find()
	opts.SetSort(bson.D{{"duration", -1}})
	sortCursor,err := episodesCollection.Find(ctx, bson.D{
		{"duration", bson.D{
			{"$gt", 24},
		}},
	}, opts)

	var episodesSorted []bson.M
	if err = sortCursor.All(ctx, &episodesSorted); err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(episodesSorted)
}

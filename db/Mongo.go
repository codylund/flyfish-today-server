package db

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type DbOperation func(db *mongo.Database)

func Run(op DbOperation) error {
    ctx := context.TODO()
      uri := "mongodb+srv://streamflow.f04mikl.mongodb.net/?authSource=%24external&authMechanism=MONGODB-X509&retryWrites=true&w=majority&tlsCertificateKeyFile=/Users/cody/Downloads/X509-cert-6006564856114604467.pem"
      
    serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
      clientOptions := options.Client().
        ApplyURI(uri).
        SetServerAPIOptions(serverAPIOptions)

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil { 
        return err
    }
    defer client.Disconnect(ctx)

    op(client.Database("Streamflow"))
    return nil
}
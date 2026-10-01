package main

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/firestore"
	"github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/projects"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		project := config.Get(ctx, "gcp:project")
		if project == "" {
			return fmt.Errorf("gcp:project is required")
		}

		region := config.Get(ctx, "gcp:region")
		if region == "" {
			region = "us-central1"
		}

		databaseID := config.Get(ctx, "firestore:databaseId")
		if databaseID == "" {
			databaseID = "(default)"
		}

		databaseType := config.Get(ctx, "firestore:databaseType")
		if databaseType == "" {
			databaseType = "FIRESTORE_NATIVE"
		}

		// Enable Firestore API
		_, err := projects.NewService(ctx, "firestore-api", &projects.ServiceArgs{
			Service:          pulumi.String("firestore.googleapis.com"),
			DisableOnDestroy: pulumi.Bool(false),
			Project:          pulumi.String(project),
		})
		if err != nil {
			return err
		}

		// Create Firestore database
		database, err := firestore.NewDatabase(ctx, "database", &firestore.DatabaseArgs{
			Name:                  pulumi.String(databaseID),
			LocationId:            pulumi.String(region),
			Type:                  pulumi.String(databaseType),
			Project:               pulumi.String(project),
			DeleteProtectionState: pulumi.String("DELETE_PROTECTION_DISABLED"),
			DeletionPolicy:        pulumi.String("DELETE"),
		})
		if err != nil {
			return err
		}

		// Export outputs
		ctx.Export("databaseName", database.Name)
		ctx.Export("databaseId", database.Name)
		ctx.Export("databaseLocation", database.LocationId)

		return nil
	})
}

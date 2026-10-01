# Pulumi Go GCP Firestore Database

This Pulumi project provisions a Google Cloud Firestore database (Native or Datastore mode).

## Resources Created

- **Firestore Database**: Native Firestore or Datastore mode database
- **API Enablement**: `firestore.googleapis.com`

## Prerequisites

- Pulumi CLI installed
- Go 1.21+
- GCP credentials configured
- A GCP project with billing enabled

## Configuration

Copy `Pulumi.dev.yaml.example` to `Pulumi.dev.yaml` and update values:

```yaml
config:
  gcp:project: your-project-id
  gcp:region: us-central1
  firestore:databaseId: "(default)"
  firestore:databaseType: FIRESTORE_NATIVE
```

### Config Options

| Key | Description | Default |
|-----|-------------|---------|
| `gcp:project` | GCP Project ID | **Required** |
| `gcp:region` | GCP Region (Free Tier: us-west1, us-central1, us-east1) | `us-central1` |
| `firestore:databaseId` | Database ID | `(default)` |
| `firestore:databaseType` | Type: FIRESTORE_NATIVE or DATASTORE_MODE | `FIRESTORE_NATIVE` |

## Deployment

```bash
# Install dependencies
go mod tidy

# Preview changes
pulumi preview

# Deploy
pulumi up
```

## Outputs

| Output | Description |
|--------|-------------|
| `databaseName` | The name of the Firestore database |
| `databaseId` | The ID of the Firestore database |
| `databaseLocation` | The location of the Firestore database |

## Database Types

- **FIRESTORE_NATIVE**: Native Firestore mode (recommended for new projects)
- **DATASTORE_MODE**: Datastore mode (for legacy Datastore compatibility)

## Free Tier Notes

Firestore Free Tier includes:
- 1 GB stored data
- 50,000 reads/day
- 20,000 writes/day
- 20,000 deletes/day

Available in Free Tier regions: us-west1, us-central1, us-east1

See [Firestore Pricing](https://cloud.google.com/firestore/pricing) for details.

## Connecting

```go
import (
	"cloud.google.com/go/firestore"
	"context"
)

ctx := context.Background()
client, err := firestore.NewClient(ctx, "your-project-id")
```

## Cleanup

```bash
pulumi destroy
```
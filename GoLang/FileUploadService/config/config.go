package config

import (
	"os"
)

// Load environment variables for S3 configurations
var (
	AWS_S3_REGION = os.Getenv("AWS_S3_REGION")
	AWS_S3_BUCKET = os.Getenv("AWS_S3_BUCKET")
)

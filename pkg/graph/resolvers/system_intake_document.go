package resolvers

import (
	"context"
	"path/filepath"

	"github.com/google/uuid"

	"github.com/cmsgov/easi-app/pkg/appcontext"
	"github.com/cmsgov/easi-app/pkg/graph/model"
	"github.com/cmsgov/easi-app/pkg/models"
	"github.com/cmsgov/easi-app/pkg/storage"
	"github.com/cmsgov/easi-app/pkg/upload"
)

// GetSystemIntakeDocumentsByRequestID fetches all documents attached to the system intake with the given ID.
func GetSystemIntakeDocumentsByRequestID(ctx context.Context, store *storage.Store, s3Client *upload.S3Client, id uuid.UUID) ([]*models.SystemIntakeDocument, error) {
	return store.GetSystemIntakeDocumentsByRequestID(ctx, id)
}

// GetURLForSystemIntakeDocument queries S3 for a presigned URL that can be used to fetch the document with the given s3Key
func GetURLForSystemIntakeDocument(s3Client *upload.S3Client, s3Key string) (string, error) {
	presignedURL, err := s3Client.NewGetPresignedURL(s3Key)
	if err != nil {
		return "", err
	}

	return presignedURL.URL, nil
}

// GetStatusForSystemIntakeDocument queries S3 for the virus-scanning status of a document with the given s3Key
func GetStatusForSystemIntakeDocument(s3Client *upload.S3Client, s3Key string) (models.SystemIntakeDocumentStatus, error) {
	avStatus, err := s3Client.TagValueForKey(s3Key, upload.AVStatusTagName)
	if err != nil {
		return "", err
	}

	// possible tag values come from virus scanning lambda
	// this is the same logic as in schema.resolvers.go's Documents() method for 508 documents
	if avStatus == "CLEAN" {
		return models.SystemIntakeDocumentStatusAvailable, nil
	} else if avStatus == "INFECTED" {
		return models.SystemIntakeDocumentStatusUnavailable, nil
	} else {
		return models.SystemIntakeDocumentStatusPending, nil
	}
}

// CreateSystemIntakeDocument uploads a document to S3, then saves its metadata to our database.
func CreateSystemIntakeDocument(ctx context.Context, store *storage.Store, s3Client *upload.S3Client, input model.CreateSystemIntakeDocumentInput) (*models.SystemIntakeDocument, error) {
	s3Key := uuid.New().String()

	existingExtension := filepath.Ext(input.FileData.Filename)
	if existingExtension != "" {
		s3Key += existingExtension
	} else {
		s3Key += fallbackExtension
	}

	err := s3Client.UploadFile(s3Key, input.FileData.File)
	if err != nil {
		return nil, err
	}

	documentDatabaseRecord := models.SystemIntakeDocument{
		SystemIntakeRequestID: input.RequestID,
		CommonDocumentType:    input.DocumentType,
		FileName:              input.FileData.Filename,
		S3Key:                 s3Key,
		Bucket:                s3Client.GetBucket(),
		// Status isn't saved in database - will be fetched from S3
		// URL isn't saved in database - will be generated by querying S3
	}
	documentDatabaseRecord.CreatedBy = appcontext.Principal(ctx).ID()
	if input.OtherTypeDescription != nil {
		documentDatabaseRecord.OtherType = *input.OtherTypeDescription
	}

	return store.CreateSystemIntakeDocument(ctx, &documentDatabaseRecord)
}

// DeleteSystemIntakeDocument deletes an existing SystemIntakeDocumet, given its ID.
//
// Does *not* delete the uploaded file from S3, following the example of TRB request documents.
func DeleteSystemIntakeDocument(ctx context.Context, store *storage.Store, id uuid.UUID) (*models.SystemIntakeDocument, error) {
	return store.DeleteSystemIntakeDocument(ctx, id)
}

package task

import (
	"github.com/onfleet/gonfleet/resource/destination"
	"github.com/onfleet/gonfleet/resource/metadata"
	"github.com/onfleet/gonfleet/resource/recipient"
)

type TaskState int

const (
	TaskStateUnassigned TaskState = 0
	TaskStateAssigned   TaskState = 1
	TaskStateActive     TaskState = 2
	TaskStateCompleted  TaskState = 3
)

type TaskCompletionEvent struct {
	Name     string               `json:"name,omitempty"`
	Time     int64                `json:"time,omitempty"`
	Location destination.Location `json:"location,omitempty"`
}

type TaskCompletionDetails struct {
	Notes                  string                `json:"notes,omitempty"`
	Success                bool                  `json:"success,omitempty"`
	Time                   *int64                `json:"time,omitempty"`
	Events                 []TaskCompletionEvent `json:"events,omitempty"`
	FailureNotes           string                `json:"failureNotes,omitempty"`
	FailureReason          string                `json:"failureReason,omitempty"`
	SignatureUploadId      *string               `json:"signatureUploadId,omitempty"`
	PhotoUploadId          *string               `json:"photoUploadId,omitempty"`
	PhotoUploadIds         *[]string             `json:"photoUploadIds,omitempty"`
	Actions                []any                 `json:"actions,omitempty"`
	FirstLocation          destination.Location  `json:"firstLocation,omitempty"`
	LastLocation           destination.Location  `json:"lastLocation,omitempty"`
	UnavailableAttachments []any                 `json:"unavailableAttachments,omitempty"`
	Distance               float64               `json:"distance,omitempty"`
}

type TaskOverrides struct {
	RecipientName                 *string `json:"recipientName,omitempty"`
	RecipientNotes                *string `json:"recipientNotes,omitempty"`
	RecipientSkipSmsNotifications *bool   `json:"recipientSkipSMSNotifications,omitempty"`
	UseMerchantForProxy           *string `json:"useMerchantForProxy,omitempty"`
}

type TaskAdditionalQuantities struct {
	QuantityA float32 `json:"quantityA,omitempty"`
	QuantityB float32 `json:"quantityB,omitempty"`
	QuantityC float32 `json:"quantityC,omitempty"`
}

type TaskIdentity struct {
	FailedScanCount int  `json:"failedScanCount,omitempty"`
	Checksum        *any `json:"checksum,omitempty"`
}

type TaskAppearance struct {
	TriangleColor *string `json:"triangleColor,omitempty"`
}

type TaskContainerType string

const (
	TaskContainerTypeWorker       TaskContainerType = "WORKER"
	TaskContainerTypeTeam         TaskContainerType = "TEAM"
	TaskContainerTypeOrganization TaskContainerType = "ORGANIZATION"
)

type TaskContainer struct {
	Type         TaskContainerType `json:"type,omitempty"`
	Worker       string            `json:"worker,omitempty"`
	Team         string            `json:"team,omitempty"`
	Organization string            `json:"organization,omitempty"`
}

type Task struct {
	ID                       string                   `json:"id,omitempty"`
	TimeCreated              int64                    `json:"timeCreated,omitempty"`
	TimeLastModified         int64                    `json:"timeLastModified,omitempty"`
	Organization             string                   `json:"organization,omitempty"`
	ShortId                  string                   `json:"shortId,omitempty"`
	TrackingUrl              string                   `json:"trackingURL,omitempty"`
	Worker                   *string                  `json:"worker,omitempty"`
	Merchant                 string                   `json:"merchant,omitempty"`
	Executor                 string                   `json:"executor,omitempty"`
	Creator                  string                   `json:"creator,omitempty"`
	Dependencies             []string                 `json:"dependencies,omitempty"`
	State                    TaskState                `json:"state,omitempty"`
	CompleteAfter            *int64                   `json:"completeAfter,omitempty"`
	CompleteBefore           *int64                   `json:"completeBefore,omitempty"`
	PickupTask               bool                     `json:"pickupTask,omitempty"`
	Notes                    string                   `json:"notes,omitempty"`
	CompletionDetails        TaskCompletionDetails    `json:"completionDetails,omitempty"`
	Feedback                 []any                    `json:"feedback,omitempty"`
	Metadata                 []metadata.Metadata      `json:"metadata,omitempty"`
	Overrides                TaskOverrides            `json:"overrides,omitempty"`
	Quantity                 int                      `json:"quantity,omitempty"`
	AdditionalQuantities     TaskAdditionalQuantities `json:"additionalQuantities,omitempty"`
	ServiceTime              float32                  `json:"serviceTime,omitempty"`
	Identity                 TaskIdentity             `json:"identity,omitempty"`
	Appearance               TaskAppearance           `json:"appearance,omitempty"`
	ScanOnlyRequiredBarcodes bool                     `json:"scanOnlyRequiredBarcodes,omitempty"`
	Container                TaskContainer            `json:"container,omitempty"`
	TrackingViewed           bool                     `json:"trackingViewed,omitempty"`
	DelayTime                *int64                   `json:"delayTime,omitempty"`
	EstimatedCompletionTime  *int64                   `json:"estimatedCompletionTime,omitempty"`
	EstimatedArrivalTime     *int64                   `json:"estimatedArrivalTime,omitempty"`
	Eta                      *int64                   `json:"eta,omitempty"`
	Recipients               []recipient.Recipient    `json:"recipients,omitempty"`
	// need destination and barcodes still
}

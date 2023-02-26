package types

type TaskState int

const (
	TaskStateUnassigned TaskState = 0
	TaskStateAssigned   TaskState = 1
	TaskStateActive     TaskState = 2
	TaskStateCompleted  TaskState = 3
)

type TaskCompletionEvent struct {
	Location Location `json:"location,omitempty"`
	Name     string   `json:"name,omitempty"`
	Time     int64    `json:"time,omitempty"`
}

type TaskCompletionDetails struct {
	Actions                []any                 `json:"actions,omitempty"`
	Distance               float64               `json:"distance,omitempty"`
	Events                 []TaskCompletionEvent `json:"events,omitempty"`
	FailureNotes           string                `json:"failureNotes,omitempty"`
	FailureReason          string                `json:"failureReason,omitempty"`
	FirstLocation          Location              `json:"firstLocation,omitempty"`
	LastLocation           Location              `json:"lastLocation,omitempty"`
	Notes                  string                `json:"notes,omitempty"`
	PhotoUploadId          *string               `json:"photoUploadId,omitempty"`
	PhotoUploadIds         *[]string             `json:"photoUploadIds,omitempty"`
	SignatureUploadId      *string               `json:"signatureUploadId,omitempty"`
	Success                bool                  `json:"success,omitempty"`
	Time                   *int64                `json:"time,omitempty"`
	UnavailableAttachments []any                 `json:"unavailableAttachments,omitempty"`
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
	Checksum        *any `json:"checksum,omitempty"`
	FailedScanCount int  `json:"failedScanCount,omitempty"`
}

type TaskAppearance struct {
	TriangleColor *string `json:"triangleColor,omitempty"`
}

type TaskContainerType string

const (
	TaskContainerTypeOrganization TaskContainerType = "ORGANIZATION"
	TaskContainerTypeTeam         TaskContainerType = "TEAM"
	TaskContainerTypeWorker       TaskContainerType = "WORKER"
)

type TaskContainer struct {
	Organization string            `json:"organization,omitempty"`
	Team         string            `json:"team,omitempty"`
	Type         TaskContainerType `json:"type,omitempty"`
	Worker       string            `json:"worker,omitempty"`
}

type Task struct {
	AdditionalQuantities TaskAdditionalQuantities `json:"additionalQuantities,omitempty"`
	Appearance           TaskAppearance           `json:"appearance,omitempty"`
	CompleteAfter        *int64                   `json:"completeAfter,omitempty"`
	CompleteBefore       *int64                   `json:"completeBefore,omitempty"`
	CompletionDetails    TaskCompletionDetails    `json:"completionDetails,omitempty"`
	Container            TaskContainer            `json:"container,omitempty"`
	Creator              string                   `json:"creator,omitempty"`
	// DelayTime is how late the task is estimated to be in seconds
	DelayTime                *float32      `json:"delayTime,omitempty"`
	Dependencies             []string      `json:"dependencies,omitempty"`
	EstimatedArrivalTime     *int64        `json:"estimatedArrivalTime,omitempty"`
	EstimatedCompletionTime  *int64        `json:"estimatedCompletionTime,omitempty"`
	Eta                      *int64        `json:"eta,omitempty"`
	Executor                 string        `json:"executor,omitempty"`
	Feedback                 []any         `json:"feedback,omitempty"`
	ID                       string        `json:"id,omitempty"`
	Identity                 TaskIdentity  `json:"identity,omitempty"`
	Merchant                 string        `json:"merchant,omitempty"`
	Metadata                 []Metadata    `json:"metadata,omitempty"`
	Notes                    string        `json:"notes,omitempty"`
	Organization             string        `json:"organization,omitempty"`
	Overrides                TaskOverrides `json:"overrides,omitempty"`
	PickupTask               bool          `json:"pickupTask,omitempty"`
	Quantity                 int           `json:"quantity,omitempty"`
	Recipients               []Recipient   `json:"recipients,omitempty"`
	ScanOnlyRequiredBarcodes bool          `json:"scanOnlyRequiredBarcodes,omitempty"`
	ServiceTime              float32       `json:"serviceTime,omitempty"`
	ShortId                  string        `json:"shortId,omitempty"`
	State                    TaskState     `json:"state,omitempty"`
	TimeCreated              int64         `json:"timeCreated,omitempty"`
	TimeLastModified         int64         `json:"timeLastModified,omitempty"`
	TrackingUrl              string        `json:"trackingURL,omitempty"`
	TrackingViewed           bool          `json:"trackingViewed,omitempty"`
	Worker                   *string       `json:"worker,omitempty"`
	// need destination and barcodes still
}

package onfleet

type Task struct {
	AdditionalQuantities     TaskAdditionalQuantities `json:"additionalQuantities"`
	Appearance               TaskAppearance           `json:"appearance"`
	Barcodes                 *TaskBarcodeContainer    `json:"barcodes,omitempty"`
	CompleteAfter            *int64                   `json:"completeAfter"`
	CompleteBefore           *int64                   `json:"completeBefore"`
	CompletionDetails        TaskCompletionDetails    `json:"completionDetails"`
	Container                *TaskContainer           `json:"container"`
	Creator                  string                   `json:"creator"`
	DelayTime                *float64                 `json:"delayTime"`
	Dependencies             []string                 `json:"dependencies"`
	Destination              Destination              `json:"destination"`
	EstimatedArrivalTime     *int64                   `json:"estimatedArrivalTime"`
	EstimatedCompletionTime  *int64                   `json:"estimatedCompletionTime"`
	Eta                      *int64                   `json:"eta"`
	Executor                 string                   `json:"executor"`
	Feedback                 []any                    `json:"feedback"`
	ID                       string                   `json:"id"`
	Identity                 TaskIdentity             `json:"identity"`
	Merchant                 string                   `json:"merchant"`
	Metadata                 []Metadata               `json:"metadata"`
	Notes                    string                   `json:"notes"`
	Organization             string                   `json:"organization"`
	Overrides                TaskOverrides            `json:"overrides"`
	PickupTask               bool                     `json:"pickupTask"`
	Quantity                 float64                  `json:"quantity"`
	Recipients               []Recipient              `json:"recipients"`
	ScanOnlyRequiredBarcodes bool                     `json:"scanOnlyRequiredBarcodes"`
	ServiceTime              float64                  `json:"serviceTime"`
	ShortId                  string                   `json:"shortId"`
	// SourceTaskId only set on cloned tasks
	SourceTaskId     string    `json:"sourceTaskId,omitempty"`
	State            TaskState `json:"state"`
	TimeCreated      int64     `json:"timeCreated"`
	TimeLastModified int64     `json:"timeLastModified"`
	TrackingUrl      string    `json:"trackingURL"`
	TrackingViewed   bool      `json:"trackingViewed"`
	Worker           *string   `json:"worker"`
}

type TaskState int

const (
	TaskStateUnassigned TaskState = 0
	TaskStateAssigned   TaskState = 1
	TaskStateActive     TaskState = 2
	TaskStateCompleted  TaskState = 3
)

type TaskCompletionEvent struct {
	Location DestinationLocation `json:"location"`
	Name     string              `json:"name"`
	Time     int64               `json:"time"`
}

type TaskCompletionDetails struct {
	Actions                []any                 `json:"actions"`
	Distance               float64               `json:"distance"`
	Events                 []TaskCompletionEvent `json:"events"`
	FailureNotes           string                `json:"failureNotes"`
	FailureReason          string                `json:"failureReason"`
	FirstLocation          DestinationLocation   `json:"firstLocation"`
	LastLocation           DestinationLocation   `json:"lastLocation"`
	Notes                  string                `json:"notes"`
	PhotoUploadId          *string               `json:"photoUploadId"`
	PhotoUploadIds         *[]string             `json:"photoUploadIds"`
	SignatureUploadId      *string               `json:"signatureUploadId"`
	Success                bool                  `json:"success"`
	Time                   *int64                `json:"time"`
	UnavailableAttachments []any                 `json:"unavailableAttachments"`
}

type TaskOverrides struct {
	RecipientName                 *string `json:"recipientName"`
	RecipientNotes                *string `json:"recipientNotes"`
	RecipientSkipSmsNotifications *bool   `json:"recipientSkipSMSNotifications"`
	UseMerchantForProxy           *string `json:"useMerchantForProxy"`
}

type TaskAdditionalQuantities struct {
	QuantityA float64 `json:"quantityA"`
	QuantityB float64 `json:"quantityB"`
	QuantityC float64 `json:"quantityC"`
}

type TaskIdentity struct {
	Checksum        *any `json:"checksum"`
	FailedScanCount int  `json:"failedScanCount"`
}

type TaskAppearance struct {
	TriangleColor *string `json:"triangleColor"`
}

type TaskContainer struct {
	Organization string        `json:"organization,omitempty"`
	Team         string        `json:"team,omitempty"`
	Type         ContainerType `json:"type"`
	Worker       string        `json:"worker,omitempty"`
}

type TaskBarcodeContainer struct {
	Captured []TaskCapturedBarcode `json:"captured"`
	Required []TaskBarcode         `json:"required"`
}

type TaskBarcode struct {
	BlockCompletion bool            `json:"blockCompletion"`
	Data            TaskBarcodeData `json:"data,omitempty"`
}

type TaskCapturedBarcode struct {
	Data         TaskBarcodeData     `json:"data"`
	ID           string              `json:"id"`
	Location     DestinationLocation `json:"location"`
	Symbology    string              `json:"symbology"`
	Time         int64               `json:"time"`
	WasRequested bool                `json:"wasRequested"`
}

type TaskBarcodeData string

type TaskParams struct {
	Appearance     *TaskAppearanceParam `json:"appearance,omitempty"`
	AutoAssign     *TaskAutoAssignParam `json:"autoAssign,omitempty"`
	CompleteAfter  int64                `json:"completeAfter,omitempty"`
	CompleteBefore int64                `json:"completeBefore,omitempty"`
	Container      *TaskContainer       `json:"container,omitempty"`
	Dependencies   []string             `json:"dependencies,omitempty"`
	// Destination can string destination id or destination object onfleet.Destination
	Destination    any        `json:"destination,omitempty"`
	Executor       string     `json:"executor,omitempty"`
	Merchant       string     `json:"merchant,omitempty"`
	Metadata       []Metadata `json:"metadata,omitempty"`
	Notes          string     `json:"notes,omitempty"`
	PickupTask     bool       `json:"pickupTask"`
	Quantity       float64    `json:"quantity,omitempty"`
	RecipientName  string     `json:"recipientName,omitempty"`
	RecipientNotes string     `json:"recipientNotes,omitempty"`
	// Recipients can be slice of string recipient ids or recipient objects []onfleet.Recipient
	Recipients                    any                              `json:"recipients,omitempty"`
	RecipientSkipSmsNotifications bool                             `json:"recipientSkipSMSNotifications,omitempty"`
	Requirements                  *TaskCompletionRequirementsParam `json:"requirements,omitempty"`
	ScanOnlyRequiredBarcodes      bool                             `json:"scanOnlyRequiredBarcodes,omitempty"`
	ServiceTime                   float64                          `json:"serviceTime,omitempty"`
	UseMerchantForProxy           bool                             `json:"useMerchantForProxy,omitempty"`
}

type TaskAutoAssignMode string

const (
	TaskAutoAssignModeDistance TaskAutoAssignMode = "distance"
	TaskAutoAssignModeLoad     TaskAutoAssignMode = "load"
)

type TaskAutoAssignParam struct {
	ConsiderDependencies bool               `json:"considerDependencies,omitempty"`
	ExcludedWorkerIds    []string           `json:"excludedWorkerIds,omitempty"`
	MaxAssignedTaskCount int                `json:"maxAssignedTaskCount,omitempty"`
	Mode                 TaskAutoAssignMode `json:"mode"`
	Team                 string             `json:"team,omitempty"`
}

type TaskCompletionRequirementsParam struct {
	MinimumAge int  `json:"minimumAge,omitempty"`
	Notes      bool `json:"notes,omitempty"`
	Photo      bool `json:"photo,omitempty"`
	Signature  bool `json:"signature,omitempty"`
}

type TaskAppearanceParam struct {
	TriangleColor string `json:"triangleColor"`
}

type TaskBatchCreateParams struct {
	Tasks []TaskParams `json:"tasks"`
}

type TaskBatchCreateResponse struct {
	Tasks  []Task                 `json:"tasks"`
	Errors []TaskBatchCreateError `json:"errors"`
}

type TaskBatchCreateError struct {
	Error RequestErrorMessage `json:"error"`
	Task  TaskParams          `json:"task"`
}

type TaskForceCompletionParams struct {
	CompletionDetails TaskForceCompletionDetailsParam `json:"completionDetails"`
}

type TaskForceCompletionDetailsParam struct {
	Success bool   `json:"success"`
	Notes   string `json:"notes,omitempty"`
}

type TaskCloneParams struct {
	IncludeBarcodes     bool                     `json:"includeBarcodes"`
	IncludeDependencies bool                     `json:"includeDependencies"`
	IncludeMetadata     bool                     `json:"includeMetadata"`
	Overrides           *TaskCloneOverridesParam `json:"overrides,omitempty"`
}

type TaskCloneOverridesParam struct {
	CompleteAfter  int64 `json:"completeAfter,omitempty"`
	CompleteBefore int64 `json:"completeBefore,omitempty"`
	// Destination can string destination id or destination object onfleet.Destination
	Destination any        `json:"destination,omitempty"`
	Metadata    []Metadata `json:"metadata,omitempty"`
	Notes       string     `json:"notes,omitempty"`
	PickupTask  bool       `json:"pickupTask"`
	// Recipients can be slice of string recipient ids or recipient objects []onfleet.Recipient
	Recipients  any     `json:"recipients,omitempty"`
	ServiceTime float64 `json:"serviceTime,omitempty"`
}

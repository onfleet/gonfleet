package onfleet

// Onfleet Task.
// Reference https://docs.onfleet.com/reference/tasks.
type Task struct {
	AdditionalQuantities     TaskAdditionalQuantities `json:"additionalQuantities"`
	Appearance               TaskAppearance           `json:"appearance"`
	Barcodes                 TaskBarcodeContainer     `json:"barcodes,omitempty"`
	CompleteAfter            *int64                   `json:"completeAfter"`
	CompleteBefore           *int64                   `json:"completeBefore"`
	CompletionDetails        TaskCompletionDetails    `json:"completionDetails"`
	Container                TaskContainer            `json:"container"`
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
	Quantity                 int                      `json:"quantity"`
	Recipients               []Recipient              `json:"recipients"`
	ScanOnlyRequiredBarcodes bool                     `json:"scanOnlyRequiredBarcodes"`
	ServiceTime              float32                  `json:"serviceTime"`
	ShortId                  string                   `json:"shortId"`
	State                    TaskState                `json:"state"`
	TimeCreated              int64                    `json:"timeCreated"`
	TimeLastModified         int64                    `json:"timeLastModified"`
	TrackingUrl              string                   `json:"trackingURL"`
	TrackingViewed           bool                     `json:"trackingViewed"`
	Worker                   *string                  `json:"worker"`
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
	QuantityA float32 `json:"quantityA"`
	QuantityB float32 `json:"quantityB"`
	QuantityC float32 `json:"quantityC"`
}

type TaskIdentity struct {
	Checksum        *any `json:"checksum"`
	FailedScanCount int  `json:"failedScanCount"`
}

type TaskAppearance struct {
	TriangleColor *string `json:"triangleColor"`
}

type TaskContainer struct {
	Organization string        `json:"organization"`
	Team         string        `json:"team"`
	Type         ContainerType `json:"type"`
	Worker       string        `json:"worker"`
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

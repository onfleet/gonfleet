package task

import (
	"net/http"

	"github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/netwrk"
)

type Client struct {
	apiKey       string
	rlHttpClient *netwrk.RlHttpClient
	url          string
	call         netwrk.Caller
}

func Plug(apiKey string, rlHttpClient *netwrk.RlHttpClient, url string, call netwrk.Caller) *Client {
	return &Client{
		apiKey:       apiKey,
		rlHttpClient: rlHttpClient,
		url:          url,
		call:         call,
	}
}

// Reference https://docs.onfleet.com/reference/get-single-task
func (c *Client) Get(taskId string) (onfleet.Task, error) {
	task := onfleet.Task{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		[]string{taskId},
		nil,
		nil,
		&task,
	)
	return task, err
}

// Reference https://docs.onfleet.com/reference/get-single-task-by-shortid
func (c *Client) GetByShortId(taskShortId string) (onfleet.Task, error) {
	task := onfleet.Task{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		[]string{"shortId", taskShortId},
		nil,
		nil,
		&task,
	)
	return task, err
}

// Reference https://docs.onfleet.com/reference/list-tasks
func (c *Client) List(params onfleet.TaskListQueryParams) (onfleet.TasksPaginated, error) {
	paginatedTasks := onfleet.TasksPaginated{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		[]string{"all"},
		params,
		nil,
		&paginatedTasks,
	)
	return paginatedTasks, err
}

// Reference https://docs.onfleet.com/reference/querying-by-metadata
func (c *Client) ListWithMetadataQuery(metadata []onfleet.Metadata) ([]onfleet.Task, error) {
	tasks := []onfleet.Task{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPost,
		c.url,
		[]string{"metadata"},
		nil,
		metadata,
		&tasks,
	)
	return tasks, err
}

// Reference https://docs.onfleet.com/reference/create-task
func (c *Client) Create(params onfleet.TaskParams) (onfleet.Task, error) {
	task := onfleet.Task{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPost,
		c.url,
		nil,
		nil,
		params,
		&task,
	)
	return task, err
}

// Reference https://docs.onfleet.com/reference/create-tasks-in-batch
func (c *Client) BatchCreate(params onfleet.TaskBatchCreateParams) (onfleet.TaskBatchCreateResponse, error) {
	batchTasks := onfleet.TaskBatchCreateResponse{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPost,
		c.url,
		[]string{"batch"},
		nil,
		params,
		&batchTasks,
	)
	return batchTasks, err
}

// Reference https://docs.onfleet.com/reference/create-tasks-in-batch-async
func (c *Client) BatchCreateAsync(params onfleet.TaskBatchCreateParams) (onfleet.TaskBatchCreateResponseAsync, error) {
	batchRes := onfleet.TaskBatchCreateResponseAsync{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPost,
		c.url,
		[]string{"batch-async"},
		nil,
		params,
		&batchRes,
	)
	return batchRes, err
}

// Reference https://docs.onfleet.com/reference/batch-job-status
func (c *Client) GetBatchJobStatus(batchJobId string) (onfleet.TaskBatchStatusResponseAsync, error) {
	batchStatus := onfleet.TaskBatchStatusResponseAsync{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodGet,
		c.url,
		[]string{"batch", batchJobId},
		nil,
		nil,
		&batchStatus,
	)
	return batchStatus, err
}

// Reference https://docs.onfleet.com/reference/update-task
func (c *Client) Update(taskId string, params onfleet.TaskParams) (onfleet.Task, error) {
	task := onfleet.Task{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPut,
		c.url,
		[]string{taskId},
		nil,
		params,
		&task,
	)
	return task, err
}

// Reference https://docs.onfleet.com/reference/complete-task
func (c *Client) ForceComplete(taskId string, params onfleet.TaskForceCompletionParams) error {
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPost,
		c.url,
		[]string{taskId, "complete"},
		nil,
		params,
		nil,
	)
	return err
}

// Reference https://docs.onfleet.com/reference/clone-task
func (c *Client) Clone(taskId string, params *onfleet.TaskCloneParams) (onfleet.Task, error) {
	task := onfleet.Task{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPost,
		c.url,
		[]string{taskId, "clone"},
		nil,
		params,
		&task,
	)
	return task, err
}

// Reference https://docs.onfleet.com/reference/delete-task
func (c *Client) Delete(taskId string) error {
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodDelete,
		c.url,
		[]string{taskId},
		nil,
		nil,
		nil,
	)
	return err
}

// Reference https://docs.onfleet.com/reference/automatically-assign-list-of-tasks
func (c *Client) AutoAssignMulti(params onfleet.TaskAutoAssignMultiParams) (onfleet.TaskAutoAssignMultiResponse, error) {
	autoAssignMulti := onfleet.TaskAutoAssignMultiResponse{}
	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPost,
		c.url,
		[]string{"autoAssign"},
		nil,
		params,
		&autoAssignMulti,
	)
	return autoAssignMulti, err
}

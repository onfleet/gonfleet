package testingutil

import (
	"testing"
)


// GetStringPtr returns a pointer to a string (useful for optional fields)
func GetStringPtr(s string) *string {
	return &s
}

// GetIntPtr returns a pointer to an int
func GetIntPtr(i int) *int {
	return &i
}

// GetInt64Ptr returns a pointer to an int64
func GetInt64Ptr(i int64) *int64 {
	return &i
}

// GetFloat64Ptr returns a pointer to a float64
func GetFloat64Ptr(f float64) *float64 {
	return &f
}

// GetBoolPtr returns a pointer to a bool
func GetBoolPtr(b bool) *bool {
	return &b
}

// RunTableTest runs table-driven tests with a consistent pattern
func RunTableTest(t *testing.T, testName string, tests []TableTestCase, testFunc func(*testing.T, TableTestCase)) {
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			testFunc(t, tt)
		})
	}
}

// TableTestCase represents a single test case in a table-driven test
type TableTestCase struct {
	Name     string
	Input    interface{}
	Expected interface{}
	Error    string
	Setup    func()
	Cleanup  func()
}

// SetupTest performs common test setup operations
func SetupTest(t *testing.T) *MockHTTPClient {
	t.Helper()
	return NewMockHTTPClient(t)
}

// CleanupTest performs common test cleanup operations
func CleanupTest(t *testing.T, mockClient *MockHTTPClient) {
	t.Helper()
	if mockClient != nil {
		mockClient.Reset()
	}
}
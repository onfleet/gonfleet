package testingutil

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

// AssertEqual compares two values for equality and fails the test if they don't match
func AssertEqual(t *testing.T, expected, actual interface{}) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}

// AssertNotEqual compares two values and fails the test if they match
func AssertNotEqual(t *testing.T, notExpected, actual interface{}) {
	t.Helper()
	if reflect.DeepEqual(notExpected, actual) {
		t.Errorf("Expected values to be different, but both were %+v", actual)
	}
}

// AssertNil checks if a value is nil and fails the test if it's not
func AssertNil(t *testing.T, value interface{}) {
	t.Helper()
	if value != nil {
		t.Errorf("Expected nil, got %+v", value)
	}
}

// AssertNotNil checks if a value is not nil and fails the test if it is
func AssertNotNil(t *testing.T, value interface{}) {
	t.Helper()
	if value == nil {
		t.Error("Expected non-nil value, got nil")
	}
}

// AssertError checks if an error occurred and fails the test if it didn't
func AssertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

// AssertNoError checks if no error occurred and fails the test if one did
func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// AssertErrorContains checks if an error contains a specific substring
func AssertErrorContains(t *testing.T, err error, substring string) {
	t.Helper()
	if err == nil {
		t.Error("Expected an error, got nil")
		return
	}
	if !strings.Contains(err.Error(), substring) {
		t.Errorf("Expected error to contain '%s', got '%s'", substring, err.Error())
	}
}

// AssertTrue checks if a condition is true
func AssertTrue(t *testing.T, condition bool) {
	t.Helper()
	if !condition {
		t.Error("Expected condition to be true, got false")
	}
}

// AssertFalse checks if a condition is false
func AssertFalse(t *testing.T, condition bool) {
	t.Helper()
	if condition {
		t.Error("Expected condition to be false, got true")
	}
}

// AssertStringContains checks if a string contains a substring
func AssertStringContains(t *testing.T, str, substring string) {
	t.Helper()
	if !strings.Contains(str, substring) {
		t.Errorf("Expected string '%s' to contain '%s'", str, substring)
	}
}

// AssertLen checks if a slice or map has the expected length
func AssertLen(t *testing.T, obj interface{}, expectedLen int) {
	t.Helper()
	val := reflect.ValueOf(obj)
	actualLen := val.Len()
	if actualLen != expectedLen {
		t.Errorf("Expected length %d, got %d", expectedLen, actualLen)
	}
}

// AssertTimeWithin checks if a time is within a duration of an expected time
func AssertTimeWithin(t *testing.T, expected time.Time, actual time.Time, delta time.Duration) {
	t.Helper()
	diff := actual.Sub(expected)
	if diff < 0 {
		diff = -diff
	}
	if diff > delta {
		t.Errorf("Expected time within %v of %v, got %v (diff: %v)", delta, expected, actual, diff)
	}
}


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
package Backend_Challenge_test

import (
	"fmt"
	"go_study/rest_api"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	// Test setup: set environment variables
	fmt.Println("Setting up environment variable for tests")
	os.Setenv("TEST_ENV", "true")

	// Run the tests
	code := m.Run()

	// Test teardown: clean up environment variables
	fmt.Println("Cleaning up environment variable after tests")
	os.Unsetenv("TEST_ENV")

	// Exit with the test result code
	os.Exit(code)
}

func TestEcho(t *testing.T) {
	rest_api.SetCSVData([][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}})
	request := httptest.NewRequest("GET", "/echo", nil)
	responseRecorder := httptest.NewRecorder()
	rest_api.Echo(responseRecorder, request)

	expected := "1,2,3\n4,5,6\n7,8,9"
	if responseRecorder.Body.String() != expected {
		t.Errorf("expected %q but got %q", expected, responseRecorder.Body.String())
	}
}

func TestInvert(t *testing.T) {
	rest_api.SetCSVData([][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}})
	request := httptest.NewRequest("GET", "/invert", nil)
	responseRecorder := httptest.NewRecorder()

	rest_api.Invert(responseRecorder, request)

	expected := "1,4,7\n2,5,8\n3,6,9"
	if responseRecorder.Body.String() != expected {
		t.Errorf("expected %q but got %q", expected, responseRecorder.Body.String())
	}
}

func TestFlatten(t *testing.T) {
	rest_api.SetCSVData([][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}})
	request := httptest.NewRequest("GET", "/flatten", nil)
	responseRecorder := httptest.NewRecorder()

	rest_api.Flatten(responseRecorder, request)

	expected := "1,2,3,4,5,6,7,8,9"
	if responseRecorder.Body.String() != expected {
		t.Errorf("expected %q but got %q", expected, responseRecorder.Body.String())
	}
}

func TestSum(t *testing.T) {
	rest_api.SetCSVData([][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}})
	request := httptest.NewRequest("GET", "/sum", nil)
	responseRecorder := httptest.NewRecorder()

	rest_api.Sum(responseRecorder, request)

	expected := "45"
	if strings.TrimSpace(responseRecorder.Body.String()) != expected {
		t.Errorf("expected %q but got %q", expected, responseRecorder.Body.String())
	}
}

func TestMultiply(t *testing.T) {
	rest_api.SetCSVData([][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}})
	request := httptest.NewRequest("GET", "/multiply", nil)
	responseRecorder := httptest.NewRecorder()

	rest_api.Multiply(responseRecorder, request)

	expected := "362880"
	if strings.TrimSpace(responseRecorder.Body.String()) != expected {
		t.Errorf("expected %q but got %q", expected, responseRecorder.Body.String())
	}
}

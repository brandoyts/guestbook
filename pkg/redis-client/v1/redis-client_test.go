package redisClient

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRedisClient is a mock implementation of the RedisClient interface
type MockRedisClient struct {
	mock.Mock
}

func (m *MockRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	args := m.Called(ctx, key, value, expiration)
	return args.Get(0).(*redis.StatusCmd)
}

func (m *MockRedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	args := m.Called(ctx, key)
	return args.Get(0).(*redis.StringCmd)
}

func TestSet(t *testing.T) {
	// Create a mock redis client
	mockClient := new(MockRedisClient)

	// Prepare mock for Set command
	mockClient.On("Set", context.Background(), "test-key", "test-value", time.Duration(0)).Return(&redis.StatusCmd{})

	// Inject the mock client into SetClient
	SetClient(mockClient)

	// Call Set function
	Set("test-key", "test-value")

	// Assert that the Set method was called with the expected arguments
	mockClient.AssertExpectations(t)
}

func TestGet(t *testing.T) {
	// Create a mock redis client
	mockClient := new(MockRedisClient)

	// Prepare mock for Get command
	cmd := redis.NewStringCmd(context.Background(), "test-key")
	cmd.SetVal("test-value") // Setting the value that we expect in the result

	// Prepare mock for Get command to return the command with the value
	mockClient.On("Get", context.Background(), "test-key").Return(cmd)

	// Inject the mock client into SetClient
	SetClient(mockClient)

	// Call Get function
	value := Get("test-key")

	// Assert that the value returned is as expected
	assert.Equal(t, "test-value", value)

	// Assert that the Get method was called with the expected arguments
	mockClient.AssertExpectations(t)
}

func TestXxx(t *testing.T) {
	getRedisClient()
}

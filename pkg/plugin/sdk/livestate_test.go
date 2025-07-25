// Copyright 2025 The PipeCD Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sdk

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	config "github.com/pipe-cd/pipecd/pkg/configv1"
	"github.com/pipe-cd/pipecd/pkg/model"
	"github.com/pipe-cd/pipecd/pkg/plugin/api/v1alpha1/common"
	"github.com/pipe-cd/pipecd/pkg/plugin/api/v1alpha1/livestate"
)

type mockLivestatePlugin struct {
	result *GetLivestateResponse
	err    error
}

func (m *mockLivestatePlugin) GetLivestate(ctx context.Context, config ConfigNone, targets DeployTargetsNone, input *GetLivestateInput[struct{}]) (*GetLivestateResponse, error) {
	return m.result, m.err
}

func newTestLivestatePluginServer(t *testing.T, plugin *mockLivestatePlugin) *LivestatePluginServer[struct{}, struct{}, struct{}] {
	return &LivestatePluginServer[struct{}, struct{}, struct{}]{
		base: plugin,
		commonFields: commonFields[struct{}, struct{}]{
			logger: zaptest.NewLogger(t),
			config: &config.PipedPlugin{
				Name: "mockLivestatePlugin",
			},
			deployTargets: map[string]*DeployTarget[struct{}]{
				"target1": {
					Name: "target1",
					Labels: map[string]string{
						"key1": "value1",
					},
				},
			},
		},
	}
}

func TestLivestatePluginServer_GetLivestate(t *testing.T) {
	t.Parallel()

	validConfig := strings.TrimSpace(`
apiVersion: pipecd.dev/v1beta1
kind: Appilcation
spec: {}
`)

	tests := []struct {
		name           string
		request        *livestate.GetLivestateRequest
		result         *GetLivestateResponse
		err            error
		expectedStatus codes.Code
		expectErr      bool
	}{
		{
			name: "success",
			request: &livestate.GetLivestateRequest{
				PipedId:         "piped1",
				ApplicationId:   "app1",
				ApplicationName: "app1",
				DeployTargets:   []string{"target1"},
				DeploySource: &common.DeploymentSource{
					ApplicationDirectory:      "app-dir",
					CommitHash:                "commit-hash",
					ApplicationConfig:         []byte(validConfig),
					ApplicationConfigFilename: "app-config-filename",
				},
			},
			result: &GetLivestateResponse{
				LiveState: ApplicationLiveState{
					Resources: []ResourceState{
						{
							ID:   "resource1",
							Name: "Resource 1",
						},
					},
				},
				SyncState: ApplicationSyncState{
					Status: ApplicationSyncStateSynced,
				},
			},
			expectedStatus: codes.OK,
		},
		{
			name: "failure when deploy target not found",
			request: &livestate.GetLivestateRequest{
				PipedId:         "piped1",
				ApplicationId:   "app1",
				ApplicationName: "app1",
				DeployTargets:   []string{"target2"},
			},
			result:         &GetLivestateResponse{},
			expectErr:      true,
			expectedStatus: codes.Internal,
		},
		{
			name:           "error",
			request:        &livestate.GetLivestateRequest{},
			result:         &GetLivestateResponse{},
			err:            errors.New("some error"),
			expectErr:      true,
			expectedStatus: codes.Internal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			plugin := &mockLivestatePlugin{
				result: tt.result,
				err:    tt.err,
			}
			server := newTestLivestatePluginServer(t, plugin)

			response, err := server.GetLivestate(context.Background(), tt.request)
			if (err != nil) != tt.expectErr {
				t.Fatalf("unexpected error: %v", err)
			}

			if status.Code(err) != tt.expectedStatus {
				t.Errorf("expected status %v, got %v", tt.expectedStatus, status.Code(err))
			}

			if response != nil && response.GetApplicationLiveState().GetResources()[0].GetId() != tt.result.LiveState.Resources[0].ID {
				t.Errorf("expected resource ID %v, got %v", tt.result.LiveState.Resources[0].ID, response.GetApplicationLiveState().GetResources()[0].GetId())
			}
		})
	}
}

func TestApplicationLiveState_toModel(t *testing.T) {
	t.Parallel()

	now := time.Now()
	tests := []struct {
		name     string
		input    ApplicationLiveState
		expected *model.ApplicationLiveState
	}{
		{
			name: "convert ApplicationLiveState to model with healthy resource",
			input: ApplicationLiveState{
				Resources: []ResourceState{
					{
						ID:           "resource1",
						Name:         "Resource 1",
						HealthStatus: ResourceHealthStateHealthy,
						CreatedAt:    now,
					},
				},
			},
			expected: &model.ApplicationLiveState{
				Resources: []*model.ResourceState{
					{
						Id:           "resource1",
						Name:         "Resource 1",
						PluginName:   "test-plugin",
						HealthStatus: model.ResourceState_HEALTHY,
						CreatedAt:    now.Unix(),
						UpdatedAt:    now.Unix(),
					},
				},
				HealthStatus: model.ApplicationLiveState_HEALTHY,
			},
		},
		{
			name: "convert ApplicationLiveState to model with unhealthy resource",
			input: ApplicationLiveState{
				Resources: []ResourceState{
					{
						ID:           "resource1",
						Name:         "Resource 1",
						HealthStatus: ResourceHealthStateUnhealthy,
						CreatedAt:    now,
					},
				},
			},
			expected: &model.ApplicationLiveState{
				Resources: []*model.ResourceState{
					{
						Id:           "resource1",
						Name:         "Resource 1",
						PluginName:   "test-plugin",
						HealthStatus: model.ResourceState_UNHEALTHY,
						CreatedAt:    now.Unix(),
						UpdatedAt:    now.Unix(),
					},
				},
				HealthStatus: model.ApplicationLiveState_OTHER,
			},
		},
		{
			name: "convert ApplicationLiveState to model with unknown health status resource",
			input: ApplicationLiveState{
				Resources: []ResourceState{
					{
						ID:           "resource1",
						Name:         "Resource 1",
						HealthStatus: ResourceHealthStateUnknown,
						CreatedAt:    now,
					},
				},
			},
			expected: &model.ApplicationLiveState{
				Resources: []*model.ResourceState{
					{
						Id:           "resource1",
						Name:         "Resource 1",
						PluginName:   "test-plugin",
						HealthStatus: model.ResourceState_UNKNOWN,
						CreatedAt:    now.Unix(),
						UpdatedAt:    now.Unix(),
					},
				},
				HealthStatus: model.ApplicationLiveState_UNKNOWN,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := tt.input.toModel("test-plugin", now)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestResourceState_toModel(t *testing.T) {
	t.Parallel()

	now := time.Now()
	tests := []struct {
		name     string
		input    ResourceState
		expected *model.ResourceState
	}{
		{
			name: "convert ResourceState to model",
			input: ResourceState{
				ID:        "resource1",
				Name:      "Resource 1",
				CreatedAt: now,
			},
			expected: &model.ResourceState{
				Id:         "resource1",
				Name:       "Resource 1",
				PluginName: "test-plugin",
				CreatedAt:  now.Unix(),
				UpdatedAt:  now.Unix(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := tt.input.toModel("test-plugin", now)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestApplicationSyncState_toModel(t *testing.T) {
	t.Parallel()

	now := time.Now()
	tests := []struct {
		name     string
		input    ApplicationSyncState
		expected *model.ApplicationSyncState
	}{
		{
			name: "convert ApplicationSyncState to model",
			input: ApplicationSyncState{
				Status:      ApplicationSyncStateSynced,
				ShortReason: "All resources are synced",
				Reason:      "No differences found",
			},
			expected: &model.ApplicationSyncState{
				Status:      model.ApplicationSyncStatus_SYNCED,
				ShortReason: "All resources are synced",
				Reason:      "No differences found",
				Timestamp:   now.Unix(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := tt.input.toModel(now)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestApplicationHealthStatus_toModel(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ApplicationHealthStatus
		expected model.ApplicationLiveState_Status
	}{
		{
			name:     "convert ApplicationHealthStateHealthy to model",
			input:    ApplicationHealthStateHealthy,
			expected: model.ApplicationLiveState_HEALTHY,
		},
		{
			name:     "convert ApplicationHealthStateOther to model",
			input:    ApplicationHealthStateOther,
			expected: model.ApplicationLiveState_OTHER,
		},
		{
			name:     "convert ApplicationHealthStateUnknown to model",
			input:    ApplicationHealthStateUnknown,
			expected: model.ApplicationLiveState_UNKNOWN,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := tt.input.toModel()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestResourceHealthStatus_toModel(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ResourceHealthStatus
		expected model.ResourceState_HealthStatus
	}{
		{
			name:     "convert ResourceHealthStateHealthy to model",
			input:    ResourceHealthStateHealthy,
			expected: model.ResourceState_HEALTHY,
		},
		{
			name:     "convert ResourceHealthStateUnhealthy to model",
			input:    ResourceHealthStateUnhealthy,
			expected: model.ResourceState_UNHEALTHY,
		},
		{
			name:     "convert ResourceHealthStateUnknown to model",
			input:    ResourceHealthStateUnknown,
			expected: model.ResourceState_UNKNOWN,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := tt.input.toModel()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestApplicationSyncStatus_toModel(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ApplicationSyncStatus
		expected model.ApplicationSyncStatus
	}{
		{
			name:     "convert ApplicationSyncStateSynced to model",
			input:    ApplicationSyncStateSynced,
			expected: model.ApplicationSyncStatus_SYNCED,
		},
		{
			name:     "convert ApplicationSyncStateOutOfSync to model",
			input:    ApplicationSyncStateOutOfSync,
			expected: model.ApplicationSyncStatus_OUT_OF_SYNC,
		},
		{
			name:     "convert ApplicationSyncStateInvalidConfig to model",
			input:    ApplicationSyncStateInvalidConfig,
			expected: model.ApplicationSyncStatus_INVALID_CONFIG,
		},
		{
			name:     "convert ApplicationSyncStateUnknown to model",
			input:    ApplicationSyncStateUnknown,
			expected: model.ApplicationSyncStatus_UNKNOWN,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := tt.input.toModel()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestApplicationLiveState_healthStatus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		resources []ResourceState
		expected  ApplicationHealthStatus
	}{
		{
			name: "all healthy",
			resources: []ResourceState{
				{HealthStatus: ResourceHealthStateHealthy},
			},
			expected: ApplicationHealthStateHealthy,
		},
		{
			name: "one unhealthy, none unknown",
			resources: []ResourceState{
				{HealthStatus: ResourceHealthStateHealthy},
				{HealthStatus: ResourceHealthStateUnhealthy},
			},
			expected: ApplicationHealthStateOther,
		},
		{
			name: "one unknown, others healthy",
			resources: []ResourceState{
				{HealthStatus: ResourceHealthStateHealthy},
				{HealthStatus: ResourceHealthStateUnknown},
			},
			expected: ApplicationHealthStateUnknown,
		},
		{
			name: "one unknown, one unhealthy",
			resources: []ResourceState{
				{HealthStatus: ResourceHealthStateUnhealthy},
				{HealthStatus: ResourceHealthStateUnknown},
			},
			expected: ApplicationHealthStateUnknown,
		},
		{
			name:      "empty resources",
			resources: []ResourceState{},
			expected:  ApplicationHealthStateHealthy,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ApplicationLiveState{Resources: tt.resources}
			result := s.healthStatus()
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// Copyright 2024 The PipeCD Authors.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.12
// source: pkg/model/planpreview.proto

package model

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PlanPreviewCommandResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandId string `protobuf:"bytes,1,opt,name=command_id,json=commandId,proto3" json:"command_id,omitempty"`
	// The Piped that handles command.
	PipedId string `protobuf:"bytes,2,opt,name=piped_id,json=pipedId,proto3" json:"piped_id,omitempty"`
	// Web URL to the piped page.
	// This is only filled before returning to the client.
	PipedUrl string                          `protobuf:"bytes,3,opt,name=piped_url,json=pipedUrl,proto3" json:"piped_url,omitempty"`
	Results  []*ApplicationPlanPreviewResult `protobuf:"bytes,4,rep,name=results,proto3" json:"results,omitempty"`
	// Error while handling command.
	Error     string `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
	PipedName string `protobuf:"bytes,6,opt,name=piped_name,json=pipedName,proto3" json:"piped_name,omitempty"`
}

func (x *PlanPreviewCommandResult) Reset() {
	*x = PlanPreviewCommandResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_planpreview_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlanPreviewCommandResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanPreviewCommandResult) ProtoMessage() {}

func (x *PlanPreviewCommandResult) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_planpreview_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanPreviewCommandResult.ProtoReflect.Descriptor instead.
func (*PlanPreviewCommandResult) Descriptor() ([]byte, []int) {
	return file_pkg_model_planpreview_proto_rawDescGZIP(), []int{0}
}

func (x *PlanPreviewCommandResult) GetCommandId() string {
	if x != nil {
		return x.CommandId
	}
	return ""
}

func (x *PlanPreviewCommandResult) GetPipedId() string {
	if x != nil {
		return x.PipedId
	}
	return ""
}

func (x *PlanPreviewCommandResult) GetPipedUrl() string {
	if x != nil {
		return x.PipedUrl
	}
	return ""
}

func (x *PlanPreviewCommandResult) GetResults() []*ApplicationPlanPreviewResult {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *PlanPreviewCommandResult) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *PlanPreviewCommandResult) GetPipedName() string {
	if x != nil {
		return x.PipedName
	}
	return ""
}

type ApplicationPlanPreviewResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Application information.
	ApplicationId   string `protobuf:"bytes,1,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	ApplicationName string `protobuf:"bytes,2,opt,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`
	// Web URL to the application page.
	// This is only filled before returning to the client.
	ApplicationUrl string `protobuf:"bytes,3,opt,name=application_url,json=applicationUrl,proto3" json:"application_url,omitempty"`
	// Deprecated: Do not use in pipedv1
	//
	// Deprecated: Do not use.
	ApplicationKind      ApplicationKind   `protobuf:"varint,4,opt,name=application_kind,json=applicationKind,proto3,enum=model.ApplicationKind" json:"application_kind,omitempty"`
	ApplicationDirectory string            `protobuf:"bytes,5,opt,name=application_directory,json=applicationDirectory,proto3" json:"application_directory,omitempty"`
	PipedId              string            `protobuf:"bytes,9,opt,name=piped_id,json=pipedId,proto3" json:"piped_id,omitempty"`
	ProjectId            string            `protobuf:"bytes,10,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Labels               map[string]string `protobuf:"bytes,11,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Target commit information.
	HeadBranch string `protobuf:"bytes,20,opt,name=head_branch,json=headBranch,proto3" json:"head_branch,omitempty"`
	HeadCommit string `protobuf:"bytes,21,opt,name=head_commit,json=headCommit,proto3" json:"head_commit,omitempty"`
	// Planpreview result.
	SyncStrategy SyncStrategy `protobuf:"varint,30,opt,name=sync_strategy,json=syncStrategy,proto3,enum=model.SyncStrategy" json:"sync_strategy,omitempty"`
	// Deprecated: Use plugin_plan_results instead in pipedv1.
	//
	// Deprecated: Do not use.
	PlanSummary []byte `protobuf:"bytes,31,opt,name=plan_summary,json=planSummary,proto3" json:"plan_summary,omitempty"`
	// Deprecated: Use plugin_plan_results instead in pipedv1.
	//
	// Deprecated: Do not use.
	PlanDetails []byte `protobuf:"bytes,32,opt,name=plan_details,json=planDetails,proto3" json:"plan_details,omitempty"`
	// Mark if no change were detected.
	NoChange          bool                       `protobuf:"varint,33,opt,name=no_change,json=noChange,proto3" json:"no_change,omitempty"`
	PluginPlanResults []*PluginPlanPreviewResult `protobuf:"bytes,34,rep,name=plugin_plan_results,json=pluginPlanResults,proto3" json:"plugin_plan_results,omitempty"`
	// Note: In pipedv1, this will not be empty. It will be "<unknown>" if plugins are not successfully loaded.
	DeploymentPluginNames []string `protobuf:"bytes,35,rep,name=deployment_plugin_names,json=deploymentPluginNames,proto3" json:"deployment_plugin_names,omitempty"`
	// Error while building planpreview result.
	Error     string `protobuf:"bytes,40,opt,name=error,proto3" json:"error,omitempty"`
	CreatedAt int64  `protobuf:"varint,90,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *ApplicationPlanPreviewResult) Reset() {
	*x = ApplicationPlanPreviewResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_planpreview_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationPlanPreviewResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationPlanPreviewResult) ProtoMessage() {}

func (x *ApplicationPlanPreviewResult) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_planpreview_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationPlanPreviewResult.ProtoReflect.Descriptor instead.
func (*ApplicationPlanPreviewResult) Descriptor() ([]byte, []int) {
	return file_pkg_model_planpreview_proto_rawDescGZIP(), []int{1}
}

func (x *ApplicationPlanPreviewResult) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *ApplicationPlanPreviewResult) GetApplicationName() string {
	if x != nil {
		return x.ApplicationName
	}
	return ""
}

func (x *ApplicationPlanPreviewResult) GetApplicationUrl() string {
	if x != nil {
		return x.ApplicationUrl
	}
	return ""
}

// Deprecated: Do not use.
func (x *ApplicationPlanPreviewResult) GetApplicationKind() ApplicationKind {
	if x != nil {
		return x.ApplicationKind
	}
	return ApplicationKind_KUBERNETES
}

func (x *ApplicationPlanPreviewResult) GetApplicationDirectory() string {
	if x != nil {
		return x.ApplicationDirectory
	}
	return ""
}

func (x *ApplicationPlanPreviewResult) GetPipedId() string {
	if x != nil {
		return x.PipedId
	}
	return ""
}

func (x *ApplicationPlanPreviewResult) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *ApplicationPlanPreviewResult) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *ApplicationPlanPreviewResult) GetHeadBranch() string {
	if x != nil {
		return x.HeadBranch
	}
	return ""
}

func (x *ApplicationPlanPreviewResult) GetHeadCommit() string {
	if x != nil {
		return x.HeadCommit
	}
	return ""
}

func (x *ApplicationPlanPreviewResult) GetSyncStrategy() SyncStrategy {
	if x != nil {
		return x.SyncStrategy
	}
	return SyncStrategy_AUTO
}

// Deprecated: Do not use.
func (x *ApplicationPlanPreviewResult) GetPlanSummary() []byte {
	if x != nil {
		return x.PlanSummary
	}
	return nil
}

// Deprecated: Do not use.
func (x *ApplicationPlanPreviewResult) GetPlanDetails() []byte {
	if x != nil {
		return x.PlanDetails
	}
	return nil
}

func (x *ApplicationPlanPreviewResult) GetNoChange() bool {
	if x != nil {
		return x.NoChange
	}
	return false
}

func (x *ApplicationPlanPreviewResult) GetPluginPlanResults() []*PluginPlanPreviewResult {
	if x != nil {
		return x.PluginPlanResults
	}
	return nil
}

func (x *ApplicationPlanPreviewResult) GetDeploymentPluginNames() []string {
	if x != nil {
		return x.DeploymentPluginNames
	}
	return nil
}

func (x *ApplicationPlanPreviewResult) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ApplicationPlanPreviewResult) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type PluginPlanPreviewResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PluginName string `protobuf:"bytes,1,opt,name=plugin_name,json=pluginName,proto3" json:"plugin_name,omitempty"`
	// Name of the deploy target.
	DeployTarget string `protobuf:"bytes,2,opt,name=deploy_target,json=deployTarget,proto3" json:"deploy_target,omitempty"`
	PlanSummary  []byte `protobuf:"bytes,3,opt,name=plan_summary,json=planSummary,proto3" json:"plan_summary,omitempty"`
	PlanDetails  []byte `protobuf:"bytes,4,opt,name=plan_details,json=planDetails,proto3" json:"plan_details,omitempty"`
	// The language to render the details like "diff","hcl".
	// If this is empty, "diff" will be used by default.
	DiffLanguage string `protobuf:"bytes,5,opt,name=diff_language,json=diffLanguage,proto3" json:"diff_language,omitempty"`
}

func (x *PluginPlanPreviewResult) Reset() {
	*x = PluginPlanPreviewResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_planpreview_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginPlanPreviewResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginPlanPreviewResult) ProtoMessage() {}

func (x *PluginPlanPreviewResult) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_planpreview_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginPlanPreviewResult.ProtoReflect.Descriptor instead.
func (*PluginPlanPreviewResult) Descriptor() ([]byte, []int) {
	return file_pkg_model_planpreview_proto_rawDescGZIP(), []int{2}
}

func (x *PluginPlanPreviewResult) GetPluginName() string {
	if x != nil {
		return x.PluginName
	}
	return ""
}

func (x *PluginPlanPreviewResult) GetDeployTarget() string {
	if x != nil {
		return x.DeployTarget
	}
	return ""
}

func (x *PluginPlanPreviewResult) GetPlanSummary() []byte {
	if x != nil {
		return x.PlanSummary
	}
	return nil
}

func (x *PluginPlanPreviewResult) GetPlanDetails() []byte {
	if x != nil {
		return x.PlanDetails
	}
	return nil
}

func (x *PluginPlanPreviewResult) GetDiffLanguage() string {
	if x != nil {
		return x.DiffLanguage
	}
	return ""
}

var File_pkg_model_planpreview_proto protoreflect.FileDescriptor

var file_pkg_model_planpreview_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70,
	0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x01, 0x0a, 0x18, 0x50, 0x6c, 0x61, 0x6e, 0x50, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x26, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x08, 0x70, 0x69,
	0x70, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x70, 0x69, 0x70, 0x65, 0x64, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x69, 0x70, 0x65, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x69, 0x70, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x3d, 0x0a, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x6c, 0x61, 0x6e, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x69, 0x70, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x69, 0x70, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0xcd, 0x07, 0x0a, 0x1c, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x6c, 0x61, 0x6e, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x2e, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x32, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x4d, 0x0a,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x69, 0x6e,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x42,
	0x0a, 0x18, 0x01, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0f, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x3c, 0x0a, 0x15,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x14, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x22, 0x0a, 0x08, 0x70, 0x69,
	0x70, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x70, 0x69, 0x70, 0x65, 0x64, 0x49, 0x64, 0x12, 0x26,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x47, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x50, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12,
	0x28, 0x0a, 0x0b, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x68,
	0x65, 0x61, 0x64, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x28, 0x0a, 0x0b, 0x68, 0x65, 0x61,
	0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x68, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x12, 0x38, 0x0a, 0x0d, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52,
	0x0c, 0x73, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x25, 0x0a,
	0x0c, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x1f, 0x20,
	0x01, 0x28, 0x0c, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x6e, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x12, 0x25, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x18, 0x20, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0b,
	0x70, 0x6c, 0x61, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6e,
	0x6f, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x21, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x6e, 0x6f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x4e, 0x0a, 0x13, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18,
	0x22, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x11, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x6c, 0x61,
	0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x17, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x18, 0x23, 0x20, 0x03, 0x28, 0x09, 0x52, 0x15, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x26, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22,
	0x02, 0x20, 0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x39,
	0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x09, 0x22,
	0xd3, 0x01, 0x0a, 0x17, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x50, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x28, 0x0a, 0x0b, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6c,
	0x61, 0x6e, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0b, 0x70, 0x6c, 0x61, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x66, 0x66, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x69, 0x66, 0x66, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x2d, 0x63, 0x64, 0x2f, 0x70, 0x69, 0x70, 0x65,
	0x63, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_model_planpreview_proto_rawDescOnce sync.Once
	file_pkg_model_planpreview_proto_rawDescData = file_pkg_model_planpreview_proto_rawDesc
)

func file_pkg_model_planpreview_proto_rawDescGZIP() []byte {
	file_pkg_model_planpreview_proto_rawDescOnce.Do(func() {
		file_pkg_model_planpreview_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_model_planpreview_proto_rawDescData)
	})
	return file_pkg_model_planpreview_proto_rawDescData
}

var file_pkg_model_planpreview_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_model_planpreview_proto_goTypes = []interface{}{
	(*PlanPreviewCommandResult)(nil),     // 0: model.PlanPreviewCommandResult
	(*ApplicationPlanPreviewResult)(nil), // 1: model.ApplicationPlanPreviewResult
	(*PluginPlanPreviewResult)(nil),      // 2: model.PluginPlanPreviewResult
	nil,                                  // 3: model.ApplicationPlanPreviewResult.LabelsEntry
	(ApplicationKind)(0),                 // 4: model.ApplicationKind
	(SyncStrategy)(0),                    // 5: model.SyncStrategy
}
var file_pkg_model_planpreview_proto_depIdxs = []int32{
	1, // 0: model.PlanPreviewCommandResult.results:type_name -> model.ApplicationPlanPreviewResult
	4, // 1: model.ApplicationPlanPreviewResult.application_kind:type_name -> model.ApplicationKind
	3, // 2: model.ApplicationPlanPreviewResult.labels:type_name -> model.ApplicationPlanPreviewResult.LabelsEntry
	5, // 3: model.ApplicationPlanPreviewResult.sync_strategy:type_name -> model.SyncStrategy
	2, // 4: model.ApplicationPlanPreviewResult.plugin_plan_results:type_name -> model.PluginPlanPreviewResult
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_model_planpreview_proto_init() }
func file_pkg_model_planpreview_proto_init() {
	if File_pkg_model_planpreview_proto != nil {
		return
	}
	file_pkg_model_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_model_planpreview_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlanPreviewCommandResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_model_planpreview_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationPlanPreviewResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_model_planpreview_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginPlanPreviewResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_model_planpreview_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_model_planpreview_proto_goTypes,
		DependencyIndexes: file_pkg_model_planpreview_proto_depIdxs,
		MessageInfos:      file_pkg_model_planpreview_proto_msgTypes,
	}.Build()
	File_pkg_model_planpreview_proto = out.File
	file_pkg_model_planpreview_proto_rawDesc = nil
	file_pkg_model_planpreview_proto_goTypes = nil
	file_pkg_model_planpreview_proto_depIdxs = nil
}

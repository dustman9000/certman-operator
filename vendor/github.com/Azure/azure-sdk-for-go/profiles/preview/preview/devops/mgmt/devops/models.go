// +build go1.9

// Copyright 2020 Microsoft Corporation
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

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package devops

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/devops/mgmt/2019-07-01-preview/devops"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CodeRepositoryType = original.CodeRepositoryType

const (
	GitHub  CodeRepositoryType = original.GitHub
	VstsGit CodeRepositoryType = original.VstsGit
)

type InputDataType = original.InputDataType

const (
	InputDataTypeAuthorization InputDataType = original.InputDataTypeAuthorization
	InputDataTypeBool          InputDataType = original.InputDataTypeBool
	InputDataTypeInt           InputDataType = original.InputDataTypeInt
	InputDataTypeSecureString  InputDataType = original.InputDataTypeSecureString
	InputDataTypeString        InputDataType = original.InputDataTypeString
)

type Authorization = original.Authorization
type BaseClient = original.BaseClient
type BootstrapConfiguration = original.BootstrapConfiguration
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type CodeRepository = original.CodeRepository
type InputDescriptor = original.InputDescriptor
type InputValue = original.InputValue
type Operation = original.Operation
type OperationDisplayValue = original.OperationDisplayValue
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type OrganizationReference = original.OrganizationReference
type Pipeline = original.Pipeline
type PipelineListResult = original.PipelineListResult
type PipelineListResultIterator = original.PipelineListResultIterator
type PipelineListResultPage = original.PipelineListResultPage
type PipelineProperties = original.PipelineProperties
type PipelineTemplate = original.PipelineTemplate
type PipelineTemplateDefinition = original.PipelineTemplateDefinition
type PipelineTemplateDefinitionListResult = original.PipelineTemplateDefinitionListResult
type PipelineTemplateDefinitionListResultIterator = original.PipelineTemplateDefinitionListResultIterator
type PipelineTemplateDefinitionListResultPage = original.PipelineTemplateDefinitionListResultPage
type PipelineTemplateDefinitionsClient = original.PipelineTemplateDefinitionsClient
type PipelineUpdateParameters = original.PipelineUpdateParameters
type PipelinesClient = original.PipelinesClient
type PipelinesCreateOrUpdateFuture = original.PipelinesCreateOrUpdateFuture
type ProjectReference = original.ProjectReference
type Resource = original.Resource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPipelineListResultIterator(page PipelineListResultPage) PipelineListResultIterator {
	return original.NewPipelineListResultIterator(page)
}
func NewPipelineListResultPage(getNextPage func(context.Context, PipelineListResult) (PipelineListResult, error)) PipelineListResultPage {
	return original.NewPipelineListResultPage(getNextPage)
}
func NewPipelineTemplateDefinitionListResultIterator(page PipelineTemplateDefinitionListResultPage) PipelineTemplateDefinitionListResultIterator {
	return original.NewPipelineTemplateDefinitionListResultIterator(page)
}
func NewPipelineTemplateDefinitionListResultPage(getNextPage func(context.Context, PipelineTemplateDefinitionListResult) (PipelineTemplateDefinitionListResult, error)) PipelineTemplateDefinitionListResultPage {
	return original.NewPipelineTemplateDefinitionListResultPage(getNextPage)
}
func NewPipelineTemplateDefinitionsClient(subscriptionID string) PipelineTemplateDefinitionsClient {
	return original.NewPipelineTemplateDefinitionsClient(subscriptionID)
}
func NewPipelineTemplateDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) PipelineTemplateDefinitionsClient {
	return original.NewPipelineTemplateDefinitionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPipelinesClient(subscriptionID string) PipelinesClient {
	return original.NewPipelinesClient(subscriptionID)
}
func NewPipelinesClientWithBaseURI(baseURI string, subscriptionID string) PipelinesClient {
	return original.NewPipelinesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCodeRepositoryTypeValues() []CodeRepositoryType {
	return original.PossibleCodeRepositoryTypeValues()
}
func PossibleInputDataTypeValues() []InputDataType {
	return original.PossibleInputDataTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
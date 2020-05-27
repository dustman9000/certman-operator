package consumptionapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/consumption/mgmt/2018-05-31/consumption"
)

// PriceSheetClientAPI contains the set of methods on the PriceSheetClient type.
type PriceSheetClientAPI interface {
	Get(ctx context.Context, expand string, skiptoken string, top *int32) (result consumption.PriceSheetResult, err error)
	GetByBillingPeriod(ctx context.Context, billingPeriodName string, expand string, skiptoken string, top *int32) (result consumption.PriceSheetResult, err error)
}

var _ PriceSheetClientAPI = (*consumption.PriceSheetClient)(nil)

// UsageDetailsClientAPI contains the set of methods on the UsageDetailsClient type.
type UsageDetailsClientAPI interface {
	List(ctx context.Context, expand string, filter string, skiptoken string, top *int32, apply string) (result consumption.UsageDetailsListResultPage, err error)
	ListComplete(ctx context.Context, expand string, filter string, skiptoken string, top *int32, apply string) (result consumption.UsageDetailsListResultIterator, err error)
	ListByBillingAccount(ctx context.Context, billingAccountID string, expand string, filter string, skiptoken string, top *int32, apply string) (result consumption.UsageDetailsListResultPage, err error)
	ListByBillingAccountComplete(ctx context.Context, billingAccountID string, expand string, filter string, skiptoken string, top *int32, apply string) (result consumption.UsageDetailsListResultIterator, err error)
	ListByBillingPeriod(ctx context.Context, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (result consumption.UsageDetailsListResultPage, err error)
	ListByBillingPeriodComplete(ctx context.Context, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (result consumption.UsageDetailsListResultIterator, err error)
	ListByDepartment(ctx context.Context, departmentID string, expand string, filter string, skiptoken string, top *int32, apply string) (result consumption.UsageDetailsListResultPage, err error)
	ListByDepartmentComplete(ctx context.Context, departmentID string, expand string, filter string, skiptoken string, top *int32, apply string) (result consumption.UsageDetailsListResultIterator, err error)
	ListByEnrollmentAccount(ctx context.Context, enrollmentAccountID string, expand string, filter string, skiptoken string, top *int32, apply string) (result consumption.UsageDetailsListResultPage, err error)
	ListByEnrollmentAccountComplete(ctx context.Context, enrollmentAccountID string, expand string, filter string, skiptoken string, top *int32, apply string) (result consumption.UsageDetailsListResultIterator, err error)
	ListForBillingPeriodByBillingAccount(ctx context.Context, billingAccountID string, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (result consumption.UsageDetailsListResultPage, err error)
	ListForBillingPeriodByBillingAccountComplete(ctx context.Context, billingAccountID string, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (result consumption.UsageDetailsListResultIterator, err error)
	ListForBillingPeriodByDepartment(ctx context.Context, departmentID string, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (result consumption.UsageDetailsListResultPage, err error)
	ListForBillingPeriodByDepartmentComplete(ctx context.Context, departmentID string, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (result consumption.UsageDetailsListResultIterator, err error)
	ListForBillingPeriodByEnrollmentAccount(ctx context.Context, enrollmentAccountID string, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (result consumption.UsageDetailsListResultPage, err error)
	ListForBillingPeriodByEnrollmentAccountComplete(ctx context.Context, enrollmentAccountID string, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (result consumption.UsageDetailsListResultIterator, err error)
}

var _ UsageDetailsClientAPI = (*consumption.UsageDetailsClient)(nil)

// ForecastsClientAPI contains the set of methods on the ForecastsClient type.
type ForecastsClientAPI interface {
	List(ctx context.Context, filter string) (result consumption.ForecastsListResult, err error)
}

var _ ForecastsClientAPI = (*consumption.ForecastsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result consumption.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result consumption.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*consumption.OperationsClient)(nil)
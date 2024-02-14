// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package job

import (
	"context"

	"github.com/gomodule/redigo/redis"
)

const (
	// Check in data placeholder for saving data space
	redundantCheckInData = "[REDUNDANT]"
)

// Tracker is designed to track the life cycle of the job described by the stats
// The status change is linear and then has strict preorder and successor
// Check should be enforced before switching
//
// Pending is default status when creating job, so no need to switch
type Tracker interface {
	// Save the job stats which tracked by this tracker to the backend
	//
	// Return:
	//   none nil error returned if any issues happened
	Save() error

	// Load the job stats which tracked by this tracker with the backend data
	//
	// Return:
	//   none nil error returned if any issues happened
	Load() error

	// Get the job stats which tracked by this tracker
	//
	// Returns:
	//  *models.Info : job stats data
	Job() *Stats

	// Update the properties of the job stats
	//
	// fieldAndValues ...interface{} : One or more properties being updated
	//
	// Returns:
	//  error if update failed
	Update(fieldAndValues ...interface{}) error

	// NumericID returns the numeric ID of periodic job.
	// Please pay attention, this only for periodic job.
	NumericID() (int64, error)

	// Mark the periodic job execution to done by update the score
	// of the relation between its periodic policy and execution to -1.
	PeriodicExecutionDone() error

	// Check in message
	CheckIn(message string) error

	// Update status with retry enabled
	UpdateStatusWithRetry(targetStatus Status) error

	// The current status of job
	Status() (Status, error)

	// Switch status to running
	Run() error

	// Switch status to stopped
	Stop() error

	// Switch the status to error
	Fail() error

	// Switch the status to success
	Succeed() error

	// Reset the status to `pending`
	Reset() error

	// Fire status hook to report the current status
	FireHook() error
}

type basicTracker struct {
	namespace string
	context   context.Context
	pool      *redis.Pool
	jobID     string
	jobStats  *Stats
	callback  HookCallback
	retryList *list
}

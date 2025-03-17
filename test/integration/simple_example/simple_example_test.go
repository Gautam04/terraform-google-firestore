// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package firestore_resource

import (
	"fmt"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/gcloud"
	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/tft"
	"github.com/stretchr/testify/assert"
)

func TestSimpleExample(t *testing.T) {
	example := tft.NewTFBlueprintTest(t)

	example.DefineVerify(func(assert *assert.Assertions) {
		example.DefaultVerify(assert)

		projectId := example.GetTFSetupStringOutput("project_id")
		databaseId := example.GetStringOutput("database_id")
		databaseIdPrefix := fmt.Sprintf("projects/%s/databases/", projectId)
		databaseName := strings.TrimPrefix(databaseId, databaseIdPrefix)
		databaseInfo := gcloud.Run(
			t,
			"firestore databases describe",
			gcloud.WithCommonArgs([]string{"--project", projectId, "--database", databaseName, "--format", "json"}),
		)

		assert.Equal(
			databaseId,
			databaseInfo.Get("name").String(),
			"Database ID does not match.",
		)

		assert.Equal(
			"FIRESTORE_NATIVE",
			databaseInfo.Get("type").String(),
			"Database type does not match.",
		)

		assert.Equal(
			"us-central1",
			databaseInfo.Get("locationId").String(),
			"Location ID does not match.",
		)

		assert.Equal(
			"POINT_IN_TIME_RECOVERY_DISABLED",
			databaseInfo.Get("pointInTimeRecoveryEnablement").String(),
			"Location ID does not match.",
		)

	})
	example.Test()
}

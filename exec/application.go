/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package exec

import (
	"fmt"
	"path"

	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/chaosblade-io/chaosblade-spec-go/util"

	"github.com/chaosblade-io/chaosblade-exec-cri/version"
)

var JvmSpecFileForYaml = ""

// getJvmModels returns java experiment specs
func getJvmModels() []spec.ExpModelCommandSpec {
	var jvmSpecFile = path.Join(util.GetYamlHome(), fmt.Sprintf("chaosblade-jvm-spec-%s.yaml", version.BladeVersion))
	if JvmSpecFileForYaml != "" {
		jvmSpecFile = JvmSpecFileForYaml
	}
	modelCommandSpecs := make([]spec.ExpModelCommandSpec, 0)
	models, err := util.ParseSpecsToModel(jvmSpecFile, nil)
	if err != nil {
		return modelCommandSpecs
	}
	for idx := range models.Models {
		modelCommandSpecs = append(modelCommandSpecs, &models.Models[idx])
	}
	return modelCommandSpecs
}

/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package parsers

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v2"

	"path/filepath"

	"github.com/apache/incubator-openwhisk-client-go/whisk"
	"github.com/apache/incubator-openwhisk-wskdeploy/utils"
	"github.com/apache/incubator-openwhisk-wskdeploy/wskderrors"
	"github.com/apache/incubator-openwhisk-wskdeploy/wskenv"
	"github.com/apache/incubator-openwhisk-wskdeploy/wski18n"
	"github.com/apache/incubator-openwhisk-wskdeploy/wskprint"
)

const (
	API                 = "API"
	HTTPS               = "https://"
	HTTP                = "http://"
	API_VERSION         = "v1"
	WEB                 = "web"
	PATH_SEPARATOR      = "/"
	DEFAULT_PACKAGE     = "default"
	NATIVE_DOCKER_IMAGE = "openwhisk/dockerskeleton"
)

// Read existing manifest file or create new if none exists
func ReadOrCreateManifest() (*YAML, error) {
	maniyaml := YAML{}

	if _, err := os.Stat(utils.ManifestFileNameYaml); err == nil {
		dat, _ := ioutil.ReadFile(utils.ManifestFileNameYaml)
		err := NewYAMLParser().Unmarshal(dat, &maniyaml)
		if err != nil {
			return &maniyaml, wskderrors.NewFileReadError(utils.ManifestFileNameYaml, err.Error())
		}
	}
	return &maniyaml, nil
}

// Serialize manifest to local file
func Write(manifest *YAML, filename string) error {
	output, err := NewYAMLParser().marshal(manifest)
	if err != nil {
		return wskderrors.NewYAMLFileFormatError(filename, err.Error())
	}

	f, err := os.Create(filename)
	if err != nil {
		return wskderrors.NewFileReadError(filename, err.Error())
	}
	defer f.Close()

	f.Write(output)
	return nil
}

func (dm *YAMLParser) Unmarshal(input []byte, manifest *YAML) error {
	err := yaml.UnmarshalStrict(input, manifest)
	if err != nil {
		return err
	}
	return nil
}

func (dm *YAMLParser) marshal(manifest *YAML) (output []byte, err error) {
	data, err := yaml.Marshal(manifest)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (dm *YAMLParser) ParseManifest(manifestPath string) (*YAML, error) {
	mm := NewYAMLParser()
	maniyaml := YAML{}

	content, err := utils.Read(manifestPath)
	if err != nil {
		return &maniyaml, wskderrors.NewFileReadError(manifestPath, err.Error())
	}

	err = mm.Unmarshal(content, &maniyaml)
	if err != nil {
		return &maniyaml, wskderrors.NewYAMLParserErr(manifestPath, err)
	}
	maniyaml.Filepath = manifestPath
	manifest := ReadEnvVariable(&maniyaml)

	return manifest, nil
}

func (dm *YAMLParser) composeInputsOrOutputs(inputs map[string]Parameter, manifestFilePath string) (whisk.KeyValueArr, error) {
	var errorParser error
	keyValArr := make(whisk.KeyValueArr, 0)
	for name, param := range inputs {
		var keyVal whisk.KeyValue
		keyVal.Key = name
		keyVal.Value, errorParser = ResolveParameter(name, &param, manifestFilePath)
		if errorParser != nil {
			return nil, errorParser
		}
		if keyVal.Value != nil {
			keyValArr = append(keyValArr, keyVal)
		}
	}
	return keyValArr, nil
}

func (dm *YAMLParser) composeAnnotations(annotations map[string]interface{}) whisk.KeyValueArr {
	listOfAnnotations := make(whisk.KeyValueArr, 0)
	for name, value := range annotations {
		var keyVal whisk.KeyValue
		keyVal.Key = name
		keyVal.Value = wskenv.InterpolateStringWithEnvVar(value)
		listOfAnnotations = append(listOfAnnotations, keyVal)
	}
	return listOfAnnotations
}

func (dm *YAMLParser) ComposeDependenciesFromAllPackages(manifest *YAML, projectPath string, filePath string, managedAnnotations whisk.KeyValue) (map[string]utils.DependencyRecord, error) {
	dependencies := make(map[string]utils.DependencyRecord)
	packages := make(map[string]Package)

	if len(manifest.Packages) != 0 {
		packages = manifest.Packages
	} else {
		packages = manifest.GetProject().Packages
	}

	for n, p := range packages {
		d, err := dm.ComposeDependencies(p, projectPath, filePath, n, managedAnnotations)
		if err == nil {
			for k, v := range d {
				dependencies[k] = v
			}
		} else {
			return nil, err
		}
	}
	return dependencies, nil
}

func (dm *YAMLParser) ComposeDependencies(pkg Package, projectPath string, filePath string, packageName string, managedAnnotations whisk.KeyValue) (map[string]utils.DependencyRecord, error) {

	depMap := make(map[string]utils.DependencyRecord)
	for key, dependency := range pkg.Dependencies {
		version := dependency.Version
		if len(version) == 0 {
			version = YAML_VALUE_BRANCH_MASTER
		}

		location := dependency.Location

		isBinding := false
		if utils.LocationIsBinding(location) {
			if !strings.HasPrefix(location, PATH_SEPARATOR) {
				location = PATH_SEPARATOR + dependency.Location
			}
			isBinding = true
		} else if utils.LocationIsGithub(location) {

			// TODO() define const for the protocol prefix, etc.
			if !strings.HasPrefix(location, HTTPS) && !strings.HasPrefix(location, HTTP) {
				location = HTTPS + dependency.Location
				location = wskenv.InterpolateStringWithEnvVar(location).(string)
			}
			isBinding = false
		} else {
			// TODO() create new named error in wskerrors package
			return nil, errors.New(wski18n.T(wski18n.ID_ERR_DEPENDENCY_UNKNOWN_TYPE))
		}

		inputs, err := dm.composeInputsOrOutputs(dependency.Inputs, filePath)
		if err != nil {
			return nil, err
		}

		annotations := dm.composeAnnotations(dependency.Annotations)

		if utils.Flags.Managed || utils.Flags.Sync {
			annotations = append(annotations, managedAnnotations)
		}

		packDir := path.Join(projectPath, strings.Title(YAML_KEY_PACKAGES))
		depName := packageName + ":" + key
		depMap[depName] = utils.NewDependencyRecord(packDir, packageName, location, version, inputs, annotations, isBinding)
	}

	return depMap, nil
}

func (dm *YAMLParser) ComposeAllPackages(manifest *YAML, filePath string, managedAnnotations whisk.KeyValue) (map[string]*whisk.Package, error) {
	packages := map[string]*whisk.Package{}
	manifestPackages := make(map[string]Package)

	if len(manifest.Packages) != 0 {
		manifestPackages = manifest.Packages
	} else {
		manifestPackages = manifest.GetProject().Packages
	}

	if len(manifestPackages) == 0 {
		warningString := wski18n.T(
			wski18n.ID_WARN_PACKAGES_NOT_FOUND_X_path_X,
			map[string]interface{}{
				wski18n.KEY_PATH: manifest.Filepath})
		wskprint.PrintOpenWhiskWarning(warningString)
	}

	// Compose each package found in manifest
	for n, p := range manifestPackages {
		s, err := dm.ComposePackage(p, n, filePath, managedAnnotations)

		if err == nil {
			packages[n] = s
		} else {
			return nil, err
		}
	}

	return packages, nil
}

func (dm *YAMLParser) ComposePackage(pkg Package, packageName string, filePath string, managedAnnotations whisk.KeyValue) (*whisk.Package, error) {
	pag := &whisk.Package{}
	pag.Name = packageName
	//The namespace for this package is absent, so we use default guest here.
	pag.Namespace = pkg.Namespace
	pub := false
	pag.Publish = &pub

	//Version is a mandatory value
	//If it is an empty string, it will be set to default value
	//And print an warning message
	// TODO(#673) implement STRICT flag
	if pkg.Version == "" {
		warningString := wski18n.T(
			wski18n.ID_WARN_KEY_MISSING_X_key_X_value_X,
			map[string]interface{}{
				wski18n.KEY_KEY:   wski18n.PACKAGE_VERSION,
				wski18n.KEY_VALUE: DEFAULT_PACKAGE_VERSION})
		wskprint.PrintOpenWhiskWarning(warningString)

		warningString = wski18n.T(
			wski18n.ID_WARN_KEYVALUE_NOT_SAVED_X_key_X,
			map[string]interface{}{wski18n.KEY_KEY: wski18n.PACKAGE_VERSION})

		wskprint.PrintOpenWhiskWarning(warningString)
		pkg.Version = DEFAULT_PACKAGE_VERSION
	}

	//License is a mandatory value
	//set license to unknown if it is an empty string
	//And print an warning message
	// TODO(#673) implement STRICT flag
	if pkg.License == "" {
		warningString := wski18n.T(
			wski18n.ID_WARN_KEY_MISSING_X_key_X_value_X,
			map[string]interface{}{
				wski18n.KEY_KEY:   wski18n.PACKAGE_LICENSE,
				wski18n.KEY_VALUE: DEFAULT_PACKAGE_LICENSE})
		wskprint.PrintOpenWhiskWarning(warningString)

		warningString = wski18n.T(
			wski18n.ID_WARN_KEYVALUE_NOT_SAVED_X_key_X,
			map[string]interface{}{wski18n.KEY_KEY: wski18n.PACKAGE_VERSION})

		wskprint.PrintOpenWhiskWarning(warningString)

		pkg.License = DEFAULT_PACKAGE_LICENSE
	} else {
		utils.CheckLicense(pkg.License)
	}

	//set parameters
	inputs, err := dm.composeInputsOrOutputs(pkg.Inputs, filePath)
	if err != nil {
		return nil, err
	}
	if len(inputs) > 0 {
		pag.Parameters = inputs
	}

	// set Package Annotations
	listOfAnnotations := dm.composeAnnotations(pkg.Annotations)
	if len(listOfAnnotations) > 0 {
		pag.Annotations = append(pag.Annotations, listOfAnnotations...)
	}

	// add Managed Annotations if this is Managed Deployment
	if utils.Flags.Managed || utils.Flags.Sync {
		pag.Annotations = append(pag.Annotations, managedAnnotations)
	}

	// "default" package is a reserved package name
	// and in this case wskdeploy deploys openwhisk entities under
	// /namespace instead of /namespace/package
	if strings.ToLower(pag.Name) == DEFAULT_PACKAGE {
		wskprint.PrintlnOpenWhiskInfo(wski18n.T(wski18n.ID_MSG_DEFAULT_PACKAGE))
		// when a package is marked public with "Public: true" in manifest file
		// the package is visible to anyone and created with publish state
		// set to shared otherwise publish state is set to private
	} else if pkg.Public {
		warningMsg := wski18n.T(wski18n.ID_WARN_PACKAGE_IS_PUBLIC_X_package_X,
			map[string]interface{}{
				wski18n.KEY_PACKAGE: pag.Name})
		wskprint.PrintlnOpenWhiskWarning(warningMsg)
		pag.Publish = &(pkg.Public)
	}

	return pag, nil
}

func (dm *YAMLParser) ComposeSequencesFromAllPackages(namespace string, mani *YAML, manifestFilePath string, managedAnnotations whisk.KeyValue) ([]utils.ActionRecord, error) {
	var sequences []utils.ActionRecord = make([]utils.ActionRecord, 0)
	manifestPackages := make(map[string]Package)

	if len(mani.Packages) != 0 {
		manifestPackages = mani.Packages
	} else {
		manifestPackages = mani.GetProject().Packages
	}

	for n, p := range manifestPackages {
		s, err := dm.ComposeSequences(namespace, p.Sequences, n, manifestFilePath, managedAnnotations)
		if err == nil {
			sequences = append(sequences, s...)
		} else {
			return nil, err
		}
	}
	return sequences, nil
}

func (dm *YAMLParser) ComposeSequences(namespace string, sequences map[string]Sequence, packageName string, manifestFilePath string, managedAnnotations whisk.KeyValue) ([]utils.ActionRecord, error) {
	var listOfSequences []utils.ActionRecord = make([]utils.ActionRecord, 0)
	var errorParser error

	for key, sequence := range sequences {
		wskaction := new(whisk.Action)
		wskaction.Exec = new(whisk.Exec)
		wskaction.Exec.Kind = YAML_KEY_SEQUENCE
		actionList := strings.Split(sequence.Actions, ",")

		var components []string
		for _, a := range actionList {
			act := strings.TrimSpace(a)
			if !strings.ContainsRune(act, []rune(PATH_SEPARATOR)[0]) && !strings.HasPrefix(act, packageName+PATH_SEPARATOR) &&
				strings.ToLower(packageName) != DEFAULT_PACKAGE {
				act = path.Join(packageName, act)
			}
			components = append(components, path.Join(PATH_SEPARATOR+namespace, act))
		}

		wskaction.Exec.Components = components
		wskaction.Name = key
		pub := false
		wskaction.Publish = &pub
		wskaction.Namespace = namespace

		annotations := dm.composeAnnotations(sequence.Annotations)
		if len(annotations) > 0 {
			wskaction.Annotations = annotations
		}

		// appending managed annotations if its a managed deployment
		if utils.Flags.Managed || utils.Flags.Sync {
			wskaction.Annotations = append(wskaction.Annotations, managedAnnotations)
		}

		// Web Export
		// Treat sequence as a web action, a raw HTTP web action, or as a standard action based on web-export;
		// when web-export is set to yes | true, treat sequence as a web action,
		// when web-export is set to raw, treat sequence as a raw HTTP web action,
		// when web-export is set to no | false, treat sequence as a standard action
		if len(sequence.Web) != 0 {
			wskaction.Annotations, errorParser = utils.WebAction(manifestFilePath, wskaction.Name, sequence.Web, wskaction.Annotations, false)
			if errorParser != nil {
				return nil, errorParser
			}
		}

		record := utils.ActionRecord{Action: wskaction, Packagename: packageName, Filepath: key}
		listOfSequences = append(listOfSequences, record)
	}
	return listOfSequences, nil
}

func (dm *YAMLParser) ComposeActionsFromAllPackages(manifest *YAML, filePath string, managedAnnotations whisk.KeyValue) ([]utils.ActionRecord, error) {
	var actions []utils.ActionRecord = make([]utils.ActionRecord, 0)
	manifestPackages := make(map[string]Package)

	if len(manifest.Packages) != 0 {
		manifestPackages = manifest.Packages
	} else {
		manifestPackages = manifest.GetProject().Packages
	}

	for n, p := range manifestPackages {
		a, err := dm.ComposeActions(filePath, p.Actions, n, managedAnnotations)
		if err == nil {
			actions = append(actions, a...)
		} else {
			return nil, err
		}
	}
	return actions, nil
}

func (dm *YAMLParser) validateActionCode(manifestFilePath string, action Action) error {
	// Check if action.Function is specified with action.Code
	// with action.Code, action.Function is not allowed
	// with action.Code, action.Runtime should be specified
	if len(action.Function) != 0 {
		err := wski18n.T(wski18n.ID_ERR_ACTION_INVALID_X_action_X,
			map[string]interface{}{
				wski18n.KEY_ACTION: action.Name})
		return wskderrors.NewYAMLFileFormatError(manifestFilePath, err)
	}
	if len(action.Runtime) == 0 {
		err := wski18n.T(wski18n.ID_ERR_ACTION_MISSING_RUNTIME_WITH_CODE_X_action_X,
			map[string]interface{}{
				wski18n.KEY_ACTION: action.Name})
		return wskderrors.NewYAMLFileFormatError(manifestFilePath, err)
	}
	return nil
}

func (dm *YAMLParser) readActionCode(manifestFilePath string, action Action) (*whisk.Exec, error) {
	exec := new(whisk.Exec)
	if err := dm.validateActionCode(manifestFilePath, action); err != nil {
		return nil, err
	}
	// validate runtime from the manifest file
	// error out if the specified runtime is not valid or not supported
	// even if runtime is invalid, deploy action with specified runtime in strict mode
	if utils.Flags.Strict {
		exec.Kind = action.Runtime
	} else if utils.CheckExistRuntime(action.Runtime, utils.SupportedRunTimes) {
		exec.Kind = action.Runtime
	} else if len(utils.DefaultRunTimes[action.Runtime]) != 0 {
		exec.Kind = utils.DefaultRunTimes[action.Runtime]
	} else {
		err := wski18n.T(wski18n.ID_ERR_RUNTIME_INVALID_X_runtime_X_action_X,
			map[string]interface{}{
				wski18n.KEY_RUNTIME: action.Runtime,
				wski18n.KEY_ACTION:  action.Name})
		return nil, wskderrors.NewYAMLFileFormatError(manifestFilePath, err)
	}
	exec.Code = &(action.Code)
	// we can specify the name of the action entry point using main
	if len(action.Main) != 0 {
		exec.Main = action.Main
	}
	return exec, nil
}

func (dm *YAMLParser) validateActionFunction(manifestFileName string, action Action, ext string, kind string) error {
	// produce an error when a runtime could not be derived from the action file extension
	// and its not explicitly specified in the manifest YAML file
	// and action source is not a zip file
	if len(action.Runtime) == 0 && len(action.Docker) == 0 && !action.Native {
		if ext == utils.ZIP_FILE_EXTENSION {
			errMessage := wski18n.T(wski18n.ID_ERR_RUNTIME_INVALID_X_runtime_X_action_X,
				map[string]interface{}{
					wski18n.KEY_RUNTIME: utils.RUNTIME_NOT_SPECIFIED,
					wski18n.KEY_ACTION:  action.Name})
			return wskderrors.NewInvalidRuntimeError(errMessage,
				manifestFileName,
				action.Name,
				utils.RUNTIME_NOT_SPECIFIED,
				utils.ListOfSupportedRuntimes(utils.SupportedRunTimes))
		} else if len(kind) == 0 {
			errMessage := wski18n.T(wski18n.ID_ERR_RUNTIME_ACTION_SOURCE_NOT_SUPPORTED_X_ext_X_action_X,
				map[string]interface{}{
					wski18n.KEY_EXTENSION: ext,
					wski18n.KEY_ACTION:    action.Name})
			return wskderrors.NewInvalidRuntimeError(errMessage,
				manifestFileName,
				action.Name,
				utils.RUNTIME_NOT_SPECIFIED,
				utils.ListOfSupportedRuntimes(utils.SupportedRunTimes))
		}
	}
	return nil
}

func (dm *YAMLParser) readActionFunction(manifestFilePath string, manifestFileName string, action Action) (string, *whisk.Exec, error) {
	var actionFilePath string
	var zipFileName string
	exec := new(whisk.Exec)

	// check if action function is pointing to an URL
	// we do not support if function is pointing to remote directory
	// therefore error out if there is a combination of http/https ending in a directory
	if strings.HasPrefix(action.Function, HTTP) || strings.HasPrefix(action.Function, HTTPS) {
		if len(path.Ext(action.Function)) == 0 {
			err := wski18n.T(wski18n.ID_ERR_ACTION_FUNCTION_REMOTE_DIR_NOT_SUPPORTED_X_action_X_url_X,
				map[string]interface{}{
					wski18n.KEY_ACTION: action.Name,
					wski18n.KEY_URL:    action.Function})
			return actionFilePath, nil, wskderrors.NewYAMLFileFormatError(manifestFilePath, err)
		}
		actionFilePath = action.Function
	} else {
		actionFilePath = strings.TrimRight(manifestFilePath, manifestFileName) + action.Function
	}

	if utils.IsDirectory(actionFilePath) {
		zipFileName = actionFilePath + "." + utils.ZIP_FILE_EXTENSION
		err := utils.NewZipWritter(actionFilePath, zipFileName).Zip()
		if err != nil {
			return actionFilePath, nil, err
		}
		defer os.Remove(zipFileName)
		actionFilePath = zipFileName
	}

	action.Function = actionFilePath

	// determine extension of the given action source file
	ext := filepath.Ext(actionFilePath)
	// drop the "." from file extension
	if len(ext) > 0 && ext[0] == '.' {
		ext = ext[1:]
	}

	// determine default runtime for the given file extension
	var kind string
	r := utils.FileExtensionRuntimeKindMap[ext]
	kind = utils.DefaultRunTimes[r]
	if err := dm.validateActionFunction(manifestFileName, action, ext, kind); err != nil {
		return actionFilePath, nil, err
	}
	exec.Kind = kind

	dat, err := utils.Read(actionFilePath)
	if err != nil {
		return actionFilePath, nil, err
	}
	code := string(dat)
	if ext == utils.ZIP_FILE_EXTENSION || ext == utils.JAR_FILE_EXTENSION {
		code = base64.StdEncoding.EncodeToString([]byte(dat))
	}
	exec.Code = &code

	/*
	*  Action.Runtime
	*  Perform few checks if action runtime is specified in manifest YAML file
	*  (1) Check if specified runtime is one of the supported runtimes by OpenWhisk server
	*  (2) Check if specified runtime is consistent with action source file extensions
	*  Set the action runtime to match with the source file extension, if wskdeploy is not invoked in strict mode
	 */
	if len(action.Runtime) != 0 {
		if utils.CheckExistRuntime(action.Runtime, utils.SupportedRunTimes) {
			// for zip actions, rely on the runtimes from the manifest file as it can not be derived from the action source file extension
			// pick runtime from manifest file if its supported by OpenWhisk server
			if ext == utils.ZIP_FILE_EXTENSION {
				exec.Kind = action.Runtime
			} else {
				if utils.CheckRuntimeConsistencyWithFileExtension(ext, action.Runtime) {
					exec.Kind = action.Runtime
				} else {
					warnStr := wski18n.T(wski18n.ID_ERR_RUNTIME_MISMATCH_X_runtime_X_ext_X_action_X,
						map[string]interface{}{
							wski18n.KEY_RUNTIME:   action.Runtime,
							wski18n.KEY_EXTENSION: ext,
							wski18n.KEY_ACTION:    action.Name})
					wskprint.PrintOpenWhiskWarning(warnStr)

					// even if runtime is not consistent with file extension, deploy action with specified runtime in strict mode
					if utils.Flags.Strict {
						exec.Kind = action.Runtime
					} else {
						warnStr := wski18n.T(wski18n.ID_WARN_RUNTIME_CHANGED_X_runtime_X_action_X,
							map[string]interface{}{
								wski18n.KEY_RUNTIME: exec.Kind,
								wski18n.KEY_ACTION:  action.Name})
						wskprint.PrintOpenWhiskWarning(warnStr)
					}
				}
			}
		} else {
			warnStr := wski18n.T(wski18n.ID_ERR_RUNTIME_INVALID_X_runtime_X_action_X,
				map[string]interface{}{
					wski18n.KEY_RUNTIME: action.Runtime,
					wski18n.KEY_ACTION:  action.Name})
			wskprint.PrintOpenWhiskWarning(warnStr)

			if ext == utils.ZIP_FILE_EXTENSION {
				// for zip action, error out if specified runtime is not supported by
				// OpenWhisk server
				return actionFilePath, nil, wskderrors.NewInvalidRuntimeError(warnStr,
					manifestFileName,
					action.Name,
					action.Runtime,
					utils.ListOfSupportedRuntimes(utils.SupportedRunTimes))
			} else {
				warnStr := wski18n.T(wski18n.ID_WARN_RUNTIME_CHANGED_X_runtime_X_action_X,
					map[string]interface{}{
						wski18n.KEY_RUNTIME: exec.Kind,
						wski18n.KEY_ACTION:  action.Name})
				wskprint.PrintOpenWhiskWarning(warnStr)
			}

		}
	}
	// we can specify the name of the action entry point using main
	if len(action.Main) != 0 {
		exec.Main = action.Main
	}

	return actionFilePath, exec, nil
}

func (dm *YAMLParser) composeActionExec(manifestFilePath string, manifestFileName string, action Action) (string, *whisk.Exec, error) {
	var actionFilePath string
	exec := new(whisk.Exec)
	var err error

	if len(action.Code) != 0 {
		exec, err = dm.readActionCode(manifestFilePath, action)
		if err != nil {
			return actionFilePath, nil, err
		}
	}
	if len(action.Function) != 0 {
		actionFilePath, exec, err = dm.readActionFunction(manifestFilePath, manifestFileName, action)
		if err != nil {
			return actionFilePath, nil, err
		}
	}

	// when an action has Docker image specified,
	// set exec.Kind to "blackbox" and
	// set exec.Image to specified image e.g. dockerhub/image
	// when an action Native is set to true,
	// set exec.Image to openwhisk/skeleton
	if len(action.Docker) != 0 || action.Native {
		exec.Kind = utils.BLACKBOX
		if action.Native {
			exec.Image = NATIVE_DOCKER_IMAGE
		} else {
			exec.Image = action.Docker
		}
	}

	return actionFilePath, exec, err
}

func (dm *YAMLParser) validateActionLimits(limits Limits) {
	// TODO() use LIMITS_UNSUPPORTED in yamlparser to enumerate through instead of hardcoding
	// emit warning errors if these limits are not nil
	utils.NotSupportLimits(limits.ConcurrentActivations, LIMIT_VALUE_CONCURRENT_ACTIVATIONS)
	utils.NotSupportLimits(limits.UserInvocationRate, LIMIT_VALUE_USER_INVOCATION_RATE)
	utils.NotSupportLimits(limits.CodeSize, LIMIT_VALUE_CODE_SIZE)
	utils.NotSupportLimits(limits.ParameterSize, LIMIT_VALUE_PARAMETER_SIZE)
}

func (dm *YAMLParser) composeActionLimits(limits Limits) *whisk.Limits {
	dm.validateActionLimits(limits)
	wsklimits := new(whisk.Limits)
	for _, t := range LIMITS_SUPPORTED {
		switch t {
		case LIMIT_VALUE_TIMEOUT:
			if utils.LimitsTimeoutValidation(limits.Timeout) {
				wsklimits.Timeout = limits.Timeout
			} else {
				warningString := wski18n.T(wski18n.ID_WARN_LIMIT_IGNORED_X_limit_X,
					map[string]interface{}{wski18n.KEY_LIMIT: LIMIT_VALUE_TIMEOUT})
				wskprint.PrintOpenWhiskWarning(warningString)
			}
		case LIMIT_VALUE_MEMORY_SIZE:
			if utils.LimitsMemoryValidation(limits.Memory) {
				wsklimits.Memory = limits.Memory
			} else {
				warningString := wski18n.T(wski18n.ID_WARN_LIMIT_IGNORED_X_limit_X,
					map[string]interface{}{wski18n.KEY_LIMIT: LIMIT_VALUE_MEMORY_SIZE})
				wskprint.PrintOpenWhiskWarning(warningString)

			}
		case LIMIT_VALUE_LOG_SIZE:
			if utils.LimitsLogsizeValidation(limits.Logsize) {
				wsklimits.Logsize = limits.Logsize
			} else {
				warningString := wski18n.T(wski18n.ID_WARN_LIMIT_IGNORED_X_limit_X,
					map[string]interface{}{wski18n.KEY_LIMIT: LIMIT_VALUE_LOG_SIZE})
				wskprint.PrintOpenWhiskWarning(warningString)

			}
		}
	}
	if wsklimits.Timeout != nil || wsklimits.Memory != nil || wsklimits.Logsize != nil {
		return wsklimits
	}
	return nil
}

func (dm *YAMLParser) validateActionWebFlag(action Action) {
	if len(action.Web) != 0 && len(action.Webexport) != 0 {
		warningString := wski18n.T(wski18n.ID_WARN_ACTION_WEB_X_action_X,
			map[string]interface{}{wski18n.KEY_ACTION: action.Name})
		wskprint.PrintOpenWhiskWarning(warningString)
	}
}

func (dm *YAMLParser) ComposeActions(manifestFilePath string, actions map[string]Action, packageName string, managedAnnotations whisk.KeyValue) ([]utils.ActionRecord, error) {

	var errorParser error
	var listOfActions []utils.ActionRecord = make([]utils.ActionRecord, 0)
	splitManifestFilePath := strings.Split(manifestFilePath, string(PATH_SEPARATOR))
	manifestFileName := splitManifestFilePath[len(splitManifestFilePath)-1]

	for actionName, action := range actions {
		var actionFilePath string

		// update the action (of type Action) to set its name
		// here key name is the action name
		action.Name = actionName

		// Create action data object from client library
		wskaction := new(whisk.Action)

		//set action.Function to action.Location
		//because Location is deprecated in Action entity
		if len(action.Function) == 0 && len(action.Location) != 0 {
			action.Function = action.Location
		}

		actionFilePath, wskaction.Exec, errorParser = dm.composeActionExec(manifestFilePath, manifestFileName, action)
		if errorParser != nil {
			return nil, errorParser
		}

		// Action.Inputs
		listOfInputs, err := dm.composeInputsOrOutputs(action.Inputs, manifestFilePath)
		if err != nil {
			return nil, err
		}
		if len(listOfInputs) > 0 {
			wskaction.Parameters = listOfInputs
		}

		// Action.Outputs
		// TODO{} add outputs as annotations (work to discuss officially supporting for compositions)
		listOfOutputs, err := dm.composeInputsOrOutputs(action.Outputs, manifestFilePath)
		if err != nil {
			return nil, err
		}
		if len(listOfOutputs) > 0 {
			//wskaction.Annotations = listOfOutputs
		}

		// Action.Annotations
		if listOfAnnotations := dm.composeAnnotations(action.Annotations); len(listOfAnnotations) > 0 {
			wskaction.Annotations = append(wskaction.Annotations, listOfAnnotations...)
		}

		// add managed annotations if its marked as managed deployment
		if utils.Flags.Managed || utils.Flags.Sync {
			wskaction.Annotations = append(wskaction.Annotations, managedAnnotations)
		}

		// Web Export
		// Treat ACTION as a web action, a raw HTTP web action, or as a standard action based on web-export;
		// when web-export is set to yes | true, treat action as a web action,
		// when web-export is set to raw, treat action as a raw HTTP web action,
		// when web-export is set to no | false, treat action as a standard action
		dm.validateActionWebFlag(action)
		if len(action.GetWeb()) != 0 {
			wskaction.Annotations, errorParser = utils.WebAction(manifestFilePath, action.Name, action.GetWeb(), wskaction.Annotations, false)
			if errorParser != nil {
				return listOfActions, errorParser
			}
		}

		// Action.Limits
		if action.Limits != nil {
			if wsklimits := dm.composeActionLimits(*(action.Limits)); wsklimits != nil {
				wskaction.Limits = wsklimits
			}
		}
		// Conductor Action
		if action.Conductor {
			wskaction.Annotations = append(wskaction.Annotations, utils.ConductorAction())
		}

		wskaction.Name = actionName
		pub := false
		wskaction.Publish = &pub

		record := utils.ActionRecord{Action: wskaction, Packagename: packageName, Filepath: actionFilePath}
		listOfActions = append(listOfActions, record)
	}

	return listOfActions, nil

}

func (dm *YAMLParser) ComposeTriggersFromAllPackages(manifest *YAML, filePath string, managedAnnotations whisk.KeyValue) ([]*whisk.Trigger, error) {
	var triggers []*whisk.Trigger = make([]*whisk.Trigger, 0)
	manifestPackages := make(map[string]Package)

	if len(manifest.Packages) != 0 {
		manifestPackages = manifest.Packages
	} else {
		manifestPackages = manifest.GetProject().Packages
	}

	for _, p := range manifestPackages {
		t, err := dm.ComposeTriggers(filePath, p, managedAnnotations)
		if err == nil {
			triggers = append(triggers, t...)
		} else {
			return nil, err
		}
	}
	return triggers, nil
}

func (dm *YAMLParser) ComposeTriggers(filePath string, pkg Package, managedAnnotations whisk.KeyValue) ([]*whisk.Trigger, error) {
	var errorParser error
	var listOfTriggers []*whisk.Trigger = make([]*whisk.Trigger, 0)

	for _, trigger := range pkg.GetTriggerList() {
		wsktrigger := new(whisk.Trigger)
		wsktrigger.Name = wskenv.ConvertSingleName(trigger.Name)
		wsktrigger.Namespace = trigger.Namespace
		pub := false
		wsktrigger.Publish = &pub

		// print warning information when .Source key's value is not empty
		if len(trigger.Source) != 0 {
			warningString := wski18n.T(
				wski18n.ID_WARN_KEY_DEPRECATED_X_oldkey_X_filetype_X_newkey_X,
				map[string]interface{}{
					wski18n.KEY_OLD:       YAML_KEY_SOURCE,
					wski18n.KEY_NEW:       YAML_KEY_FEED,
					wski18n.KEY_FILE_TYPE: wski18n.MANIFEST_FILE})
			wskprint.PrintOpenWhiskWarning(warningString)
		}
		if len(trigger.Feed) == 0 {
			trigger.Feed = trigger.Source
		}

		// replacing env. variables here in the trigger feed name
		// to support trigger feed with $READ_FROM_ENV_TRIGGER_FEED
		trigger.Feed = wskenv.InterpolateStringWithEnvVar(trigger.Feed).(string)

		keyValArr := make(whisk.KeyValueArr, 0)
		if len(trigger.Feed) != 0 {
			var keyVal whisk.KeyValue
			keyVal.Key = YAML_KEY_FEED
			keyVal.Value = trigger.Feed
			keyValArr = append(keyValArr, keyVal)
			wsktrigger.Annotations = keyValArr
		}

		inputs, err := dm.composeInputsOrOutputs(trigger.Inputs, filePath)
		if err != nil {
			return nil, errorParser
		}
		if len(inputs) > 0 {
			wsktrigger.Parameters = inputs
		}

		listOfAnnotations := dm.composeAnnotations(trigger.Annotations)
		if len(listOfAnnotations) > 0 {
			wsktrigger.Annotations = append(wsktrigger.Annotations, listOfAnnotations...)
		}

		// add managed annotations if its a managed deployment
		if utils.Flags.Managed || utils.Flags.Sync {
			wsktrigger.Annotations = append(wsktrigger.Annotations, managedAnnotations)
		}

		listOfTriggers = append(listOfTriggers, wsktrigger)
	}
	return listOfTriggers, nil
}

func (dm *YAMLParser) ComposeRulesFromAllPackages(manifest *YAML, managedAnnotations whisk.KeyValue) ([]*whisk.Rule, error) {
	var rules []*whisk.Rule = make([]*whisk.Rule, 0)
	manifestPackages := make(map[string]Package)

	if len(manifest.Packages) != 0 {
		manifestPackages = manifest.Packages
	} else {
		manifestPackages = manifest.GetProject().Packages
	}

	for n, p := range manifestPackages {
		r, err := dm.ComposeRules(p, n, managedAnnotations)
		if err == nil {
			rules = append(rules, r...)
		} else {
			return nil, err
		}
	}
	return rules, nil
}

func (dm *YAMLParser) ComposeRules(pkg Package, packageName string, managedAnnotations whisk.KeyValue) ([]*whisk.Rule, error) {
	var rules []*whisk.Rule = make([]*whisk.Rule, 0)

	for _, rule := range pkg.GetRuleList() {
		wskrule := new(whisk.Rule)
		wskrule.Name = wskenv.ConvertSingleName(rule.Name)
		//wskrule.Namespace = rule.Namespace
		pub := false
		wskrule.Publish = &pub
		wskrule.Trigger = wskenv.ConvertSingleName(rule.Trigger)
		wskrule.Action = wskenv.ConvertSingleName(rule.Action)
		act := strings.TrimSpace(wskrule.Action.(string))
		if !strings.ContainsRune(act, []rune(PATH_SEPARATOR)[0]) && !strings.HasPrefix(act, packageName+PATH_SEPARATOR) &&
			strings.ToLower(packageName) != DEFAULT_PACKAGE {
			act = path.Join(packageName, act)
		}
		wskrule.Action = act
		listOfAnnotations := dm.composeAnnotations(rule.Annotations)
		if len(listOfAnnotations) > 0 {
			wskrule.Annotations = append(wskrule.Annotations, listOfAnnotations...)
		}

		// add managed annotations if its a managed deployment
		if utils.Flags.Managed || utils.Flags.Sync {
			wskrule.Annotations = append(wskrule.Annotations, managedAnnotations)
		}

		rules = append(rules, wskrule)
	}
	return rules, nil
}

func (dm *YAMLParser) ComposeApiRecordsFromAllPackages(client *whisk.Config, manifest *YAML) ([]*whisk.ApiCreateRequest, error) {
	var requests []*whisk.ApiCreateRequest = make([]*whisk.ApiCreateRequest, 0)
	manifestPackages := make(map[string]Package)

	if len(manifest.Packages) != 0 {
		manifestPackages = manifest.Packages
	} else {
		manifestPackages = manifest.GetProject().Packages
	}

	for packageName, p := range manifestPackages {
		r, err := dm.ComposeApiRecords(client, packageName, p, manifest.Filepath)
		if err == nil {
			requests = append(requests, r...)
		} else {
			return nil, err
		}
	}
	return requests, nil
}

/*
 * read API section from manifest file:
 * apis: # List of APIs
 *     hello-world: #API name
 *	/hello: #gateway base path
 *	    /world:   #gateway rel path
 *		greeting: get #action name: gateway method
 *
 * compose APIDoc structure from the manifest:
 * {
 *	"apidoc":{
 *      	"namespace":<namespace>,
 *      	"gatewayBasePath":"/hello",
 *      	"gatewayPath":"/world",
 *      	"gatewayMethod":"GET",
 *      	"action":{
 *         		"name":"hello",
 *			"namespace":"guest",
 *			"backendMethod":"GET",
 *			"backendUrl":<url>,
 *			"authkey":<auth>
 *		}
 * 	}
 * }
 */
func (dm *YAMLParser) ComposeApiRecords(client *whisk.Config, packageName string, pkg Package, manifestPath string) ([]*whisk.ApiCreateRequest, error) {
	var requests []*whisk.ApiCreateRequest = make([]*whisk.ApiCreateRequest, 0)

	if pkg.Apis != nil && len(pkg.Apis) != 0 {
		// verify APIGW_ACCESS_TOKEN is set before composing APIs
		// until this point, we dont know whether APIs are specified in manifest or not
		if len(client.ApigwAccessToken) == 0 {
			return nil, wskderrors.NewWhiskClientInvalidConfigError(wski18n.ID_MSG_CONFIG_MISSING_APIGW_ACCESS_TOKEN)
		}
	}

	for apiName, apiDoc := range pkg.Apis {
		for gatewayBasePath, gatewayBasePathMap := range apiDoc {
			// append "/" to the gateway base path if its missing
			if !strings.HasPrefix(gatewayBasePath, PATH_SEPARATOR) {
				gatewayBasePath = PATH_SEPARATOR + gatewayBasePath
			}
			for gatewayRelPath, gatewayRelPathMap := range gatewayBasePathMap {
				// append "/" to the gateway relative path if its missing
				if !strings.HasPrefix(gatewayRelPath, PATH_SEPARATOR) {
					gatewayRelPath = PATH_SEPARATOR + gatewayRelPath
				}
				for actionName, gatewayMethod := range gatewayRelPathMap {
					// verify that the action is defined under actions sections
					if _, ok := pkg.Actions[actionName]; !ok {
						return nil, wskderrors.NewYAMLFileFormatError(manifestPath,
							wski18n.T(wski18n.ID_ERR_API_MISSING_ACTION_X_action_X_api_X,
								map[string]interface{}{
									wski18n.KEY_ACTION: actionName,
									wski18n.KEY_API:    apiName}))
					} else {
						// verify that the action is defined as web action
						// web or web-export set to any of [true, yes, raw]
						a := pkg.Actions[actionName]
						if !utils.IsWebAction(a.GetWeb()) {
							warningString := wski18n.T(wski18n.ID_WARN_API_MISSING_WEB_ACTION_X_action_X_api_X,
								map[string]interface{}{
									wski18n.KEY_ACTION: actionName,
									wski18n.KEY_API:    apiName})
							wskprint.PrintOpenWhiskWarning(warningString)
							if a.Annotations == nil {
								a.Annotations = make(map[string]interface{}, 0)
							}
							a.Annotations[utils.WEB_EXPORT_ANNOT] = true
							pkg.Actions[actionName] = a
						}
						request := new(whisk.ApiCreateRequest)
						request.ApiDoc = new(whisk.Api)
						request.ApiDoc.GatewayBasePath = gatewayBasePath
						// is API verb is valid, it must be one of (GET, PUT, POST, DELETE)
						request.ApiDoc.GatewayRelPath = gatewayRelPath
						if _, ok := whisk.ApiVerbs[strings.ToUpper(gatewayMethod)]; !ok {
							return nil, wskderrors.NewInvalidAPIGatewayMethodError(manifestPath,
								gatewayBasePath+gatewayRelPath,
								gatewayMethod,
								dm.getGatewayMethods())
						}
						request.ApiDoc.GatewayMethod = strings.ToUpper(gatewayMethod)
						request.ApiDoc.Namespace = client.Namespace
						request.ApiDoc.ApiName = apiName
						request.ApiDoc.Id = strings.Join([]string{API, request.ApiDoc.Namespace, request.ApiDoc.GatewayRelPath}, ":")
						// set action of an API Doc
						request.ApiDoc.Action = new(whisk.ApiAction)
						if packageName == DEFAULT_PACKAGE {
							request.ApiDoc.Action.Name = actionName
						} else {
							request.ApiDoc.Action.Name = packageName + PATH_SEPARATOR + actionName
						}
						url := []string{HTTPS + client.Host, strings.ToLower(API),
							API_VERSION, WEB, client.Namespace, packageName,
							actionName + "." + utils.HTTP_FILE_EXTENSION}
						request.ApiDoc.Action.Namespace = client.Namespace
						request.ApiDoc.Action.BackendUrl = strings.Join(url, PATH_SEPARATOR)
						request.ApiDoc.Action.BackendMethod = gatewayMethod
						request.ApiDoc.Action.Auth = client.AuthToken
						// add a newly created ApiCreateRequest object to a list of requests
						requests = append(requests, request)
					}
				}
			}
		}
	}
	return requests, nil
}

func (dm *YAMLParser) getGatewayMethods() []string {
	methods := []string{}
	for k := range whisk.ApiVerbs {
		methods = append(methods, k)
	}
	return methods
}

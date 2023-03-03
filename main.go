package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	xfnv1alpha1 "github.com/crossplane/crossplane/apis/apiextensions/fn/io/v1alpha1"
	"sigs.k8s.io/yaml"

	// redisV1 "github.com/vshn/component-appcat/apis/vshn/v1"
	redisV1 "xfnTest/redis"
)

func main() {
	funcIO := xfnv1alpha1.FunctionIO{}

	x, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(x, &funcIO)
	if err != nil {
		log.Fatal(err)
	}

	redisXR := &redisV1.VSHNRedis{}
	err = json.Unmarshal(funcIO.Observed.Composite.Resource.Raw, redisXR)
	if err != nil {
		log.Fatal(err)
	}

	if _, ok := redisXR.GetAnnotations()["appcat.io/effectiveVersion"]; !ok {
		if redisXR.ObjectMeta.Annotations == nil {
			redisXR.ObjectMeta.Annotations = map[string]string{}
		}
		redisXR.ObjectMeta.Annotations["appcat.io/effectiveVersion"] = redisXR.Spec.Parameters.Service.Version
	}

	rawData, err := json.Marshal(redisXR)
	if err != nil {
		log.Fatal(err)
	}

	funcIO.Desired.Composite.Resource.Raw = rawData

	funcIO.Results = append(funcIO.Results, xfnv1alpha1.Result{
		Severity: xfnv1alpha1.SeverityNormal,
		Message:  "I was reconciled",
	})

	finalYaml, err := yaml.Marshal(funcIO)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(finalYaml))

}

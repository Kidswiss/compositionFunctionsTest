package main

import (
	"encoding/json"
	"log"
	"os"

	xfnv1alpha1 "github.com/crossplane/crossplane/apis/apiextensions/fn/io/v1alpha1"
	// redisV1 "github.com/vshn/component-appcat/apis/vshn/v1"
	redisV1 "xfnTest/redis"
)

func main() {
	funcIO := xfnv1alpha1.FunctionIO{}

	err := json.NewDecoder(os.Stdin).Decode(&funcIO)
	if err != nil {
		log.Fatal(err)
	}

	redisXR := &redisV1.VSHNRedis{}
	err = json.Unmarshal(funcIO.Observed.Composite.Resource.Raw, redisXR)
	if err != nil {
		log.Fatal(err)
	}

	if redisXR.Status.EffectiveVersion == "" {
		redisXR.Status.EffectiveVersion = redisXR.Spec.Parameters.Service.Version
	}

	rawData, err := json.Marshal(redisXR)
	if err != nil {
		log.Fatal(err)
	}

	funcIO.Desired.Composite.Resource.Raw = rawData

	json.NewEncoder(os.Stdout).Encode(funcIO)

}

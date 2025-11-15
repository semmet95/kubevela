//go:build mage

package main

import (
	"encoding/json"
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type K3D mg.Namespace

const (
	defaultClusterName = "vela-test"
	serverNodes        = 1
	agentNodes         = 1
)

func (K3D) Create() {
	if checkClusterExists(defaultClusterName) {
		deleteCluster(defaultClusterName)
	}

	output, err := sh.Output(
		"k3d", "cluster", "create", defaultClusterName, "--servers", fmt.Sprint(serverNodes), "--agents", fmt.Sprint(agentNodes),
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf(output)
}

func deleteCluster(name string) {
	output, err := sh.Output("k3d", "cluster", "delete", name)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}

func checkClusterExists(name string) bool {
	output, err := sh.Output("k3d", "cluster", "list", "-o", "json")
	if err != nil {
		panic(err)
	}

	var clusters []map[string]interface{}
	if err := json.Unmarshal([]byte(output), &clusters); err != nil {
		panic(err)
	}

	for _, cluster := range clusters {
		if cluster["name"].(string) == name {
			return true
		}
	}

	return false
}

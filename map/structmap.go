package main

import "fmt"

// RolloutDefinitionSpec defines the desired state of RolloutDefinition
type RolloutDefinitionSpec struct {
	ControlResource ControlResource `json:"controlResource"`
	Path            Path            `json:"path,omitempty"`
}

// ControlResource defines the controlled resource
type ControlResource struct {
	APIVersion string `json:"apiVersion"`
	Resource   string `json:"resource"`
}

// Path indicates the path
type Path struct {
	SpecPath   SpecPath   `json:"specPath,omitempty"`
	StatusPath StatusPath `json:"statusPath,omitempty"`
}

// SpecPath indicates the spec path
type SpecPath struct {
	// Paused indicates the paused path of controlled workload
	Paused string `json:"paused,omitempty"`
	// Partition indicates the partition path of controlled workload.
	Partition string `json:"partition,omitempty"`
	// MaxUnavailable indicates the maxUnavailable path of controlled workload.
	MaxUnavailable string `json:"maxUnavailable,omitempty"`
}

// StatusPath indicates the status path
type StatusPath struct {
	// ObservedGeneration indicates the observedGeneration path of controlled workload
	ObservedGeneration string `json:"observedGeneration,omitempty"`
	// Replicas indicates the replicas path of controlled workload
	Replicas string `json:"replicas"`
	// ReadyReplicas indicates the readyReplicas path of controlled workload
	ReadyReplicas string `json:"readyReplicas,omitempty"`
	// CurrentReplicas indicates the currentReplicas path of controlled workload
	CurrentReplicas string `json:"currentReplicas,omitempty"`
	// UpdatedReplicas indicates the updatedReplicas path of controlled workload
	UpdatedReplicas string `json:"updatedReplicas,omitempty"`
	// Represents the latest available observations of current state.
	Conditions string `json:"conditions,omitempty"`
}

type data struct {
	name string
}

func main() {
	m := map[string]*data{"x": {"one"}}

	m["x"] = &data{"two"}
	m["x"].name = "three"
	m["x"].name = "four"
	fmt.Println(m["x"]) //prints: &{four}

	resource := ControlResource{"version", "res"}
	fmt.Printf("resource: %s", resource)
	fmt.Println()
	fmt.Printf("map: %v", m["x"])
}

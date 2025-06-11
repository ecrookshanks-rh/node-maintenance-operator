package controllers

import (
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ContainsString checks if the string array contains the given string.
func ContainsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

// RemoveString removes the given string from the string array if exists.
func RemoveString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return result
}

// GetPodNameList returns a list of pod identifiers in the form "namespace/name".
func GetPodNameList(pods []corev1.Pod) (result []string) {
	for _, pod := range pods {
		result = append(result, client.ObjectKeyFromObject(&pod).String())
	}
	return result
}

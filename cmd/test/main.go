package main

import (
	"context"
	"time"

	v1net "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/retry"
)

func main() {
	// Load kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		panic(err)
	}

	// Create Kubernetes client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// Create NetworkPolicy client
	policyClient := clientset.NetworkingV1().NetworkPolicies("server")

	// Create NetworkPolicy object
	policy := &v1net.NetworkPolicy{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "deny-server-egress",
			
			Namespace: "server",
		},
		Spec: v1net.NetworkPolicySpec{
			PodSelector: metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "server",
				},
			},
			PolicyTypes: []v1net.PolicyType{
				v1net.PolicyTypeEgress,
			},
			Egress: []v1net.NetworkPolicyEgressRule{},
		},
	}

	// Create NetworkPolicy
	_, err = policyClient.Create(context.Background(), policy, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}

	// Wait for 60 seconds
	time.Sleep(60 * time.Second)

	// Delete NetworkPolicy
	err = retry.RetryOnConflict(retry.DefaultBackoff, func() error {
		return policyClient.Delete(context.Background(), policy.Name, metav1.DeleteOptions{
			GracePeriodSeconds: new(int64),
		})
	})
	if err != nil {
		panic(err)
	}
}

package kube

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"path/filepath"
	"strconv"

	"ext-github.swm.de/SWM/rancher-sources/kubestatus/internal/models"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

type Kube struct {
	Client        *kubernetes.Clientset
	MetricsClient *versioned.Clientset
	Config        *rest.Config
}

// NewKube creates a new Kube object
func NewKube(accesstype string, kubeConfigPath string) (*Kube, error) {
	fmt.Println("accesstype: " + accesstype)
	k := Kube{}

	if accesstype == "incluster" {
		fmt.Println("Starting in cluster")
		config, err := rest.InClusterConfig()
		k.Config = config
		if err != nil {
			panic(err.Error())
		}
		// creates the clientset
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			return nil, err
		}

		k.Client = clientset
	} else if accesstype == "outcluster" {
		var kubeconfig *string

		if kubeConfigPath != "" {
			kubeconfig = flag.String("kubeconfig", kubeConfigPath, "absolute path to the kubeconfig file")
		} else {
			home := homedir.HomeDir()
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "path to the kubeconfig file")
		}
		flag.Parse()

		// use the current context in kubeconfig
		config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			panic(err.Error())
		}
		k.Config = config

		// create the clientset
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}
		k.Client = clientset
	} else {
		return nil, fmt.Errorf("wrong access type " + accesstype)
	}

	return &k, nil

}

func NewMetricsClient(config *rest.Config) (*versioned.Clientset, error) {
	// Metrics Client erstellen
	metricsClient, err := versioned.NewForConfig(config)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error creating metrics client: %v", err))
	}
	return metricsClient, nil

}

func (k *Kube) GetNodeCount() (int, error) {
	nodes, err := k.Client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return -1, err
	}
	return len(nodes.Items), nil
}

func (k *Kube) GetStatus() (*models.ClusterStatus, error) {

	var clusterStatus models.ClusterStatus

	// count the nodes
	nodes, err := k.Client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		//return c.String(http.StatusInternalServerError, fmt.Sprintf("Error fetching nodes: %v", err))
		return nil, errors.New(fmt.Sprintf("Error fetching nodes: %v", err))

	}
	var totalCPUUsed, totalRAMUsed, totalCPUCapacity, totalRAMCapacity resource.Quantity

	// fetch the resources from each node and summarize
	metricsClient, err := NewMetricsClient(k.Config)
	if err != nil {
		return nil, err
	}
	for _, node := range nodes.Items {
		metrics, err := metricsClient.MetricsV1beta1().NodeMetricses().Get(context.TODO(), node.Name, metav1.GetOptions{})
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error fetching metrics: %v", err))
		}

		totalCPUUsed.Add(metrics.Usage[corev1.ResourceCPU])
		totalRAMUsed.Add(metrics.Usage[corev1.ResourceMemory])

		totalCPUCapacity.Add(node.Status.Capacity[corev1.ResourceCPU])
		totalRAMCapacity.Add(node.Status.Capacity[corev1.ResourceMemory])
	}

	clusterStatus.NodeCount = len(nodes.Items)

	clusterStatus.CPU.Used = strconv.FormatInt(totalCPUUsed.ScaledValue(-3), 10)
	clusterStatus.CPU.Capacity = strconv.FormatInt(totalCPUCapacity.ScaledValue(-3), 10)
	clusterStatus.Memory.Used = strconv.FormatInt(totalRAMUsed.ScaledValue(resource.Mega), 10)
	clusterStatus.Memory.Capacity = strconv.FormatInt(totalRAMCapacity.ScaledValue(resource.Mega), 10)

	return &clusterStatus, nil

}

package kube

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"regexp"
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

func (k *Kube) FetchMemoryUsage() (*[]byte, error) {

	var pods models.PodMetricsList
	data, err := k.Client.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/pods").DoRaw(context.TODO())
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	err = json.Unmarshal(data, &pods)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	//convert podmetrics do data

	namespacemap := make(map[string]models.Namespace)

	for _, m := range pods.Items {

		var pod models.Pod
		pod.Name = m.Metadata.Name

		// wenn der namespace schon gefunden wurde
		if _, ok := namespacemap[m.Metadata.Namespace]; ok {

		} else {
			var namespaces models.Namespace
			namespaces.Name = m.Metadata.Namespace
			namespacemap[m.Metadata.Namespace] = namespaces
		}

		for _, containerData := range m.Containers {

			containerName := containerData.Name
			memString := containerData.Usage.Memory
			cpuString := containerData.Usage.CPU
			rx := regexp.MustCompile("[0-9]*") // regex Compilieren

			bytea := rx.Find([]byte(memString)) // byte[] mit allen matches des regex
			dstring := string(bytea[:])         // string aus dem byte[]
			sumMem, _ := strconv.Atoi(dstring)  // int aus dem string machen

			bytea = rx.Find([]byte(cpuString)) // byte[] mit allen matches des regex
			dstring = string(bytea[:])         // string aus dem byte[]
			sumCpu, _ := strconv.Atoi(dstring) // int aus dem string machen

			var container = models.Container{
				Name: containerName,
				MEM:  sumMem,
				CPU:  sumCpu,
			}

			pod.Containers = append(pod.Containers, container)

		}
		namespace := namespacemap[m.Metadata.Namespace]
		namespace.Pods = append(namespace.Pods, pod)
		namespacemap[m.Metadata.Namespace] = namespace

	}

	var Payload models.Data

	for _, element := range namespacemap {
		Payload.Namespaces = append(Payload.Namespaces, element)
	}

	/* 	datafile, err := json.MarshalIndent(Payload, "", "\t")
	   	if err != nil {
	   		log.Println(err.Error())
	   		return nil, err
	   	} */

	// write json file
	/* 	err = os.WriteFile("static/data.json", datafile, 0644)
	   	if err != nil {
	   		log.Println(err.Error())
	   		return nil, err
	   	}
	*/
	returndata, _ := json.MarshalIndent(Payload, "", "\t")
	return &returndata, nil

}

func (k *Kube) FetchCpuUsage() (*[]byte, error) {

	var pods models.PodMetricsList
	data, err := k.Client.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/pods").DoRaw(context.TODO())
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	err = json.Unmarshal(data, &pods)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	//convert podmetrics do data

	namespacemap := make(map[string]models.Namespace)

	for _, m := range pods.Items {

		var pod models.Pod
		pod.Name = m.Metadata.Name

		// wenn der namespace schon gefunden wurde
		if _, ok := namespacemap[m.Metadata.Namespace]; ok {

		} else {
			var namespaces models.Namespace
			namespaces.Name = m.Metadata.Namespace
			namespacemap[m.Metadata.Namespace] = namespaces
		}

		for _, containerData := range m.Containers {

			containerName := containerData.Name
			memString := containerData.Usage.Memory
			cpuString := containerData.Usage.CPU
			rx := regexp.MustCompile("[0-9]*") // regex Compilieren

			bytea := rx.Find([]byte(memString)) // byte[] mit allen matches des regex
			dstring := string(bytea[:])         // string aus dem byte[]
			sumMem, _ := strconv.Atoi(dstring)  // int aus dem string machen

			bytea = rx.Find([]byte(cpuString)) // byte[] mit allen matches des regex
			dstring = string(bytea[:])         // string aus dem byte[]
			sumCpu, _ := strconv.Atoi(dstring) // int aus dem string machen

			var container = models.Container{
				Name: containerName,
				MEM:  sumMem,
				CPU:  sumCpu,
			}

			pod.Containers = append(pod.Containers, container)

		}
		namespace := namespacemap[m.Metadata.Namespace]
		namespace.Pods = append(namespace.Pods, pod)
		namespacemap[m.Metadata.Namespace] = namespace

	}

	var Payload models.Data

	for _, element := range namespacemap {
		Payload.Namespaces = append(Payload.Namespaces, element)
	}
	/*
		datafile, err := json.MarshalIndent(Payload, "", "\t")
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	*/
	/* 	// write json file
	   	err = os.WriteFile("static/data.json", datafile, 0644)
	   	if err != nil {
	   		log.Println(err.Error())
	   		return nil, err
	   	} */

	returndata, _ := json.MarshalIndent(Payload, "", "\t")
	return &returndata, nil

}

func (k *Kube) FetchLimitsAndReuests() (*[]byte, error) {

	var pods models.PodList
	data, err := k.Client.RESTClient().Get().AbsPath("api/v1/pods").DoRaw(context.TODO())
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	err = json.Unmarshal(data, &pods)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	// convert Podlist to data.json

	namespacemap := make(map[string]models.Namespace)

	for _, item := range pods.Items {

		var pod models.Pod
		pod.Name = item.Metadata.Name

		// füge den Namespace der Namespacemap hinzu
		if _, ok := namespacemap[item.Metadata.Namespace]; ok {

		} else {
			var namespaces models.Namespace
			namespaces.Name = item.Metadata.Namespace
			namespacemap[item.Metadata.Namespace] = namespaces
		}

		for _, containerData := range item.Spec.Containers {
			containerName := containerData.Name

			// Angaben sind in mCPU oder MByte daher umwandeln
			memRequestedString := containerData.Resources.Requests.Memory
			cpuRequestedString := containerData.Resources.Requests.CPU
			memLimitedString := containerData.Resources.Limits.Memory
			cpuLimitedString := containerData.Resources.Limits.CPU

			// Setze leere Strings auf "0"
			if memRequestedString == "" {
				memRequestedString = "0Mi"
			}
			if cpuRequestedString == "" {
				cpuRequestedString = "0m"
			}
			if memLimitedString == "" {
				memLimitedString = "0Mi"
			}
			if cpuLimitedString == "" {
				cpuLimitedString = "0m"
			}

			//log.Println(containerName + ": " + memRequestedString + ", " + memRequestedString + ", " + memRequestedString + ", " + cpuLimitedString)

			rx := regexp.MustCompile("[0-9]*")           // regex Compilieren
			bytea := rx.Find([]byte(memRequestedString)) // byte[] mit allen matches des regex
			dstring := string(bytea[:])                  // string aus dem byte[]
			memRequested, _ := strconv.Atoi(dstring)     // int aus dem string machen

			// Wenn Angabe in Gi dann umrechnen in Mb
			rxsearch := regexp.MustCompile("Gi")
			res := rxsearch.MatchString(memRequestedString)
			//res, _ := regexp.MatchString(`Gi`, memRequestedString)
			if res {
				memRequested = memRequested * 1024
			}

			bytea = rx.Find([]byte(cpuRequestedString)) // byte[] mit allen matches des regex
			dstring = string(bytea[:])                  // string aus dem byte[]
			cpuRequested, _ := strconv.Atoi(dstring)    // int aus dem string machen

			// Wenn nicht in MiliCPU dann umrechnen
			res, _ = regexp.MatchString(`m`, cpuRequestedString)
			if !res {
				cpuRequested = cpuRequested * 1000
			}

			bytea = rx.Find([]byte(memLimitedString)) // byte[] mit allen matches des regex
			dstring = string(bytea[:])                // string aus dem byte[]
			memLimited, _ := strconv.Atoi(dstring)    // int aus dem string machen

			// Wenn Angabe in Gi dann umrechnen in Mb
			res, _ = regexp.MatchString(`Gi`, memLimitedString)
			if res {
				memLimited = memLimited * 1024
			}

			bytea = rx.Find([]byte(cpuLimitedString)) // byte[] mit allen matches des regex
			dstring = string(bytea[:])                // string aus dem byte[]
			cpuLimited, _ := strconv.Atoi(dstring)    // int aus dem string machen

			// Wenn nicht in MiliCPU dann umrechnen
			res, _ = regexp.MatchString(`m`, cpuLimitedString)
			if !res {
				cpuLimited = cpuLimited * 1000
			}

			var container = models.Container{
				Name:       containerName,
				MEMRequest: memRequested,
				CPURequest: cpuRequested,
				MEMLimit:   memLimited,
				CPULimit:   cpuLimited,
			}
			//log.Println(container)

			pod.Containers = append(pod.Containers, container)
		}
		namespace := namespacemap[item.Metadata.Namespace]
		namespace.Pods = append(namespace.Pods, pod)
		namespacemap[item.Metadata.Namespace] = namespace
	}
	var Payload models.Data

	for _, element := range namespacemap {
		Payload.Namespaces = append(Payload.Namespaces, element)
	}

	/* 	datafile, err := json.MarshalIndent(Payload, "", "\t")
	   	if err != nil {
	   		log.Println(err.Error())
	   		return nil, err
	   	} */

	/* 	// write json file
	   	err = ioutil.WriteFile("static/data.json", datafile, 0644)
	   	if err != nil {
	   		log.Println(err.Error())
	   		return nil, err
	   	} */

	returndata, _ := json.MarshalIndent(Payload, "", "\t")
	return &returndata, nil
}

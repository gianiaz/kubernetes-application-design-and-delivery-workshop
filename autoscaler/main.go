package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/retry"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type LoadMsg struct {
	Load float32
}

func main() {
	deployment_name := os.Getenv("DEPLOYMENT_NAME")
	service := os.Getenv("TARGET_SERVICE")
	port := os.Getenv("TARGET_PORT")
	namespace := os.Getenv("NAMESPACE")
	min, err := strconv.Atoi(os.Getenv("MIN_REPLICAS"))
	if err != nil {
		panic(err.Error())
	}
	max, err := strconv.Atoi(os.Getenv("MAX_REPLICAS"))
	if err != nil {
		panic(err.Error())
	}
	delta_t, err := strconv.Atoi(os.Getenv("SECONDS"))
	var desired_replicas int32
	if err != nil {
		panic(err.Error())
	}
	endpoint := fmt.Sprintf("http://%s:%s/load", service, port)
	log.Printf("target %s, namespace %s, deployment %s, min %v, max %v, sleep %v", endpoint, namespace, deployment_name, min, max, delta_t)
	config, err := rest.InClusterConfig()
	var load LoadMsg
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	deploymentsClient := clientset.AppsV1().Deployments(namespace)
	for {
		log.Printf("Let's try contact %s", endpoint)
		resp, err := http.Get(endpoint)
		if err != nil {
			log.Printf("Error connectig to %s, will retry in %v seconds", endpoint, delta_t)
		} else {
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			json.Unmarshal(body, &load)
			retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
				result, err := deploymentsClient.Get(deployment_name, metav1.GetOptions{})
				if errors.IsNotFound(err) {
					log.Printf("Deployment %s not found\n", deployment_name)
				} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
					log.Printf("Error getting deployment %v\n", statusError.ErrStatus.Message)
				} else if err != nil {
					panic(err.Error())
				}
				desired_replicas = int32(float32(min) + load.Load*float32(max-min))
				log.Printf("Load: %v, desired replicas: %v, actual replicas: %v", load.Load, desired_replicas, *result.Spec.Replicas)
				result.Spec.Replicas = &desired_replicas
				_, updateErr := deploymentsClient.Update(result)
				return updateErr
			})
			if retryErr != nil {
				panic(fmt.Errorf("Update failed: %v", retryErr))
			}
		}
		time.Sleep(time.Duration(delta_t) * time.Second)
	}
}

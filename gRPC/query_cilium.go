package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

type forwardCnt struct {
	ingressValueOld   int64
	egressValueOld    int64
	ingressValueNew   int64
	egressValueNew    int64
	ingressValueDelta int64
	egressValueDelta  int64
}

var shortPodPattern = regexp.MustCompile("^(.+?)(-[a-z0-9]+){1,2}$")

func shortenPodName(name string) string {
	return shortPodPattern.ReplaceAllString(name, "${1}")
}

func getForwardCnt() {
	forwardCntMap := make(map[string]*forwardCnt)

	for {
		resp, err := http.Get(`http://10.0.0.252:39000/api/v1/query?query=cilium_forward_count_total{pod_template_generation="4",instance!="10.0.0.252:9090"}`)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}

		// fmt.Println(string(body))
		// err = ioutil.WriteFile("/home1/root/DeathStarBench/gRPC/out.json", body, 0644)
		// fmt.Println(len(string(body)))
		jsonStr := string(body)
		metricsLen := gjson.Get(jsonStr, "data.result.#").Int()
		// fmt.Println(metricsLen)
		// fmt.Println(reflect.TypeOf(metricsLen))

		var i int64
		for i = 0; i < metricsLen; i++ {
			directionPathStr := fmt.Sprintf("data.result.%v.metric.direction", i)
			directionStr := gjson.Get(jsonStr, directionPathStr).String()

			valuePathStr := fmt.Sprintf("data.result.%v.value", i)
			valueArray := gjson.Get(jsonStr, valuePathStr).Array()

			instancePathStr := fmt.Sprintf("data.result.%v.metric.instance", i)
			instanceStr := gjson.Get(jsonStr, instancePathStr).String()
			splitPos := strings.Index(instanceStr, ":")
			instanceStr = instanceStr[:splitPos]

			if _, exist := forwardCntMap[instanceStr]; !exist {
				forwardCntMap[instanceStr] = new(forwardCnt)
				// forwardCntMap[instanceStr].ingressValue = make([]float64, 2)
				// forwardCntMap[instanceStr].egressValue = make([]float64, 2)
			}

			if directionStr == "INGRESS" {
				if forwardCntMap[instanceStr].ingressValueOld == 0 {
					forwardCntMap[instanceStr].ingressValueOld = valueArray[1].Int()
					forwardCntMap[instanceStr].ingressValueNew = valueArray[1].Int()
				} else {
					forwardCntMap[instanceStr].ingressValueOld = forwardCntMap[instanceStr].ingressValueNew
					forwardCntMap[instanceStr].ingressValueNew = valueArray[1].Int()
					forwardCntMap[instanceStr].ingressValueDelta = forwardCntMap[instanceStr].ingressValueNew - forwardCntMap[instanceStr].ingressValueOld
				}
			} else if directionStr == "EGRESS" {
				if forwardCntMap[instanceStr].egressValueOld == 0 {
					forwardCntMap[instanceStr].egressValueOld = valueArray[1].Int()
					forwardCntMap[instanceStr].egressValueNew = valueArray[1].Int()
				} else {
					forwardCntMap[instanceStr].egressValueOld = forwardCntMap[instanceStr].egressValueNew
					forwardCntMap[instanceStr].egressValueNew = valueArray[1].Int()
					forwardCntMap[instanceStr].egressValueDelta = forwardCntMap[instanceStr].egressValueNew - forwardCntMap[instanceStr].egressValueOld
				}
			}
		}

		for k, v := range forwardCntMap {
			fmt.Printf("node: %v:\n\tingress packets: %v,\n\tegress packets: %v\n", k, v.ingressValueDelta, v.egressValueDelta)
		}
		time.Sleep(time.Duration(2) * time.Second)
		fmt.Println("============\n\n")
	}
}

func getHubbleFlow() {
	resp, err := http.Get(`http://10.0.0.252:39000/api/v1/query?query=hubble_port_distribution_total`)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	// err = ioutil.WriteFile("/home1/root/DeathStarBench/gRPC/out.json", body, 0644)
	fmt.Println(len(string(body)))
}

func main() {
	log.Println("begin")
	// getForwardCnt()
	getHubbleFlow()
}

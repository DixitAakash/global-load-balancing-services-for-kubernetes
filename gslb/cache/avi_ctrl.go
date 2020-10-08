/*
* [2013] - [2020] Avi Networks Incorporated
* All Rights Reserved.
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*   http://www.apache.org/licenses/LICENSE-2.0
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package cache

import (
	"errors"
	"os"
	"sync"

	"github.com/avinetworks/amko/gslb/gslbutils"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/vmware/load-balancer-and-ingress-services-for-kubernetes/pkg/utils"
)

var aviClientInstance *utils.AviRestClientPool

var clientOnce sync.Once

// SharedAviClients initializes a pool of connections to the avi controller
func SharedAviClients() *utils.AviRestClientPool {
	clientOnce.Do(func() {
		var err error

		ctrlCfg := gslbutils.GetAviConfig()
		if ctrlCfg.Username == "" || ctrlCfg.Password == "" || ctrlCfg.IPAddr == "" || ctrlCfg.Version == "" {
			utils.AviLog.Fatal("AVI Controller information is missing, update them in kubernetes secret or via environment variable.")
		}
		os.Setenv("CTRL_VERSION", ctrlCfg.Version)
		aviClientInstance, err = utils.NewAviRestClientPool(gslbutils.NumRestWorkers, ctrlCfg.IPAddr, ctrlCfg.Username, ctrlCfg.Password)
		if err != nil {
			utils.AviLog.Errorf("AVI Controller Initialization failed, %s", err)
		}
	})
	return aviClientInstance
}

func IsAviSiteLeader() (bool, error) {
	aviRestClientPool := SharedAviClients()
	if len(aviRestClientPool.AviClient) < 1 {
		gslbutils.Errf("no avi clients initialized, returning")
		return false, errors.New("no avi clients initialized")
	}

	aviClient := aviRestClientPool.AviClient[0]
	clusterUuid, err := GetClusterUuid(aviClient)
	if err != nil {
		gslbutils.Errf("error in finding controller cluster uuid: %s", err.Error())
		return false, err
	}

	gslbLeaderUuid, err := GetGslbLeaderUuid(aviClient)
	if err != nil {
		gslbutils.Errf("error in finding the GSLB leader's uuid: %s", err.Error())
		return false, errors.New("error in finding the GSLB leader's uuid")
	}
	if clusterUuid == gslbLeaderUuid {
		return true, nil
	}
	return false, nil
}

func GetClusterUuid(client *clients.AviClient) (string, error) {
	var clusterIntf interface{}

	uri := "/api/cluster"

	err := client.AviSession.Get(uri, &clusterIntf)
	if err != nil {
		gslbutils.Logf("object: ControllerCluster, msg: Cluster get URI %s returned error %s", uri, err.Error())
		return "", err
	}

	if clusterIntf == nil {
		gslbutils.Logf("object: ControllerCluster, msg: Cluster get URI %s returned %v type %T",
			uri, clusterIntf, clusterIntf)
		return "", errors.New("unexpected response for get cluster")
	}
	gslbutils.Debugf("object: ControllerCluster, msg: Cluster get URI %s returned a cluster", uri)

	cluster, ok := clusterIntf.(map[string]interface{})
	if !ok {
		gslbutils.Warnf("resp: %v, msg: response can't be parsed to map[string]interface", clusterIntf)
		return "", errors.New("response can't be parsed to map[string]interface")
	}
	name, ok := cluster["name"].(string)
	if !ok {
		gslbutils.Warnf("resp: %v, msg: name not present in response", clusterIntf)
		return "", errors.New("name not present in the cluster response")
	}
	clusterUUID, ok := cluster["uuid"].(string)
	if !ok {
		gslbutils.Warnf("resp: %v, msg: uuid not present in response", clusterIntf)
		return "", errors.New("uuid not present in the cluster response")
	}

	gslbutils.Logf("object: ControllerCluster, name: %s, uuid: %s, msg: fetched uuid for cluster", name, clusterUUID)
	return clusterUUID, nil
}

func GetGslbLeaderUuid(client *clients.AviClient) (string, error) {
	var resp interface{}

	uri := "/api/gslb"

	err := client.AviSession.Get(uri, &resp)
	if err != nil {
		gslbutils.Logf("object: GslbConfig, msg: gslb get URI %s returned error %s", uri, err.Error())
		return "", err
	}

	restResp, ok := resp.(map[string]interface{})
	if !ok {
		gslbutils.Logf("object: GslbConfig, msg: gslb get URI %s returned %v type %T",
			uri, resp, restResp)
		return "", errors.New("unexpected response for get gslb")
	}
	gslbutils.Debugf("object: GslbConfig, msg: gslb get URI %s returned %v count", uri, restResp["count"])
	results, ok := restResp["results"].([]interface{})

	if !ok {
		gslbutils.Logf("object: GslbConfig, msg: results not of type []interface{} instead of type %T",
			restResp["results"])
		return "", errors.New("results not of type []interface{}")
	}

	if len(results) == 0 {
		gslbutils.Logf("object: GslbConfig, msg: results length is zero, probably controller not a part of gslb config")
		return "", errors.New("no results for uri " + uri)
	}
	// results[0] contains the GSLB information
	gslbIntf := results[0]
	gslbConfig := gslbIntf.(map[string]interface{})
	leaderUUID, ok := gslbConfig["leader_cluster_uuid"].(string)
	if !ok {
		gslbutils.Warnf("resp: %v, msg: leader_cluster_uuid not present in response", gslbIntf)
		return "", errors.New("gslb_leader_uuid not present in gslb response")
	}

	gslbutils.Logf("object: GslbConfig, leader_cluster_uuid: %s, msg: fetched leader_cluster_uuid for gslb",
		leaderUUID)
	return leaderUUID, nil
}

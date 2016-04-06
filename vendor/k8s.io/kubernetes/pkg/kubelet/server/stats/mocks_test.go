/*
Copyright 2016 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package stats

import "github.com/stretchr/testify/mock"

import cadvisorapi "github.com/google/cadvisor/info/v1"
import cadvisorapiv2 "github.com/google/cadvisor/info/v2"
import "k8s.io/kubernetes/pkg/api"
import "k8s.io/kubernetes/pkg/kubelet/cm"

import "k8s.io/kubernetes/pkg/types"
import "k8s.io/kubernetes/pkg/volume"

// DO NOT EDIT
// GENERATED BY mockery

type MockStatsProvider struct {
	mock.Mock
}

// GetContainerInfo provides a mock function with given fields: podFullName, uid, containerName, req
func (_m *MockStatsProvider) GetContainerInfo(podFullName string, uid types.UID, containerName string, req *cadvisorapi.ContainerInfoRequest) (*cadvisorapi.ContainerInfo, error) {
	ret := _m.Called(podFullName, uid, containerName, req)

	var r0 *cadvisorapi.ContainerInfo
	if rf, ok := ret.Get(0).(func(string, types.UID, string, *cadvisorapi.ContainerInfoRequest) *cadvisorapi.ContainerInfo); ok {
		r0 = rf(podFullName, uid, containerName, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cadvisorapi.ContainerInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, types.UID, string, *cadvisorapi.ContainerInfoRequest) error); ok {
		r1 = rf(podFullName, uid, containerName, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetContainerInfoV2 provides a mock function with given fields: name, options
func (_m *MockStatsProvider) GetContainerInfoV2(name string, options cadvisorapiv2.RequestOptions) (map[string]cadvisorapiv2.ContainerInfo, error) {
	ret := _m.Called(name, options)

	var r0 map[string]cadvisorapiv2.ContainerInfo
	if rf, ok := ret.Get(0).(func(string, cadvisorapiv2.RequestOptions) map[string]cadvisorapiv2.ContainerInfo); ok {
		r0 = rf(name, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]cadvisorapiv2.ContainerInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, cadvisorapiv2.RequestOptions) error); ok {
		r1 = rf(name, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRawContainerInfo provides a mock function with given fields: containerName, req, subcontainers
func (_m *MockStatsProvider) GetRawContainerInfo(containerName string, req *cadvisorapi.ContainerInfoRequest, subcontainers bool) (map[string]*cadvisorapi.ContainerInfo, error) {
	ret := _m.Called(containerName, req, subcontainers)

	var r0 map[string]*cadvisorapi.ContainerInfo
	if rf, ok := ret.Get(0).(func(string, *cadvisorapi.ContainerInfoRequest, bool) map[string]*cadvisorapi.ContainerInfo); ok {
		r0 = rf(containerName, req, subcontainers)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*cadvisorapi.ContainerInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *cadvisorapi.ContainerInfoRequest, bool) error); ok {
		r1 = rf(containerName, req, subcontainers)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPodByName provides a mock function with given fields: namespace, name
func (_m *MockStatsProvider) GetPodByName(namespace string, name string) (*api.Pod, bool) {
	ret := _m.Called(namespace, name)

	var r0 *api.Pod
	if rf, ok := ret.Get(0).(func(string, string) *api.Pod); ok {
		r0 = rf(namespace, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.Pod)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string, string) bool); ok {
		r1 = rf(namespace, name)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetNode provides a mock function with given fields:
func (_m *MockStatsProvider) GetNode() (*api.Node, error) {
	ret := _m.Called()

	var r0 *api.Node
	if rf, ok := ret.Get(0).(func() *api.Node); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodeConfig provides a mock function with given fields:
func (_m *MockStatsProvider) GetNodeConfig() cm.NodeConfig {
	ret := _m.Called()

	var r0 cm.NodeConfig
	if rf, ok := ret.Get(0).(func() cm.NodeConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(cm.NodeConfig)
	}

	return r0
}

// DockerImagesFsInfo provides a mock function with given fields:
func (_m *MockStatsProvider) DockerImagesFsInfo() (cadvisorapiv2.FsInfo, error) {
	ret := _m.Called()

	var r0 cadvisorapiv2.FsInfo
	if rf, ok := ret.Get(0).(func() cadvisorapiv2.FsInfo); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(cadvisorapiv2.FsInfo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RootFsInfo provides a mock function with given fields:
func (_m *MockStatsProvider) RootFsInfo() (cadvisorapiv2.FsInfo, error) {
	ret := _m.Called()

	var r0 cadvisorapiv2.FsInfo
	if rf, ok := ret.Get(0).(func() cadvisorapiv2.FsInfo); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(cadvisorapiv2.FsInfo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListVolumesForPod provides a mock function with given fields: podUID
func (_m *MockStatsProvider) ListVolumesForPod(podUID types.UID) (map[string]volume.Volume, bool) {
	ret := _m.Called(podUID)

	var r0 map[string]volume.Volume
	if rf, ok := ret.Get(0).(func(types.UID) map[string]volume.Volume); ok {
		r0 = rf(podUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]volume.Volume)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(types.UID) bool); ok {
		r1 = rf(podUID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetPods provides a mock function with given fields:
func (_m *MockStatsProvider) GetPods() []*api.Pod {
	ret := _m.Called()

	var r0 []*api.Pod
	if rf, ok := ret.Get(0).(func() []*api.Pod); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*api.Pod)
		}
	}

	return r0
}

//Licensed to the Apache Software Foundation (ASF) under one or more
//contributor license agreements.  See the NOTICE file distributed with
//this work for additional information regarding copyright ownership.
//The ASF licenses this file to You under the Apache License, Version 2.0
//(the "License"); you may not use this file except in compliance with
//the License.  You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package authority

import (
	"github.com/apache/dubbo-kubernetes/api/ca"
	"github.com/apache/dubbo-kubernetes/pkg/authority/server"
	core_runtime "github.com/apache/dubbo-kubernetes/pkg/core/runtime"
	"github.com/pkg/errors"
)

func Setup(rt core_runtime.Runtime) error {
	if !rt.Config().KubeConfig.IsKubernetesConnected {
		return nil
	}
	server := server.NewServer(rt.Config())
	if rt.Config().KubeConfig.InPodEnv {
		server.CertClient = rt.CertStorage().GetCertClient()
		server.CertStorage = rt.CertStorage()
	}
	if err := RegisterCertificateService(rt, server); err != nil {
		return errors.Wrap(err, "CertificateService register failed")
	}

	if err := rt.Add(server); err != nil {
		return errors.Wrap(err, "Add Authority Component failed")
	}
	return nil
}

func RegisterCertificateService(rt core_runtime.Runtime, service *server.AuthorityService) error {
	ca.RegisterAuthorityServiceServer(rt.GrpcServer().PlainServer, service)
	ca.RegisterAuthorityServiceServer(rt.GrpcServer().SecureServer, service)
	return nil
}

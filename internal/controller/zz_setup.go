/*
Copyright 2021 The Crossplane Authors.

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

package controller

import (
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/crossplane-runtime/pkg/logging"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/crossplane/terrajet/pkg/terraform"

	ip "github.com/crossplane-contrib/provider-jet-hcloud/internal/controller/floating/ip"
	ipassignment "github.com/crossplane-contrib/provider-jet-hcloud/internal/controller/floating/ipassignment"
	certificate "github.com/crossplane-contrib/provider-jet-hcloud/internal/controller/hcloud/certificate"
	firewall "github.com/crossplane-contrib/provider-jet-hcloud/internal/controller/hcloud/firewall"
	network "github.com/crossplane-contrib/provider-jet-hcloud/internal/controller/hcloud/network"
	rdns "github.com/crossplane-contrib/provider-jet-hcloud/internal/controller/hcloud/rdns"
	server "github.com/crossplane-contrib/provider-jet-hcloud/internal/controller/hcloud/server"
	snapshot "github.com/crossplane-contrib/provider-jet-hcloud/internal/controller/hcloud/snapshot"
	volume "github.com/crossplane-contrib/provider-jet-hcloud/internal/controller/hcloud/volume"
	balancer "github.com/crossplane-contrib/provider-jet-hcloud/internal/controller/load/balancer"
	balancernetwork "github.com/crossplane-contrib/provider-jet-hcloud/internal/controller/load/balancernetwork"
	balancerservice "github.com/crossplane-contrib/provider-jet-hcloud/internal/controller/load/balancerservice"
	balancertarget "github.com/crossplane-contrib/provider-jet-hcloud/internal/controller/load/balancertarget"
	certificatemanaged "github.com/crossplane-contrib/provider-jet-hcloud/internal/controller/managed/certificate"
	route "github.com/crossplane-contrib/provider-jet-hcloud/internal/controller/network/route"
	subnet "github.com/crossplane-contrib/provider-jet-hcloud/internal/controller/network/subnet"
	group "github.com/crossplane-contrib/provider-jet-hcloud/internal/controller/placement/group"
	providerconfig "github.com/crossplane-contrib/provider-jet-hcloud/internal/controller/providerconfig"
	networkserver "github.com/crossplane-contrib/provider-jet-hcloud/internal/controller/server/network"
	key "github.com/crossplane-contrib/provider-jet-hcloud/internal/controller/ssh/key"
	certificateuploaded "github.com/crossplane-contrib/provider-jet-hcloud/internal/controller/uploaded/certificate"
	attachment "github.com/crossplane-contrib/provider-jet-hcloud/internal/controller/volume/attachment"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter, ps terraform.SetupFn, ws *terraform.WorkspaceStore, cfg *tjconfig.Provider, concurrency int) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter, terraform.SetupFn, *terraform.WorkspaceStore, *tjconfig.Provider, int) error{
		ip.Setup,
		ipassignment.Setup,
		certificate.Setup,
		firewall.Setup,
		network.Setup,
		rdns.Setup,
		server.Setup,
		snapshot.Setup,
		volume.Setup,
		balancer.Setup,
		balancernetwork.Setup,
		balancerservice.Setup,
		balancertarget.Setup,
		certificatemanaged.Setup,
		route.Setup,
		subnet.Setup,
		group.Setup,
		providerconfig.Setup,
		networkserver.Setup,
		key.Setup,
		certificateuploaded.Setup,
		attachment.Setup,
	} {
		if err := setup(mgr, l, wl, ps, ws, cfg, concurrency); err != nil {
			return err
		}
	}
	return nil
}

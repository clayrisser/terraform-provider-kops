# kops_cluster

| attribute | type | optional/required | computed |
| --- | --- | --- | --- |
| `name` | String | Required |  |
| `channel` | String | Optional |  |
| `addons` | List([AddonSpec](./AddonSpec.generated.md)) | Optional |  |
| `config_base` | String | Optional | Computed |
| `cloud_provider` | String | Required |  |
| `container_runtime` | String | Optional |  |
| `kubernetes_version` | String | Optional |  |
| `subnet` | List([ClusterSubnetSpec](./ClusterSubnetSpec.generated.md)) | Required |  |
| `project` | String | Optional |  |
| `master_public_name` | String | Optional | Computed |
| `master_internal_name` | String | Optional | Computed |
| `network_cidr` | String | Optional | Computed |
| `additional_network_cidrs` | List(String) | Optional |  |
| `network_id` | String | Required |  |
| `topology` | [TopologySpec](./TopologySpec.generated.md) | Required |  |
| `secret_store` | String | Optional |  |
| `key_store` | String | Optional |  |
| `config_store` | String | Optional |  |
| `dns_zone` | String | Optional |  |
| `additional_sans` | List(String) | Optional |  |
| `cluster_dns_domain` | String | Optional |  |
| `service_cluster_ip_range` | String | Optional |  |
| `pod_cidr` | String | Optional |  |
| `non_masquerade_cidr` | String | Optional | Computed |
| `ssh_access` | List(String) | Optional |  |
| `node_port_access` | List(String) | Optional |  |
| `egress_proxy` | [EgressProxySpec](./EgressProxySpec.generated.md) | Optional |  |
| `ssh_key_name` | String | Optional |  |
| `kubernetes_api_access` | List(String) | Optional |  |
| `isolate_masters` | Bool | Optional |  |
| `update_policy` | String | Optional |  |
| `external_policies` | Map(List(String)) | Optional |  |
| `additional_policies` | Map(String) | Optional |  |
| `file_assets` | List([FileAssetSpec](./FileAssetSpec.generated.md)) | Optional |  |
| `etcd_cluster` | List([EtcdClusterSpec](./EtcdClusterSpec.generated.md)) | Required |  |
| `containerd` | [ContainerdConfig](./ContainerdConfig.generated.md) | Optional |  |
| `docker` | [DockerConfig](./DockerConfig.generated.md) | Optional |  |
| `kube_dns` | [KubeDNSConfig](./KubeDNSConfig.generated.md) | Optional |  |
| `kube_api_server` | [KubeAPIServerConfig](./KubeAPIServerConfig.generated.md) | Optional |  |
| `kube_controller_manager` | [KubeControllerManagerConfig](./KubeControllerManagerConfig.generated.md) | Optional |  |
| `external_cloud_controller_manager` | [CloudControllerManagerConfig](./CloudControllerManagerConfig.generated.md) | Optional |  |
| `kube_scheduler` | [KubeSchedulerConfig](./KubeSchedulerConfig.generated.md) | Optional |  |
| `kube_proxy` | [KubeProxyConfig](./KubeProxyConfig.generated.md) | Optional |  |
| `kubelet` | [KubeletConfigSpec](./KubeletConfigSpec.generated.md) | Optional |  |
| `master_kubelet` | [KubeletConfigSpec](./KubeletConfigSpec.generated.md) | Optional |  |
| `cloud_config` | [CloudConfiguration](./CloudConfiguration.generated.md) | Optional |  |
| `external_dns` | [ExternalDNSConfig](./ExternalDNSConfig.generated.md) | Optional |  |
| `networking` | [NetworkingSpec](./NetworkingSpec.generated.md) | Required |  |
| `api` | [AccessSpec](./AccessSpec.generated.md) | Optional |  |
| `authentication` | [AuthenticationSpec](./AuthenticationSpec.generated.md) | Optional |  |
| `authorization` | [AuthorizationSpec](./AuthorizationSpec.generated.md) | Optional |  |
| `node_authorization` | [NodeAuthorizationSpec](./NodeAuthorizationSpec.generated.md) | Optional |  |
| `cloud_labels` | Map(String) | Optional |  |
| `hooks` | List([HookSpec](./HookSpec.generated.md)) | Optional |  |
| `assets` | [Assets](./Assets.generated.md) | Optional |  |
| `iam` | [IAMSpec](./IAMSpec.generated.md) | Optional | Computed |
| `encryption_config` | Bool | Optional |  |
| `disable_subnet_tags` | Bool | Optional |  |
| `use_host_certificates` | Bool | Optional |  |
| `sysctl_parameters` | List(String) | Optional |  |
| `rolling_update` | [RollingUpdate](./RollingUpdate.generated.md) | Optional |  |
| `kube_server` | String | Optional | Computed |
| `kube_certificate_authority` | String | Optional | Computed |
| `kube_client_certificate` | String | Optional | Computed |
| `kube_client_key` | String | Optional | Computed |
| `kube_username` | String | Optional | Computed |
| `kube_password` | String | Optional | Computed |
| `instance_group` | List([InstanceGroup](./InstanceGroup.generated.md)) | Required |  |
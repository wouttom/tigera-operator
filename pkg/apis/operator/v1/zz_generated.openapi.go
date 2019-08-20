// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/tigera/operator/pkg/apis/operator/v1.APIServer":                     schema_pkg_apis_operator_v1_APIServer(ref),
		"github.com/tigera/operator/pkg/apis/operator/v1.APIServerSpec":                 schema_pkg_apis_operator_v1_APIServerSpec(ref),
		"github.com/tigera/operator/pkg/apis/operator/v1.APIServerStatus":               schema_pkg_apis_operator_v1_APIServerStatus(ref),
		"github.com/tigera/operator/pkg/apis/operator/v1.Auth":                          schema_pkg_apis_operator_v1_Auth(ref),
		"github.com/tigera/operator/pkg/apis/operator/v1.CNISpec":                       schema_pkg_apis_operator_v1_CNISpec(ref),
		"github.com/tigera/operator/pkg/apis/operator/v1.ComponentsSpec":                schema_pkg_apis_operator_v1_ComponentsSpec(ref),
		"github.com/tigera/operator/pkg/apis/operator/v1.Console":                       schema_pkg_apis_operator_v1_Console(ref),
		"github.com/tigera/operator/pkg/apis/operator/v1.ConsoleSpec":                   schema_pkg_apis_operator_v1_ConsoleSpec(ref),
		"github.com/tigera/operator/pkg/apis/operator/v1.ConsoleStatus":                 schema_pkg_apis_operator_v1_ConsoleStatus(ref),
		"github.com/tigera/operator/pkg/apis/operator/v1.Installation":                  schema_pkg_apis_operator_v1_Installation(ref),
		"github.com/tigera/operator/pkg/apis/operator/v1.InstallationSpec":              schema_pkg_apis_operator_v1_InstallationSpec(ref),
		"github.com/tigera/operator/pkg/apis/operator/v1.InstallationStatus":            schema_pkg_apis_operator_v1_InstallationStatus(ref),
		"github.com/tigera/operator/pkg/apis/operator/v1.KubeControllersSpec":           schema_pkg_apis_operator_v1_KubeControllersSpec(ref),
		"github.com/tigera/operator/pkg/apis/operator/v1.MonitoringConfiguration":       schema_pkg_apis_operator_v1_MonitoringConfiguration(ref),
		"github.com/tigera/operator/pkg/apis/operator/v1.MonitoringConfigurationSpec":   schema_pkg_apis_operator_v1_MonitoringConfigurationSpec(ref),
		"github.com/tigera/operator/pkg/apis/operator/v1.MonitoringConfigurationStatus": schema_pkg_apis_operator_v1_MonitoringConfigurationStatus(ref),
		"github.com/tigera/operator/pkg/apis/operator/v1.NodeSpec":                      schema_pkg_apis_operator_v1_NodeSpec(ref),
	}
}

func schema_pkg_apis_operator_v1_APIServer(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "APIServer is the Schema for the apiservers API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/tigera/operator/pkg/apis/operator/v1.APIServerSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/tigera/operator/pkg/apis/operator/v1.APIServerStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/tigera/operator/pkg/apis/operator/v1.APIServerSpec", "github.com/tigera/operator/pkg/apis/operator/v1.APIServerStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_operator_v1_APIServerSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "APIServerSpec defines the desired state of APIServer",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_operator_v1_APIServerStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "APIServerStatus defines the observed state of APIServer",
				Properties: map[string]spec.Schema{
					"state": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_operator_v1_Auth(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Auth defines authentication configuration.",
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type configures the type of authentication used by the manager. Default: \"Basic\"",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"authority": {
						SchemaProps: spec.SchemaProps{
							Description: "Authority configures the OAuth2/OIDC authority/issuer when using OAuth2 or OIDC login. Default: \"\"https://accounts.google.com\"",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"clientID": {
						SchemaProps: spec.SchemaProps{
							Description: "ClientId configures the OAuth2/OIDC client ID to use for OAuth2 or OIDC login.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_operator_v1_CNISpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CNISpec defines optional configuration for the CNI component.",
				Properties: map[string]spec.Schema{
					"extraEnv": {
						SchemaProps: spec.SchemaProps{
							Description: "ExtraEnv adds extra environment variables to the CNI container.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.EnvVar"),
									},
								},
							},
						},
					},
					"extraVolumes": {
						SchemaProps: spec.SchemaProps{
							Description: "ExtraVolumes configures custom volumes to be used by the CNI container.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.Volume"),
									},
								},
							},
						},
					},
					"extraVolumeMounts": {
						SchemaProps: spec.SchemaProps{
							Description: "ExtraVolumeMounts configures custom volume mounts to be used by the CNI container.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.VolumeMount"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.EnvVar", "k8s.io/api/core/v1.Volume", "k8s.io/api/core/v1.VolumeMount"},
	}
}

func schema_pkg_apis_operator_v1_ComponentsSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ComponentsSpec defines the desired state of components.",
				Properties: map[string]spec.Schema{
					"node": {
						SchemaProps: spec.SchemaProps{
							Description: "Node is optional configuration for the node component.",
							Ref:         ref("github.com/tigera/operator/pkg/apis/operator/v1.NodeSpec"),
						},
					},
					"cni": {
						SchemaProps: spec.SchemaProps{
							Description: "CNI is optional configuration for the CNI component.",
							Ref:         ref("github.com/tigera/operator/pkg/apis/operator/v1.CNISpec"),
						},
					},
					"kubeControllers": {
						SchemaProps: spec.SchemaProps{
							Description: "KubeControllers is optional configuration for the kube-controllers component.",
							Ref:         ref("github.com/tigera/operator/pkg/apis/operator/v1.KubeControllersSpec"),
						},
					},
					"apiServer": {
						SchemaProps: spec.SchemaProps{
							Description: "APIServer is optional configuration for the API server component.",
							Ref:         ref("github.com/tigera/operator/pkg/apis/operator/v1.APIServerSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/tigera/operator/pkg/apis/operator/v1.APIServerSpec", "github.com/tigera/operator/pkg/apis/operator/v1.CNISpec", "github.com/tigera/operator/pkg/apis/operator/v1.KubeControllersSpec", "github.com/tigera/operator/pkg/apis/operator/v1.NodeSpec"},
	}
}

func schema_pkg_apis_operator_v1_Console(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Console is the Schema for the consoles API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/tigera/operator/pkg/apis/operator/v1.ConsoleSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/tigera/operator/pkg/apis/operator/v1.ConsoleStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/tigera/operator/pkg/apis/operator/v1.ConsoleSpec", "github.com/tigera/operator/pkg/apis/operator/v1.ConsoleStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_operator_v1_ConsoleSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ConsoleSpec defines optional configuration for the Tigera Secure management console. Valid only for the variant 'TigeraSecureEnterprise'.",
				Properties: map[string]spec.Schema{
					"auth": {
						SchemaProps: spec.SchemaProps{
							Description: "Auth is optional authentication configuration for the Tigera Secure management console.",
							Ref:         ref("github.com/tigera/operator/pkg/apis/operator/v1.Auth"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/tigera/operator/pkg/apis/operator/v1.Auth"},
	}
}

func schema_pkg_apis_operator_v1_ConsoleStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ConsoleStatus defines the observed state of Console",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_operator_v1_Installation(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Installation is the Schema for the cores API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/tigera/operator/pkg/apis/operator/v1.InstallationSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/tigera/operator/pkg/apis/operator/v1.InstallationStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/tigera/operator/pkg/apis/operator/v1.InstallationSpec", "github.com/tigera/operator/pkg/apis/operator/v1.InstallationStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_operator_v1_InstallationSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "InstallationSpec defines the desired state of Installation.",
				Properties: map[string]spec.Schema{
					"variant": {
						SchemaProps: spec.SchemaProps{
							Description: "Variant is the product to install - one of Calico or TigeraSecureEnterprise Default: Calico",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"registry": {
						SchemaProps: spec.SchemaProps{
							Description: "Registry is the default Docker registry used for component Docker images. Default: docker.io/",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"imagePullSecrets": {
						SchemaProps: spec.SchemaProps{
							Description: "ImagePullSecrets is an array of references to Docker registry pull secrets.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.LocalObjectReference"),
									},
								},
							},
						},
					},
					"ipPools": {
						SchemaProps: spec.SchemaProps{
							Description: "IPPools contains a list of IP pools to use for allocating pod IP addresses. For now, a maximum of one IP pool is supported. Default: 192.168.0.0/16.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/tigera/operator/pkg/apis/operator/v1.IPPool"),
									},
								},
							},
						},
					},
					"cniNetDir": {
						SchemaProps: spec.SchemaProps{
							Description: "CNINetDir configures the path on the host where CNI network configuration files will be installed. Default: /etc/cni/net.d",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"cniBinDir": {
						SchemaProps: spec.SchemaProps{
							Description: "CNIBinDir configures the path on the host where CNI binaries will be installed. Default: /opt/cni/bin",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"components": {
						SchemaProps: spec.SchemaProps{
							Description: "Components specifies the configuration of components.",
							Ref:         ref("github.com/tigera/operator/pkg/apis/operator/v1.ComponentsSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/tigera/operator/pkg/apis/operator/v1.ComponentsSpec", "github.com/tigera/operator/pkg/apis/operator/v1.IPPool", "k8s.io/api/core/v1.LocalObjectReference"},
	}
}

func schema_pkg_apis_operator_v1_InstallationStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "InstallationStatus defines the observed state of Installation",
				Properties: map[string]spec.Schema{
					"variant": {
						SchemaProps: spec.SchemaProps{
							Description: "Variant is the installed product - one of Calico or TigeraSecureEnterprise",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_operator_v1_KubeControllersSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KubeControllersSpec defines optional configuration for the kube-controllers component.",
				Properties: map[string]spec.Schema{
					"extraEnv": {
						SchemaProps: spec.SchemaProps{
							Description: "ExtraEnv adds extra environment variables to the kube-controllers container.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.EnvVar"),
									},
								},
							},
						},
					},
					"extraVolumes": {
						SchemaProps: spec.SchemaProps{
							Description: "ExtraVolumes configures custom volumes to be used by the kube-controllers container.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.Volume"),
									},
								},
							},
						},
					},
					"extraVolumeMounts": {
						SchemaProps: spec.SchemaProps{
							Description: "ExtraVolumeMounts configures custom volume mounts to be used by the kube-controllers container.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.VolumeMount"),
									},
								},
							},
						},
					},
					"tolerations": {
						SchemaProps: spec.SchemaProps{
							Description: "Tolerations configures custom tolerations on the kube-controllers deployment.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.Toleration"),
									},
								},
							},
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Description: "Resources configures custom resource requirements on the kube-controllers container.",
							Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.EnvVar", "k8s.io/api/core/v1.ResourceRequirements", "k8s.io/api/core/v1.Toleration", "k8s.io/api/core/v1.Volume", "k8s.io/api/core/v1.VolumeMount"},
	}
}

func schema_pkg_apis_operator_v1_MonitoringConfiguration(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MonitoringConfiguration is the Schema for the monitoringconfigurations API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/tigera/operator/pkg/apis/operator/v1.MonitoringConfigurationSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/tigera/operator/pkg/apis/operator/v1.MonitoringConfigurationStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/tigera/operator/pkg/apis/operator/v1.MonitoringConfigurationSpec", "github.com/tigera/operator/pkg/apis/operator/v1.MonitoringConfigurationStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_operator_v1_MonitoringConfigurationSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MonitoringConfigurationSpec defines the desired state of MonitoringConfiguration",
				Properties: map[string]spec.Schema{
					"clusterName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"elasticsearch": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/tigera/operator/pkg/apis/operator/v1.ElasticConfig"),
						},
					},
				},
				Required: []string{"clusterName", "elasticsearch"},
			},
		},
		Dependencies: []string{
			"github.com/tigera/operator/pkg/apis/operator/v1.ElasticConfig"},
	}
}

func schema_pkg_apis_operator_v1_MonitoringConfigurationStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MonitoringConfigurationStatus defines the observed state of MonitoringConfiguration",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_operator_v1_NodeSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NodeSpec defines optional configuration for the node component.",
				Properties: map[string]spec.Schema{
					"maxUnavailable": {
						SchemaProps: spec.SchemaProps{
							Description: "MaxUnavailable configures the maximum number of pods that can be unavailable during a rolling update of the node daemonset. Default: 1",
							Ref:         ref("k8s.io/apimachinery/pkg/util/intstr.IntOrString"),
						},
					},
					"extraEnv": {
						SchemaProps: spec.SchemaProps{
							Description: "ExtraEnv adds extra environment variables to the node container.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.EnvVar"),
									},
								},
							},
						},
					},
					"extraVolumes": {
						SchemaProps: spec.SchemaProps{
							Description: "ExtraVolumes configures custom volumes to be used by the node daemonset.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.Volume"),
									},
								},
							},
						},
					},
					"extraVolumeMounts": {
						SchemaProps: spec.SchemaProps{
							Description: "ExtraVolumeMounts configures custom volume mounts to be used by the node container.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.VolumeMount"),
									},
								},
							},
						},
					},
					"tolerations": {
						SchemaProps: spec.SchemaProps{
							Description: "Tolerations configures custom tolerations on the node daemonset.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.Toleration"),
									},
								},
							},
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Description: "Resources configures custom resource requirements on the node container.",
							Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.EnvVar", "k8s.io/api/core/v1.ResourceRequirements", "k8s.io/api/core/v1.Toleration", "k8s.io/api/core/v1.Volume", "k8s.io/api/core/v1.VolumeMount", "k8s.io/apimachinery/pkg/util/intstr.IntOrString"},
	}
}

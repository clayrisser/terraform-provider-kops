package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceTopologySpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bastion": ComputedStruct(DataSourceBastionSpec()),
			"dns":     ComputedString(),
		},
	}

	return res
}

func ExpandDataSourceTopologySpec(in map[string]interface{}) kops.TopologySpec {
	if in == nil {
		panic("expand TopologySpec failure, in is nil")
	}
	return kops.TopologySpec{
		Bastion: func(in interface{}) *kops.BastionSpec {
			return func(in interface{}) *kops.BastionSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.BastionSpec) *kops.BastionSpec {
					return &in
				}(func(in interface{}) kops.BastionSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceBastionSpec(in[0].(map[string]interface{}))
					}
					return kops.BastionSpec{}
				}(in))
			}(in)
		}(in["bastion"]),
		DNS: func(in interface{}) kops.DNSType {
			return kops.DNSType(ExpandString(in))
		}(in["dns"]),
	}
}

func FlattenDataSourceTopologySpecInto(in kops.TopologySpec, out map[string]interface{}) {
	out["bastion"] = func(in *kops.BastionSpec) interface{} {
		return func(in *kops.BastionSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.BastionSpec) interface{} {
				return func(in kops.BastionSpec) []interface{} {
					return []interface{}{FlattenDataSourceBastionSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Bastion)
	out["dns"] = func(in kops.DNSType) interface{} {
		return FlattenString(string(in))
	}(in.DNS)
}

func FlattenDataSourceTopologySpec(in kops.TopologySpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceTopologySpecInto(in, out)
	return out
}

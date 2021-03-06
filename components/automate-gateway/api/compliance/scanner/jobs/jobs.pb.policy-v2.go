// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: components/automate-gateway/api/compliance/scanner/jobs/jobs.proto

package jobs

import policyv2 "github.com/chef/automate/components/automate-gateway/authz/policy_v2"

func init() {
	policyv2.MapMethodTo("/chef.automate.api.compliance.scanner.jobs.v1.JobsService/Create", "compliance:scanner:jobs", "compliance:scannerJobs:create", "POST", "/compliance/scanner/jobs", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Job); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "type":
					return m.Type
				case "status":
					return m.Status
				case "recurrence":
					return m.Recurrence
				case "parent_id":
					return m.ParentId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.compliance.scanner.jobs.v1.JobsService/Read", "compliance:scanner:jobs:{id}", "compliance:scannerJobs:get", "GET", "/compliance/scanner/jobs/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Id); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.compliance.scanner.jobs.v1.JobsService/Update", "compliance:scanner:jobs:{id}", "compliance:scannerJobs:update", "PUT", "/compliance/scanner/jobs/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Job); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "type":
					return m.Type
				case "status":
					return m.Status
				case "recurrence":
					return m.Recurrence
				case "parent_id":
					return m.ParentId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.compliance.scanner.jobs.v1.JobsService/Delete", "compliance:scanner:jobs:{id}", "compliance:scannerJobs:delete", "DELETE", "/compliance/scanner/jobs/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Id); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.compliance.scanner.jobs.v1.JobsService/List", "compliance:scanner:jobs", "compliance:scannerJobs:list", "POST", "/compliance/scanner/jobs/search", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.compliance.scanner.jobs.v1.JobsService/Rerun", "compliance:scanner:jobs:{id}", "compliance:scannerJobs:rerun", "GET", "/compliance/scanner/jobs/rerun/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Id); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
}

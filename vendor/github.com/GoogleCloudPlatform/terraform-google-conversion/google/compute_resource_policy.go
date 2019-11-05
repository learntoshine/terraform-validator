// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func GetComputeResourcePolicyCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/resourcePolicies/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetComputeResourcePolicyApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "compute.googleapis.com/ResourcePolicy",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "ResourcePolicy",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetComputeResourcePolicyApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeResourcePolicyName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	snapshotSchedulePolicyProp, err := expandComputeResourcePolicySnapshotSchedulePolicy(d.Get("snapshot_schedule_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("snapshot_schedule_policy"); !isEmptyValue(reflect.ValueOf(snapshotSchedulePolicyProp)) && (ok || !reflect.DeepEqual(v, snapshotSchedulePolicyProp)) {
		obj["snapshotSchedulePolicy"] = snapshotSchedulePolicyProp
	}
	regionProp, err := expandComputeResourcePolicyRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	return obj, nil
}

func expandComputeResourcePolicyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSchedule, err := expandComputeResourcePolicySnapshotSchedulePolicySchedule(original["schedule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSchedule); val.IsValid() && !isEmptyValue(val) {
		transformed["schedule"] = transformedSchedule
	}

	transformedRetentionPolicy, err := expandComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy(original["retention_policy"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRetentionPolicy); val.IsValid() && !isEmptyValue(val) {
		transformed["retentionPolicy"] = transformedRetentionPolicy
	}

	transformedSnapshotProperties, err := expandComputeResourcePolicySnapshotSchedulePolicySnapshotProperties(original["snapshot_properties"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSnapshotProperties); val.IsValid() && !isEmptyValue(val) {
		transformed["snapshotProperties"] = transformedSnapshotProperties
	}

	return transformed, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicySchedule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHourlySchedule, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule(original["hourly_schedule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHourlySchedule); val.IsValid() && !isEmptyValue(val) {
		transformed["hourlySchedule"] = transformedHourlySchedule
	}

	transformedDailySchedule, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleDailySchedule(original["daily_schedule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDailySchedule); val.IsValid() && !isEmptyValue(val) {
		transformed["dailySchedule"] = transformedDailySchedule
	}

	transformedWeeklySchedule, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule(original["weekly_schedule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWeeklySchedule); val.IsValid() && !isEmptyValue(val) {
		transformed["weeklySchedule"] = transformedWeeklySchedule
	}

	return transformed, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHoursInCycle, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleHoursInCycle(original["hours_in_cycle"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHoursInCycle); val.IsValid() && !isEmptyValue(val) {
		transformed["hoursInCycle"] = transformedHoursInCycle
	}

	transformedStartTime, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleStartTime(original["start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !isEmptyValue(val) {
		transformed["startTime"] = transformedStartTime
	}

	return transformed, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleHoursInCycle(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleStartTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleDailySchedule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDaysInCycle, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleDaysInCycle(original["days_in_cycle"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDaysInCycle); val.IsValid() && !isEmptyValue(val) {
		transformed["daysInCycle"] = transformedDaysInCycle
	}

	transformedStartTime, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleStartTime(original["start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !isEmptyValue(val) {
		transformed["startTime"] = transformedStartTime
	}

	return transformed, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleDaysInCycle(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleStartTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDayOfWeeks, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeks(original["day_of_weeks"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDayOfWeeks); val.IsValid() && !isEmptyValue(val) {
		transformed["dayOfWeeks"] = transformedDayOfWeeks
	}

	return transformed, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeks(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedStartTime, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeksStartTime(original["start_time"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !isEmptyValue(val) {
			transformed["startTime"] = transformedStartTime
		}

		transformedDay, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeksDay(original["day"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDay); val.IsValid() && !isEmptyValue(val) {
			transformed["day"] = transformedDay
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeksStartTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeksDay(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMaxRetentionDays, err := expandComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyMaxRetentionDays(original["max_retention_days"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxRetentionDays); val.IsValid() && !isEmptyValue(val) {
		transformed["maxRetentionDays"] = transformedMaxRetentionDays
	}

	transformedOnSourceDiskDelete, err := expandComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOnSourceDiskDelete(original["on_source_disk_delete"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOnSourceDiskDelete); val.IsValid() && !isEmptyValue(val) {
		transformed["onSourceDiskDelete"] = transformedOnSourceDiskDelete
	}

	return transformed, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyMaxRetentionDays(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOnSourceDiskDelete(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicySnapshotProperties(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLabels, err := expandComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesLabels(original["labels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !isEmptyValue(val) {
		transformed["labels"] = transformedLabels
	}

	transformedStorageLocations, err := expandComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesStorageLocations(original["storage_locations"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStorageLocations); val.IsValid() && !isEmptyValue(val) {
		transformed["storageLocations"] = transformedStorageLocations
	}

	transformedGuestFlush, err := expandComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesGuestFlush(original["guest_flush"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGuestFlush); val.IsValid() && !isEmptyValue(val) {
		transformed["guestFlush"] = transformedGuestFlush
	}

	return transformed, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesStorageLocations(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesGuestFlush(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicyRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}

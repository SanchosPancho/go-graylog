package terraform

import (
	"testing"

	"github.com/SanchosPancho/go-graylog/v11/graylog/testdata"
)

func TestAccLDAPSetting(t *testing.T) {
	setEnv()

	tc := &testCase{
		t:       t,
		Name:    "ldap setting",
		GetPath: "/api/system/ldap/settings",

		CreateReqBodyMap: map[string]interface{}{},
		UpdateReqBodyMap: testdata.CreateLDAPSettingMap(),
		CreatedDataPath:  "ldap_setting/create.json",
		UpdatedDataPath:  "ldap_setting/create.json",
		CreateTFPath:     "ldap_setting/create.tf",
		UpdateTFPath:     "ldap_setting/create.tf",
	}
	tc.Test()
}

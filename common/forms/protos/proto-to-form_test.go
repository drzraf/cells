package protos

import (
	"encoding/xml"
	"os"
	"sort"
	"testing"

	json "github.com/pydio/cells/x/jsonx"

	"github.com/ory/ladon"

	servicecontext "github.com/pydio/cells/common/service/context"

	"github.com/pydio/cells/common/proto/jobs"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/pydio/cells/common/forms"
	"github.com/pydio/cells/common/proto/idm"
	"github.com/pydio/cells/common/proto/tree"
)

func TestGenerateProtoToForm(t *testing.T) {
	Convey("Test role", t, func() {
		f := GenerateProtoToForm("userSingleQuery", &idm.UserSingleQuery{}, false)
		So(f.Groups[0].Fields, ShouldNotBeEmpty)
		f = GenerateProtoToForm("roleSingleQuery", &idm.RoleSingleQuery{}, false)
		So(f.Groups[0].Fields, ShouldNotBeEmpty)
		f2 := GenerateProtoToForm("roleSingleQuery", &idm.RoleSingleQuery{}, true)
		So(f2.Groups[0].Fields, ShouldNotBeEmpty)
		serial2 := f2.Serialize()
		b2, _ := xml.Marshal(serial2)
		So(string(b2), ShouldEqual, `<form><param name="fieldname" type="group_switch:fieldname" label="Field Name" group=""></param><param group_switch_name="fieldname" group_switch_value="Uuid" group_switch_label="roleSingleQuery.Uuid" name="@value" type="hidden" group="" default="Uuid"></param><param group_switch_name="fieldname" group_switch_value="Uuid" group_switch_label="roleSingleQuery.Uuid" name="Uuid" type="string" label="roleSingleQuery.Uuid" group="" replicationGroup="Uuid" replicationTitle="Uuid[]" replicationMandatory="true"></param><param group_switch_name="fieldname" group_switch_value="Label" group_switch_label="roleSingleQuery.Label" name="@value" type="hidden" group="" default="Label"></param><param group_switch_name="fieldname" group_switch_value="Label" group_switch_label="roleSingleQuery.Label" name="Label" type="string" label="roleSingleQuery.Label" group=""></param><param group_switch_name="fieldname" group_switch_value="IsTeam" group_switch_label="roleSingleQuery.IsTeam" name="@value" type="hidden" group="" default="IsTeam"></param><param group_switch_name="fieldname" group_switch_value="IsTeam" group_switch_label="roleSingleQuery.IsTeam" name="IsTeam" type="boolean" label="roleSingleQuery.IsTeam" group=""></param><param group_switch_name="fieldname" group_switch_value="IsGroupRole" group_switch_label="roleSingleQuery.IsGroupRole" name="@value" type="hidden" group="" default="IsGroupRole"></param><param group_switch_name="fieldname" group_switch_value="IsGroupRole" group_switch_label="roleSingleQuery.IsGroupRole" name="IsGroupRole" type="boolean" label="roleSingleQuery.IsGroupRole" group=""></param><param group_switch_name="fieldname" group_switch_value="IsUserRole" group_switch_label="roleSingleQuery.IsUserRole" name="@value" type="hidden" group="" default="IsUserRole"></param><param group_switch_name="fieldname" group_switch_value="IsUserRole" group_switch_label="roleSingleQuery.IsUserRole" name="IsUserRole" type="boolean" label="roleSingleQuery.IsUserRole" group=""></param><param group_switch_name="fieldname" group_switch_value="HasAutoApply" group_switch_label="roleSingleQuery.HasAutoApply" name="@value" type="hidden" group="" default="HasAutoApply"></param><param group_switch_name="fieldname" group_switch_value="HasAutoApply" group_switch_label="roleSingleQuery.HasAutoApply" name="HasAutoApply" type="boolean" label="roleSingleQuery.HasAutoApply" group=""></param><param group_switch_name="fieldname" group_switch_value="not" group_switch_label="roleSingleQuery.Not" name="@value" type="hidden" group="" default="not"></param><param group_switch_name="fieldname" group_switch_value="not" group_switch_label="roleSingleQuery.Not" name="not" type="boolean" label="roleSingleQuery.not" group=""></param></form>`)

		f = GenerateProtoToForm("aclSingleQuery", &idm.ACLSingleQuery{}, true)
		sw := f.Groups[0].Fields[0].(*forms.SwitchField)
		a := GenerateProtoToForm("aclAction", &idm.ACLAction{}, false)
		sw.Values = append(sw.Values, &forms.SwitchValue{
			Name:  "Actions",
			Value: "Actions",
			Label: "Actions",
			Fields: []forms.Field{&forms.ReplicableFields{
				Id:          "Actions",
				Title:       "Actions",
				Description: "Acl Actions",
				Mandatory:   true,
				Fields:      a.Groups[0].Fields,
			}},
		})
		serial := f.Serialize()
		b, _ := xml.Marshal(serial)
		So(string(b), ShouldEqual, `<form><param name="fieldname" type="group_switch:fieldname" label="Field Name" group=""></param><param group_switch_name="fieldname" group_switch_value="RoleIDs" group_switch_label="aclSingleQuery.RoleIDs" name="@value" type="hidden" group="" default="RoleIDs"></param><param group_switch_name="fieldname" group_switch_value="RoleIDs" group_switch_label="aclSingleQuery.RoleIDs" name="RoleIDs" type="string" label="aclSingleQuery.RoleIDs" group="" replicationGroup="RoleIDs" replicationTitle="RoleIDs[]" replicationMandatory="true"></param><param group_switch_name="fieldname" group_switch_value="WorkspaceIDs" group_switch_label="aclSingleQuery.WorkspaceIDs" name="@value" type="hidden" group="" default="WorkspaceIDs"></param><param group_switch_name="fieldname" group_switch_value="WorkspaceIDs" group_switch_label="aclSingleQuery.WorkspaceIDs" name="WorkspaceIDs" type="string" label="aclSingleQuery.WorkspaceIDs" group="" replicationGroup="WorkspaceIDs" replicationTitle="WorkspaceIDs[]" replicationMandatory="true"></param><param group_switch_name="fieldname" group_switch_value="NodeIDs" group_switch_label="aclSingleQuery.NodeIDs" name="@value" type="hidden" group="" default="NodeIDs"></param><param group_switch_name="fieldname" group_switch_value="NodeIDs" group_switch_label="aclSingleQuery.NodeIDs" name="NodeIDs" type="string" label="aclSingleQuery.NodeIDs" group="" replicationGroup="NodeIDs" replicationTitle="NodeIDs[]" replicationMandatory="true"></param><param group_switch_name="fieldname" group_switch_value="not" group_switch_label="aclSingleQuery.Not" name="@value" type="hidden" group="" default="not"></param><param group_switch_name="fieldname" group_switch_value="not" group_switch_label="aclSingleQuery.Not" name="not" type="boolean" label="aclSingleQuery.not" group=""></param><param group_switch_name="fieldname" group_switch_value="Actions" group_switch_label="Actions" name="@value" type="hidden" group="" default="Actions"></param><param group_switch_name="fieldname" group_switch_value="Actions" group_switch_label="Actions" name="Name" type="string" label="aclAction.Name" group="" replicationGroup="Actions" replicationTitle="Actions" replicationDescription="Acl Actions" replicationMandatory="true"></param><param group_switch_name="fieldname" group_switch_value="Actions" group_switch_label="Actions" name="Value" type="string" label="aclAction.Value" group="" replicationGroup="Actions" replicationTitle="Actions" replicationDescription="Acl Actions" replicationMandatory="true"></param></form>`)

		So(f.Groups[0].Fields, ShouldNotBeEmpty)
		f = GenerateProtoToForm("workspaceSingleQuery", &idm.WorkspaceSingleQuery{}, false)
		So(f.Groups[0].Fields, ShouldNotBeEmpty)
		f = GenerateProtoToForm("treeQuery", &tree.Query{}, false)
		So(f.Groups[0].Fields, ShouldNotBeEmpty)
	})
}

func SkipTestGenerateJsonLanguagesFiles(t *testing.T) {
	Convey("Generate i18n Keys", t, func() {
		keys := make(map[string]string)
		GenerateProtoToForm("userSingleQuery", &idm.UserSingleQuery{}, false, keys)
		GenerateProtoToForm("roleSingleQuery", &idm.RoleSingleQuery{}, false, keys)
		GenerateProtoToForm("workspaceSingleQuery", &idm.WorkspaceSingleQuery{}, false, keys)
		GenerateProtoToForm("aclSingleQuery", &idm.ACLSingleQuery{}, false, keys)
		GenerateProtoToForm("aclAction", &idm.ACLAction{}, false, keys)
		GenerateProtoToForm("treeQuery", &tree.Query{}, false, keys)
		GenerateProtoToForm("actionOutputSingleQuery", &jobs.ActionOutputSingleQuery{}, false, keys)
		GenerateProtoToForm("contextMetaSingleQuery", &jobs.ContextMetaSingleQuery{}, false, keys)

		ctxMeta := []string{
			servicecontext.HttpMetaRemoteAddress,
			servicecontext.HttpMetaUserAgent,
			servicecontext.HttpMetaContentType,
			servicecontext.HttpMetaProtocol,
			servicecontext.HttpMetaHost,
			servicecontext.HttpMetaPort,
			servicecontext.HttpMetaHostname,
			servicecontext.HttpMetaRequestMethod,
			servicecontext.HttpMetaRequestURI,
			servicecontext.HttpMetaCookiesString,
			servicecontext.ServerTime,
		}
		for _, k := range ctxMeta {
			keys["contextMetaField."+k] = k
		}
		for name, f := range ladon.ConditionFactories {
			condition := f()
			GenerateProtoToForm("condition"+condition.GetName(), condition, false, keys)
			keys["contextMetaCondition."+name] = name
		}

		var data = make(map[string]interface{}, len(keys))
		var kk []string
		for k, _ := range keys {
			kk = append(kk, k)
		}
		sort.Strings(kk)
		for _, k := range kk {
			data[k] = map[string]string{
				"other": keys[k],
			}
		}
		j, er := json.MarshalIndent(data, "", "  ")
		So(er, ShouldBeNil)
		tmpFile, e := os.OpenFile("./en-us.all.json", os.O_WRONLY|os.O_CREATE, 0777)
		So(e, ShouldBeNil)
		defer tmpFile.Close()
		t.Log("Printing json to " + tmpFile.Name())
		_, e = tmpFile.Write(j)
		So(e, ShouldBeNil)
	})
}

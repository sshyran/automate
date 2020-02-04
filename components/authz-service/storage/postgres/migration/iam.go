package migration

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	constants_v1 "github.com/chef/automate/components/authz-service/constants/v1"
	constants_v2 "github.com/chef/automate/components/authz-service/constants/v2"
	storage_errors "github.com/chef/automate/components/authz-service/storage"
	storage_v1 "github.com/chef/automate/components/authz-service/storage/v1"
	storage "github.com/chef/automate/components/authz-service/storage/v2"
	"github.com/chef/automate/components/automate-cli/pkg/status"
	uuid "github.com/chef/automate/lib/uuid4"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
)

const (
	enumPristine        = "init"
	enumInProgress      = "in-progress"
	enumSuccessful      = "successful"
	enumSuccessfulBeta1 = "successful-beta1"
	enumFailed          = "failed"
)

// IAM v2 default policy IDs.
const (
	AdminPolicyID  = "administrator-access"
	EditorPolicyID = "editor-access"
	ViewerPolicyID = "viewer-access"
	IngestPolicyID = "ingest-access"
)

// IAM v2 system policy IDs. These are never shown to the enduser
// so GUIDs are fine.
const (
	UniversalAccessPolicyID  = "e729c61f-c40a-4bfa-affe-2a541368169f"
	IngestProviderPolicyID   = "e166f6f9-860d-464a-a91f-be3509369f92"
	SystemPolicyID           = "1074e13b-a918-4892-98be-47a5a8b2d2b6"
	SystemLocalUsersPolicyID = "00a38187-7557-4105-92a0-48db63af4103"
	ChefManagedPolicyID      = "e62bc524-d903-4708-92de-a4435ce0252e"
)

// V1 -> IAM v2 Legacy Policy IDs.
const (
	CfgmgmtPolicyID         = "infrastructure-automation-access-legacy"
	CompliancePolicyID      = "compliance-access-legacy"
	EventsPolicyID          = "events-access-legacy"
	LegacyIngestPolicyID    = "ingest-access-legacy"
	NodesPolicyID           = "nodes-access-legacy"
	NodeManagersPolicyID    = "node-managers-access-legacy"
	SecretsPolicyID         = "secrets-access-legacy"
	TelemetryPolicyID       = "telemetry-access-legacy"
	ComplianceTokenPolicyID = "compliance-profile-access-legacy"
)

// IAM v2 well-known role IDs
const (
	OwnerRoleID        = "owner"
	EditorRoleID       = "editor"
	ViewerRoleID       = "viewer"
	IngestRoleID       = "ingest"
	ProjectOwnerRoleID = "project-owner"
)

// Subjects
const (
	// LocalAdminsTeamSubject is the member for the local admins team.
	LocalAdminsTeamSubject = "team:local:admins"

	// LocalEditorsTeamSubject is the member for the local editors team.
	LocalEditorsTeamSubject = "team:local:editors"

	// LocalViewersTeamSubject is the member for the local viewers team.
	LocalViewersTeamSubject = "team:local:viewers"
)

// IAM v2 well-known project IDs
const (
	AllProjectsID         = "~~ALL-PROJECTS~~" // must match rego file!
	AllProjectsExternalID = "*"
	UnassignedProjectID   = "(unassigned)"
)

func needsV2Migration(ctx context.Context, db *sql.DB) (bool, error) {
	var status string
	row := db.QueryRowContext(ctx, `SELECT state FROM migration_status`)
	err := row.Scan(&status)
	if err != nil {
		return true, err // shouldn't happen, migration initializes state
	}
	switch status {
	case enumPristine:
		return true, nil
	case enumSuccessful:
		return false, nil
	case enumSuccessfulBeta1:
		return false, nil
	// TODO how should we properly handle these cases? is re-running migration enough?
	case enumInProgress:
		return true, nil
	case enumFailed:
		return true, nil
	}
	return true, fmt.Errorf("unexpected migration status: %q", status)
}

// MigrateToV2 sets the V2 store to its factory defaults and then migrates
// any existing V1 policies, unless the install is already on IAM v2.
func migrateToV2(ctx context.Context, db *sql.DB) error {
	ifNotOnV2, err := needsV2Migration(ctx, db)
	if err != nil {
		return errors.Wrap(err, "could not query IAM migration state")
	}

	if ifNotOnV2 {
		for _, role := range defaultRoles() {
			if err := createRole(ctx, db, &role); err != nil {
				return errors.Wrap(err, "could not create default role")
			}
		}

		// TODO include policy in error? same above?
		for _, pol := range defaultPolicies() {
			if err := createPolicy(ctx, db, pol); err != nil {
				return errors.Wrap(err, "could not create default policy")
			}
		}

		var reports []string
		errs, err := migrateV1Policies(ctx)
		if err != nil {
			return errors.Wrap(err, "migrate v1 policies: %s")
		}
		for _, e := range errs {
			reports = append(reports, e.Error())
		}
	}

	// // we've made it!
	// var v api.Version
	// switch req.Flag {
	// case api.Flag_VERSION_2_1:
	// 	err = s.store.SuccessBeta1(ctx)
	// 	v = api.Version{Major: api.Version_V2, Minor: api.Version_V1}
	// default:
	// 	err = s.store.Success(ctx)
	// 	v = api.Version{Major: api.Version_V2, Minor: api.Version_V0}
	// }
	// if err != nil {
	// 	recordFailure()
	// 	return nil, status.Errorf(codes.Internal, "record migration status: %s", err.Error())
	// }

	// s.setVersionForInterceptorSwitch(v)
	// return &api.MigrateToV2Resp{Reports: reports}, nil

	return nil
}

/*
COPY PASTA DATABASE CODE

The below is code we've copied from our database functionality because we need
versions of the database functions needed for the migrations that do not change.
This is because this migration is run at a single point in time as part of the schema
upgrades. So this code need to be compatible with a specific schema version that never changes.
*/

// Policy-related database copy-pasta

// Policy represents a policy definition to be persisted to storage.
type Policy struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Members    []Member    `json:"members"`
	Statements []Statement `json:"statements"`
	Type       Type        `json:"type"`
	Projects   []string    `json:"projects"`
}

type Member struct {
	Name string `json:"name"`
}

type Statement struct {
	Actions   []string `json:"actions"`
	Resources []string `json:"resources"`
	Role      string   `json:"role"`
	Projects  []string `json:"projects"`
	Effect    Effect   `json:"effect"`
}

// Effect is an enum of allow or deny for use in Statements.
type Effect int

const (
	// Allow represents the allow case for a Statement Effect.
	Allow Effect = iota
	// Deny represents the deny case for a Statement Effect.
	Deny
)

func newStatement(effect Effect, role string, projects, resources, actions []string) Statement {
	return Statement{
		Effect:    effect,
		Role:      role,
		Projects:  projects,
		Actions:   actions,
		Resources: resources,
	}
}

func createPolicy(ctx context.Context, db *sql.DB, pol Policy) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Note(sr): we're using BeginTx with the context that'll be cancelled in a
	// `defer` when the function ends. This should rollback transactions that
	// haven't been committed -- what would happen when any of the following
	// `err != nil` cases return early.
	// However, I haven't played with this extensively, so there's a bit of a
	// chance that this understanding is just plain wrong.

	tx, err := db.BeginTx(ctx, nil /* use driver default */)
	if err != nil {
		return errors.Wrap(err, "begin create role tx")
	}

	// This update happens with migrations for default policies
	// projects := pol.Projects
	// if projects == nil {
	// 	projects = []string{}
	// }

	// TODO: at least for the upgrade (not legacy move), we don't need this
	// // skip project permissions check on upgrade from v1 or for chef-managed policies
	// if pol.Type == v2.Custom {
	// 	err = p.ensureNoProjectsMissingWithQuerier(ctx, tx, projects)
	// 	if err != nil {
	// 		return nil, p.processError(err)
	// 	}

	// 	err = projectassignment.AuthorizeProjectAssignment(ctx,
	// 		p.engine,
	// 		auth_context.FromContext(auth_context.FromIncomingMetadata(ctx)).Subjects,
	// 		[]string{},
	// 		projects,
	// 		false)
	// 	if err != nil {
	// 		return nil, p.processError(err)
	// 	}
	// }

	_, err = tx.ExecContext(ctx,
		// This helper should be *right* at this point in time
		`SELECT insert_iam_policy($1, $2, $3);`,
		pol.ID, pol.Name, pol.Type.String(),
	)
	if err != nil {
		return err
	}

	// Cascading drop any existing members.
	// TODO might need for legacy migration?
	// _, err := q.ExecContext(ctx,
	// 	`DELETE FROM iam_policy_members WHERE policy_id=policy_db_id($1);`, policyID)
	// if err != nil {
	// 	return err
	// }

	for _, member := range pol.Members {
		_, err := tx.ExecContext(ctx,
			"INSERT INTO iam_members (name) VALUES ($1) ON CONFLICT DO NOTHING",
			member.Name)
		if err != nil {
			return errors.Wrapf(err, "failed to upsert member %s", member.Name)
		}

		_, err = tx.ExecContext(ctx,
			`INSERT INTO iam_policy_members (policy_id, member_id)
				VALUES (policy_db_id($1), member_db_id($2)) ON CONFLICT DO NOTHING`, pol.ID, member.Name)
		return errors.Wrapf(err, "failed to upsert member link: member=%s, policy_id=%s", member.Name, pol.ID)

	}

	// if err := p.notifyPolicyChange(ctx, tx); err != nil {
	// 	return nil, p.processError(err)
	// }
	// below is the notifyPolicyChange code, ported
	// _, err := q.ExecContext(ctx, "SELECT notify_policy_change()")
	// return err
	err = tx.Commit()
	if err != nil {
		return storage_errors.NewTxCommitError(err)
	}
	return nil
}

// DefaultPolicies shipped with IAM v2, and also the set of policies to which we
// factory-reset our storage.
func defaultPolicies() []Policy {
	// admin policy statements
	s1 := newStatement(Allow, "", []string{}, []string{"*"}, []string{"*"})
	s2 := newStatement(Deny, "", []string{}, []string{"iam:policies:" + AdminPolicyID},
		[]string{"iam:policies:delete", "iam:policies:update"})

	// editor policy statements
	s3 := newStatement(Allow, EditorRoleID, []string{}, []string{"*"}, []string{})

	// viewer policy statements
	s4 := newStatement(Allow, ViewerRoleID, []string{}, []string{"*"}, []string{})

	// ingest policy statements
	s5 := newStatement(Allow, IngestRoleID, []string{}, []string{"*"}, []string{})

	admin := Member{Name: LocalAdminsTeamSubject}
	editors := Member{Name: LocalEditorsTeamSubject}
	viewers := Member{Name: LocalViewersTeamSubject}

	adminPol := Policy{
		ID:         AdminPolicyID,
		Name:       "Administrator",
		Members:    []Member{admin},
		Statements: []Statement{s1, s2},
		Type:       ChefManaged,
	}

	editorPol := Policy{
		ID:         EditorPolicyID,
		Name:       "Editors",
		Members:    []Member{editors},
		Statements: []Statement{s3},
		Type:       ChefManaged,
	}

	viewerPol := Policy{
		ID:         ViewerPolicyID,
		Name:       "Viewers",
		Members:    []Member{viewers},
		Statements: []Statement{s4},
		Type:       ChefManaged,
	}

	ingestPol := Policy{
		ID:         IngestPolicyID,
		Name:       "Ingest",
		Members:    []Member{},
		Statements: []Statement{s5},
		Type:       ChefManaged,
	}

	return []Policy{adminPol, editorPol, viewerPol, ingestPol}
}

// Role-related database copy-pasta

func createRole(ctx context.Context, db *sql.DB, role *Role) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := db.BeginTx(ctx, nil /* use driver default */)
	if err != nil {
		return errors.Wrap(err, "begin create role tx")
	}

	row := tx.QueryRowContext(ctx, `INSERT INTO iam_roles (id, name, type, actions) VALUES ($1, $2, $3, $4)
		RETURNING db_id`,
		role.ID, role.Name, role.Type.String(), pq.Array(role.Actions))
	var dbID string
	if err := row.Scan(&dbID); err != nil {
		return errors.Wrap(err, "insert role")
	}

	_, err = tx.ExecContext(ctx,
		`INSERT INTO iam_role_projects (role_id, project_id)
		SELECT $1, project_db_id(p) FROM unnest($2::TEXT[]) as p`,
		dbID, pq.Array(role.Projects))
	if err != nil {
		return errors.Wrap(err, "insert role projects")
	}

	err = tx.Commit()
	if err != nil {
		return storage_errors.NewTxCommitError(err)
	}

	return nil
}

// Type is an enum to denote custom or chef-managed policy.
type Type int

const (
	// Custom represents a policy created by the enduser.
	Custom Type = iota
	// ChefManaged represents a policy created by Chef Software.
	ChefManaged
	// System represents a policy that is only loaded directly into OPA
	// to allow Automate to function correctly without revealing Automate's
	// internal policies to the customer
	// This type is only used in the OPA cache (not in API or database)
	System
)

const (
	customTypeString  = "custom"
	managedTypeString = "chef-managed"
	systemTypeString  = "system"
)

var strValues = [...]string{
	customTypeString,
	managedTypeString,
	systemTypeString,
}

func (t Type) String() string {
	if t < Custom || t > System {
		panic(fmt.Sprintf("unknown value from iota Type on String() conversion: %d", t))
	}

	return strValues[t]
}

type Role struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Actions  []string `json:"actions"`
	Type     Type     `json:"type"`
	Projects []string `json:"projects"`
}

// DefaultRoles defines the default Chef-managed roles provided on storage reset
func defaultRoles() []Role {
	owner := Role{
		ID:      OwnerRoleID,
		Name:    "Owner",
		Actions: []string{"*"},
		Type:    ChefManaged,
	}

	editor := Role{
		ID:   EditorRoleID,
		Name: "Editor",
		Actions: []string{
			"infra:*",
			"compliance:*",
			"system:*",
			"event:*",
			"ingest:*",
			"secrets:*",
			"telemetry:*",
		},
		Type: ChefManaged,
	}

	viewer := Role{
		ID:   ViewerRoleID,
		Name: "Viewer",
		Actions: []string{
			"secrets:*:get",
			"secrets:*:list",
			"infra:*:get",
			"infra:*:list",
			"compliance:*:get",
			"compliance:*:list",
			"system:*:get",
			"system:*:list",
			"event:*:get",
			"event:*:list",
			"ingest:*:get",
			"ingest:*:list",
		},
		Type: ChefManaged,
	}

	ingest := Role{
		ID:   IngestRoleID,
		Name: "Ingest",
		Actions: []string{
			"infra:ingest:*",
			"compliance:profiles:get",
			"compliance:profiles:list",
		},
		Type: ChefManaged,
	}

	return []Role{owner, editor, viewer, ingest}
}

// all the migration code for translating v1 policies

// migrateV1Policies has two error returns: the second one is the ordinary,
// garden-variety, "something went wrong, I've given up" signal; the first one
// serves as an aggregate of errors that happened attempting to convert and
// store individual (custom) policies.
func migrateV1Policies(ctx context.Context) ([]error, error) {
	pols, err := ListPoliciesWithSubjects(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list v1 policies: %s", err.Error())
	}

	var errs []error
	for _, pol := range pols {
		adminTokenPolicy, err := checkForAdminTokenPolicy(pol)
		if err != nil {
			errs = append(errs, errors.Wrapf(err, "verify subjects %q for admin policy %q", pol.Subjects, pol.ID.String()))
			continue
		}
		if adminTokenPolicy != nil {
			if err := s.addTokenToAdminPolicy(ctx, adminTokenPolicy.Subjects[0]); err != nil {
				errs = append(errs, errors.Wrapf(err, "adding members %q for admin policy %q", pol.Subjects, pol.ID.String()))
			}
			continue // don't migrate admin policies with single token
		}
		storagePol, err := migrateV1Policy(pol)
		if err != nil {
			// collect error
			errs = append(errs, errors.Wrapf(err, "convert v1 policy %q", pol.ID.String()))
			continue // nothing to create
		}
		if storagePol == nil {
			continue // nothing to create
		}
		_, err = s.store.CreatePolicy(ctx, storagePol, true)
		switch err {
		case nil, storage_errors.ErrConflict: // ignore, continue
		default:
			errs = append(errs, errors.Wrapf(err, "store converted v1 policy %q", pol.ID.String()))
		}
	}
	return errs, nil
}

func migrateV1Policy(pol *storage_v1.Policy) (*storage.Policy, error) {
	wellKnown, err := isWellKnown(pol.ID)
	if err != nil {
		return nil, errors.Wrapf(err, "lookup v1 default policy %q", pol.ID.String())
	}

	if wellKnown {
		return legacyPolicyFromV1(pol)
	}
	return customPolicyFromV1(pol)
}

func (s *policyServer) addTokenToAdminPolicy(ctx context.Context, tok string) error {
	m, err := storage.NewMember(tok)
	if err != nil {
		return errors.Wrap(err, "format v2 member for admin team")
	}
	mems, err := s.store.AddPolicyMembers(ctx, constants_v2.AdminPolicyID, []storage.Member{m})
	if err != nil {
		return errors.Wrapf(err, "could not add members %q to admin policy", mems)
	}
	return nil
}

func checkForAdminTokenPolicy(pol *storage_v1.Policy) (*storage_v1.Policy, error) {
	if pol.Action == "*" && pol.Resource == "*" && len(pol.Subjects) == 1 && strings.HasPrefix(pol.Subjects[0], "token:") {
		return pol, nil
	}
	return nil, nil
}

func isWellKnown(id uuid.UUID) (bool, error) {
	defaultV1, err := storage_v1.DefaultPolicies()
	if err != nil {
		return false, err
	}
	_, found := defaultV1[id.String()]
	return found, nil
}

var v1PoliciesToSkip = map[string]struct{}{
	constants_v1.AdminPolicyID:                          {},
	constants_v1.ServiceInfoWildcardPolicyID:            {},
	constants_v1.AuthIntrospectionWildcardPolicyID:      {},
	constants_v1.LicenseStatusPolicyID:                  {},
	constants_v1.ReadOwnUserProfilePolicyID:             {},
	constants_v1.LocalUserSelfPolicyID:                  {},
	constants_v1.PolicyVersionPolicyID:                  {},
	constants_v1.OcErchefIngestStatusPolicyID:           {},
	constants_v1.OcErchefIngestEventsPolicyID:           {},
	constants_v1.CSNginxComplianceProfilesPolicyID:      {},
	constants_v1.CSNginxComplianceDataCollectorPolicyID: {},
	constants_v1.ApplicationsServiceGroupsPolicyID:      {},
}

var v1CfgmgmtPolicies = map[string]struct{}{
	constants_v1.CfgmgmtNodesContainerPolicyID: {},
	constants_v1.CfgmgmtNodesWildcardPolicyID:  {},
	constants_v1.CfgmgmtStatsWildcardPolicyID:  {},
}

var v1EventFeedPolicies = map[string]struct{}{
	constants_v1.EventsContainerPolicyID: {},
	constants_v1.EventsWildcardPolicyID:  {},
}

var v1NodesPolicies = map[string]struct{}{
	constants_v1.NodesContainerPolicyID: {},
	constants_v1.NodesWildcardPolicyID:  {},
}

var v1NodeManagersPolicies = map[string]struct{}{
	constants_v1.NodeManagersContainerPolicyID: {},
	constants_v1.NodeManagersWildcardPolicyID:  {},
}

var v1SecretsPolicies = map[string]struct{}{
	constants_v1.SecretsContainerPolicyID: {},
	constants_v1.SecretsWildcardPolicyID:  {},
}

var v1ComplianceTokenPolicies = map[string]struct{}{
	constants_v1.ComplianceTokenReadProfilesPolicyID:   {},
	constants_v1.ComplianceTokenSearchProfilesPolicyID: {},
	constants_v1.ComplianceTokenUploadProfilesPolicyID: {},
}

// nolint: gocyclo
func legacyPolicyFromV1(pol *storage_v1.Policy) (*storage.Policy, error) {
	if _, found := v1PoliciesToSkip[pol.ID.String()]; found {
		return nil, nil
	}
	noProjects := []string{}

	// there's three cfgmgmt policies (which had been deletable) that are now
	// mapped into one:
	//  CfgmgmtNodesContainerPolicyID
	//  CfgmgmtNodesWildcardPolicyID
	//  CfgmgmtStatsWildcardPolicyID
	if _, found := v1CfgmgmtPolicies[pol.ID.String()]; found {
		cfgmgmtStatement, err := storage.NewStatement(storage.Allow, "", []string{},
			[]string{"*"}, []string{"infra:*"})
		if err != nil {
			return nil, errors.Wrap(err, "format v2 statement (cfgmgmt)")
		}
		member, err := storage.NewMember("user:*")
		if err != nil {
			return nil, errors.Wrap(err, "format v2 member (cfgmgmt)")
		}
		cfgmgmtPolicy, err := storage.NewPolicy(constants_v2.CfgmgmtPolicyID,
			"[Legacy] Infrastructure Automation Access",
			storage.Custom, []storage.Member{member},
			[]storage.Statement{cfgmgmtStatement}, noProjects)
		if err != nil {
			return nil, errors.Wrap(err, "format v2 policy (cfgmgmt)")
		}
		return &cfgmgmtPolicy, nil
	}

	if pol.ID.String() == constants_v1.ComplianceWildcardPolicyID {
		complianceStatement, err := storage.NewStatement(storage.Allow, "", []string{},
			[]string{"*"}, []string{"compliance:*"})
		if err != nil {
			return nil, errors.Wrap(err, "format v2 statement (compliance)")
		}
		member, err := storage.NewMember("user:*")
		if err != nil {
			return nil, errors.Wrap(err, "format v2 member (compliance)")
		}
		compliancePolicy, err := storage.NewPolicy(constants_v2.CompliancePolicyID,
			"[Legacy] Compliance Access",
			storage.Custom, []storage.Member{member}, []storage.Statement{complianceStatement}, noProjects)
		if err != nil {
			return nil, errors.Wrap(err, "format v2 policy (cfgmgmt)")
		}
		return &compliancePolicy, nil
	}

	if _, found := v1EventFeedPolicies[pol.ID.String()]; found {
		eventsStatement, err := storage.NewStatement(storage.Allow, "", []string{},
			[]string{"*"}, []string{"event:*"})
		if err != nil {
			return nil, errors.Wrap(err, "format v2 statement (events)")
		}
		member, err := storage.NewMember("user:*")
		if err != nil {
			return nil, errors.Wrap(err, "format v2 member (events)")
		}
		eventsPolicy, err := storage.NewPolicy(constants_v2.EventsPolicyID,
			"[Legacy] Events Access",
			storage.Custom, []storage.Member{member}, []storage.Statement{eventsStatement}, noProjects)
		if err != nil {
			return nil, errors.Wrap(err, "format v2 policy (events)")
		}
		return &eventsPolicy, nil
	}

	if pol.ID.String() == constants_v1.IngestWildcardPolicyID {
		ingestStatement, err := storage.NewStatement(storage.Allow, "", []string{},
			[]string{"*"}, []string{"infra:ingest:*"})
		if err != nil {
			return nil, errors.Wrap(err, "format v2 statement (ingest)")
		}
		member, err := storage.NewMember("token:*")
		if err != nil {
			return nil, errors.Wrap(err, "format v2 member (ingest)")
		}

		ingestPolicy, err := storage.NewPolicy(constants_v2.LegacyIngestPolicyID,
			"[Legacy] Ingest Access",
			storage.Custom, []storage.Member{member}, []storage.Statement{ingestStatement}, noProjects)
		if err != nil {
			return nil, errors.Wrap(err, "format v2 policy (ingest)")
		}
		return &ingestPolicy, nil
	}

	if _, found := v1NodesPolicies[pol.ID.String()]; found {
		nodesStatement, err := storage.NewStatement(storage.Allow, "", []string{},
			[]string{"*"}, []string{"infra:nodes:*"})
		if err != nil {
			return nil, errors.Wrap(err, "format v2 statement (nodes)")
		}
		member, err := storage.NewMember("user:*")
		if err != nil {
			return nil, errors.Wrap(err, "format v2 member (nodes)")
		}
		nodesPolicy, err := storage.NewPolicy(
			constants_v2.NodesPolicyID,
			"[Legacy] Nodes Access",
			storage.Custom,
			[]storage.Member{member},
			[]storage.Statement{nodesStatement},
			noProjects,
		)
		if err != nil {
			return nil, errors.Wrap(err, "format v2 policy (nodes)")
		}
		return &nodesPolicy, nil
	}

	if _, found := v1NodeManagersPolicies[pol.ID.String()]; found {
		nodeManagersStatement, err := storage.NewStatement(storage.Allow, "", []string{},
			[]string{"*"}, []string{"infra:nodeManagers:*"})
		if err != nil {
			return nil, errors.Wrap(err, "format v2 statement (nodemanagers)")
		}
		member, err := storage.NewMember("user:*")
		if err != nil {
			return nil, errors.Wrap(err, "format v2 member (nodemanagers)")
		}
		nodeManagersPolicy, err := storage.NewPolicy(
			constants_v2.NodeManagersPolicyID,
			"[Legacy] Node Managers Access",
			storage.Custom,
			[]storage.Member{member},
			[]storage.Statement{nodeManagersStatement},
			noProjects,
		)
		if err != nil {
			return nil, errors.Wrap(err, "format v2 policy (nodemanagers)")
		}
		return &nodeManagersPolicy, nil
	}

	if _, found := v1SecretsPolicies[pol.ID.String()]; found {
		secretsStatement, err := storage.NewStatement(storage.Allow, "", []string{},
			[]string{"*"}, []string{"secrets:*"})
		if err != nil {
			return nil, errors.Wrap(err, "format v2 statement (secrets)")
		}
		member, err := storage.NewMember("user:*")
		if err != nil {
			return nil, errors.Wrap(err, "format v2 member (secrets)")
		}
		secretsPolicy, err := storage.NewPolicy(
			constants_v2.SecretsPolicyID,
			"[Legacy] Secrets Access",
			storage.Custom,
			[]storage.Member{member},
			[]storage.Statement{secretsStatement},
			noProjects,
		)
		if err != nil {
			return nil, errors.Wrap(err, "format v2 policy (secrets)")
		}
		return &secretsPolicy, nil
	}

	if pol.ID.String() == constants_v1.TelemetryConfigPolicyID {
		telemetryStatement, err := storage.NewStatement(storage.Allow, "", []string{},
			[]string{"*"}, []string{"system:telemetryConfig:*"})
		if err != nil {
			return nil, errors.Wrap(err, "format v2 statement (telemetry)")
		}
		member, err := storage.NewMember("user:*")
		if err != nil {
			return nil, errors.Wrap(err, "format v2 member (telemetry)")
		}
		telemetryPolicy, err := storage.NewPolicy(
			constants_v2.TelemetryPolicyID,
			"[Legacy] Telemetry Access",
			storage.Custom,
			[]storage.Member{member},
			[]storage.Statement{telemetryStatement},
			noProjects,
		)
		if err != nil {
			return nil, errors.Wrap(err, "format v2 policy (telemetry)")
		}
		return &telemetryPolicy, nil
	}

	if _, found := v1ComplianceTokenPolicies[pol.ID.String()]; found {
		complianceTokenStatement, err := storage.NewStatement(storage.Allow, "", []string{},
			[]string{"*"}, []string{"compliance:profiles:*"})
		if err != nil {
			return nil, errors.Wrap(err, "format v2 statement (compliance token)")
		}
		member, err := storage.NewMember("token:*")
		if err != nil {
			return nil, errors.Wrap(err, "format v2 member (compliance token)")
		}
		complianceTokenPolicy, err := storage.NewPolicy(
			constants_v2.ComplianceTokenPolicyID,
			"[Legacy] Compliance Profile Access",
			storage.Custom,
			[]storage.Member{member},
			[]storage.Statement{complianceTokenStatement},
			noProjects,
		)
		if err != nil {
			return nil, errors.Wrap(err, "format v2 policy (compliance token)")
		}
		return &complianceTokenPolicy, nil
	}

	return nil, errors.New("unknown \"well-known\" policy")
}

func customPolicyFromV1(pol *storage_v1.Policy) (*storage.Policy, error) {
	name := fmt.Sprintf("%s (custom)", pol.ID.String())

	// TODO: If we encounter an unknown action can we just be less permissive with a warning?
	// AKA just use []string{"*"} instead of failing the migration?
	action, err := convertV1Action(pol.Action, pol.Resource)
	if err != nil {
		return nil, errors.Wrap(err, "could not derive v2 action")
	}

	resource, err := convertV1Resource(pol.Resource)
	if err != nil {
		return nil, errors.Wrap(err, "could not derive v2 resource")
	}

	// Note: v1 only had (custom) allow policies
	statement, err := storage.NewStatement(storage.Allow, "", []string{}, []string{resource}, action)
	if err != nil {
		return nil, errors.Wrap(err, "format v2 statement")
	}

	members := make([]storage.Member, len(pol.Subjects))
	for i, subject := range pol.Subjects {
		memberInt, err := storage.NewMember(subject)
		if err != nil {
			return nil, errors.Wrap(err, "format v2 member")
		}
		members[i] = memberInt
	}

	policy, err := storage.NewPolicy(
		pol.ID.String(),
		name,
		storage.Custom,
		members,
		[]storage.Statement{statement},
		[]string{})
	if err != nil {
		return nil, errors.Wrap(err, "format v2 policy")
	}

	return &policy, nil
}

// Basically implements "Current Resource" -> "New Resource Names (RFR)" column of
// https://docs.google.com/spreadsheets/d/1ccaYDJdMnHBfFgmNC1n2_S1YOnMJ-OgkYd8ySbb-mS0/edit#gid=363200100
func convertV1Resource(resource string) (string, error) {
	terms := strings.Split(resource, ":")
	if len(terms) == 0 {
		return "", errors.New("there was no resource passed")
	}

	if len(terms) == 1 && terms[0] == "*" {
		return "*", nil
	}

	switch terms[0] {
	case "service_info":
		switch terms[1] {
		case "version":
			return "system:service:version", nil
		case "health":
			return "system:health", nil
		}
		return "system:*", nil
	case "auth":
		terms = changeTerm(terms, "auth", "iam")
		terms = changeTerm(terms, "api_tokens", "tokens")
		return combineTermsIntoResource(terms...), nil
	case "users":
		// "users" -> "iam:usersSelf"
		terms[0] = "usersSelf"
		return combineTermsIntoResource(prefixTerms("iam", terms)...), nil
	case "auth_introspection":
		// Special case
		if terms[1] == "*" {
			return "iam:introspect", nil
		}
		terms = changeTerm(terms, "auth_introspection", "iam")
		terms = changeTerm(terms, "introspect_all", "introspect")
		terms = changeTerm(terms, "introspect_some", "introspect")
		return combineTermsIntoResource(terms...), nil
	case "cfgmgmt":
		return convertV1Cfgmgmt(terms)
	case "compliance":
		// Special case
		if resource == "compliance:profiles:market" {
			return "compliance:marketProfiles", nil
		}
		return combineTermsIntoResource(deleteTerm(terms, "storage")...), nil
	case "events":
		return convertV1Events(terms)
	case "ingest":
		// Special case: "ingest:status" -> "infra:ingest:status" (no wildcards to worry about)
		if terms[1] == "status" {
			return "infra:ingest:status", nil
		}

		// Special case: "ingest:unified_events" -> "infra:unifiedEvents" (no wildcards to worry about)
		if terms[1] == "unified_events" {
			return "infra:unifiedEvents", nil
		}

		terms = changeTerm(terms, "ingest", "infra")
		terms = changeTerm(terms, "unified_events", "unifiedEvents")
		return combineTermsIntoResource(terms...), nil
	case "license":
		if len(terms) == 1 {
			return "system:license", nil
		}
		// if len(terms) == 2 aka license:* or license:status
		return "system:status", nil
	case "nodemanagers":
		// "nodemanagers" -> "infra:nodeManagers"
		return combineTermsIntoResource(prefixTerms("infra", changeTerm(terms, "nodemanagers", "nodeManagers"))...),
			nil
	case "nodes":
		// "nodes" -> "infra:nodes"
		return combineTermsIntoResource(prefixTerms("infra", terms)...), nil
	case "secrets":
		// "secrets" -> "secrets:secrets"
		return combineTermsIntoResource(prefixTerms("secrets", terms)...), nil
	case "telemetry":
		// either telemetry:config or telemetry:* maps to system:config
		return "system:config", nil
	case "notifications":
		return resource, nil // unchanged
	case "service_groups":
		return "applications:serviceGroups", nil
	default:
		return "", fmt.Errorf("did not recognize base v1 resource term: %s", terms[0])
	}
}

func convertV1Cfgmgmt(terms []string) (string, error) {
	if terms[1] == "stats" {
		return "infra:nodes", nil
	}
	// cfgmgmt:nodes:{node_id}:runs -> infra:nodes:{node_id}
	// cfgmgmt:nodes:{node_id}:runs:{run_id}" -> infra:nodes:{node_id}
	if len(terms) >= 4 && terms[3] == "runs" {
		return combineTermsIntoResource("infra", "nodes", terms[2]), nil
	}
	// cfgmgmt:nodes:{node_id}:attribute -> infra:nodes:{node_id}
	if len(terms) >= 4 && terms[3] == "attribute" {
		return combineTermsIntoResource("infra", "nodes", terms[2]), nil
	}
	// cfgmgmt:nodes:{node_id}:* -> infra:nodes:{node_id}
	if len(terms) >= 3 && terms[2] == "node_id" {
		return combineTermsIntoResource("infra", "nodes", terms[2]), nil
	}
	terms = changeTerm(terms, "cfgmgmt", "infra")
	terms = changeTerm(terms, "marked-nodes", "markedNodes")
	return combineTermsIntoResource(terms...), nil
}

func convertV1Events(terms []string) (string, error) {
	// "events:*" -> "event:events"
	if len(terms) == 1 {
		return "event:events", nil
	}

	switch terms[1] {
	case "types", "tasks", "strings":
		return "event:events", nil
	default:
		return combineTermsIntoResource(changeTerm(terms, "event", "events")...), nil
	}
}

func changeTerm(terms []string, original, updated string) []string {
	for i, term := range terms {
		if term == original {
			terms[i] = updated
		}
	}
	return terms
}

func prefixTerms(prefix string, terms []string) []string {
	return append([]string{prefix}, terms...)
}

func deleteTerm(terms []string, original string) []string {
	for i, term := range terms {
		if term == original {
			terms = append(terms[:i], terms[i+1:]...)
		}
	}
	return terms
}

func combineTermsIntoResource(terms ...string) string {
	return strings.Join(terms, ":")
}

func convertV1Action(action string, resource string) ([]string, error) {
	terms := strings.Split(resource, ":")
	// introspect is a special case,
	// since there was only the "read" action for every different
	// API endpoint, we must assume them all.
	if terms[0] == "auth_introspection" {
		if action == "*" {
			return []string{"*"}, nil
		}
		return []string{"*:getAll", "*:getSome", "*:get"}, nil
	}

	// resource ingest:* is a special case to handle since we collapsed
	// that top-level ingest term into infra:ingest.
	// Only need to handle the * case specially though.
	if terms[0] == "ingest" && terms[1] == "*" {
		switch action {
		case "*":
			return []string{"infra:ingest:*"}, nil
		case "create":
			return []string{"infra:ingest:create", "infra:ingestUnifiedEvents:create"}, nil
		case "delete":
			return []string{"infra:ingest:delete"}, nil
		case "status":
			return []string{"infra:ingestStatus:get"}, nil
		}

		return []string{"*:getAll", "*:getSome", "*:get"}, nil
	}

	if terms[0] == "events" && action == "count" {
		return []string{"*:list"}, nil
	}

	// The rest of all all actions for other resources can
	// be generally mapped.
	switch action {
	case "read":
		return []string{"*:get", "*:list"}, nil
	case "create":
		return []string{"*:create"}, nil
	case "update":
		return []string{"*:update"}, nil
	case "delete":
		return []string{"*:delete"}, nil
	case "search":
		return []string{"*:list"}, nil
	case "rerun":
		return []string{"*:rerun"}, nil
	case "count":
		return []string{"*:get"}, nil
	case "start":
		return []string{"*:start"}, nil
	case "stop":
		return []string{"*:stop"}, nil // TODO: gdoc says stop -> ["*:start"] but pretty sure that is wrong xD
	case "configure":
		return []string{"*:update"}, nil
	case "mark-missing":
		return []string{"*:markMissing"}, nil
	case "apply":
		return []string{"*:apply"}, nil
	case "request":
		return []string{"*:request"}, nil
	case "list":
		return []string{"*:list"}, nil
	case "validate":
		return []string{"*:validate"}, nil
	case "export":
		return []string{"*:export"}, nil
	case "upload":
		return []string{"*:create"}, nil
	case "*":
		return []string{"*"}, nil
	default:
		// TODO: should we just warn the logs in this case and not crash the whole migration on the off
		// chance that the user created a policy where they mistyped an action name?
		return nil, fmt.Errorf("could not parse V1 action: %s", action)
	}
}

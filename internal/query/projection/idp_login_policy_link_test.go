package projection

import (
	"testing"

	"github.com/caos/zitadel/internal/domain"
	"github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore"
	"github.com/caos/zitadel/internal/eventstore/handler"
	"github.com/caos/zitadel/internal/eventstore/repository"
	"github.com/caos/zitadel/internal/repository/iam"
	"github.com/caos/zitadel/internal/repository/org"
)

func TestIDPLoginPolicyLinkProjection_reduces(t *testing.T) {
	type args struct {
		event func(t *testing.T) eventstore.EventReader
	}
	tests := []struct {
		name   string
		args   args
		reduce func(event eventstore.EventReader) (*handler.Statement, error)
		want   wantReduce
	}{
		{
			name: "iam.reduceAdded",
			args: args{
				event: getEvent(testEvent(
					repository.EventType(iam.LoginPolicyIDPProviderAddedEventType),
					iam.AggregateType,
					[]byte(`{
	"idpConfigId": "idp-config-id",
    "idpProviderType": 1
}`),
				), iam.IdentityProviderAddedEventMapper),
			},
			reduce: (&IDPLoginPolicyLinkProjection{}).reduceAdded,
			want: wantReduce{
				aggregateType:    iam.AggregateType,
				sequence:         15,
				previousSequence: 10,
				projection:       IDPLoginPolicyLinkTable,
				executer: &testExecuter{
					executions: []execution{
						{
							expectedStmt: "INSERT INTO zitadel.projections.idp_login_policy_links (idp_id, aggregate_id, creation_date, change_date, sequence, resource_owner, provider_type) VALUES ($1, $2, $3, $4, $5, $6, $7)",
							expectedArgs: []interface{}{
								"idp-config-id",
								"agg-id",
								anyArg{},
								anyArg{},
								uint64(15),
								"ro-id",
								domain.IdentityProviderTypeSystem,
							},
						},
					},
				},
			},
		},
		{
			name: "iam.reduceRemoved",
			args: args{
				event: getEvent(testEvent(
					repository.EventType(iam.LoginPolicyIDPProviderRemovedEventType),
					iam.AggregateType,
					[]byte(`{
	"idpConfigId": "idp-config-id",
    "idpProviderType": 1
}`),
				), iam.IdentityProviderRemovedEventMapper),
			},
			reduce: (&IDPLoginPolicyLinkProjection{}).reduceRemoved,
			want: wantReduce{
				aggregateType:    iam.AggregateType,
				sequence:         15,
				previousSequence: 10,
				projection:       IDPLoginPolicyLinkTable,
				executer: &testExecuter{
					executions: []execution{
						{
							expectedStmt: "DELETE FROM zitadel.projections.idp_login_policy_links WHERE (idp_id = $1) AND (aggregate_id = $2)",
							expectedArgs: []interface{}{
								"idp-config-id",
								"agg-id",
							},
						},
					},
				},
			},
		},
		{
			name: "iam.reduceCascadeRemoved",
			args: args{
				event: getEvent(testEvent(
					repository.EventType(iam.LoginPolicyIDPProviderCascadeRemovedEventType),
					iam.AggregateType,
					[]byte(`{
	"idpConfigId": "idp-config-id",
    "idpProviderType": 1
}`),
				), iam.IdentityProviderCascadeRemovedEventMapper),
			},
			reduce: (&IDPLoginPolicyLinkProjection{}).reduceCascadeRemoved,
			want: wantReduce{
				aggregateType:    iam.AggregateType,
				sequence:         15,
				previousSequence: 10,
				projection:       IDPLoginPolicyLinkTable,
				executer: &testExecuter{
					executions: []execution{
						{
							expectedStmt: "DELETE FROM zitadel.projections.idp_login_policy_links WHERE (idp_id = $1) AND (aggregate_id = $2)",
							expectedArgs: []interface{}{
								"idp-config-id",
								"agg-id",
							},
						},
					},
				},
			},
		},
		{
			name: "org.reduceAdded",
			args: args{
				event: getEvent(testEvent(
					repository.EventType(org.LoginPolicyIDPProviderAddedEventType),
					org.AggregateType,
					[]byte(`{
	"idpConfigId": "idp-config-id",
    "idpProviderType": 1
}`),
				), org.IdentityProviderAddedEventMapper),
			},
			reduce: (&IDPLoginPolicyLinkProjection{}).reduceAdded,
			want: wantReduce{
				aggregateType:    org.AggregateType,
				sequence:         15,
				previousSequence: 10,
				projection:       IDPLoginPolicyLinkTable,
				executer: &testExecuter{
					executions: []execution{
						{
							expectedStmt: "INSERT INTO zitadel.projections.idp_login_policy_links (idp_id, aggregate_id, creation_date, change_date, sequence, resource_owner, provider_type) VALUES ($1, $2, $3, $4, $5, $6, $7)",
							expectedArgs: []interface{}{
								"idp-config-id",
								"agg-id",
								anyArg{},
								anyArg{},
								uint64(15),
								"ro-id",
								domain.IdentityProviderTypeOrg,
							},
						},
					},
				},
			},
		},
		{
			name: "org.reduceRemoved",
			args: args{
				event: getEvent(testEvent(
					repository.EventType(org.LoginPolicyIDPProviderRemovedEventType),
					org.AggregateType,
					[]byte(`{
	"idpConfigId": "idp-config-id",
    "idpProviderType": 1
}`),
				), org.IdentityProviderRemovedEventMapper),
			},
			reduce: (&IDPLoginPolicyLinkProjection{}).reduceRemoved,
			want: wantReduce{
				aggregateType:    org.AggregateType,
				sequence:         15,
				previousSequence: 10,
				projection:       IDPLoginPolicyLinkTable,
				executer: &testExecuter{
					executions: []execution{
						{
							expectedStmt: "DELETE FROM zitadel.projections.idp_login_policy_links WHERE (idp_id = $1) AND (aggregate_id = $2)",
							expectedArgs: []interface{}{
								"idp-config-id",
								"agg-id",
							},
						},
					},
				},
			},
		},
		{
			name: "org.reduceCascadeRemoved",
			args: args{
				event: getEvent(testEvent(
					repository.EventType(org.LoginPolicyIDPProviderCascadeRemovedEventType),
					org.AggregateType,
					[]byte(`{
	"idpConfigId": "idp-config-id",
    "idpProviderType": 1
}`),
				), org.IdentityProviderCascadeRemovedEventMapper),
			},
			reduce: (&IDPLoginPolicyLinkProjection{}).reduceCascadeRemoved,
			want: wantReduce{
				aggregateType:    org.AggregateType,
				sequence:         15,
				previousSequence: 10,
				projection:       IDPLoginPolicyLinkTable,
				executer: &testExecuter{
					executions: []execution{
						{
							expectedStmt: "DELETE FROM zitadel.projections.idp_login_policy_links WHERE (idp_id = $1) AND (aggregate_id = $2)",
							expectedArgs: []interface{}{
								"idp-config-id",
								"agg-id",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event := baseEvent(t)
			got, err := tt.reduce(event)
			if _, ok := err.(errors.InvalidArgument); !ok {
				t.Errorf("no wrong event mapping: %v, got: %v", err, got)
			}

			event = tt.args.event(t)
			got, err = tt.reduce(event)
			assertReduce(t, got, err, tt.want)
		})
	}
}
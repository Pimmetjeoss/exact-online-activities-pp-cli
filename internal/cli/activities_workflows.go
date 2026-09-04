// Copyright 2026 pimmetjeoss. Licensed under Apache-2.0. See LICENSE.
// PATCH: Activities-specific discovery helpers and workflow commands generated from Exact's official Activities docs.
package cli

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	registerNovelCommand(func(root *cobra.Command, flags *rootFlags) {
		var parent *cobra.Command
		for _, sub := range root.Commands() {
			if sub.Name() == "activities" {
				parent = sub
				break
			}
		}
		if parent == nil {
			return
		}
		addNovelCommandIfAbsent(parent, newActivitiesResourcesCmd(flags))
		addNovelCommandIfAbsent(parent, newActivitiesDocsCmd(flags))
		addNovelCommandIfAbsent(parent, newActivitiesOverviewCmd(flags))
		addNovelCommandIfAbsent(parent, newActivitiesOpenItemsCmd(flags))
		addNovelCommandIfAbsent(parent, newActivitiesByAccountCmd(flags))
	})
}

type activitiesODataFlags struct {
	filter       string
	selectClause string
	orderBy      string
	top          int
	skipToken    string
	expand       string
	query        []string
}

func addActivitiesODataFlags(cmd *cobra.Command, f *activitiesODataFlags) {
	cmd.Flags().StringVar(&f.filter, "filter", "", "OData $filter expression, e.g. Account eq guid'...' or Status eq 1 (sent as $filter)")
	cmd.Flags().StringVar(&f.selectClause, "odata-select", "", "OData $select clause; use root --select for output shaping")
	cmd.Flags().StringVar(&f.orderBy, "orderby", "", "OData $orderby expression, e.g. Modified desc")
	cmd.Flags().IntVar(&f.top, "top", 0, "OData $top row limit")
	cmd.Flags().StringVar(&f.skipToken, "skiptoken", "", "OData $skiptoken continuation token")
	cmd.Flags().StringVar(&f.expand, "expand", "", "OData $expand clause")
	cmd.Flags().StringArrayVar(&f.query, "query", nil, "Extra query parameter key=value; may be repeated")
}

func (f activitiesODataFlags) toParams() map[string]string {
	params := map[string]string{}
	if f.filter != "" {
		params["$filter"] = f.filter
	}
	if f.selectClause != "" {
		params["$select"] = f.selectClause
	}
	if f.orderBy != "" {
		params["$orderby"] = f.orderBy
	}
	if f.top > 0 {
		params["$top"] = fmt.Sprint(f.top)
	}
	if f.skipToken != "" {
		params["$skiptoken"] = f.skipToken
	}
	if f.expand != "" {
		params["$expand"] = f.expand
	}
	for _, q := range f.query {
		k, v, ok := strings.Cut(q, "=")
		if ok && strings.TrimSpace(k) != "" {
			params[strings.TrimSpace(k)] = v
		}
	}
	return params
}

type activitiesResourceInfo struct {
	Name        string   `json:"name"`
	Path        string   `json:"path"`
	Methods     []string `json:"methods"`
	ReadOnly    bool     `json:"read_only"`
	Docs        string   `json:"docs"`
	Description string   `json:"description"`
}

var activitiesResources = []activitiesResourceInfo{
	{"AnnualStatements", "/api/v1/{division}/activities/AnnualStatements", []string{"GET", "POST", "PUT"}, false, "https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ActivitiesAnnualStatements", "Annual statement workflow requests with action dates and statuses."},
	{"CommunicationNotes", "/api/v1/{division}/activities/CommunicationNotes", []string{"GET", "POST"}, false, "https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ActivitiesCommunicationNotes", "Communication notes linked to accounts, contacts, campaigns and attachments."},
	{"Complaints", "/api/v1/{division}/activities/Complaints", []string{"GET", "POST"}, false, "https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ActivitiesComplaints", "Customer complaints linked to accounts, contacts, assigned users and attachments."},
	{"Events", "/api/v1/{division}/activities/Events", []string{"GET", "POST"}, false, "https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ActivitiesEvents", "Events linked to accounts, contacts, campaigns and attachments."},
	{"Fiscals", "/api/v1/{division}/activities/Fiscals", []string{"GET", "POST", "PUT"}, false, "https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ActivitiesFiscals", "Fiscal workflow requests with assigned users, action dates and statuses."},
	{"ServiceRequests", "/api/v1/{division}/activities/ServiceRequests", []string{"GET", "POST"}, false, "https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ActivitiesServiceRequests", "Service requests linked to accounts, contacts, assigned users and attachments."},
	{"Tasks", "/api/v1/{division}/activities/Tasks", []string{"GET", "POST"}, false, "https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ActivitiesTasks", "Tasks linked to accounts, contacts, action dates and attachments."},
}

func activitiesResourceByName(name string) (activitiesResourceInfo, bool) {
	norm := strings.ToLower(strings.ReplaceAll(name, "-", ""))
	for _, r := range activitiesResources {
		if strings.ToLower(r.Name) == strings.ToLower(name) || strings.ToLower(strings.ReplaceAll(r.Name, "-", "")) == norm {
			return r, true
		}
	}
	return activitiesResourceInfo{}, false
}

// pp:data-source local
func newActivitiesResourcesCmd(flags *rootFlags) *cobra.Command {
	var method string
	cmd := &cobra.Command{
		Use:         "resources",
		Short:       "List the covered Exact Online Activities resources and methods",
		Example:     "  exact-online-activities-pp-cli activities resources --json\n  exact-online-activities-pp-cli activities resources --method POST --json",
		Annotations: map[string]string{"mcp:read-only": "true"},
		RunE: func(cmd *cobra.Command, args []string) error {
			rows := make([]activitiesResourceInfo, 0, len(activitiesResources))
			for _, r := range activitiesResources {
				if method == "" {
					rows = append(rows, r)
					continue
				}
				for _, m := range r.Methods {
					if strings.EqualFold(m, method) {
						rows = append(rows, r)
						break
					}
				}
			}
			data, _ := json.Marshal(rows)
			return printOutputWithFlags(cmd.OutOrStdout(), data, flags)
		},
	}
	cmd.Flags().StringVar(&method, "method", "", "Filter by HTTP method (GET, POST, PUT)")
	return cmd
}

// pp:data-source local
func newActivitiesDocsCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:         "docs <resource>",
		Short:       "Show the official Exact docs URL and local endpoint metadata for an Activities resource",
		Example:     "  exact-online-activities-pp-cli activities docs Tasks --json",
		Annotations: map[string]string{"mcp:read-only": "true"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			r, ok := activitiesResourceByName(args[0])
			if !ok {
				return usageErr(fmt.Errorf("unknown activities resource %q", args[0]))
			}
			data, _ := json.Marshal(r)
			return printOutputWithFlags(cmd.OutOrStdout(), data, flags)
		},
	}
	return cmd
}

// pp:data-source live
func newActivitiesOverviewCmd(flags *rootFlags) *cobra.Command {
	var q activitiesODataFlags
	cmd := &cobra.Command{
		Use:         "overview <division>",
		Short:       "Fetch all main Activities lists for a compact cross-resource overview",
		Example:     "  exact-online-activities-pp-cli activities overview 123456 --top 10 --agent",
		Annotations: map[string]string{"mcp:read-only": "true"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			return activitiesMultiGet(cmd, flags, args[0], []string{"Tasks", "Events", "CommunicationNotes", "Complaints", "ServiceRequests", "AnnualStatements", "Fiscals"}, q.toParams())
		},
	}
	addActivitiesODataFlags(cmd, &q)
	return cmd
}

// pp:data-source live
func newActivitiesOpenItemsCmd(flags *rootFlags) *cobra.Command {
	var assignedTo string
	var q activitiesODataFlags
	cmd := &cobra.Command{
		Use:         "open-items <division>",
		Short:       "Fetch workflow-style open Activities: tasks, service requests, complaints, annual statements and fiscal requests",
		Example:     "  exact-online-activities-pp-cli activities open-items 123456 --assigned-to 00000000-0000-0000-0000-000000000000 --top 25 --agent",
		Annotations: map[string]string{"mcp:read-only": "true"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			params := q.toParams()
			if assignedTo != "" && params["$filter"] == "" {
				params["$filter"] = "AssignedTo eq guid'" + assignedTo + "'"
			}
			return activitiesMultiGet(cmd, flags, args[0], []string{"Tasks", "ServiceRequests", "Complaints", "AnnualStatements", "Fiscals"}, params)
		},
	}
	cmd.Flags().StringVar(&assignedTo, "assigned-to", "", "Exact user GUID; builds an AssignedTo eq guid'...' filter when --filter is absent")
	addActivitiesODataFlags(cmd, &q)
	return cmd
}

// pp:data-source live
func newActivitiesByAccountCmd(flags *rootFlags) *cobra.Command {
	var accountID string
	var q activitiesODataFlags
	cmd := &cobra.Command{
		Use:         "by-account <division>",
		Short:       "Fetch account-related Activities across notes, complaints, events, service requests and tasks",
		Example:     "  exact-online-activities-pp-cli activities by-account 123456 --account-id 00000000-0000-0000-0000-000000000000 --agent",
		Annotations: map[string]string{"mcp:read-only": "true"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			params := q.toParams()
			if accountID != "" && params["$filter"] == "" {
				params["$filter"] = "Account eq guid'" + accountID + "'"
			}
			return activitiesMultiGet(cmd, flags, args[0], []string{"CommunicationNotes", "Complaints", "Events", "ServiceRequests", "Tasks"}, params)
		},
	}
	cmd.Flags().StringVar(&accountID, "account-id", "", "Account GUID; builds an Account eq guid'...' filter when --filter is absent")
	addActivitiesODataFlags(cmd, &q)
	return cmd
}

func activitiesMultiGet(cmd *cobra.Command, flags *rootFlags, division string, resourceNames []string, params map[string]string) error {
	c, err := flags.newClient()
	if err != nil {
		return err
	}
	out := map[string]any{"division": division, "resources": map[string]any{}}
	for _, name := range resourceNames {
		r, _ := activitiesResourceByName(name)
		path := replacePathParam(r.Path, "division", division)
		data, err := c.Get(cmd.Context(), path, params)
		if err != nil {
			return classifyAPIError(cmd.OutOrStdout(), err, flags)
		}
		var parsed any
		if json.Unmarshal(data, &parsed) != nil {
			parsed = string(data)
		}
		out["resources"].(map[string]any)[name] = map[string]any{"path": path, "docs": r.Docs, "data": parsed}
	}
	data, _ := json.Marshal(out)
	return printOutputWithFlags(cmd.OutOrStdout(), data, flags)
}

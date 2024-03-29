// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// TagGroup is an object representing the database table.
type TagGroup struct { // ID
	ID int64 `boil:"id" json:"id" toml:"id" yaml:"id"`
	// ユーザーID
	UserID string `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	// グループ名
	GroupName string `boil:"group_name" json:"group_name" toml:"group_name" yaml:"group_name"`
	// 作成日時
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	// 更新日時
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *tagGroupR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tagGroupL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TagGroupColumns = struct {
	ID        string
	UserID    string
	GroupName string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	UserID:    "user_id",
	GroupName: "group_name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var TagGroupTableColumns = struct {
	ID        string
	UserID    string
	GroupName string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "tag_group.id",
	UserID:    "tag_group.user_id",
	GroupName: "tag_group.group_name",
	CreatedAt: "tag_group.created_at",
	UpdatedAt: "tag_group.updated_at",
}

// Generated where

var TagGroupWhere = struct {
	ID        whereHelperint64
	UserID    whereHelperstring
	GroupName whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "`tag_group`.`id`"},
	UserID:    whereHelperstring{field: "`tag_group`.`user_id`"},
	GroupName: whereHelperstring{field: "`tag_group`.`group_name`"},
	CreatedAt: whereHelpertime_Time{field: "`tag_group`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`tag_group`.`updated_at`"},
}

// TagGroupRels is where relationship names are stored.
var TagGroupRels = struct {
	User          string
	GroupUserTags string
}{
	User:          "User",
	GroupUserTags: "GroupUserTags",
}

// tagGroupR is where relationships are stored.
type tagGroupR struct {
	User          *User        `boil:"User" json:"User" toml:"User" yaml:"User"`
	GroupUserTags UserTagSlice `boil:"GroupUserTags" json:"GroupUserTags" toml:"GroupUserTags" yaml:"GroupUserTags"`
}

// NewStruct creates a new relationship struct
func (*tagGroupR) NewStruct() *tagGroupR {
	return &tagGroupR{}
}

func (r *tagGroupR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

func (r *tagGroupR) GetGroupUserTags() UserTagSlice {
	if r == nil {
		return nil
	}
	return r.GroupUserTags
}

// tagGroupL is where Load methods for each relationship are stored.
type tagGroupL struct{}

var (
	tagGroupAllColumns            = []string{"id", "user_id", "group_name", "created_at", "updated_at"}
	tagGroupColumnsWithoutDefault = []string{"user_id", "group_name", "created_at", "updated_at"}
	tagGroupColumnsWithDefault    = []string{"id"}
	tagGroupPrimaryKeyColumns     = []string{"id"}
	tagGroupGeneratedColumns      = []string{}
)

type (
	// TagGroupSlice is an alias for a slice of pointers to TagGroup.
	// This should almost always be used instead of []TagGroup.
	TagGroupSlice []*TagGroup
	// TagGroupHook is the signature for custom TagGroup hook methods
	TagGroupHook func(context.Context, boil.ContextExecutor, *TagGroup) error

	tagGroupQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tagGroupType                 = reflect.TypeOf(&TagGroup{})
	tagGroupMapping              = queries.MakeStructMapping(tagGroupType)
	tagGroupPrimaryKeyMapping, _ = queries.BindMapping(tagGroupType, tagGroupMapping, tagGroupPrimaryKeyColumns)
	tagGroupInsertCacheMut       sync.RWMutex
	tagGroupInsertCache          = make(map[string]insertCache)
	tagGroupUpdateCacheMut       sync.RWMutex
	tagGroupUpdateCache          = make(map[string]updateCache)
	tagGroupUpsertCacheMut       sync.RWMutex
	tagGroupUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var tagGroupAfterSelectHooks []TagGroupHook

var tagGroupBeforeInsertHooks []TagGroupHook
var tagGroupAfterInsertHooks []TagGroupHook

var tagGroupBeforeUpdateHooks []TagGroupHook
var tagGroupAfterUpdateHooks []TagGroupHook

var tagGroupBeforeDeleteHooks []TagGroupHook
var tagGroupAfterDeleteHooks []TagGroupHook

var tagGroupBeforeUpsertHooks []TagGroupHook
var tagGroupAfterUpsertHooks []TagGroupHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TagGroup) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagGroupAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TagGroup) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagGroupBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TagGroup) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagGroupAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TagGroup) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagGroupBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TagGroup) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagGroupAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TagGroup) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagGroupBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TagGroup) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagGroupAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TagGroup) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagGroupBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TagGroup) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagGroupAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTagGroupHook registers your hook function for all future operations.
func AddTagGroupHook(hookPoint boil.HookPoint, tagGroupHook TagGroupHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		tagGroupAfterSelectHooks = append(tagGroupAfterSelectHooks, tagGroupHook)
	case boil.BeforeInsertHook:
		tagGroupBeforeInsertHooks = append(tagGroupBeforeInsertHooks, tagGroupHook)
	case boil.AfterInsertHook:
		tagGroupAfterInsertHooks = append(tagGroupAfterInsertHooks, tagGroupHook)
	case boil.BeforeUpdateHook:
		tagGroupBeforeUpdateHooks = append(tagGroupBeforeUpdateHooks, tagGroupHook)
	case boil.AfterUpdateHook:
		tagGroupAfterUpdateHooks = append(tagGroupAfterUpdateHooks, tagGroupHook)
	case boil.BeforeDeleteHook:
		tagGroupBeforeDeleteHooks = append(tagGroupBeforeDeleteHooks, tagGroupHook)
	case boil.AfterDeleteHook:
		tagGroupAfterDeleteHooks = append(tagGroupAfterDeleteHooks, tagGroupHook)
	case boil.BeforeUpsertHook:
		tagGroupBeforeUpsertHooks = append(tagGroupBeforeUpsertHooks, tagGroupHook)
	case boil.AfterUpsertHook:
		tagGroupAfterUpsertHooks = append(tagGroupAfterUpsertHooks, tagGroupHook)
	}
}

// One returns a single tagGroup record from the query.
func (q tagGroupQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TagGroup, error) {
	o := &TagGroup{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for tag_group")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TagGroup records from the query.
func (q tagGroupQuery) All(ctx context.Context, exec boil.ContextExecutor) (TagGroupSlice, error) {
	var o []*TagGroup

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TagGroup slice")
	}

	if len(tagGroupAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TagGroup records in the query.
func (q tagGroupQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count tag_group rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q tagGroupQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if tag_group exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *TagGroup) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`user_id` = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// GroupUserTags retrieves all the user_tag's UserTags with an executor via group_id column.
func (o *TagGroup) GroupUserTags(mods ...qm.QueryMod) userTagQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`user_tag`.`group_id`=?", o.ID),
	)

	return UserTags(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (tagGroupL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTagGroup interface{}, mods queries.Applicator) error {
	var slice []*TagGroup
	var object *TagGroup

	if singular {
		var ok bool
		object, ok = maybeTagGroup.(*TagGroup)
		if !ok {
			object = new(TagGroup)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTagGroup)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTagGroup))
			}
		}
	} else {
		s, ok := maybeTagGroup.(*[]*TagGroup)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTagGroup)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTagGroup))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &tagGroupR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &tagGroupR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`user`),
		qm.WhereIn(`user.user_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for user")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user")
	}

	if len(userAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.TagGroups = append(foreign.R.TagGroups, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.UserID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.TagGroups = append(foreign.R.TagGroups, local)
				break
			}
		}
	}

	return nil
}

// LoadGroupUserTags allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (tagGroupL) LoadGroupUserTags(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTagGroup interface{}, mods queries.Applicator) error {
	var slice []*TagGroup
	var object *TagGroup

	if singular {
		var ok bool
		object, ok = maybeTagGroup.(*TagGroup)
		if !ok {
			object = new(TagGroup)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTagGroup)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTagGroup))
			}
		}
	} else {
		s, ok := maybeTagGroup.(*[]*TagGroup)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTagGroup)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTagGroup))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &tagGroupR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &tagGroupR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`user_tag`),
		qm.WhereIn(`user_tag.group_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load user_tag")
	}

	var resultSlice []*UserTag
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice user_tag")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on user_tag")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user_tag")
	}

	if len(userTagAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.GroupUserTags = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userTagR{}
			}
			foreign.R.Group = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.GroupID) {
				local.R.GroupUserTags = append(local.R.GroupUserTags, foreign)
				if foreign.R == nil {
					foreign.R = &userTagR{}
				}
				foreign.R.Group = local
				break
			}
		}
	}

	return nil
}

// SetUser of the tagGroup to the related item.
// Sets o.R.User to related.
// Adds o to related.R.TagGroups.
func (o *TagGroup) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `tag_group` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, tagGroupPrimaryKeyColumns),
	)
	values := []interface{}{related.UserID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.UserID
	if o.R == nil {
		o.R = &tagGroupR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			TagGroups: TagGroupSlice{o},
		}
	} else {
		related.R.TagGroups = append(related.R.TagGroups, o)
	}

	return nil
}

// AddGroupUserTags adds the given related objects to the existing relationships
// of the tag_group, optionally inserting them as new records.
// Appends related to o.R.GroupUserTags.
// Sets related.R.Group appropriately.
func (o *TagGroup) AddGroupUserTags(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*UserTag) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.GroupID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `user_tag` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"group_id"}),
				strmangle.WhereClause("`", "`", 0, userTagPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.GroupID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &tagGroupR{
			GroupUserTags: related,
		}
	} else {
		o.R.GroupUserTags = append(o.R.GroupUserTags, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userTagR{
				Group: o,
			}
		} else {
			rel.R.Group = o
		}
	}
	return nil
}

// SetGroupUserTags removes all previously related items of the
// tag_group replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Group's GroupUserTags accordingly.
// Replaces o.R.GroupUserTags with related.
// Sets related.R.Group's GroupUserTags accordingly.
func (o *TagGroup) SetGroupUserTags(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*UserTag) error {
	query := "update `user_tag` set `group_id` = null where `group_id` = ?"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.GroupUserTags {
			queries.SetScanner(&rel.GroupID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Group = nil
		}
		o.R.GroupUserTags = nil
	}

	return o.AddGroupUserTags(ctx, exec, insert, related...)
}

// RemoveGroupUserTags relationships from objects passed in.
// Removes related items from R.GroupUserTags (uses pointer comparison, removal does not keep order)
// Sets related.R.Group.
func (o *TagGroup) RemoveGroupUserTags(ctx context.Context, exec boil.ContextExecutor, related ...*UserTag) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.GroupID, nil)
		if rel.R != nil {
			rel.R.Group = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("group_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.GroupUserTags {
			if rel != ri {
				continue
			}

			ln := len(o.R.GroupUserTags)
			if ln > 1 && i < ln-1 {
				o.R.GroupUserTags[i] = o.R.GroupUserTags[ln-1]
			}
			o.R.GroupUserTags = o.R.GroupUserTags[:ln-1]
			break
		}
	}

	return nil
}

// TagGroups retrieves all the records using an executor.
func TagGroups(mods ...qm.QueryMod) tagGroupQuery {
	mods = append(mods, qm.From("`tag_group`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`tag_group`.*"})
	}

	return tagGroupQuery{q}
}

// FindTagGroup retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTagGroup(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*TagGroup, error) {
	tagGroupObj := &TagGroup{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `tag_group` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, tagGroupObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from tag_group")
	}

	if err = tagGroupObj.doAfterSelectHooks(ctx, exec); err != nil {
		return tagGroupObj, err
	}

	return tagGroupObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TagGroup) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no tag_group provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tagGroupColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tagGroupInsertCacheMut.RLock()
	cache, cached := tagGroupInsertCache[key]
	tagGroupInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tagGroupAllColumns,
			tagGroupColumnsWithDefault,
			tagGroupColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(tagGroupType, tagGroupMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tagGroupType, tagGroupMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `tag_group` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `tag_group` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `tag_group` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, tagGroupPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into tag_group")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == tagGroupMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for tag_group")
	}

CacheNoHooks:
	if !cached {
		tagGroupInsertCacheMut.Lock()
		tagGroupInsertCache[key] = cache
		tagGroupInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TagGroup.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TagGroup) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	tagGroupUpdateCacheMut.RLock()
	cache, cached := tagGroupUpdateCache[key]
	tagGroupUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tagGroupAllColumns,
			tagGroupPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update tag_group, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `tag_group` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, tagGroupPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tagGroupType, tagGroupMapping, append(wl, tagGroupPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update tag_group row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for tag_group")
	}

	if !cached {
		tagGroupUpdateCacheMut.Lock()
		tagGroupUpdateCache[key] = cache
		tagGroupUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q tagGroupQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for tag_group")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for tag_group")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TagGroupSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tagGroupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `tag_group` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tagGroupPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in tagGroup slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all tagGroup")
	}
	return rowsAff, nil
}

var mySQLTagGroupUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TagGroup) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no tag_group provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tagGroupColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTagGroupUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	tagGroupUpsertCacheMut.RLock()
	cache, cached := tagGroupUpsertCache[key]
	tagGroupUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			tagGroupAllColumns,
			tagGroupColumnsWithDefault,
			tagGroupColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			tagGroupAllColumns,
			tagGroupPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert tag_group, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`tag_group`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `tag_group` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(tagGroupType, tagGroupMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tagGroupType, tagGroupMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for tag_group")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == tagGroupMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(tagGroupType, tagGroupMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for tag_group")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for tag_group")
	}

CacheNoHooks:
	if !cached {
		tagGroupUpsertCacheMut.Lock()
		tagGroupUpsertCache[key] = cache
		tagGroupUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single TagGroup record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TagGroup) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TagGroup provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tagGroupPrimaryKeyMapping)
	sql := "DELETE FROM `tag_group` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from tag_group")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for tag_group")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q tagGroupQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no tagGroupQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from tag_group")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for tag_group")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TagGroupSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(tagGroupBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tagGroupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `tag_group` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tagGroupPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from tagGroup slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for tag_group")
	}

	if len(tagGroupAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TagGroup) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTagGroup(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TagGroupSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TagGroupSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tagGroupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `tag_group`.* FROM `tag_group` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tagGroupPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TagGroupSlice")
	}

	*o = slice

	return nil
}

// TagGroupExists checks if the TagGroup row exists.
func TagGroupExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `tag_group` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if tag_group exists")
	}

	return exists, nil
}

// Exists checks if the TagGroup row exists.
func (o *TagGroup) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TagGroupExists(ctx, exec, o.ID)
}

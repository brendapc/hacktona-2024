// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// Business is an object representing the database table.
type Business struct {
	ID           int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name         string `boil:"name" json:"name" toml:"name" yaml:"name"`
	CNPJ         string `boil:"cnpj" json:"cnpj" toml:"cnpj" yaml:"cnpj"`
	Address      string `boil:"address" json:"address" toml:"address" yaml:"address"`
	Phone        string `boil:"phone" json:"phone" toml:"phone" yaml:"phone"`
	Email        string `boil:"email" json:"email" toml:"email" yaml:"email"`
	PasswordHash string `boil:"password_hash" json:"password_hash" toml:"password_hash" yaml:"password_hash"`
	Needs        string `boil:"needs" json:"needs" toml:"needs" yaml:"needs"`

	R *businessR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L businessL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BusinessColumns = struct {
	ID           string
	Name         string
	CNPJ         string
	Address      string
	Phone        string
	Email        string
	PasswordHash string
	Needs        string
}{
	ID:           "id",
	Name:         "name",
	CNPJ:         "cnpj",
	Address:      "address",
	Phone:        "phone",
	Email:        "email",
	PasswordHash: "password_hash",
	Needs:        "needs",
}

var BusinessTableColumns = struct {
	ID           string
	Name         string
	CNPJ         string
	Address      string
	Phone        string
	Email        string
	PasswordHash string
	Needs        string
}{
	ID:           "business.id",
	Name:         "business.name",
	CNPJ:         "business.cnpj",
	Address:      "business.address",
	Phone:        "business.phone",
	Email:        "business.email",
	PasswordHash: "business.password_hash",
	Needs:        "business.needs",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod   { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod  { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) ILIKE(x string) qm.QueryMod  { return qm.Where(w.field+" ILIKE ?", x) }
func (w whereHelperstring) NILIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT ILIKE ?", x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var BusinessWhere = struct {
	ID           whereHelperint
	Name         whereHelperstring
	CNPJ         whereHelperstring
	Address      whereHelperstring
	Phone        whereHelperstring
	Email        whereHelperstring
	PasswordHash whereHelperstring
	Needs        whereHelperstring
}{
	ID:           whereHelperint{field: "\"business\".\"id\""},
	Name:         whereHelperstring{field: "\"business\".\"name\""},
	CNPJ:         whereHelperstring{field: "\"business\".\"cnpj\""},
	Address:      whereHelperstring{field: "\"business\".\"address\""},
	Phone:        whereHelperstring{field: "\"business\".\"phone\""},
	Email:        whereHelperstring{field: "\"business\".\"email\""},
	PasswordHash: whereHelperstring{field: "\"business\".\"password_hash\""},
	Needs:        whereHelperstring{field: "\"business\".\"needs\""},
}

// BusinessRels is where relationship names are stored.
var BusinessRels = struct {
}{}

// businessR is where relationships are stored.
type businessR struct {
}

// NewStruct creates a new relationship struct
func (*businessR) NewStruct() *businessR {
	return &businessR{}
}

// businessL is where Load methods for each relationship are stored.
type businessL struct{}

var (
	businessAllColumns            = []string{"id", "name", "cnpj", "address", "phone", "email", "password_hash", "needs"}
	businessColumnsWithoutDefault = []string{"name", "cnpj", "address", "phone", "email", "password_hash", "needs"}
	businessColumnsWithDefault    = []string{"id"}
	businessPrimaryKeyColumns     = []string{"id"}
	businessGeneratedColumns      = []string{}
)

type (
	// BusinessSlice is an alias for a slice of pointers to Business.
	// This should almost always be used instead of []Business.
	BusinessSlice []*Business
	// BusinessHook is the signature for custom Business hook methods
	BusinessHook func(context.Context, boil.ContextExecutor, *Business) error

	businessQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	businessType                 = reflect.TypeOf(&Business{})
	businessMapping              = queries.MakeStructMapping(businessType)
	businessPrimaryKeyMapping, _ = queries.BindMapping(businessType, businessMapping, businessPrimaryKeyColumns)
	businessInsertCacheMut       sync.RWMutex
	businessInsertCache          = make(map[string]insertCache)
	businessUpdateCacheMut       sync.RWMutex
	businessUpdateCache          = make(map[string]updateCache)
	businessUpsertCacheMut       sync.RWMutex
	businessUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var businessAfterSelectMu sync.Mutex
var businessAfterSelectHooks []BusinessHook

var businessBeforeInsertMu sync.Mutex
var businessBeforeInsertHooks []BusinessHook
var businessAfterInsertMu sync.Mutex
var businessAfterInsertHooks []BusinessHook

var businessBeforeUpdateMu sync.Mutex
var businessBeforeUpdateHooks []BusinessHook
var businessAfterUpdateMu sync.Mutex
var businessAfterUpdateHooks []BusinessHook

var businessBeforeDeleteMu sync.Mutex
var businessBeforeDeleteHooks []BusinessHook
var businessAfterDeleteMu sync.Mutex
var businessAfterDeleteHooks []BusinessHook

var businessBeforeUpsertMu sync.Mutex
var businessBeforeUpsertHooks []BusinessHook
var businessAfterUpsertMu sync.Mutex
var businessAfterUpsertHooks []BusinessHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Business) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range businessAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Business) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range businessBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Business) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range businessAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Business) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range businessBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Business) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range businessAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Business) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range businessBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Business) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range businessAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Business) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range businessBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Business) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range businessAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBusinessHook registers your hook function for all future operations.
func AddBusinessHook(hookPoint boil.HookPoint, businessHook BusinessHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		businessAfterSelectMu.Lock()
		businessAfterSelectHooks = append(businessAfterSelectHooks, businessHook)
		businessAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		businessBeforeInsertMu.Lock()
		businessBeforeInsertHooks = append(businessBeforeInsertHooks, businessHook)
		businessBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		businessAfterInsertMu.Lock()
		businessAfterInsertHooks = append(businessAfterInsertHooks, businessHook)
		businessAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		businessBeforeUpdateMu.Lock()
		businessBeforeUpdateHooks = append(businessBeforeUpdateHooks, businessHook)
		businessBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		businessAfterUpdateMu.Lock()
		businessAfterUpdateHooks = append(businessAfterUpdateHooks, businessHook)
		businessAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		businessBeforeDeleteMu.Lock()
		businessBeforeDeleteHooks = append(businessBeforeDeleteHooks, businessHook)
		businessBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		businessAfterDeleteMu.Lock()
		businessAfterDeleteHooks = append(businessAfterDeleteHooks, businessHook)
		businessAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		businessBeforeUpsertMu.Lock()
		businessBeforeUpsertHooks = append(businessBeforeUpsertHooks, businessHook)
		businessBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		businessAfterUpsertMu.Lock()
		businessAfterUpsertHooks = append(businessAfterUpsertHooks, businessHook)
		businessAfterUpsertMu.Unlock()
	}
}

// One returns a single business record from the query.
func (q businessQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Business, error) {
	o := &Business{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for business")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Business records from the query.
func (q businessQuery) All(ctx context.Context, exec boil.ContextExecutor) (BusinessSlice, error) {
	var o []*Business

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Business slice")
	}

	if len(businessAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Business records in the query.
func (q businessQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count business rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q businessQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if business exists")
	}

	return count > 0, nil
}

// Businesses retrieves all the records using an executor.
func Businesses(mods ...qm.QueryMod) businessQuery {
	mods = append(mods, qm.From("\"business\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"business\".*"})
	}

	return businessQuery{q}
}

// FindBusiness retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBusiness(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Business, error) {
	businessObj := &Business{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"business\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, businessObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from business")
	}

	if err = businessObj.doAfterSelectHooks(ctx, exec); err != nil {
		return businessObj, err
	}

	return businessObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Business) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no business provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(businessColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	businessInsertCacheMut.RLock()
	cache, cached := businessInsertCache[key]
	businessInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			businessAllColumns,
			businessColumnsWithDefault,
			businessColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(businessType, businessMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(businessType, businessMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"business\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"business\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into business")
	}

	if !cached {
		businessInsertCacheMut.Lock()
		businessInsertCache[key] = cache
		businessInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Business.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Business) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	businessUpdateCacheMut.RLock()
	cache, cached := businessUpdateCache[key]
	businessUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			businessAllColumns,
			businessPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update business, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"business\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, businessPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(businessType, businessMapping, append(wl, businessPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update business row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for business")
	}

	if !cached {
		businessUpdateCacheMut.Lock()
		businessUpdateCache[key] = cache
		businessUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q businessQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for business")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for business")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BusinessSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), businessPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"business\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, businessPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in business slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all business")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Business) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no business provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(businessColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	businessUpsertCacheMut.RLock()
	cache, cached := businessUpsertCache[key]
	businessUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			businessAllColumns,
			businessColumnsWithDefault,
			businessColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			businessAllColumns,
			businessPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert business, could not build update column list")
		}

		ret := strmangle.SetComplement(businessAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(businessPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert business, could not build conflict column list")
			}

			conflict = make([]string, len(businessPrimaryKeyColumns))
			copy(conflict, businessPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"business\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(businessType, businessMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(businessType, businessMapping, ret)
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
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert business")
	}

	if !cached {
		businessUpsertCacheMut.Lock()
		businessUpsertCache[key] = cache
		businessUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Business record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Business) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Business provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), businessPrimaryKeyMapping)
	sql := "DELETE FROM \"business\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from business")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for business")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q businessQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no businessQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from business")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for business")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BusinessSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(businessBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), businessPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"business\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, businessPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from business slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for business")
	}

	if len(businessAfterDeleteHooks) != 0 {
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
func (o *Business) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBusiness(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BusinessSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BusinessSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), businessPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"business\".* FROM \"business\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, businessPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BusinessSlice")
	}

	*o = slice

	return nil
}

// BusinessExists checks if the Business row exists.
func BusinessExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"business\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if business exists")
	}

	return exists, nil
}

// Exists checks if the Business row exists.
func (o *Business) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return BusinessExists(ctx, exec, o.ID)
}

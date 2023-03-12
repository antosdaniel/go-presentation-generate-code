// Code generated by SQLBoiler 4.14.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// Payroll is an object representing the database table.
type Payroll struct {
	ID       string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	TenantID string    `boil:"tenant_id" json:"tenant_id" toml:"tenant_id" yaml:"tenant_id"`
	Payday   time.Time `boil:"payday" json:"payday" toml:"payday" yaml:"payday"`

	R *payrollR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L payrollL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PayrollColumns = struct {
	ID       string
	TenantID string
	Payday   string
}{
	ID:       "id",
	TenantID: "tenant_id",
	Payday:   "payday",
}

var PayrollTableColumns = struct {
	ID       string
	TenantID string
	Payday   string
}{
	ID:       "payrolls.id",
	TenantID: "payrolls.tenant_id",
	Payday:   "payrolls.payday",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
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

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var PayrollWhere = struct {
	ID       whereHelperstring
	TenantID whereHelperstring
	Payday   whereHelpertime_Time
}{
	ID:       whereHelperstring{field: "\"payrolls\".\"id\""},
	TenantID: whereHelperstring{field: "\"payrolls\".\"tenant_id\""},
	Payday:   whereHelpertime_Time{field: "\"payrolls\".\"payday\""},
}

// PayrollRels is where relationship names are stored.
var PayrollRels = struct {
	Payslips string
}{
	Payslips: "Payslips",
}

// payrollR is where relationships are stored.
type payrollR struct {
	Payslips PayslipSlice `boil:"Payslips" json:"Payslips" toml:"Payslips" yaml:"Payslips"`
}

// NewStruct creates a new relationship struct
func (*payrollR) NewStruct() *payrollR {
	return &payrollR{}
}

func (r *payrollR) GetPayslips() PayslipSlice {
	if r == nil {
		return nil
	}
	return r.Payslips
}

// payrollL is where Load methods for each relationship are stored.
type payrollL struct{}

var (
	payrollAllColumns            = []string{"id", "tenant_id", "payday"}
	payrollColumnsWithoutDefault = []string{"id", "tenant_id", "payday"}
	payrollColumnsWithDefault    = []string{}
	payrollPrimaryKeyColumns     = []string{"id"}
	payrollGeneratedColumns      = []string{}
)

type (
	// PayrollSlice is an alias for a slice of pointers to Payroll.
	// This should almost always be used instead of []Payroll.
	PayrollSlice []*Payroll

	payrollQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	payrollType                 = reflect.TypeOf(&Payroll{})
	payrollMapping              = queries.MakeStructMapping(payrollType)
	payrollPrimaryKeyMapping, _ = queries.BindMapping(payrollType, payrollMapping, payrollPrimaryKeyColumns)
	payrollInsertCacheMut       sync.RWMutex
	payrollInsertCache          = make(map[string]insertCache)
	payrollUpdateCacheMut       sync.RWMutex
	payrollUpdateCache          = make(map[string]updateCache)
	payrollUpsertCacheMut       sync.RWMutex
	payrollUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single payroll record from the query.
func (q payrollQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Payroll, error) {
	o := &Payroll{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for payrolls")
	}

	return o, nil
}

// All returns all Payroll records from the query.
func (q payrollQuery) All(ctx context.Context, exec boil.ContextExecutor) (PayrollSlice, error) {
	var o []*Payroll

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Payroll slice")
	}

	return o, nil
}

// Count returns the count of all Payroll records in the query.
func (q payrollQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count payrolls rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q payrollQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if payrolls exists")
	}

	return count > 0, nil
}

// Payslips retrieves all the payslip's Payslips with an executor.
func (o *Payroll) Payslips(mods ...qm.QueryMod) payslipQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"payslips\".\"payroll_id\"=?", o.ID),
	)

	return Payslips(queryMods...)
}

// LoadPayslips allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (payrollL) LoadPayslips(ctx context.Context, e boil.ContextExecutor, singular bool, maybePayroll interface{}, mods queries.Applicator) error {
	var slice []*Payroll
	var object *Payroll

	if singular {
		var ok bool
		object, ok = maybePayroll.(*Payroll)
		if !ok {
			object = new(Payroll)
			ok = queries.SetFromEmbeddedStruct(&object, &maybePayroll)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybePayroll))
			}
		}
	} else {
		s, ok := maybePayroll.(*[]*Payroll)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybePayroll)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybePayroll))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &payrollR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &payrollR{}
			}

			for _, a := range args {
				if a == obj.ID {
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
		qm.From(`payslips`),
		qm.WhereIn(`payslips.payroll_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load payslips")
	}

	var resultSlice []*Payslip
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice payslips")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on payslips")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for payslips")
	}

	if singular {
		object.R.Payslips = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &payslipR{}
			}
			foreign.R.Payroll = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.PayrollID {
				local.R.Payslips = append(local.R.Payslips, foreign)
				if foreign.R == nil {
					foreign.R = &payslipR{}
				}
				foreign.R.Payroll = local
				break
			}
		}
	}

	return nil
}

// AddPayslips adds the given related objects to the existing relationships
// of the payroll, optionally inserting them as new records.
// Appends related to o.R.Payslips.
// Sets related.R.Payroll appropriately.
func (o *Payroll) AddPayslips(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Payslip) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.PayrollID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"payslips\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"payroll_id"}),
				strmangle.WhereClause("\"", "\"", 2, payslipPrimaryKeyColumns),
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

			rel.PayrollID = o.ID
		}
	}

	if o.R == nil {
		o.R = &payrollR{
			Payslips: related,
		}
	} else {
		o.R.Payslips = append(o.R.Payslips, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &payslipR{
				Payroll: o,
			}
		} else {
			rel.R.Payroll = o
		}
	}
	return nil
}

// Payrolls retrieves all the records using an executor.
func Payrolls(mods ...qm.QueryMod) payrollQuery {
	mods = append(mods, qm.From("\"payrolls\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"payrolls\".*"})
	}

	return payrollQuery{q}
}

// FindPayroll retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPayroll(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Payroll, error) {
	payrollObj := &Payroll{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"payrolls\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, payrollObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from payrolls")
	}

	return payrollObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Payroll) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no payrolls provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(payrollColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	payrollInsertCacheMut.RLock()
	cache, cached := payrollInsertCache[key]
	payrollInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			payrollAllColumns,
			payrollColumnsWithDefault,
			payrollColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(payrollType, payrollMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(payrollType, payrollMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"payrolls\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"payrolls\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into payrolls")
	}

	if !cached {
		payrollInsertCacheMut.Lock()
		payrollInsertCache[key] = cache
		payrollInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Payroll.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Payroll) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	payrollUpdateCacheMut.RLock()
	cache, cached := payrollUpdateCache[key]
	payrollUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			payrollAllColumns,
			payrollPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update payrolls, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"payrolls\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, payrollPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(payrollType, payrollMapping, append(wl, payrollPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update payrolls row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for payrolls")
	}

	if !cached {
		payrollUpdateCacheMut.Lock()
		payrollUpdateCache[key] = cache
		payrollUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q payrollQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for payrolls")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for payrolls")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PayrollSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), payrollPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"payrolls\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, payrollPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in payroll slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all payroll")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Payroll) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no payrolls provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(payrollColumnsWithDefault, o)

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

	payrollUpsertCacheMut.RLock()
	cache, cached := payrollUpsertCache[key]
	payrollUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			payrollAllColumns,
			payrollColumnsWithDefault,
			payrollColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			payrollAllColumns,
			payrollPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert payrolls, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(payrollPrimaryKeyColumns))
			copy(conflict, payrollPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"payrolls\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(payrollType, payrollMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(payrollType, payrollMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert payrolls")
	}

	if !cached {
		payrollUpsertCacheMut.Lock()
		payrollUpsertCache[key] = cache
		payrollUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Payroll record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Payroll) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Payroll provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), payrollPrimaryKeyMapping)
	sql := "DELETE FROM \"payrolls\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from payrolls")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for payrolls")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q payrollQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no payrollQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from payrolls")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for payrolls")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PayrollSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), payrollPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"payrolls\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, payrollPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from payroll slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for payrolls")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Payroll) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPayroll(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PayrollSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PayrollSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), payrollPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"payrolls\".* FROM \"payrolls\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, payrollPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PayrollSlice")
	}

	*o = slice

	return nil
}

// PayrollExists checks if the Payroll row exists.
func PayrollExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"payrolls\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if payrolls exists")
	}

	return exists, nil
}

// Exists checks if the Payroll row exists.
func (o *Payroll) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return PayrollExists(ctx, exec, o.ID)
}
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent/item"
	"github.com/hay-kot/homebox/backend/internal/data/ent/maintenanceentry"
	"github.com/hay-kot/homebox/backend/internal/data/ent/predicate"
)

// MaintenanceEntryQuery is the builder for querying MaintenanceEntry entities.
type MaintenanceEntryQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.MaintenanceEntry
	withItem   *ItemQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MaintenanceEntryQuery builder.
func (meq *MaintenanceEntryQuery) Where(ps ...predicate.MaintenanceEntry) *MaintenanceEntryQuery {
	meq.predicates = append(meq.predicates, ps...)
	return meq
}

// Limit adds a limit step to the query.
func (meq *MaintenanceEntryQuery) Limit(limit int) *MaintenanceEntryQuery {
	meq.limit = &limit
	return meq
}

// Offset adds an offset step to the query.
func (meq *MaintenanceEntryQuery) Offset(offset int) *MaintenanceEntryQuery {
	meq.offset = &offset
	return meq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (meq *MaintenanceEntryQuery) Unique(unique bool) *MaintenanceEntryQuery {
	meq.unique = &unique
	return meq
}

// Order adds an order step to the query.
func (meq *MaintenanceEntryQuery) Order(o ...OrderFunc) *MaintenanceEntryQuery {
	meq.order = append(meq.order, o...)
	return meq
}

// QueryItem chains the current query on the "item" edge.
func (meq *MaintenanceEntryQuery) QueryItem() *ItemQuery {
	query := &ItemQuery{config: meq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := meq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := meq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(maintenanceentry.Table, maintenanceentry.FieldID, selector),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, maintenanceentry.ItemTable, maintenanceentry.ItemColumn),
		)
		fromU = sqlgraph.SetNeighbors(meq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MaintenanceEntry entity from the query.
// Returns a *NotFoundError when no MaintenanceEntry was found.
func (meq *MaintenanceEntryQuery) First(ctx context.Context) (*MaintenanceEntry, error) {
	nodes, err := meq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{maintenanceentry.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (meq *MaintenanceEntryQuery) FirstX(ctx context.Context) *MaintenanceEntry {
	node, err := meq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MaintenanceEntry ID from the query.
// Returns a *NotFoundError when no MaintenanceEntry ID was found.
func (meq *MaintenanceEntryQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = meq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{maintenanceentry.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (meq *MaintenanceEntryQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := meq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MaintenanceEntry entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MaintenanceEntry entity is found.
// Returns a *NotFoundError when no MaintenanceEntry entities are found.
func (meq *MaintenanceEntryQuery) Only(ctx context.Context) (*MaintenanceEntry, error) {
	nodes, err := meq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{maintenanceentry.Label}
	default:
		return nil, &NotSingularError{maintenanceentry.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (meq *MaintenanceEntryQuery) OnlyX(ctx context.Context) *MaintenanceEntry {
	node, err := meq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MaintenanceEntry ID in the query.
// Returns a *NotSingularError when more than one MaintenanceEntry ID is found.
// Returns a *NotFoundError when no entities are found.
func (meq *MaintenanceEntryQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = meq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{maintenanceentry.Label}
	default:
		err = &NotSingularError{maintenanceentry.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (meq *MaintenanceEntryQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := meq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MaintenanceEntries.
func (meq *MaintenanceEntryQuery) All(ctx context.Context) ([]*MaintenanceEntry, error) {
	if err := meq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return meq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (meq *MaintenanceEntryQuery) AllX(ctx context.Context) []*MaintenanceEntry {
	nodes, err := meq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MaintenanceEntry IDs.
func (meq *MaintenanceEntryQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := meq.Select(maintenanceentry.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (meq *MaintenanceEntryQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := meq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (meq *MaintenanceEntryQuery) Count(ctx context.Context) (int, error) {
	if err := meq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return meq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (meq *MaintenanceEntryQuery) CountX(ctx context.Context) int {
	count, err := meq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (meq *MaintenanceEntryQuery) Exist(ctx context.Context) (bool, error) {
	if err := meq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return meq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (meq *MaintenanceEntryQuery) ExistX(ctx context.Context) bool {
	exist, err := meq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MaintenanceEntryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (meq *MaintenanceEntryQuery) Clone() *MaintenanceEntryQuery {
	if meq == nil {
		return nil
	}
	return &MaintenanceEntryQuery{
		config:     meq.config,
		limit:      meq.limit,
		offset:     meq.offset,
		order:      append([]OrderFunc{}, meq.order...),
		predicates: append([]predicate.MaintenanceEntry{}, meq.predicates...),
		withItem:   meq.withItem.Clone(),
		// clone intermediate query.
		sql:    meq.sql.Clone(),
		path:   meq.path,
		unique: meq.unique,
	}
}

// WithItem tells the query-builder to eager-load the nodes that are connected to
// the "item" edge. The optional arguments are used to configure the query builder of the edge.
func (meq *MaintenanceEntryQuery) WithItem(opts ...func(*ItemQuery)) *MaintenanceEntryQuery {
	query := &ItemQuery{config: meq.config}
	for _, opt := range opts {
		opt(query)
	}
	meq.withItem = query
	return meq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MaintenanceEntry.Query().
//		GroupBy(maintenanceentry.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (meq *MaintenanceEntryQuery) GroupBy(field string, fields ...string) *MaintenanceEntryGroupBy {
	grbuild := &MaintenanceEntryGroupBy{config: meq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := meq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return meq.sqlQuery(ctx), nil
	}
	grbuild.label = maintenanceentry.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.MaintenanceEntry.Query().
//		Select(maintenanceentry.FieldCreatedAt).
//		Scan(ctx, &v)
func (meq *MaintenanceEntryQuery) Select(fields ...string) *MaintenanceEntrySelect {
	meq.fields = append(meq.fields, fields...)
	selbuild := &MaintenanceEntrySelect{MaintenanceEntryQuery: meq}
	selbuild.label = maintenanceentry.Label
	selbuild.flds, selbuild.scan = &meq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a MaintenanceEntrySelect configured with the given aggregations.
func (meq *MaintenanceEntryQuery) Aggregate(fns ...AggregateFunc) *MaintenanceEntrySelect {
	return meq.Select().Aggregate(fns...)
}

func (meq *MaintenanceEntryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range meq.fields {
		if !maintenanceentry.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if meq.path != nil {
		prev, err := meq.path(ctx)
		if err != nil {
			return err
		}
		meq.sql = prev
	}
	return nil
}

func (meq *MaintenanceEntryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MaintenanceEntry, error) {
	var (
		nodes       = []*MaintenanceEntry{}
		_spec       = meq.querySpec()
		loadedTypes = [1]bool{
			meq.withItem != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MaintenanceEntry).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MaintenanceEntry{config: meq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, meq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := meq.withItem; query != nil {
		if err := meq.loadItem(ctx, query, nodes, nil,
			func(n *MaintenanceEntry, e *Item) { n.Edges.Item = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (meq *MaintenanceEntryQuery) loadItem(ctx context.Context, query *ItemQuery, nodes []*MaintenanceEntry, init func(*MaintenanceEntry), assign func(*MaintenanceEntry, *Item)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*MaintenanceEntry)
	for i := range nodes {
		fk := nodes[i].ItemID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(item.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "item_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (meq *MaintenanceEntryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := meq.querySpec()
	_spec.Node.Columns = meq.fields
	if len(meq.fields) > 0 {
		_spec.Unique = meq.unique != nil && *meq.unique
	}
	return sqlgraph.CountNodes(ctx, meq.driver, _spec)
}

func (meq *MaintenanceEntryQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := meq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (meq *MaintenanceEntryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   maintenanceentry.Table,
			Columns: maintenanceentry.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: maintenanceentry.FieldID,
			},
		},
		From:   meq.sql,
		Unique: true,
	}
	if unique := meq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := meq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, maintenanceentry.FieldID)
		for i := range fields {
			if fields[i] != maintenanceentry.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := meq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := meq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := meq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := meq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (meq *MaintenanceEntryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(meq.driver.Dialect())
	t1 := builder.Table(maintenanceentry.Table)
	columns := meq.fields
	if len(columns) == 0 {
		columns = maintenanceentry.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if meq.sql != nil {
		selector = meq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if meq.unique != nil && *meq.unique {
		selector.Distinct()
	}
	for _, p := range meq.predicates {
		p(selector)
	}
	for _, p := range meq.order {
		p(selector)
	}
	if offset := meq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := meq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MaintenanceEntryGroupBy is the group-by builder for MaintenanceEntry entities.
type MaintenanceEntryGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (megb *MaintenanceEntryGroupBy) Aggregate(fns ...AggregateFunc) *MaintenanceEntryGroupBy {
	megb.fns = append(megb.fns, fns...)
	return megb
}

// Scan applies the group-by query and scans the result into the given value.
func (megb *MaintenanceEntryGroupBy) Scan(ctx context.Context, v any) error {
	query, err := megb.path(ctx)
	if err != nil {
		return err
	}
	megb.sql = query
	return megb.sqlScan(ctx, v)
}

func (megb *MaintenanceEntryGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range megb.fields {
		if !maintenanceentry.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := megb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := megb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (megb *MaintenanceEntryGroupBy) sqlQuery() *sql.Selector {
	selector := megb.sql.Select()
	aggregation := make([]string, 0, len(megb.fns))
	for _, fn := range megb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(megb.fields)+len(megb.fns))
		for _, f := range megb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(megb.fields...)...)
}

// MaintenanceEntrySelect is the builder for selecting fields of MaintenanceEntry entities.
type MaintenanceEntrySelect struct {
	*MaintenanceEntryQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mes *MaintenanceEntrySelect) Aggregate(fns ...AggregateFunc) *MaintenanceEntrySelect {
	mes.fns = append(mes.fns, fns...)
	return mes
}

// Scan applies the selector query and scans the result into the given value.
func (mes *MaintenanceEntrySelect) Scan(ctx context.Context, v any) error {
	if err := mes.prepareQuery(ctx); err != nil {
		return err
	}
	mes.sql = mes.MaintenanceEntryQuery.sqlQuery(ctx)
	return mes.sqlScan(ctx, v)
}

func (mes *MaintenanceEntrySelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(mes.fns))
	for _, fn := range mes.fns {
		aggregation = append(aggregation, fn(mes.sql))
	}
	switch n := len(*mes.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		mes.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		mes.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := mes.sql.Query()
	if err := mes.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
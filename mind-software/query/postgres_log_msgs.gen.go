// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"mind-software/pgmodels"
)

func newPostgresLogMsg(db *gorm.DB, opts ...gen.DOOption) postgresLogMsg {
	_postgresLogMsg := postgresLogMsg{}

	_postgresLogMsg.postgresLogMsgDo.UseDB(db, opts...)
	_postgresLogMsg.postgresLogMsgDo.UseModel(&pgmodels.PostgresLogMsg{})

	tableName := _postgresLogMsg.postgresLogMsgDo.TableName()
	_postgresLogMsg.ALL = field.NewAsterisk(tableName)
	_postgresLogMsg.ID = field.NewUint64(tableName, "id")
	_postgresLogMsg.ObjID = field.NewUint64(tableName, "obj_id")
	_postgresLogMsg.Timestamp = field.NewUint64(tableName, "id")
	_postgresLogMsg.Log = field.NewString(tableName, "id")
	_postgresLogMsg.CreatedAt = field.NewUint64(tableName, "created_at")
	_postgresLogMsg.UpdatedAt = field.NewUint64(tableName, "updated_at")
	_postgresLogMsg.DeletedAt = field.NewField(tableName, "deleted_at")

	_postgresLogMsg.fillFieldMap()

	return _postgresLogMsg
}

type postgresLogMsg struct {
	postgresLogMsgDo

	ALL       field.Asterisk
	ID        field.Uint64
	ObjID     field.Uint64
	Timestamp field.Uint64
	Log       field.String
	CreatedAt field.Uint64
	UpdatedAt field.Uint64
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (p postgresLogMsg) Table(newTableName string) *postgresLogMsg {
	p.postgresLogMsgDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p postgresLogMsg) As(alias string) *postgresLogMsg {
	p.postgresLogMsgDo.DO = *(p.postgresLogMsgDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *postgresLogMsg) updateTableName(table string) *postgresLogMsg {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint64(table, "id")
	p.ObjID = field.NewUint64(table, "obj_id")
	p.Timestamp = field.NewUint64(table, "id")
	p.Log = field.NewString(table, "id")
	p.CreatedAt = field.NewUint64(table, "created_at")
	p.UpdatedAt = field.NewUint64(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")

	p.fillFieldMap()

	return p
}

func (p *postgresLogMsg) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *postgresLogMsg) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 7)
	p.fieldMap["id"] = p.ID
	p.fieldMap["obj_id"] = p.ObjID
	p.fieldMap["id"] = p.Timestamp
	p.fieldMap["id"] = p.Log
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
}

func (p postgresLogMsg) clone(db *gorm.DB) postgresLogMsg {
	p.postgresLogMsgDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p postgresLogMsg) replaceDB(db *gorm.DB) postgresLogMsg {
	p.postgresLogMsgDo.ReplaceDB(db)
	return p
}

type postgresLogMsgDo struct{ gen.DO }

type IPostgresLogMsgDo interface {
	gen.SubQuery
	Debug() IPostgresLogMsgDo
	WithContext(ctx context.Context) IPostgresLogMsgDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPostgresLogMsgDo
	WriteDB() IPostgresLogMsgDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPostgresLogMsgDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPostgresLogMsgDo
	Not(conds ...gen.Condition) IPostgresLogMsgDo
	Or(conds ...gen.Condition) IPostgresLogMsgDo
	Select(conds ...field.Expr) IPostgresLogMsgDo
	Where(conds ...gen.Condition) IPostgresLogMsgDo
	Order(conds ...field.Expr) IPostgresLogMsgDo
	Distinct(cols ...field.Expr) IPostgresLogMsgDo
	Omit(cols ...field.Expr) IPostgresLogMsgDo
	Join(table schema.Tabler, on ...field.Expr) IPostgresLogMsgDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPostgresLogMsgDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPostgresLogMsgDo
	Group(cols ...field.Expr) IPostgresLogMsgDo
	Having(conds ...gen.Condition) IPostgresLogMsgDo
	Limit(limit int) IPostgresLogMsgDo
	Offset(offset int) IPostgresLogMsgDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPostgresLogMsgDo
	Unscoped() IPostgresLogMsgDo
	Create(values ...*pgmodels.PostgresLogMsg) error
	CreateInBatches(values []*pgmodels.PostgresLogMsg, batchSize int) error
	Save(values ...*pgmodels.PostgresLogMsg) error
	First() (*pgmodels.PostgresLogMsg, error)
	Take() (*pgmodels.PostgresLogMsg, error)
	Last() (*pgmodels.PostgresLogMsg, error)
	Find() ([]*pgmodels.PostgresLogMsg, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pgmodels.PostgresLogMsg, err error)
	FindInBatches(result *[]*pgmodels.PostgresLogMsg, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*pgmodels.PostgresLogMsg) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPostgresLogMsgDo
	Assign(attrs ...field.AssignExpr) IPostgresLogMsgDo
	Joins(fields ...field.RelationField) IPostgresLogMsgDo
	Preload(fields ...field.RelationField) IPostgresLogMsgDo
	FirstOrInit() (*pgmodels.PostgresLogMsg, error)
	FirstOrCreate() (*pgmodels.PostgresLogMsg, error)
	FindByPage(offset int, limit int) (result []*pgmodels.PostgresLogMsg, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPostgresLogMsgDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p postgresLogMsgDo) Debug() IPostgresLogMsgDo {
	return p.withDO(p.DO.Debug())
}

func (p postgresLogMsgDo) WithContext(ctx context.Context) IPostgresLogMsgDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p postgresLogMsgDo) ReadDB() IPostgresLogMsgDo {
	return p.Clauses(dbresolver.Read)
}

func (p postgresLogMsgDo) WriteDB() IPostgresLogMsgDo {
	return p.Clauses(dbresolver.Write)
}

func (p postgresLogMsgDo) Session(config *gorm.Session) IPostgresLogMsgDo {
	return p.withDO(p.DO.Session(config))
}

func (p postgresLogMsgDo) Clauses(conds ...clause.Expression) IPostgresLogMsgDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p postgresLogMsgDo) Returning(value interface{}, columns ...string) IPostgresLogMsgDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p postgresLogMsgDo) Not(conds ...gen.Condition) IPostgresLogMsgDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p postgresLogMsgDo) Or(conds ...gen.Condition) IPostgresLogMsgDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p postgresLogMsgDo) Select(conds ...field.Expr) IPostgresLogMsgDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p postgresLogMsgDo) Where(conds ...gen.Condition) IPostgresLogMsgDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p postgresLogMsgDo) Order(conds ...field.Expr) IPostgresLogMsgDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p postgresLogMsgDo) Distinct(cols ...field.Expr) IPostgresLogMsgDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p postgresLogMsgDo) Omit(cols ...field.Expr) IPostgresLogMsgDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p postgresLogMsgDo) Join(table schema.Tabler, on ...field.Expr) IPostgresLogMsgDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p postgresLogMsgDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPostgresLogMsgDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p postgresLogMsgDo) RightJoin(table schema.Tabler, on ...field.Expr) IPostgresLogMsgDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p postgresLogMsgDo) Group(cols ...field.Expr) IPostgresLogMsgDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p postgresLogMsgDo) Having(conds ...gen.Condition) IPostgresLogMsgDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p postgresLogMsgDo) Limit(limit int) IPostgresLogMsgDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p postgresLogMsgDo) Offset(offset int) IPostgresLogMsgDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p postgresLogMsgDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPostgresLogMsgDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p postgresLogMsgDo) Unscoped() IPostgresLogMsgDo {
	return p.withDO(p.DO.Unscoped())
}

func (p postgresLogMsgDo) Create(values ...*pgmodels.PostgresLogMsg) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p postgresLogMsgDo) CreateInBatches(values []*pgmodels.PostgresLogMsg, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p postgresLogMsgDo) Save(values ...*pgmodels.PostgresLogMsg) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p postgresLogMsgDo) First() (*pgmodels.PostgresLogMsg, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*pgmodels.PostgresLogMsg), nil
	}
}

func (p postgresLogMsgDo) Take() (*pgmodels.PostgresLogMsg, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*pgmodels.PostgresLogMsg), nil
	}
}

func (p postgresLogMsgDo) Last() (*pgmodels.PostgresLogMsg, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*pgmodels.PostgresLogMsg), nil
	}
}

func (p postgresLogMsgDo) Find() ([]*pgmodels.PostgresLogMsg, error) {
	result, err := p.DO.Find()
	return result.([]*pgmodels.PostgresLogMsg), err
}

func (p postgresLogMsgDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*pgmodels.PostgresLogMsg, err error) {
	buf := make([]*pgmodels.PostgresLogMsg, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p postgresLogMsgDo) FindInBatches(result *[]*pgmodels.PostgresLogMsg, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p postgresLogMsgDo) Attrs(attrs ...field.AssignExpr) IPostgresLogMsgDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p postgresLogMsgDo) Assign(attrs ...field.AssignExpr) IPostgresLogMsgDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p postgresLogMsgDo) Joins(fields ...field.RelationField) IPostgresLogMsgDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p postgresLogMsgDo) Preload(fields ...field.RelationField) IPostgresLogMsgDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p postgresLogMsgDo) FirstOrInit() (*pgmodels.PostgresLogMsg, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*pgmodels.PostgresLogMsg), nil
	}
}

func (p postgresLogMsgDo) FirstOrCreate() (*pgmodels.PostgresLogMsg, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*pgmodels.PostgresLogMsg), nil
	}
}

func (p postgresLogMsgDo) FindByPage(offset int, limit int) (result []*pgmodels.PostgresLogMsg, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p postgresLogMsgDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p postgresLogMsgDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p postgresLogMsgDo) Delete(models ...*pgmodels.PostgresLogMsg) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *postgresLogMsgDo) withDO(do gen.Dao) *postgresLogMsgDo {
	p.DO = *do.(*gen.DO)
	return p
}
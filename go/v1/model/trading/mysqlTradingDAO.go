package trading

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/fkmhrk/OpenInvoice/v1/entity"
	"github.com/fkmhrk/OpenInvoice/v1/model"
	"github.com/fkmhrk/OpenInvoice/v1/model/db"
	"github.com/fkmhrk/OpenInvoice/v1/model/logger"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

const (
	tradingTableName     = "trading"
	sqlSelectAllTradings = "SELECT id,company_id,title_type,subject," +
		"work_from,work_to,total," +
		"quotation_date,quotation_number," +
		"bill_date,bill_number," +
		"delivery_date,delivery_number," +
		"tax_rate,assignee,product,memo," +
		"created_time, modified_time " +
		"FROM " + tradingTableName + " "
	sqlSelectTradingsByUser = sqlSelectAllTradings +
		"WHERE assignee=? AND deleted <> 1 ORDER BY id ASC"
	sqlSelectTradingList = sqlSelectAllTradings +
		"WHERE deleted <> 1 ORDER BY modified_time DESC"
	sqlSelectTradingByID = sqlSelectAllTradings +
		"WHERE id=? AND deleted <> 1 LIMIT 1"
	sqlInsertTrading = "INSERT INTO " + tradingTableName +
		"(id,company_id,subject,title_type," +
		"work_from,work_to,total," +
		"quotation_date,quotation_number," +
		"bill_date,bill_number," +
		"delivery_date,delivery_number," +
		"tax_rate,assignee,product,memo," +
		"created_time,modified_time,deleted)" +
		"VALUES(?,?,?,?," +
		"?,?,?," +
		"?,''," +
		"?,''," +
		"?,''," +
		"?,?,?,?," +
		"unix_timestamp(),unix_timestamp(),0)"
	sqlUpdateTrading = "UPDATE " + tradingTableName + " " +
		"SET company_id=?,title_type=?,subject=?," +
		"work_from=?,work_to=?,total=?," +
		"quotation_date=?,quotation_number=?," +
		"bill_date=?,bill_number=?," +
		"delivery_date=?,delivery_number=?," +
		"tax_rate=?,assignee=?,product=?,memo=?," +
		"modified_time=unix_timestamp() " +
		"WHERE id=? AND deleted <> 1"
	sqlSoftDeleteTrading = "UPDATE " + tradingTableName + " " +
		"SET deleted=1 " +
		"WHERE id=? AND deleted <> 1"

	sqlSelectTradingID = "SELECT num FROM trading_id WHERE date=?"
	sqlInsertTradingID = "INSERT INTO trading_id" +
		"(date,num) VALUES(?,1)"
	sqlUpdateTradingID = "UPDATE trading_id " +
		"SET num=? WHERE date=?"

	tradingItemTableName      = "trading_item"
	sqlSelectTradingItemsByID = "SELECT id,sort_order,subject,unit_price,amount," +
		"degree,tax_type,memo " +
		"FROM " + tradingItemTableName + " " +
		"WHERE trading_id=? AND deleted <> 1 ORDER BY sort_order ASC"
	sqlInsertTradingItem = "INSERT INTO " + tradingItemTableName +
		"(id,trading_id,sort_order,subject," +
		"unit_price,amount,degree," +
		"tax_type,memo," +
		"created_time,modified_time,deleted)" +
		"VALUES(?,?,?,?," +
		"?,?,?," +
		"?,?," +
		"unix_timestamp(),unix_timestamp(),0)"
	sqlUpdateTradingItem = "UPDATE " + tradingItemTableName + " " +
		"SET sort_order=?,subject=?," +
		"unit_price=?,amount=?,degree=?," +
		"tax_type=?,memo=?," +
		"modified_time=unix_timestamp() " +
		"WHERE id=? AND trading_id=? AND deleted <> 1"
	sqlSoftDeleteItem = "UPDATE " + tradingItemTableName + " " +
		"SET deleted=1 " +
		"WHERE id=? AND trading_id=? AND deleted <> 1"
	sqlSoftDeleteAllItem = "UPDATE " + tradingItemTableName + " " +
		"SET deleted=1 " +
		"WHERE trading_id=? AND deleted <> 1"
)

type tradingDAO struct {
	connection *db.Connection
	logger     logger.Logger
}

// New creates instance
func New(connection *db.Connection, logger logger.Logger) model.Trading {
	return &tradingDAO{
		connection: connection,
		logger:     logger,
	}
}

func (d *tradingDAO) GetList() ([]*entity.Trading, error) {
	db := d.connection.Connect()
	st, err := db.Prepare(sqlSelectTradingList)
	if err != nil {
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*entity.Trading
	for rows.Next() {
		item := d.scanTrading(rows)
		list = append(list, &item)
	}
	return list, nil
}

func (d *tradingDAO) GetListByUser(userId string) ([]*entity.Trading, error) {
	db := d.connection.Connect()
	st, err := db.Prepare(sqlSelectTradingsByUser)
	if err != nil {
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query(userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*entity.Trading
	for rows.Next() {
		item := d.scanTrading(rows)
		list = append(list, &item)
	}
	return list, nil

}

func (d *tradingDAO) GetById(id string) (*entity.Trading, error) {
	db := d.connection.Connect()
	st, err := db.Prepare(sqlSelectTradingByID)
	if err != nil {
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	item := d.scanTrading(rows)
	return &item, nil
}

func (d *tradingDAO) Create(companyId, subject string, titleType int, workFrom, workTo, total, quotationDate, billDate, deliveryDate int64, taxRate float32, assignee, product, memo string) (*entity.Trading, error) {
	tr, err := d.connection.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Rollback()

	st, err := tr.Prepare(sqlInsertTrading)
	if err != nil {
		return nil, err
	}
	defer st.Close()

	id, err := InsertWithUUID(32, func(id string) error {
		_, err = st.Exec(id, companyId, subject, titleType,
			workFrom, workTo, total,
			quotationDate,
			billDate,
			deliveryDate,
			taxRate, assignee, product, memo)
		return err
	})
	if err != nil {
		return nil, err
	}

	tr.Commit()

	return &entity.Trading{
		Id:            id,
		CompanyId:     companyId,
		Subject:       subject,
		TitleType:     titleType,
		WorkFrom:      workFrom,
		WorkTo:        workTo,
		QuotationDate: quotationDate,
		BillDate:      billDate,
		TaxRate:       taxRate,
		AssigneeId:    assignee,
		Product:       product,
	}, nil
}

func (d *tradingDAO) Update(trading entity.Trading) (*entity.Trading, error) {
	tr, err := d.connection.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Rollback()

	st, err := tr.Prepare(sqlUpdateTrading)
	if err != nil {
		return nil, err
	}
	defer st.Close()

	_, err = st.Exec(trading.CompanyId, trading.TitleType, trading.Subject,
		trading.WorkFrom, trading.WorkTo, trading.Total,
		trading.QuotationDate, trading.QuotationNumber,
		trading.BillDate, trading.BillNumber,
		trading.DeliveryDate, trading.DeliveryNumber,
		trading.TaxRate, trading.AssigneeId, trading.Product, trading.Memo,
		trading.Id)
	if err != nil {
		return nil, err
	}

	tr.Commit()

	return &trading, nil
}

func (d *tradingDAO) Delete(id string) error {
	tr, err := d.connection.Begin()
	if err != nil {
		return err
	}
	defer tr.Rollback()
	// deletes item
	err = d.softAllDeleteItem(tr, id)
	if err != nil {
		return err
	}
	// delete trading
	err = d.softDelete(tr, id)
	if err != nil {
		return err
	}
	tr.Commit()
	return nil
}

func (d *tradingDAO) GetItemsById(tradingId string) ([]*entity.TradingItem, error) {
	db := d.connection.Connect()
	st, err := db.Prepare(sqlSelectTradingItemsByID)
	if err != nil {
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query(tradingId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*entity.TradingItem
	var id, subject, degree, memo string
	var sortOrder, unitPrice, taxType int
	var amount float64
	for rows.Next() {
		rows.Scan(&id, &sortOrder, &subject, &unitPrice, &amount,
			&degree, &taxType, &memo)

		list = append(list, &entity.TradingItem{
			Id:        id,
			TradingId: tradingId,
			SortOrder: sortOrder,
			Subject:   subject,
			UnitPrice: unitPrice,
			Amount:    amount,
			Degree:    degree,
			TaxType:   taxType,
			Memo:      memo,
		})
	}
	return list, nil

}

func (d *tradingDAO) CreateItem(tradingId, subject, degree, memo string, sortOrder, unitPrice int, amount float64, taxType int) (*entity.TradingItem, error) {
	tr, err := d.connection.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Rollback()

	st, err := tr.Prepare(sqlInsertTradingItem)
	if err != nil {
		return nil, err
	}
	defer st.Close()

	// generate ID
	var id string
	for i := 0; i < 10; i++ {
		id = generateUUID(32)
		_, err = st.Exec(id, tradingId, sortOrder, subject, unitPrice, amount,
			degree, taxType, memo)
		if err == nil {
			break
		}
		id = ""
		if err2, ok := err.(*mysql.MySQLError); ok {
			if err2.Number != 1062 {
				return nil, err2
			}
		} else {
			return nil, err
		}
	}
	if len(id) == 0 {
		return nil, errors.New("Failed to create")
	}

	tr.Commit()

	return &entity.TradingItem{
		Id:        id,
		TradingId: tradingId,
		SortOrder: sortOrder,
		Subject:   subject,
		UnitPrice: unitPrice,
		Amount:    amount,
		Degree:    degree,
		TaxType:   taxType,
		Memo:      memo,
	}, nil
}

func (d *tradingDAO) UpdateItem(id, tradingId, subject, degree, memo string, sortOrder, unitPrice int, amount float64, taxType int) (*entity.TradingItem, error) {
	tr, err := d.connection.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Rollback()

	st, err := tr.Prepare(sqlUpdateTradingItem)
	if err != nil {
		return nil, err
	}
	defer st.Close()

	// execute
	_, err = st.Exec(sortOrder, subject, unitPrice, amount,
		degree, taxType, memo, id, tradingId)
	if err != nil {
		return nil, err
	}

	tr.Commit()

	return &entity.TradingItem{
		Id:        id,
		TradingId: tradingId,
		SortOrder: sortOrder,
		Subject:   subject,
		UnitPrice: unitPrice,
		Amount:    amount,
		Degree:    degree,
		TaxType:   taxType,
		Memo:      memo,
	}, nil
}

func (d *tradingDAO) SoftDeleteItem(id, tradingId string) error {
	tr, err := d.connection.Begin()
	if err != nil {
		return err
	}
	defer tr.Rollback()

	st, err := tr.Prepare(sqlSoftDeleteItem)
	if err != nil {
		return err
	}
	defer st.Close()

	// execute
	_, err = st.Exec(id, tradingId)
	if err != nil {
		return err
	}

	tr.Commit()
	return nil
}

func (d *tradingDAO) generateNextId(tr *sql.Tx, date string) (string, error) {
	num, err := d.getId(tr, date)
	if err != nil {
		return "", err
	}

	if num == -1 {
		err = d.insertId(tr, date)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s001", date), nil
	} else {
		num += 1
		err = d.updateId(tr, date, num)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s%03d", date, num), nil
	}
}

func (d *tradingDAO) getId(tr *sql.Tx, date string) (int, error) {
	st, err := tr.Prepare(sqlSelectTradingID)
	if err != nil {
		return -1, err
	}
	defer st.Close()

	rows, err := st.Query(date)
	if err != nil {
		return -1, err
	}
	defer rows.Close()

	if !rows.Next() {
		return -1, nil
	}

	var num int
	rows.Scan(&num)
	return num, nil
}

func (d *tradingDAO) insertId(tr *sql.Tx, date string) error {
	st, err := tr.Prepare(sqlInsertTradingID)
	if err != nil {
		return err
	}
	defer st.Close()

	_, err = st.Exec(date)
	return err
}

func (d *tradingDAO) updateId(tr *sql.Tx, date string, num int) error {
	st, err := tr.Prepare(sqlUpdateTradingID)
	if err != nil {
		return err
	}
	defer st.Close()

	_, err = st.Exec(num, date)
	return err
}

func (d *tradingDAO) scanTrading(rows *sql.Rows) entity.Trading {
	var id, companyId, subject, product string
	var titleType int
	var taxRate float32
	var assignee, quotationNumber, billNumber, deliveryNumber, memo string
	var workFrom, workTo, total, quotationDate, billDate, deliveryDate, created, modified int64

	err := rows.Scan(&id, &companyId, &titleType, &subject,
		&workFrom, &workTo, &total,
		&quotationDate, &quotationNumber,
		&billDate, &billNumber,
		&deliveryDate, &deliveryNumber,
		&taxRate, &assignee, &product, &memo,
		&created, &modified)
	if err != nil {
		d.logger.Errorf("Failed to scan trading :%s", err)
	}

	return entity.Trading{
		Id:              id,
		CompanyId:       companyId,
		TitleType:       titleType,
		Subject:         subject,
		WorkFrom:        workFrom,
		WorkTo:          workTo,
		Total:           total,
		QuotationDate:   quotationDate,
		QuotationNumber: quotationNumber,
		BillDate:        billDate,
		BillNumber:      billNumber,
		DeliveryDate:    deliveryDate,
		DeliveryNumber:  deliveryNumber,
		TaxRate:         taxRate,
		AssigneeId:      assignee,
		Product:         product,
		Memo:            memo,
		CreatedTime:     created,
		ModifiedTime:    modified,
	}
}

func (d *tradingDAO) softDelete(tr *sql.Tx, tradingId string) error {
	st, err := tr.Prepare(sqlSoftDeleteTrading)
	if err != nil {
		return err
	}
	defer st.Close()

	// execute
	_, err = st.Exec(tradingId)
	if err != nil {
		return err
	}
	return nil
}

func (d *tradingDAO) softAllDeleteItem(tr *sql.Tx, tradingId string) error {
	st, err := tr.Prepare(sqlSoftDeleteAllItem)
	if err != nil {
		return err
	}
	defer st.Close()

	// execute
	_, err = st.Exec(tradingId)
	if err != nil {
		return err
	}
	return nil
}

func generateUUID(length int) string {
	id1 := uuid.New().String()
	id2 := uuid.New().String()
	id := strings.Replace(id1+id2, "-", "", -1)
	return id[:length]
}

func InsertWithUUID(l int, f func(id string) error) (string, error) {
	for i := 0; i < 10; i++ {
		id := generateUUID(l)
		err := f(id)
		if err == nil {
			return id, nil
		}
		if err2, ok := err.(*mysql.MySQLError); ok {
			if err2.Number != 1062 {
				return "", err2
			}
		} else {
			return "", err
		}

	}
	return "", errors.New("Failed to insert 10 times.")
}

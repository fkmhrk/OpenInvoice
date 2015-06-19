package impl

import (
	m "../"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

const (
	select_trading = "SELECT id,company_id,title_type,subject," +
		"work_from,work_to,total," +
		"quotation_date,quotation_number," +
		"bill_date,bill_number," +
		"tax_rate,assignee,product," +
		"created_time, modified_time FROM trading"
)

type tradingDAO struct {
	connection *Connection
	logger     m.Logger
}

func NewTradingDAO(connection *Connection, logger m.Logger) *tradingDAO {
	return &tradingDAO{
		connection: connection,
		logger:     logger,
	}
}

func (d *tradingDAO) GetListByUser(userId string) ([]*m.Trading, error) {
	db := d.connection.Connect()
	st, err := db.Prepare(select_trading +
		" WHERE assignee=? AND deleted <> 1 ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query(userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*m.Trading
	for rows.Next() {
		item := d.scanTrading(rows)
		list = append(list, &item)
	}
	return list, nil

}

func (d *tradingDAO) GetById(id, userId string) (*m.Trading, error) {
	db := d.connection.Connect()
	st, err := db.Prepare(select_trading +
		" WHERE id=? AND assignee=? AND deleted <> 1 LIMIT 1")
	if err != nil {
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query(id, userId)
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

func (d *tradingDAO) Create(companyId, subject string, titleType int, workFrom, workTo, total, quotationDate, billDate int64, taxRate float32, assignee, product string) (*m.Trading, error) {
	tr, err := d.connection.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Rollback()

	st, err := tr.Prepare("INSERT INTO trading(" +
		"id,company_id,subject,title_type," +
		"work_from,work_to,total," +
		"quotation_date,quotation_number," +
		"bill_date,bill_number," +
		"tax_rate,assignee,product," +
		"created_time,modified_time,deleted)" +
		"VALUES(?,?,?,?," +
		"?,?,?," +
		"?,''," +
		"?,''," +
		"?,?,?," +
		"unix_timestamp(now()),unix_timestamp(now()),0)")
	if err != nil {
		return nil, err
	}
	defer st.Close()

	id, err := insertWithUUID(32, func(id string) error {
		_, err = st.Exec(id, companyId, subject, titleType,
			workFrom, workTo, total,
			quotationDate,
			billDate,
			taxRate, assignee, product)
		return err
	})
	if err != nil {
		return nil, err
	}

	tr.Commit()

	return &m.Trading{
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

func (d *tradingDAO) Update(trading m.Trading) (*m.Trading, error) {
	tr, err := d.connection.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Rollback()

	st, err := tr.Prepare("UPDATE trading SET " +
		"company_id=?,title_type=?,subject=?," +
		"work_from=?,work_to=?,total=?," +
		"quotation_date=?,quotation_number=?," +
		"bill_date=?,bill_number=?," +
		"tax_rate=?,assignee=?,product=?," +
		"modified_time=unix_timestamp(now()) " +
		"WHERE id=? AND deleted <> 1")
	if err != nil {
		return nil, err
	}
	defer st.Close()

	_, err = st.Exec(trading.CompanyId, trading.TitleType, trading.Subject,
		trading.WorkFrom, trading.WorkTo, trading.Total,
		trading.QuotationDate, trading.QuotationNumber,
		trading.BillDate, trading.BillNumber,
		trading.TaxRate, trading.AssigneeId, trading.Product, trading.Id)
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

func (d *tradingDAO) GetItemsById(tradingId string) ([]*m.TradingItem, error) {
	db := d.connection.Connect()
	st, err := db.Prepare("SELECT id,sort_order,subject,unit_price,amount," +
		"degree,tax_type,memo FROM trading_item " +
		"WHERE trading_id=? AND deleted <> 1 ORDER BY sort_order ASC")
	if err != nil {
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query(tradingId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*m.TradingItem
	var id, subject, degree, memo string
	var sortOrder, unitPrice, amount, taxType int
	for rows.Next() {
		rows.Scan(&id, &sortOrder, &subject, &unitPrice, &amount,
			&degree, &taxType, &memo)

		list = append(list, &m.TradingItem{
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

func (d *tradingDAO) CreateItem(tradingId, subject, degree, memo string, sortOrder, unitPrice, amount, taxType int) (*m.TradingItem, error) {
	tr, err := d.connection.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Rollback()

	st, err := tr.Prepare("INSERT INTO trading_item(" +
		"id,trading_id,sort_order,subject," +
		"unit_price,amount,degree," +
		"tax_type,memo," +
		"created_time,modified_time,deleted)" +
		"VALUES(?,?,?,?," +
		"?,?,?," +
		"?,?," +
		"unix_timestamp(now()),unix_timestamp(now()),0)")
	if err != nil {
		return nil, err
	}
	defer st.Close()

	// generate ID
	var id string
	for i := 0; i < 10; i++ {
		id = generateId(32)
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

	return &m.TradingItem{
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

func (d *tradingDAO) UpdateItem(id, tradingId, subject, degree, memo string, sortOrder, unitPrice, amount, taxType int) (*m.TradingItem, error) {
	tr, err := d.connection.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Rollback()

	st, err := tr.Prepare("UPDATE trading_item SET " +
		"sort_order=?,subject=?," +
		"unit_price=?,amount=?,degree=?," +
		"tax_type=?,memo=?," +
		"modified_time=unix_timestamp(now()) " +
		"WHERE id=? AND trading_id=? AND deleted <> 1")
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

	return &m.TradingItem{
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

	st, err := tr.Prepare("UPDATE trading_item SET " +
		"deleted=1 " +
		"WHERE id=? AND trading_id=? AND deleted <> 1")
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
	st, err := tr.Prepare("SELECT num FROM trading_id WHERE date=?")
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
	st, err := tr.Prepare("INSERT INTO trading_id" +
		"(date,num) VALUES(?,1)")
	if err != nil {
		return err
	}
	defer st.Close()

	_, err = st.Exec(date)
	return err
}

func (d *tradingDAO) updateId(tr *sql.Tx, date string, num int) error {
	st, err := tr.Prepare("UPDATe trading_id " +
		"SET num=? WHERE date=?")
	if err != nil {
		return err
	}
	defer st.Close()

	_, err = st.Exec(num, date)
	return err
}

func (d *tradingDAO) scanTrading(rows *sql.Rows) m.Trading {
	var id, companyId, subject, product string
	var titleType int
	var taxRate float32
	var assignee, quotationNumber, billNumber string
	var workFrom, workTo, total, quotationDate, billDate, created, modified int64

	err := rows.Scan(&id, &companyId, &titleType, &subject,
		&workFrom, &workTo, &total,
		&quotationDate, &quotationNumber,
		&billDate, &billNumber,
		&taxRate, &assignee, &product,
		&created, &modified)
	if err != nil {
		d.logger.Errorf("Failed to scan trading :%s", err)
	}

	return m.Trading{
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
		TaxRate:         taxRate,
		AssigneeId:      assignee,
		Product:         product,
		CreatedTime:     created,
		ModifiedTime:    modified,
	}
}

func (d *tradingDAO) softDelete(tr *sql.Tx, tradingId string) error {
	st, err := tr.Prepare("UPDATE trading SET " +
		"deleted=1 " +
		"WHERE id=? AND deleted <> 1")
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
	st, err := tr.Prepare("UPDATE trading_item SET " +
		"deleted=1 " +
		"WHERE trading_id=? AND deleted <> 1")
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

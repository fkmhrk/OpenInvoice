package impl

import (
	m "../"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

type tradingDAO struct {
	connection *Connection
}

func NewTradingDAO(connection *Connection) *tradingDAO {
	return &tradingDAO{
		connection: connection,
	}
}

func (d *tradingDAO) GetListByUser(userId string) ([]*m.Trading, error) {
	db := d.connection.Connect()
	st, err := db.Prepare("SELECT id,company_id,title_type,subject," +
		"work_from,work_to,quotation_date,bill_date," +
		"tax_rate,product," +
		"created_time, modified_time FROM trading " +
		"WHERE assignee=? AND deleted <> 1 ORDER BY id ASC")
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
	var id, companyId, subject, product string
	var titleType int
	var taxRate float32
	var workFrom, workTo, quotationDate, billDate, created, modified int64
	for rows.Next() {
		rows.Scan(&id, &companyId, &titleType, &subject,
			&workFrom, &workTo, &quotationDate,
			&billDate, &taxRate, &product,
			&created, &modified)

		list = append(list, &m.Trading{
			Id:            id,
			CompanyId:     companyId,
			TitleType:     titleType,
			Subject:       subject,
			WorkFrom:      workFrom,
			WorkTo:        workTo,
			QuotationDate: quotationDate,
			BillDate:      billDate,
			TaxRate:       taxRate,
			AssigneeId:    userId,
			Product:       product,
			CreatedTime:   created,
			ModifiedTime:  modified,
		})
	}
	return list, nil

}

func (d *tradingDAO) GetById(id, userId string) (*m.Trading, error) {
	db := d.connection.Connect()
	st, err := db.Prepare("SELECT company_id,title_type,subject," +
		"work_from,work_to,quotation_date,bill_date," +
		"tax_rate,product," +
		"created_time, modified_time FROM trading " +
		"WHERE id=? AND assignee=? AND deleted <> 1 LIMIT 1")
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

	var companyId, subject, product string
	var titleType int
	var taxRate float32
	var workFrom, workTo, quotationDate, billDate, created, modified int64
	rows.Scan(&companyId, &titleType, &subject, &workFrom, &workTo,
		&quotationDate, &billDate, &taxRate, &product, &created, &modified)

	return &m.Trading{
		Id:            id,
		CompanyId:     companyId,
		TitleType:     titleType,
		Subject:       subject,
		WorkFrom:      workFrom,
		WorkTo:        workTo,
		QuotationDate: quotationDate,
		BillDate:      billDate,
		TaxRate:       taxRate,
		AssigneeId:    userId,
		Product:       product,
		CreatedTime:   created,
		ModifiedTime:  modified,
	}, nil
}

func (d *tradingDAO) Create(date, companyId, subject string, titleType int, workFrom, workTo, quotationDate, billDate int64, taxRate float32, assignee, product string) (*m.Trading, error) {
	tr, err := d.connection.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Rollback()
	// generate ID
	id, err := d.generateNextId(tr, date)
	if err != nil {
		return nil, err
	}

	st, err := tr.Prepare("INSERT INTO trading(" +
		"id,company_id,subject,title_type," +
		"work_from,work_to,quotation_date,bill_date," +
		"tax_rate,assignee,product," +
		"created_time,modified_time,deleted)" +
		"VALUES(?,?,?,?," +
		"?,?,?,?," +
		"?,?,?," +
		"unix_timestamp(now()),unix_timestamp(now()),0)")
	if err != nil {
		return nil, err
	}
	defer st.Close()

	_, err = st.Exec(id, companyId, subject, titleType,
		workFrom, workTo, quotationDate, billDate,
		taxRate, assignee, product)
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

func (d *tradingDAO) Update(id, companyId, subject string, titleType int, workFrom, workTo, quotationDate, billDate int64, taxRate float32, assignee, product string) (*m.Trading, error) {
	tr, err := d.connection.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Rollback()

	st, err := tr.Prepare("UPDATE trading SET " +
		"company_id=?,title_type=?,subject=?," +
		"work_from=?,work_to=?,quotation_date=?,bill_date=?," +
		"tax_rate=?,assignee=?,product=?," +
		"modified_time=unix_timestamp(now()) " +
		"WHERE id=? AND deleted <> 1")
	if err != nil {
		return nil, err
	}
	defer st.Close()

	_, err = st.Exec(companyId, titleType, subject, workFrom, workTo,
		quotationDate, billDate, taxRate,
		assignee, product, id)
	if err != nil {
		return nil, err
	}

	tr.Commit()

	return &m.Trading{
		Id:            id,
		CompanyId:     companyId,
		Subject:       subject,
		WorkFrom:      workFrom,
		WorkTo:        workTo,
		QuotationDate: quotationDate,
		BillDate:      billDate,
		TaxRate:       taxRate,
		AssigneeId:    assignee,
		Product:       product,
	}, nil
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

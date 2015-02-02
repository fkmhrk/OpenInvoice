package impl

import (
	m "../"
	_ "errors"
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
	st, err := db.Prepare("SELECT id,company_id,subject," +
		"work_from,work_to,product," +
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
	var workFrom, workTo, created, modified int64
	for rows.Next() {
		rows.Scan(&id, &companyId, &subject, &workFrom, &workTo,
			&product, &created, &modified)

		list = append(list, &m.Trading{
			Id:           id,
			CompanyId:    companyId,
			Subject:      subject,
			WorkFrom:     workFrom,
			WorkTo:       workTo,
			AssigneeId:   userId,
			Product:      product,
			CreatedTime:  created,
			ModifiedTime: modified,
		})
	}
	return list, nil

}

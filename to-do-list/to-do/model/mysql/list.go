package mysql

import (
  "database/sql"
  "r24xtech.net/to-do/model"
)

type ListModel struct {
  DB *sql.DB
}

func (m *ListModel) Insert(item string) error {
  stmt := `INSERT INTO list (item, created) VALUES(?, UTC_TIMESTAMP())`
  _, err := m.DB.Exec(stmt, item)
  if err != nil {
    return err
  }
  return nil
}

func (m *ListModel) Delete(id int) error {
  stmt := `DELETE FROM list WHERE id = ?`
  _, err := m.DB.Exec(stmt, id)
  if err != nil {
    return err
  }
  return nil
}

func (m *ListModel) Latest() ([]*model.ItemList, error) {
  stmt := `SELECT id, item, created FROM list`
  rows, err := m.DB.Query(stmt)
  if err != nil {
    return nil, err
  }
  defer rows.Close()
  items := []*model.ItemList{}
  for rows.Next() {
    item := &model.ItemList{}
    err = rows.Scan(&item.ID, &item.Item, &item.Created)
    if err != nil {
      return nil, err
    }
    items = append(items, item)
  }
  if err = rows.Err(); err != nil {
    return nil, err
  }
  return items, nil
}

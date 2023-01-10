package dao

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"goadmin/dao/mysql"
)

type Permission struct {
	Id              string         `json:"id"  db:"id"`
	Pid             string         `json:"pid" db:"pid"`
	Name            string         `json:"name" db:"name"`
	Type            int            `json:"type" db:"type"`
	PermissionValue sql.NullString `json:"permission_value" db:"permission_value"`
	Path            sql.NullString `json:"path" db:"path"`
	Component       sql.NullString `json:"component" db:"component"`
	Icon            sql.NullString `json:"icon" db:"icon"`
	Status          sql.NullInt64  `json:"status" db:"status"`
	Level           int            `json:"level" db:"level"`
	Children        []*Permission  `json:"children" db:"children"`
	IsSelect        bool           `json:"is_select" db:"is_select"`
	IsDeleted       bool           `json:"is_deleted" db:"is_deleted"`
	GmtCreate       string         `json:"gmt_create" db:"gmt_create"`
	GmtModified     string         `json:"gmt_modified" db:"gmt_modified"`
}

type RolePermission struct {
	Id           string `json:"id"  db:"id"`
	RoleId       string `json:"role_id"  db:"role_id"`
	PermissionId string `json:"permission_id" db:"permission_id"`
	IsDeleted    int    `json:"is_deleted" db:"is_deleted"`
	GmtCreate    string `json:"gmt_create" db:"gmt_create"`
	GmtModified  string `json:"gmt_modified" db:"gmt_modified"`
}

func GetAllMenu() ([]*Permission, error) {

	db := mysql.GetDb()
	sqlstr := "select id,pid,name,type,permission_value,path,component,icon,status,is_deleted from acl_permission "
	list := make([]*Permission, 0, 10)
	err := db.Select(&list, sqlstr)
	if err != nil {
		return nil, err
	}
	return list, err
}

func GetPermissionByID(id string) ([]string, error) {

	ids := []string{}
	db := mysql.GetDb()
	sqlstr := "select id from acl_permission where pid=?"
	err := db.Select(&ids, sqlstr, id)
	if err != nil {
		return nil, err
	}
	return ids, err
}

func DelPermissionByids(ids []string) error {
	db := mysql.GetDb()
	query, args, err := sqlx.In("delete from acl_permission where id in (?)", ids)
	if err != nil {
		return err
	}
	_, err = db.Exec(query, args...)
	if err != nil {
		return err
	}
	return err
}

func SaveRolepermissionIds(rpids []*RolePermission) error {
	db := mysql.GetDb()
	_, err := db.NamedExec("INSERT INTO acl_role_permission (id,role_id,permission_id,is_deleted,gmt_create,gmt_modified) "+
		"VALUES (:id,:role_id,:permission_id,:is_deleted,:gmt_create,:gmt_modified)", rpids)
	if err != nil {
		return err
	}
	return nil
}

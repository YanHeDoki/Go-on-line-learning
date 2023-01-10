package handles

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goadmin/constants"
	"goadmin/dao"
	"goadmin/utils"
	"net/http"
	"time"
)

func buildPermission(p []*dao.Permission) []*dao.Permission {

	list := make([]*dao.Permission, 0, len(p))

	for _, v := range p {
		if v.Pid == "0" {
			v.Level = 1
			list = append(list, selectChildren(v, p))
		}
	}
	return list
}
func selectChildren(node *dao.Permission, p []*dao.Permission) *dao.Permission {

	node.Children = make([]*dao.Permission, 0, 5)

	for _, v := range p {
		if node.Id == v.Pid {
			v.Level = node.Level + 1
			node.Children = append(node.Children, selectChildren(v, p))
		}
	}
	return node
}

func GetAllMenu(c *gin.Context) {

	resp := constants.ReCode{}
	allMenu, err := dao.GetAllMenu()
	if err != nil {
		fmt.Println("GetAllMenu err", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}

	list := buildPermission(allMenu)

	resp.Ok()
	resp.Data = gin.H{"children": list}
	c.JSON(http.StatusOK, resp)
}

func DelAllMenu(c *gin.Context) {

	id := c.Param("id")
	resp := constants.ReCode{}
	ids, err := DelPermission(id)
	if err != nil {
		fmt.Println("DelPermission err", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = dao.DelPermissionByids(ids)
	if err != nil {
		fmt.Println("DelPermissionByids err", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	c.JSON(http.StatusOK, resp)
}

func DelPermission(id string) ([]string, error) {
	ids, err := SelectDelPermissionIds(id)
	if err != nil {
		return nil, err
	}
	ids = append(ids, id)
	return ids, nil
}

func SelectDelPermissionIds(id string) ([]string, error) {
	res := []string{}
	ids, err := dao.GetPermissionByID(id)
	if err != nil {
		return nil, err
	}
	for _, v := range ids {
		//res = append(res, v)
		//permissionIds, _ := SelectDelPermissionIds(v)
		//res = append(res, permissionIds...)
		permissionIds, _ := SelectDelPermissionIds(v)
		res = append(res, append([]string{v}, permissionIds...)...)
	}
	return res, err
}

func DoAssign(c *gin.Context) {
	rid := c.PostForm("roleId")
	pids := c.PostFormArray("permissionId")
	resp := constants.ReCode{}
	err := SaveRolepermissionIds(rid, pids)
	if err != nil {
		fmt.Println("SaveRolepermissionIds err", err)
		resp.Err()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Ok()
	c.JSON(http.StatusOK, resp)
}

func SaveRolepermissionIds(rid string, pids []string) error {

	list := make([]*dao.RolePermission, 0, 10)
	date := time.Now()
	datestr := date.Format("2006-01-02 15:04:05")
	for _, v := range pids {
		p := &dao.RolePermission{
			Id:           utils.GenID(),
			RoleId:       rid,
			PermissionId: v,
			IsDeleted:    0,
			GmtCreate:    datestr,
			GmtModified:  datestr,
		}
		list = append(list, p)
	}
	return dao.SaveRolepermissionIds(list)
}

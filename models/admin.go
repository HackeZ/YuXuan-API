package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"

	"YuXuanAPI/utils"
)

// YxAdmin is Admin of YuXuan-Shop
type YxAdmin struct {
	ID          int       `orm:"column(_id);auto"`
	LoginName   string    `orm:"column(login_name);size(255)"`
	Password    string    `orm:"column(password);size(32)"`
	Salt        string    `orm:"column(salt);size(16)"`
	PhoneNumber string    `orm:"column(phone_number);size(16);null"`
	Email       string    `orm:"column(email);size(32);null"`
	CreateDate  time.Time `orm:"column(create_date);type(datetime);null"`
}

// TableName Retuen TableName
func (t *YxAdmin) TableName() string {
	return "yx_admin"
}

func init() {
	orm.RegisterModel(new(YxAdmin))
}

// AddYxAdmin Insert a New YxAdmin into datebase and returns
// last inserted Id on success.
func AddYxAdmin(m *YxAdmin) (id int64, err error) {
	o := orm.NewOrm()
	m.Salt = utils.GetSalt()
	m.Password = utils.MD5(m.Password + m.Salt)
	id, err = o.Insert(m)
	return
}

// GetYxAdminByID retrieves YxAdmin by id. Retuens error if
// Id does't exist.
func GetYxAdminByID(id int) (v *YxAdmin, err error) {
	o := orm.NewOrm()
	v = &YxAdmin{ID: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllYxAdmin retrieves all YxAdmin matches certain condition. Returns empty list if
// no records exist
func GetAllYxAdmin(query map[string]string, fields, sortby, order []string, offset, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(YxAdmin))
	// query k = v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k := strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []YxAdmin
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			//trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateYxAdminByID updates YxAdmin by Id and returns error if
// the record to be updated doesn't exist
func UpdateYxAdminByID(m *YxAdmin) (err error) {
	o := orm.NewOrm()
	v := YxAdmin{ID: m.ID}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteYxAdmin deletes YxAdmin by Id and returns error if
// the record to be deleted doesn't exist
func DeleteYxAdmin(id int) (err error) {
	o := orm.NewOrm()
	v := YxAdmin{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&YxAdmin{ID: id}); err == nil {
			fmt.Println("Number of records deleted in datebase: ", num)
		}
	}
	return
}

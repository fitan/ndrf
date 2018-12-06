package ndrf

import (
	"github.com/jinzhu/gorm"
	"reflect"
	"strings"
)
type Mapping struct {
	PageOpen bool
	PageMaxLimit int
	PageDefaultLimit int
}

type ActionSqler interface {
	UseDb(db *gorm.DB)
	GetSql(requestOj interface{}, models interface{}) (err error)
	GetPageSql(mapping *Mapping,mappingForm *MappingForm,requestOj interface{}, models interface{}) (err error)
	GetIdSql(id int, model interface{}) (err error)
	PutIdSql(id int, requestOj interface{}, model interface{}) (err error)
	PostSql(model interface{}) (err error)
	DeleteIdSql(id int, model interface{}) (err error)
}


type ActionSql struct {
	db *gorm.DB
}

func (this *ActionSql) UseDb(db *gorm.DB)  {
	this.db = db
}

func (this *ActionSql) GetSql(requestOj interface{}, models interface{}) (err error) {
	refT := reflect.TypeOf(requestOj).Elem()
	refV := reflect.ValueOf(requestOj).Elem()
	querySql := []string{}
	queryList := []interface{}{}
	for i:=0; i<refT.NumField(); i++ {
		form := refT.Field(i).Tag.Get("form")
		if form == "" {
			continue
		} else {
			value := refV.Field(i).Interface()
			if value == "" {
				continue
			}
			switch refT.Field(i).Tag.Get("query") {
			case "", "=":
				querySql = append(querySql, form + " = ? ")
				queryList = append(queryList, value)
			case "@":
				querySql = append(querySql, form + " LIKE ? ")
				queryList = append(queryList, "%" + value.(string) + "%")
			case "^":
				querySql = append(querySql, form + " LIKE ? ")
				queryList = append(queryList, value.(string) + "%")
			case ">":
				querySql = append(querySql, form + " > ? ")
				queryList = append(queryList, value)
			case "<":
				querySql = append(querySql, form + " < ? ")
				queryList = append(queryList, value)
			default:
				querySql = append(querySql, form + " = ? ")
				queryList = append(queryList, value)
			}
		}
	}
	sql := strings.Join(querySql, " AND ")
	if sql == "" {
		err = this.db.Find(models).Error
		return
	}
	err = this.db.Where(sql, queryList...).Find(models).Error
	return
}

func (this *ActionSql) GetPageSql(mapping *Mapping,mappingForm *MappingForm, requestOj interface{}, models interface{}) (err error) {
	refT := reflect.TypeOf(requestOj).Elem()
	refV := reflect.ValueOf(requestOj).Elem()
	querySql := []string{}
	queryList := []interface{}{}
	for i:=0; i<refT.NumField(); i++ {
		form := refT.Field(i).Tag.Get("form")
		if form == "" {
			continue
		} else {
			value := refV.Field(i).Interface()
			if value == "" {
				continue
			}
			switch refT.Field(i).Tag.Get("query") {
			case "", "=":
				querySql = append(querySql, form + " = ? ")
				queryList = append(queryList, value)
			case "@":
				querySql = append(querySql, form + " LIKE ? ")
				queryList = append(queryList, "%" + value.(string) + "%")
			case "^":
				querySql = append(querySql, form + " LIKE ? ")
				queryList = append(queryList, value.(string) + "%")
			case ">":
				querySql = append(querySql, form + " > ? ")
				queryList = append(queryList, value)
			case "<":
				querySql = append(querySql, form + " < ? ")
				queryList = append(queryList, value)
			default:
				querySql = append(querySql, form + " = ? ")
				queryList = append(queryList, value)
			}
		}
	}
	sql := strings.Join(querySql, " AND ")
	if mappingForm.Limit == 0 {
		mappingForm.Limit = mapping.PageDefaultLimit
	}

	if mappingForm.Limit > mapping.PageMaxLimit {
		mappingForm.Limit = mapping.PageMaxLimit
	}

	if sql == "" {
		err = this.db.Limit(mappingForm.Limit).Offset((mappingForm.Page -1)*mappingForm.Limit).Find(models).Error
		return
	}
	err = this.db.Limit(mappingForm.Limit).Offset((mappingForm.Page -1)*mappingForm.Limit).Where(sql, queryList...).Find(models).Error
	return
}

func (this *ActionSql) GetIdSql(id int, model interface{}) (err error) {
	err = this.db.First(model, id).Error
	return
}

func (this *ActionSql) PutIdSql(id int, requestOj interface{}, model interface{}) (err error) {
	err = this.db.Model(model).Where("ID = ?", id).Updates(requestOj).Error
	if err != nil {
		return
	}
	err = this.db.First(model, id).Error
	return
}

func (this *ActionSql) PostSql(model interface{}) (err error) {
	err = this.db.Create(model).Error
	return
}

func (this *ActionSql) DeleteIdSql(id int, model interface{}) (err error) {
	err = this.db.First(model, id).Error
	if err != nil {
		return
	}
	err = this.db.Where("ID = ?", id).Delete(model).Error
	return
}

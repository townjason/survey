package model

import (
	. "api/database/mysql"
	"api/util/log"
	"database/sql"
)

type QuestionnaireQrcodeRecordAdmin struct {
	Id             int64     `table:"id"`
	QrcodeRecordId int64     `table:"qrcode_record_id"`
	AdminId        int64     `table:"admin_id"`
	AdminName      string
}

func (model *QuestionnaireQrcodeRecordAdmin) SetId(id int64) *QuestionnaireQrcodeRecordAdmin {
	model.Id = id
	return model
}

func (model *QuestionnaireQrcodeRecordAdmin) SetQrcodeRecordId(qrcodeRecordId int64) *QuestionnaireQrcodeRecordAdmin {
	model.QrcodeRecordId = qrcodeRecordId
	return model
}

func (model *QuestionnaireQrcodeRecordAdmin) SetAdminId(adminId int64) *QuestionnaireQrcodeRecordAdmin {
	model.AdminId = adminId
	return model
}

func (model *QuestionnaireQrcodeRecordAdmin) Create() (int64, error) {
	return Model(model).Insert()
}

func (model *QuestionnaireQrcodeRecordAdmin) Update(columns []string) error {
	return Model(model).Where("id", "=", model.Id).Update(columns)
}

func (model *QuestionnaireQrcodeRecordAdmin) Delete() error {
	table := Model(model)
	if model.Id > 0 {
		table.Where("id", "=", model.Id)
	}
	if model.QrcodeRecordId > 0 {
		table.Where("qrcode_record_id", "=", model.QrcodeRecordId)
	}
	if model.AdminId > 0 {
		table.Where("admin_id", "=", model.AdminId)
	}
	return table.Delete()
}

func (model *QuestionnaireQrcodeRecordAdmin) QueryAll(option func(*QuestionnaireQrcodeRecordAdmin)) {
	table := Model(model)

	if model.Id > 0 {
		table.Where("id", "=", model.Id)
	}
	if model.QrcodeRecordId > 0 {
		table.Where("qrcode_record_id", "=", model.QrcodeRecordId)
	}
	if model.AdminId > 0 {
		table.Where("admin_id", "=", model.AdminId)
	}

	table.Select([]string{"id", "qrcode_record_id", "admin_id"}).
		Get(func(rows *sql.Rows) (isBreak bool) {
			err := rows.Scan(&model.Id, &model.QrcodeRecordId, &model.AdminId)
			log.Error(err)
			if err == nil && option != nil {
				option(model)
			}
			return
		})
}

func (model *QuestionnaireQrcodeRecordAdmin) QueryOne() *QuestionnaireQrcodeRecordAdmin {
	table := Model(model)

	if model.Id > 0 {
		table.Where("id", "=", model.Id)
	}
	if model.QrcodeRecordId > 0 {
		table.Where("qrcode_record_id", "=", model.QrcodeRecordId)
	}
	if model.AdminId > 0 {
		table.Where("admin_id", "=", model.AdminId)
	}

	log.Error(table.Select([]string{"id", "qrcode_record_id", "admin_id"}).
		Find().Scan(&model.Id, &model.QrcodeRecordId, &model.AdminId))
	return model
}

func (model *QuestionnaireQrcodeRecordAdmin) QueryAllWithName(option func(*QuestionnaireQrcodeRecordAdmin)) {
	table := Model(model)

	if model.Id > 0 {
		table.Where("id", "=", model.Id)
	}
	if model.QrcodeRecordId > 0 {
		table.Where("qrcode_record_id", "=", model.QrcodeRecordId)
	}
	if model.AdminId > 0 {
		table.Where("admin_id", "=", model.AdminId)
	}

	table.Select([]string{"id", "qrcode_record_id", "admin_id", "admins.nick_name"}).
		InnerJoin("admins", "admins.id", "=", "questionnaire_qrcode_record_admin.admin_id").
		Get(func(rows *sql.Rows) (isBreak bool) {
			err := rows.Scan(&model.Id, &model.QrcodeRecordId, &model.AdminId, &model.AdminName)
			log.Error(err)
			if err == nil && option != nil {
				option(model)
			}
			return
		})
}

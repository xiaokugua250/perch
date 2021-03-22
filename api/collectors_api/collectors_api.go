package collectors_api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	database "perch/database/mysql"

	"perch/web/metric"
	"perch/web/model"
	"perch/web/model/applications"
)

/**

创建collectors

*/
func CollectoresRegisterHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			application    applications.Applications
			isAlreadyExist int64

			err error
		)
		response.Kind = "application"
		if err = json.NewDecoder(r.Body).Decode(&application); err != nil {
			response.Code = http.StatusBadRequest
			return err
		}

		if err = database.MysqlDb.Model(applications.Applications{}).Where("name=?", application.Name).Count(&isAlreadyExist).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = application
			return err
		}
		if isAlreadyExist > 0 {
			return errors.New(fmt.Sprintf(" application with name %s already exits..."))
		}
		if err = database.MysqlDb.Create(&application).Error; err != nil {
			return err
		}

		response.Spec = application

		response.Code = http.StatusOK
		response.Total = 1
		response.Message = "create applications successfully !!!"
		return nil
	})
}

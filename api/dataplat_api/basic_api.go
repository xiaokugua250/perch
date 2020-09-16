package dataplat

import (
	"context"
	_ "fmt"
	"net/http"
	database "perch/database/mysql"
	"perch/web/metric"
	"perch/web/model"
	resource "perch/web/model/resources"
)

func PlatDataResourcesHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			resourceArticles []resource.ResourceArticle
			err         error
		)
		response.Kind = "articles"

		if err = database.MySQL_DB.First(&resourceArticles).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = err.Error()
			return err
		}
		if err = database.MySQL_DB.Model(resource.ResourceArticle{}).Count(&response.Total).Error;err!= nil{
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = err.Error()
			return err
		}


		response.Code = http.StatusOK
		response.Spec = resourceArticles
		response.Message = " get resources articles successfully !!!"
		return nil
	})
}




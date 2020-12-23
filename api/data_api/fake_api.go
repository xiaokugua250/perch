/**
API
假数据生成  https://godoc.org/github.com/icrowley/fake
ref
*/
package data_api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/icrowley/fake"
	"perch/web/metric"
	"perch/web/model"
)

func GetFakeUsersHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			//resourceDocs []resource.ResourceBlogs
			//err          error
			userInfo = make(map[string]string)
		)
		response.Kind = "fake users"

		userInfo["user_name"] = fake.UserName()
		userInfo["passwd"] = fake.Password(8, 16, true, true, true)
		userInfo["title"] = fake.Title()
		userInfo["user_agent"] = fake.UserAgent()
		userInfo["gender"] = fake.Gender()
		userInfo["email"] = fake.EmailAddress()
		userInfo["company"] = fake.Company()
		userInfo["country"] = fake.Country()
		userInfo["job_title"] = fake.JobTitle()
		userInfo["ipv4"] = fake.IPv4()
		userInfo["ipv6"] = fake.IPv6()
		userInfo["addrees"] = fake.Continent() + "-" + fake.Country() + "-" + fake.State() + "-" + fake.City() + "-" + fake.StreetAddress()

		response.Code = http.StatusOK
		response.Spec = userInfo
		response.Message = " get fake user info successfully !!!"
		return nil
	})
}

func GetFakeEmailsHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			//resourceDocs []resource.ResourceBlogs
			//err          error
			emailInfo = make(map[string]string)
		)
		response.Kind = "fake emails"

		emailInfo["email_address"] = fake.EmailAddress()
		emailInfo["email_body"] = fake.EmailBody()
		emailInfo["email_subject"] = fake.EmailSubject()

		response.Code = http.StatusOK
		response.Spec = emailInfo
		response.Message = " get fake email info successfully !!!"
		return nil
	})
}

func GetFakeCrediCardHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			//resourceDocs []resource.ResourceBlogs
			//err          error
			creditCardInfo = make(map[string]string)
		)
		response.Kind = "fake credicard"

		creditCardInfo["credit_type"] = fake.CreditCardType()
		creditCardInfo["credit_num"] = fake.CreditCardNum("")

		response.Code = http.StatusOK
		response.Spec = creditCardInfo
		response.Message = " get fake  creditCardInfo  successfully !!!"
		return nil
	})
}

func GetFakeIPHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			//resourceDocs []resource.ResourceBlogs
			//err          error
			IpInfo = make(map[string]string)
		)
		response.Kind = "fake users"

		IpInfo["ipv4"] = fake.IPv4()
		IpInfo["ipv6"] = fake.IPv6()
		IpInfo["domain_name"] = fake.DomainName()
		IpInfo["domain_zone"] = fake.DomainZone()

		response.Code = http.StatusOK
		response.Spec = IpInfo
		response.Message = " get fake ip info successfully !!!"
		return nil
	})
}

func GetFakeLocHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			//resourceDocs []resource.ResourceBlogs
			//err          error
			locationInfo = make(map[string]string)
		)
		response.Kind = "fake location"

		locationInfo["cpuntry"] = fake.Country()
		locationInfo["city"] = fake.City()
		locationInfo["street_address"] = fake.StreetAddress()
		locationInfo["street"] = fake.Street()

		locationInfo["latitude"] = fmt.Sprintf("%f", fake.Latitude())
		locationInfo["latitude_degrees"] = fmt.Sprintf("%d", fake.LatitudeDegrees())
		locationInfo["latitude_direction"] = fake.LatitudeDirection()
		locationInfo["latitude_minutes"] = fmt.Sprintf("%d", fake.LatitudeMinutes())
		locationInfo["latitude_seconds"] = strconv.Itoa( fake.LatitudeSeconds())
		locationInfo["longitude_drees"] = fmt.Sprintf("%d", fake.LongitudeDegrees())
		locationInfo["longitude"] = fmt.Sprintf("%f", fake.Longitude())
		locationInfo["longitude_direction"] = fake.LongitudeDirection()
		locationInfo["longitude_minutes"] = fmt.Sprintf("%d", fake.LongitudeMinutes())
		locationInfo["longitude_second"] = fmt.Sprintf("%d", fake.LongitudeSeconds())

		response.Code = http.StatusOK
		response.Spec = locationInfo
		response.Message = " get fake location info successfully !!!"
		return nil
	})
}
func GetFakeTimesHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			//resourceDocs []resource.ResourceBlogs
			//err          error
			timeInfo = make(map[string]string)
		)
		response.Kind = "fake users"

		timeInfo["day"] = strconv.Itoa(fake.Day())
		timeInfo["weekday"] = fake.WeekDay()
		timeInfo["weekdaynum"] = strconv.Itoa(fake.WeekdayNum())
		timeInfo["dayshort"] = fake.WeekDayShort()
		timeInfo["month"] = fake.Month()
		timeInfo["month_num"] = strconv.Itoa(fake.MonthNum())
		timeInfo["month_short"] = fake.MonthShort()
		timeInfo["year"] = strconv.Itoa(fake.Year(0, time.Now().Year()))
		timeInfo["now"] = time.Now().String()
		timeInfo["now_unix"] = strconv.FormatInt(time.Now().Unix(), 10)

		response.Code = http.StatusOK
		response.Spec = timeInfo
		response.Message = " get fake time info successfully !!!"
		return nil
	})
}

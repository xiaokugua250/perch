// refer
package fake

import (
	"fmt"
	"github.com/icrowley/fake"
	"strconv"
	"time"
)

//var log = logrus.New()

func GenFakeUsers() map[string]string {
	userInfo := make(map[string]string)
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

	return userInfo

}

func GenFakeEmails() map[string]string {
	fakeEmail := make(map[string]string)
	fakeEmail["email_address"] = fake.EmailAddress()
	fakeEmail["email_body"] = fake.EmailBody()
	fakeEmail["email_subject"] = fake.EmailSubject()

	return fakeEmail

}

func GenFakeIps() map[string]string {
	fakeIps := make(map[string]string)
	fakeIps["ipv4"] = fake.IPv4()
	fakeIps["ipv6"] = fake.IPv6()
	fakeIps["domain_name"] = fake.DomainName()
	fakeIps["domain_zone"] = fake.DomainZone()
	return fakeIps

}

func GenFakeLocs() map[string]string {
	locationInfo := make(map[string]string)
	locationInfo["cpuntry"] = fake.Country()
	locationInfo["city"] = fake.City()
	locationInfo["street_address"] = fake.StreetAddress()
	locationInfo["street"] = fake.Street()

	locationInfo["latitude"] = fmt.Sprintf("%f", fake.Latitude())
	locationInfo["latitude_degrees"] = fmt.Sprintf("%d", fake.LatitudeDegrees())
	locationInfo["latitude_direction"] = fake.LatitudeDirection()
	locationInfo["latitude_minutes"] = fmt.Sprintf("%d", fake.LatitudeMinutes())
	locationInfo["latitude_seconds"] = strconv.Itoa(fake.LatitudeSeconds())
	locationInfo["longitude_drees"] = fmt.Sprintf("%d", fake.LongitudeDegrees())
	locationInfo["longitude"] = fmt.Sprintf("%f", fake.Longitude())
	locationInfo["longitude_direction"] = fake.LongitudeDirection()
	locationInfo["longitude_minutes"] = fmt.Sprintf("%d", fake.LongitudeMinutes())
	locationInfo["longitude_second"] = fmt.Sprintf("%d", fake.LongitudeSeconds())

	return locationInfo

}
func GenFakeCreditCard() map[string]string {
	fakeCredit := make(map[string]string)
	fakeCredit["credit_type"] = fake.CreditCardType()
	fakeCredit["credit_num"] = fake.CreditCardNum("")
	return fakeCredit

}
func GenFakeTimes() map[string]string {
	timeInfo := make(map[string]string)
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
	return timeInfo

}

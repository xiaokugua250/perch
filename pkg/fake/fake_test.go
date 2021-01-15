// refer
package fake

import (
	"fmt"
	"github.com/icrowley/fake"
	database "perch/database/mysql"
	model "perch/web/model/resources"
	"strconv"
	"sync"
	"testing"
	"time"
)

//var log = logrus.New()

func TestGenFakeUsers(t *testing.T) {
	var (
		blog  model.ResourceBlogs
		blogs []model.ResourceBlogs
		err   error
		wg    sync.WaitGroup
	)
	err = database.InitMySQLDB()
	if err != nil {
		fmt.Print(err)
	}

	for i := 1; i <= 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			blog.Title = fake.Title()
			blog.Name = fake.Sentence() + fmt.Sprintf("%s", strconv.Itoa(i))
			blog.Author = fake.UserName()
			blog.Category = fake.Words()
			blog.ContentMd = fake.SentencesN(5)
			blog.ContentHtml = fake.SentencesN(10)
			blog.Link = "http://" + fake.Sentence() + ".com"
			blog.CreatedAt = time.Now().Unix()
			blog.UpdatedAt = time.Now().Unix() + 1000
			blogs = append(blogs, blog)

		}(i)
	}
	wg.Wait()
	if err = database.MysqlDb.Create(&blogs).Error; err != nil {
		fmt.Print(err)
	}
}

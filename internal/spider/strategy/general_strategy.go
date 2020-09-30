/**
通用策略，比如时间，
*/
package strategy

import (
	"time"
)

/**
带随机延迟的时间策略
*/
func GenTimeStrategyWithDelay(timezone string, delay time.Duration, randomDelay int) error {

	var (
		err error
	)
	if timezone == "" {
		timezone = time.Now().Location().String()
	}
	_, err = time.LoadLocation(timezone)
	if err != nil {
		return err
	}

	/*randValue := rand.Intn(randomDelay )
	durationTime := time.Second*randValue
	time.Now().Add(durationTime)
	time.NewTicker(randValue*time.Second)*/

	return nil

}

/**
时间处理包
*/
package utils

import "time"

func TimeUtils_SystemTime() time.Time {
	return time.Now()
}

func TimeUtils_DelayDuration(timeStamp time.Time, duration time.Duration) time.Time {
	return timeStamp.Add(duration)
}

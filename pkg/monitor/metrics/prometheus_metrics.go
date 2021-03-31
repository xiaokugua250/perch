package metrics

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	MyTestCounter = prometheus.NewCounter(prometheus.CounterOpts{
		//因为Name不可以重复，所以建议规则为："部门名_业务名_模块名_标量名_类型"
		Name: "my_test_counter",	//唯一id，不可重复Register()，可以Unregister()
		Help: "my test counter",	//对此Counter的描述
	})
	MyTestGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "my_test_gauge",
		Help: "my test gauge",
	})
	MyTestHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "my_test_histogram",
		Help: "my test histogram",
		Buckets: prometheus.LinearBuckets(20, 5, 5), //第一个桶20起，每个桶间隔5，共5个桶。 所以20, 25, 30, 35, 40
	})
	MyTestSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "my_test_summary",
		Help: "my test summary",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001}, //返回五分数， 九分数， 九九分数
	})
)


func init(){
	prometheus.MustRegister(MyTestCounter)
	prometheus.MustRegister(MyTestGauge)
	prometheus.MustRegister(MyTestHistogram)
	prometheus.MustRegister(MyTestSummary)
}

/**
	example
  go func(){
	var i float64
          for {
		i++
		MyTestCounter.Add(10000)   //每次加常量
		MyTestGauge.Add(i)		   //每次加增量
		MyTestHistogram.Observe(30 + math.Floor(120*math.Sin(float64(i)*0.1))/10)  //每次观察一个18 - 42的量
		MyTestSummary.Observe(30 + math.Floor(120*math.Sin(float64(i)*0.1))/10)

              time.Sleep(time.Second)
          }
      }()
 */
// 系统指标
func SystemPrometheusMetric(){

}


/**
应用监控指标
 */
func ApplicationPrometheusMetric(){

}

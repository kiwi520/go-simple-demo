package main

import(
	//"github.com/robfig/cron"
	"github.com/tour/jobs"

	//"fmt"
)

//初始化
func init() {
	//jobs.Canyonfreepark()
	jobs.Canyoninlotinfopark()
}

func main(){
	//c := cron.New()
	//c.AddFunc("@hourly",      func() { fmt.Println("Every hour") })
	//c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	//c.Start()
}
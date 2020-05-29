package main

import "fmt"
import "time"
import "github.com/robfig/cron"

func main() {
	t := time.Now()
	fmt.Println(t)
	// fmt.Println(int(t.Weekday()))

	// fmt.Println(int(t.Year()))
	// fmt.Println(int(t.Month()))
	// fmt.Println(int(t.Day()))

	// fmt.Println((t.Year()))
	// fmt.Println((t.Month()))
	// fmt.Println((t.Day()))

	// fmt.Println((t.Location()))
	// fmt.Println((t.Zone()))
	// fmt.Println((t.Date()))
	// fmt.Println((t.Clock()))
	// fmt.Println((t.YearDay()))

	// fmt.Println((t.Hour()))
	// fmt.Println((t.Minute()))
	// fmt.Println((t.Second()))
	// fmt.Println((t.Nanosecond()))
	// fmt.Println(t.Add(1000*1000*1000))
	// fmt.Println(t.AddDate(1,1,1))

	// timeLayout := "2006-01-02 15:04:05"
	// fmt.Println(t.Format(timeLayout))
	// tt,_:=time.Parse(timeLayout,"2019-01-01 00:00:00");
	// fmt.Println(tt);

	// fmt.Println("11")
	// time.Sleep(time.Second*10)
	// fmt.Println("11")

	// timer1 := time.NewTimer(time.Second * 2)
	// t1 := time.Now()
	// fmt.Println("====", t1)

	// t2 := <-timer1.C
	// fmt.Println("****", t2)

	// fmt.Println(t.Truncate(1 * time.Hour))
	// fmt.Println(t.Round(1 * time.Hour))

	// fmt.Println(t.Truncate(1 * time.Minute))
	// fmt.Println(t.Round(1 * time.Minute))


	i := 0
    c := cron.New()
    spec := "*/5 * * * * ?"
    c.AddFunc(spec, func() {
        i++
        fmt.Println("cron running:", i)
    })
    c.Start()

	select{}


}

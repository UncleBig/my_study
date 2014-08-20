// rrd.go

package main

import (
	"fmt"
	"github.com/ziutek/rrd"
	//"io/ioutil"
	//"testing"
	"time"
)

func main() {
	const (
		dbfile    = "rrd/test.rrd"
		step      = 1
		heartbeat = 2 * step
	)

	c := rrd.NewCreator(dbfile, time.Now(), step)
	c.RRA("AVERAGE", 0.5, 1, 600)
	c.RRA("AVERAGE", 0.5, 4, 600)
	c.RRA("AVERAGE", 0.5, 24, 600)
	//c.DS("cnt", "COUNTER", heartbeat, 0, 100)
	c.DS("load1", "GAUGE", heartbeat, 0, 60)
	c.DS("load5", "GAUGE", heartbeat, 0, 60)
	c.DS("load15", "GAUGE", heartbeat, 0, 60)
	err := c.Create(true)
	if err != nil {
		fmt.Printf("%s", err)
	}

	// Update
	u := rrd.NewUpdater(dbfile)

	for i := 0; i < 10; i++ {
		fmt.Printf("hit\n")
		time.Sleep(step * time.Second)
		err := u.Update(time.Now(), 0.5*float64(i), 1*float64(i), 1.5*float64(i))
		if err != nil {
			fmt.Printf("%s", err)
		}
	}
	/*
		for i := 10; i < 20; i++ {
			fmt.Printf("hit\n")
			time.Sleep(step * time.Second)
			u.Cache(time.Now(), 2*float64(i))
		}
		err = u.Update()
		if err != nil {
			fmt.Printf("%s", err)
		}
	*/
	/*
		u := rrd.NewUpdater(dbfile)
		err := u.Update(time.Now(), 5, 1.5*7)
		if err != nil {
			fmt.Printf("%q", err)
		}
	*/

	g := rrd.NewGrapher()
	g.SetTitle("Disk Used")
	g.SetVLabel("Use Percent")
	g.SetSize(800, 300)
	g.SetLowerLimit(0.00)
	g.SetUpperLimit(100.00)
	g.SetWatermark("danale")
	g.Def("load1", dbfile, "load1", "AVERAGE")
	g.Def("load5", dbfile, "load5", "AVERAGE")
	g.Def("load15", dbfile, "load15", "AVERAGE")
	//g.Def("v2", dbfile, "cnt", "AVERAGE")
	g.VDef("last1", "load1,LAST")
	g.VDef("last5", "load5,LAST")
	g.VDef("last15", "load15,LAST")
	//g.VDef("avg2", "v2,AVERAGE")
	g.Line(1, "load1", "ff0000", "load1")
	g.Line(1, "load5", "00ff00", "load5")
	g.Line(1, "load15", "0000ff", "load15")
	//g.Area("v2", "0000ff", "var 2")
	g.GPrintT("last1", "Graph last %c")
	g.GPrintT("last5", "Graph last %c")
	g.GPrintT("last15", "Graph last %c")
	//g.GPrint("avg2", "avg2=%lf")
	g.PrintT("last1", "Graph last %c")
	g.PrintT("last5", "Graph last %c")
	g.PrintT("last15", "Graph last %c")

	//g.Print("avg2", "avg2=%lf")

	now := time.Now()
	fmt.Print(now)
	i, err := g.SaveGraph("rrd/disk.png", now.Add(-20*time.Minute), now)
	fmt.Printf("%+v\n", i)
	if err != nil {
		fmt.Printf("%s", err)
	}

}

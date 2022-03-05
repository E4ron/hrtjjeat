package main

import (
	"awesomeProject2/db"
	"awesomeProject2/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

type Valuate struct {
	Current float64
	Min     float64
	Max     float64
	Name    string
}

var valuates map[string]Valuate

func main() {
	valuates = make(map[string]Valuate)

	db.InitLogger()

	opt := setting.Load("setting.json")

	router := gin.Default()
	router.LoadHTMLGlob("template/*")
	router.Static("assets", "assets")

	router.GET("/", index)
	router.GET("/course", func(c *gin.Context) {
		fmt.Println(valuates)
		c.JSON(200, valuates)

	})
	router.PUT("/course", addValuate)
	go UpdateRate()

	_ = router.Run(opt.Address + ":" + opt.Port)
}

func index(c *gin.Context) {
	c.HTML(200, "index", nil)
}

func UpdateRate() {
	random := rand.New(rand.NewSource(time.Now().Unix()))

	for {
		for k, v := range valuates {
			v.Current = v.Min + random.Float64()*(v.Max)
			valuates[k] = v

			time.Sleep(time.Duration(2000))
		}
	}

}

func addValuate(c *gin.Context) {
	var v Valuate
	e := c.BindJSON(&v)
	if e != nil {
		fmt.Println(e)
		c.JSON(400, nil)
		return
	}
	valuates[v.Name] = v
	c.JSON(200, nil)
}

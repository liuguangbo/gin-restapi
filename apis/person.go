package apis

import (
	"fmt"
	. "gin-restapi/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func AddPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	p := Person{FirstName: firstName, LastName: lastName}

	ra, err := p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func GetPersonsApi(c *gin.Context) {

	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")
	p := Person{FirstName: firstName, LastName: lastName}
	persons := make([]Person, 0)
	persons, err := p.GetPersons()
	if err != nil {
		log.Fatalln(err)
	}
	//msg := fmt.Sprintf("querypersons successful %s", persons)
	c.JSON(http.StatusOK, gin.H{
		"msg": persons,
	})

}

func GetPersonApi(c *gin.Context) {
	id := c.Param("id")
	b, error := strconv.Atoi(id)
	if error != nil {
		fmt.Println("字符串转换成整数失败")
	}
	p := Person{Id: b}

	person, err := p.GetPerson()
	if err != nil {
		log.Fatalln(err)
	}
	//msg := fmt.Sprintf("querypersons successful %s", persons)
	c.JSON(http.StatusOK, gin.H{
		"msg": person,
	})

}

func ModPersonApi(c *gin.Context) {
	d := c.Param("id")
	id, err := strconv.Atoi(d)
	if err != nil {
		fmt.Println("字符串转换成整数失败")
	}
	person := Person{Id: id}
	err = c.Bind(&person)
	if err != nil {
		log.Fatalln(err)
	}

	ra, err := person.ModPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("Update person %d successful %d", person.Id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})

}

func DelPersonApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	person := Person{Id: id}
	ra, err := person.DelPerson()
	if err != nil {
		log.Fatal(err)
	}
	msg := fmt.Sprintf("Delete person %d successful %d", person.Id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

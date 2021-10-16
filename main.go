package main

import (
	"encoding/hex"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/renecouto/logu/models"
)

func checkCsrf(c *gin.Context) {
	cookie, err := c.Cookie("csrfToken")
	form := c.PostForm("csrfToken")
	if err != nil || form != cookie {
		log.Println("csrf tokens dont match !!!", cookie, "vs", form)
		c.AbortWithStatus(400)
	}
}

func checkCsrfJson(c *gin.Context, form string) {
	cookie, err := c.Cookie("csrfToken")
	if err != nil || form != cookie {
		log.Println("csrf tokens dont match !!!", cookie, "vs", form)
		c.AbortWithStatus(400)
	}
}

func userIdFromCookie(c *gin.Context, ir *models.ItemsRepository) (int, error) {
	u, err := c.Cookie("user")
	if err != nil {
		c.AbortWithStatus(400)
		log.Println("got err:", err)
		return 0, err
	}
	return ir.GetUserByUsername(u).Id, nil
}

func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

type Controller struct {
	itemsRepo *models.ItemsRepository
}

func (*Controller) GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login", gin.H{})
}

func (ctl *Controller) PostLogin(c *gin.Context) {
	username := c.PostForm("user")

	if username != "" {
		onDb := ctl.itemsRepo.GetUserByUsername(username)
		if onDb != nil {
			c.SetCookie("user", username, 1000, "/", "localhost", false, true)
			c.Redirect(http.StatusFound, "/")
			return
		}
	}
	c.AbortWithStatusJSON(401, gin.H{"code": "user does not exist"})

}

func (*Controller) PostLogout(c *gin.Context) {
	c.SetCookie("user", "", 1000, "/", "localhost", false, true)
	c.Redirect(http.StatusFound, "/")
}

func (ctl *Controller) GetIndex(c *gin.Context) {
	c.SetSameSite(http.SameSiteStrictMode)
	csrfToken := GenerateSecureToken(128)
	c.SetCookie("csrfToken", csrfToken, 1000, "/", "localhost", false, true)

	u, err := c.Cookie("user")
	if u == "" {
		c.Redirect(http.StatusFound, "/login")
	}
	if err != nil {
		panic(err)
	}
	user := ctl.itemsRepo.GetUserByUsername(u)
	log.Println("user: ", user)
	if user == nil {
		c.AbortWithStatus(500)
		return
	}
	date := c.Query("date")
	var parsed time.Time
	if date == "" {
		parsed = time.Now()
	} else {
		p, err := time.Parse(time.RFC3339, date+"T11:45:26.371Z")

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(400)
			return
		} else {
			parsed = p
		}
	}

	c.HTML(http.StatusOK, "index.html", gin.H{"itens": ctl.itemsRepo.GetAllForDateAndUser(parsed, user.Id), "user": u, "date": parsed, "csrfToken": csrfToken})
}

func (ctl *Controller) CreateEvent(c *gin.Context) {
	checkCsrf(c)
	userId, err := userIdFromCookie(c, ctl.itemsRepo)
	if err != nil {
		return
	}
	e := new(models.Event)
	h, err := strconv.Atoi(c.PostForm("ScheduledForHour"))
	if err != nil {
		c.AbortWithStatus(400)
	}
	m, err := strconv.Atoi(c.PostForm("ScheduledForMinute"))
	if err != nil {
		c.AbortWithStatus(400)
	}
	year, month, day := time.Now().Date()
	timeRebuilt := time.Date(year, month, day, h, m, 0, 0, utc)
	c.Bind(e)
	e.ScheduledFor = timeRebuilt
	e.CreatedAt = time.Now()
	e.User = userId
	ctl.itemsRepo.AddEvent(*e)
	c.Redirect(http.StatusFound, "/")
}

func (ctl *Controller) CreateTask(c *gin.Context) {
	checkCsrf(c)
	userId, err := userIdFromCookie(c, ctl.itemsRepo)
	if err != nil {
		log.Println("got err", err)
		return
	}
	e := new(models.Task)
	c.Bind(e)
	e.CreatedAt = time.Now()
	e.User = userId
	ctl.itemsRepo.AddTask(*e)
	c.Redirect(http.StatusFound, "/")
}

func (ctl *Controller) UpdateTask(c *gin.Context) {
	checkCsrfJson(c, c.GetHeader("csrfToken"))
	userId, err := userIdFromCookie(c, ctl.itemsRepo)
	if err != nil {
		return
	}
	e := new(models.Task)

	err = c.ShouldBindJSON(e)
	if err != nil {
		log.Println("got err with binding: ", err)
	}
	task := ctl.itemsRepo.GetTaskById(userId, e.Id)
	task.Done = e.Done
	log.Println(ctl.itemsRepo.GetTaskById(userId, e.Id))
	c.JSON(200, gin.H{"status": "OK"})
}

func (ctl *Controller) CreateNote(c *gin.Context) {
	checkCsrf(c)
	userId, err := userIdFromCookie(c, ctl.itemsRepo)
	if err != nil {
		return
	}
	e := new(models.Note)
	c.Bind(e)
	e.CreatedAt = time.Now()
	e.User = userId
	ctl.itemsRepo.AddNote(*e)
	c.Redirect(http.StatusFound, "/")
}

func getUtc() *time.Location {
	utc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}
	return utc
}

var utc = getUtc()

var renegade = models.User{Id: 1, Username: "renegade", FullName: "renato Cortes"}

func SetupData(itemsRepo *models.ItemsRepository) {
	itemsRepo.AddEvent(models.Event{Description: "comprar coisas", ScheduledFor: time.Now(), CreatedAt: time.Now(), User: renegade.Id})
	itemsRepo.AddTask(models.Task{Description: "regar as plantas", CreatedAt: time.Now(), User: renegade.Id})
	itemsRepo.AddNote(models.Note{Description: "hoje deve ser um bom dia para meditar, será que faço isso?", CreatedAt: time.Now(), User: renegade.Id})
	itemsRepo.AddUser(renegade)
}
func run_website() {

	r := gin.Default()
	r.LoadHTMLGlob("web/templates/**/*")

	itemsRepo := models.NewItemsRepository()
	SetupData(&itemsRepo)
	r.Static("/assets", "./assets")
	ctl := Controller{itemsRepo: &itemsRepo}
	r.GET("/login", ctl.GetLogin)
	r.POST("/login", ctl.PostLogin)
	r.POST("/logout", ctl.PostLogout)
	r.GET("/", ctl.GetIndex)
	r.POST("/create-event", ctl.CreateEvent)
	r.POST("/create-task", ctl.CreateTask)

	r.PUT("/update-task", ctl.UpdateTask)
	r.POST("/create-note", ctl.CreateNote)
	r.Run()
}

func main() {
	run_website()
}

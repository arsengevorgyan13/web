package main

import (
	"database/sql"
	_ "database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"
	_ "strconv"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type indexPage struct {
	Header             []headerData
	TopBlock           []topBlockData
	PageButtons        []pageButtonsData
	FeaturedPostsTitle string
	FeaturedPosts      []featuredPostsData
	RecentPagesTitle   string
	RecentPages        []recentPagesData
	Footer             []footerData
}

type postPage struct {
	PostHeader []postHeaderData
	Post       []postData
	PostFooter []postFooterData
}

type postHeaderData struct {
	Escape               string
	PostHeaderNavButtons []postHeaderNavButtonsData
}

type postHeaderNavButtonsData struct {
	Home       string
	Categories string
	About      string
	Contact    string
}

type postData struct {
	Image    string `db:"image_url"`
	Title    string `db:"title"`
	Subtitle string `db:"subtitle"`
	Content  string `db:"content"`
}

type postFooterData struct {
	Background         string
	PostFooterContacts []postFooterContactsData
	PostFooterBottom   []postFooterBottomData
}

type postFooterContactsData struct {
	Title  string
	Button string
}

type postFooterBottomData struct {
	Escape            string
	PostFooterButtons []postFooterButtonsData
}

type postFooterButtonsData struct {
	Home       string
	Categories string
	About      string
	Contact    string
}

type headerData struct {
	Image            string
	Escape           string
	HeaderNavButtons []headerNavButtonsData
}

type headerNavButtonsData struct {
	Home       string
	Categories string
	About      string
	Contact    string
}

type topBlockData struct {
	Title  string
	Text   string
	Button string
}

type pageButtonsData struct {
	PageButtonsList []pageButtonsListData
}

type pageButtonsListData struct {
	Nature      string
	Photography string
	Relaxation  string
	Vacation    string
	Travel      string
	Adventure   string
}

type featuredPostsData struct {
	PostID      string `db:"post_id"`
	Image       string `db:"image_url"`
	Title       string `db:"title"`
	Description string `db:"subtitle"`
	AuthorImage string `db:"author_img"`
	AuthorName  string `db:"author_name"`
	Date        string `db:"publish_date"`
	PostURL     string
}

type recentPagesData struct {
	PostID      string `db:"post_id"`
	Image       string `db:"image_url"`
	Title       string `db:"title"`
	Description string `db:"subtitle"`
	AuthorImage string `db:"author_img"`
	AuthorName  string `db:"author_name"`
	Date        string `db:"publish_date"`
	PostURL     string
}

type footerData struct {
	Background     string
	FooterContacts []footerContactsData
	FooterBottom   []footerBottomData
}

type footerContactsData struct {
	Title  string
	Button string
}

type footerBottomData struct {
	Escape        string
	FooterButtons []footerButtonsData
}

type footerButtonsData struct {
	Home       string
	Categories string
	About      string
	Contact    string
}

func Footer() []footerData {
	return []footerData{
		{
			Background:     "../static/image/footer.svg",
			FooterContacts: FooterContacts(),
			FooterBottom:   FooterBottom(),
		},
	}
}

func FooterContacts() []footerContactsData {
	return []footerContactsData{
		{
			Title:  "Stay in Touch",
			Button: "Submit",
		},
	}
}

func FooterBottom() []footerBottomData {
	return []footerBottomData{
		{
			Escape:        "../static/image/Escape.svg",
			FooterButtons: FooterButtons(),
		},
	}
}

func FooterButtons() []footerButtonsData {
	return []footerButtonsData{
		{
			Home:       "HOME",
			Categories: "CATEGORIES",
			About:      "ABOUT",
			Contact:    "CONTACT",
		},
	}
}

func PostHeader() []postHeaderData {
	return []postHeaderData{
		{
			Escape:               "../static/image/blackescape.svg",
			PostHeaderNavButtons: PostHeaderNavButtons(),
		},
	}
}

func PostHeaderNavButtons() []postHeaderNavButtonsData {
	return []postHeaderNavButtonsData{
		{
			Home:       "HOME",
			Categories: "CATEGORIES",
			About:      "ABOUT",
			Contact:    "CONTACT",
		},
	}
}

func PostFooter() []postFooterData {
	return []postFooterData{
		{
			Background:         "../static/image/footer.svg",
			PostFooterContacts: PostFooterContacts(),
			PostFooterBottom:   PostFooterBottom(),
		},
	}
}

func PostFooterContacts() []postFooterContactsData {
	return []postFooterContactsData{
		{
			Title:  "Stay in Touch",
			Button: "Submit",
		},
	}
}

func PostFooterBottom() []postFooterBottomData {
	return []postFooterBottomData{
		{
			Escape:            "../static/image/Escape.svg",
			PostFooterButtons: PostFooterButtons(),
		},
	}
}

func PostFooterButtons() []postFooterButtonsData {
	return []postFooterButtonsData{
		{
			Home:       "HOME",
			Categories: "CATEGORIES",
			About:      "ABOUT",
			Contact:    "CONTACT",
		},
	}
}

func Header() []headerData {
	return []headerData{
		{
			Image:            "../static/image/dust.png",
			Escape:           "../static/image/Escape.svg",
			HeaderNavButtons: HeaderNavButtons(),
		},
	}
}

func HeaderNavButtons() []headerNavButtonsData {
	return []headerNavButtonsData{
		{
			Home:       "HOME",
			Categories: "CATEGORIES",
			About:      "ABOUT",
			Contact:    "CONTACT",
		},
	}
}

func TopBlock() []topBlockData {
	return []topBlockData{
		{
			Title:  "Lets do it together.",
			Text:   "We travel the world in search of stories. Come along for the ride.",
			Button: "View Latest Posts",
		},
	}
}

func PageButtons() []pageButtonsData {
	return []pageButtonsData{
		{
			PageButtonsList: PageButtonsList(),
		},
	}
}

func PageButtonsList() []pageButtonsListData {
	return []pageButtonsListData{
		{
			Nature:      "Nature",
			Photography: "Photography",
			Relaxation:  "Relaxation",
			Vacation:    "Vacation",
			Travel:      "Travel",
			Adventure:   "Adventure",
		},
	}
}

/*func index(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := indexPage{
		Header:        Header(),
		TopBlock:      TopBlock(),
		PageButtons:   PageButtons(),
		FeaturedPosts: FeaturedPosts(),
		RecentPages:   RecentPages(),
		Footer:        Footer(),
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}*/

/*func post(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/post.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := postPage{
		PostHeader:   PostHeader(),
		PostTopBlock: PostTopBlock(),
		PostImage:    "../static/image/POLAR.png",
		Pharagraphs:  Pharagraphs(),
		PostFooter:   PostFooter(),
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}*/

/*func post(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) { //order
	return func(w http.ResponseWriter, r *http.Request) {
		pos3tIDStr := mux.Vars(r)["po3stID"] // Получаем orderID в виде строки из параметров урла

		pos3tID, err := strconv.Atoi(post3IDStr) // Конвертируем строку orderID в число
		if err != nil {
			http.Error(w, "Invalid order id", 403)
			log.Println(err)
			return
		}

		post, err := postByID(db, post3ID)
		if err != nil {
			if err == sql.ErrNoRows {
				// sql.ErrNoRows возвращается, когда в запросе к базе не было ничего найдено
				// В таком случае мы возвращем 404 (not found) и пишем в тело, что ордер не найден
				http.Error(w, "Order not found", 404)
				log.Println(err)
				return
			}

			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		ts, err := template.ParseFiles("pages/post.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		data := postPage{
			PostHeader:   PostHeader(),
			PostTopBlock: PostTopBlock(),
			PostImage:    "../static/image/POLAR.png",
			Pharagraphs:  Pharagraphs(),
			PostFooter:   PostFooter(),
		}

		err = ts.Execute(w, data)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		log.Println("Request completed successfully")
	}
}*/

func post(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		postIDStr := mux.Vars(r)["postID"]

		postID, err := strconv.Atoi(postIDStr)
		if err != nil {
			http.Error(w, "Invalid post id", 403)
			log.Println(err)
			return
		}

		post, err := postByID(db, postID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Post not found", 404)
				log.Println(err)
				return
			}

			http.Error(w, "Server Error", 500)
			log.Println(err)
			return
		}

		ts, err := template.ParseFiles("pages/post.html") // Второстепенная страница блога
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
			return                                      // Bыполнение ф-ии
		}

		data := postPage{
			PostHeader: PostHeader(),
			Post:       post,
			PostFooter: PostFooter(),
		}

		err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}
	}
}

func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		featuredposts, err := featuredPosts(db)
		if err != nil {
			http.Error(w, "Error1", 500)
			log.Println(err)
			return
		}

		recentpages, err := recentPages(db)
		if err != nil {
			http.Error(w, "Error2", 500)
			log.Println(err)
			return
		}

		ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err.Error())                    // Используем стандартный логгер для вывода ошибки в консоль
			return                                      // Выполнение ф-ии
		}

		data := indexPage{
			Header:             Header(),
			TopBlock:           TopBlock(),
			PageButtons:        PageButtons(),
			FeaturedPostsTitle: "Featured Posts",
			FeaturedPosts:      featuredposts,
			RecentPagesTitle:   "Recent Pages",
			RecentPages:        recentpages,
			Footer:             Footer(),
		}

		err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		log.Println("Request completed successfully")
	}
}

func featuredPosts(db *sqlx.DB) ([]featuredPostsData, error) {
	const query = `
		SELECT
			post_id,
			image_url,
			title,
			subtitle,
			author_img,
			author_name,
			publish_date
		FROM 
			post
		WHERE 
			featured = 1
	`
	// Составляем SQL-запрос для получения записей для секции featured-posts

	var featuredposts []featuredPostsData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&featuredposts, query) // Делаем запрос в базу данных
	if err != nil {                         // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	for _, post := range featuredposts {
		post.PostURL = "/post/" + post.PostID
	}

	return featuredposts, nil
}

func recentPages(db *sqlx.DB) ([]recentPagesData, error) {
	const query = `
		SELECT
			post_id,
			image_url,
			title,
			subtitle,
			author_img,
			author_name,
			publish_date
		FROM
			post
		WHERE 
			featured = 0
	` // Составляем SQL-запрос для получения записей для секции featured-posts

	var recentpages []recentPagesData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&recentpages, query) // Делаем запрос в базу данных
	if err != nil {                       // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	for _, post := range recentpages {
		post.PostURL = "/post/" + post.PostID
	}

	return recentpages, nil
}

func postByID(db *sqlx.DB, postID int) ([]postData, error) {
	const query = `
		SELECT
			image_url,
			title,
			subtitle,
			content
		FROM
			post
		WHERE
			post_id = ?
	`
	// В SQL-запросе добавились параметры, как в шаблоне. ? означает параметр, который мы передаем в запрос ниже

	var post []postData

	// Обязательно нужно передать в параметрах orderID
	err := db.Select(&post, query, postID)
	if err != nil {
		return nil, err
	}

	return post, nil
}

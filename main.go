package main

import (
	"fmt"
	"github.com/labstack/echo"
	"io"
	"net/http"
	"os"
)

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func main() {
	e := echo.New()

	e.POST("/user/login", login)
	e.POST("/upload", fileUpload)
	//e.POST("/upload", multipartFileUpload)

	e.Logger.Fatal(e.Start(":1323"))
}

func login(c echo.Context) error {
	u := new(User)

	if err := c.Bind(u); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Could not login")
	}

	return c.JSON(http.StatusOK, u)
}

//func multipartFileUpload(c echo.Context) error {
//	name := c.FormValue("name")
//	email := c.FormValue("email")
//
//	form, err := c.MultipartForm()
//	if err != nil {
//		return err
//	}
//	files := form.File["files"]
//
//	for _, file := range files {
//		// Source
//		src, err := file.Open()
//		if err != nil {
//			return err
//		}
//		defer src.Close()
//
//		// Destination
//		dst, err := os.Create(file.Filename)
//		if err != nil {
//			return err
//		}
//		defer dst.Close()
//
//		// Copy
//		if _, err = io.Copy(dst, src); err != nil {
//			return err
//		}
//	}
//
//	return  c.String(http.StatusOK, name + email + "File uploaded")
//}

func fileUpload(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")

	// form file
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Filename: ", file.Filename)

	// open file
	src, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer src.Close()

	// destination
	dst, err := os.Create(file.Filename)  // create a file in the current folder with the name of the uploaded file
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer dst.Close()
	fmt.Println("Dst: ", dst.Name())

	// copy
	if _, err = io.Copy(dst, src); err != nil {  // copy from sorce to destination
		return err
	}

	return c.String(http.StatusOK, name + ", " + email + ", your file was uploaded successfully")
	//return c.String(http.StatusOK, "File uploaded")
}

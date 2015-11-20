package controllers

import (
	// "bytes"
	// "fmt"
	// "image"
	_ "image/jpeg"
	_ "image/png"

	// "github.com/revel/samples/upload/app/routes"

	"github.com/revel/revel"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"encoding/csv"
	"fmt"
	"os"

	// "path/filepath"
)

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

type Single struct {
	App
}

func (c *Single) Upload() revel.Result {
	return c.Render()
}

func (c *Single) HandleUpload(csvFile []byte) revel.Result {

	return c.RenderJson(fmt.Printf("+%v\n", c.Request.FormFile("csvFile")));

	// for aa , vv : = range c.Params.Files["csvFile"]  {  // File herein shall be consistent and view.
 //        // Create buffer
 //        buf : =  new ( bytes . Buffer )
 //        // Create a tmpfile and assemble your multipart from there (not tested)
 //        w : = multipart . NewWriter ( buf )
 //        // Create file field
 //        FSRC , _ , _ : = c . Request . FormFile ( "csvFile" )
 //    	fdst , ERR : = OS . OpenFile ( "eqap/" + vv . Filename , OS . O_CREATE | OS . O_WRONLY ,  677 )

 //        if ERR ! =  nil  {
 //            fmt . Println ( "d" )
 //            return  nil
 //        }
 //        defer fdst . Close ()
 //        // Write file field from file to upload
 //        _ , ERR = IO . Copy ( fdst , FSRC )
 //        if ERR ! =  nil  {
 //            fmt . Println ( "e" )
 //            return  nil
 //        }
 //        FSRC . Close ()
 //        w . Close ()
 //    }

	return c.RenderJson(c.Request.FormFile("csvFile"));
	return c.RenderJson(c.Params.Files["csvFile"]);

	// dir, err := filepath.Abs(filepath.Dir(c.Params.Files["csvFile"][0].Filename))
 //    if err != nil {
 //            c.RenderJson(err)
 //    }
 //    return c.RenderJson(dir);

	// connect DB
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		return c.RenderJson(err);
	}
	defer db.Close()

	// ping DB
	err = db.Ping()
	if err != nil {
		return c.RenderJson(err);
	}


	db.Exec("TRUNCATE TABLE tmp_marketing_costs_facebook;");
	_,err = db.Exec("LOAD DATA INFILE '/Users/macbookpro/work/HotelQuickly-Facebook-for-import.csv' INTO TABLE tmp_marketing_costs_facebook FIELDS TERMINATED BY ',' ENCLOSED BY '\"' LINES TERMINATED BY '\n' IGNORE 1 ROWS;")
    if err != nil {
        return c.RenderJson(err);
    }

    return c.RenderJson("Success!!");

	// start Processing
	csvfile, err := os.Open(c.Params.Files["csvFile"][0].Filename)

     if err != nil {
             return c.RenderJson(err);
     }

     defer csvfile.Close()

     reader := csv.NewReader(csvfile)

     reader.FieldsPerRecord = -1 // see the Reader struct information below

     rawCSVdata, err := reader.ReadAll()

     if err != nil {
             fmt.Println(err)
             os.Exit(1)
     }

     // sanity check, display to standard output
     for _, each := range rawCSVdata {
             //fmt.Printf("email : %s and timestamp : %s\n", each[0], each[1])
     		return c.RenderJson(each[0]);
     }

    return c.RenderJson(c.Params.Files["csvFile"][0].Filename);

	// return c.RenderJson(FileInfo{
	// 	//ContentType: c.Params.Files["csvFile"][0].Header.Get("Content-Type"),
	// 	Filename:    c.Params.Files["csvFile"][0].Filename,
	// 	// RealFormat:  format,
	// 	//Resolution:  fmt.Sprintf("%dx%d", conf.Width, conf.Height),
	// 	Size:        len(csvFile),
	// 	Status:      "Successfully uploaded",
	// })
}


// func (c *Single) HandleUpload(avatar []byte) revel.Result {
// 	// Validation rules.
// 	c.Validation.Required(avatar)
// 	c.Validation.MinSize(avatar, 2*KB).
// 		Message("Minimum a file size of 2KB expected")
// 	c.Validation.MaxSize(avatar, 2*MB).
// 		Message("File cannot be larger than 2MB")

// 	// Check format of the file.
// 	conf, format, err := image.DecodeConfig(bytes.NewReader(avatar))
// 	c.Validation.Required(err == nil).Key("avatar").
// 		Message("Incorrect file format")
// 	c.Validation.Required(format == "jpeg" || format == "png").Key("avatar").
// 		Message("JPEG or PNG file format is expected")

// 	// Check resolution.
// 	c.Validation.Required(conf.Height >= 150 && conf.Width >= 150).Key("avatar").
// 		Message("Minimum allowed resolution is 150x150px")

// 	// Handle errors.
// 	if c.Validation.HasErrors() {
// 		c.Validation.Keep()
// 		c.FlashParams()
// 		return c.Redirect(routes.Single.Upload())
// 	}

// 	return c.RenderJson(FileInfo{
// 		ContentType: c.Params.Files["avatar"][0].Header.Get("Content-Type"),
// 		Filename:    c.Params.Files["avatar"][0].Filename,
// 		RealFormat:  format,
// 		Resolution:  fmt.Sprintf("%dx%d", conf.Width, conf.Height),
// 		Size:        len(avatar),
// 		Status:      "Successfully uploaded",
// 	})
// }

package main

    import (
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/s3"
        "github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io/ioutil"
        "fmt"
        "log"
        "os"
	"encoding/json"
    )
    func check(e error) {
    if e != nil {
        panic(e)
    }
}
    func main() {
        // NOTE: you need to store your AWS credentials in ~/.aws/credentials

        // 1) Define your bucket and item names
        bucket := "demo57"
        item  := "test.json"

        // 2) Create an AWS session
        sess, _ := session.NewSession(&aws.Config{
                Region: aws.String("us-east-2")},
        )

        // 3) Create a new AWS S3 downloader 
        downloader := s3manager.NewDownloader(sess)


        // 4) Download the item from the bucket. If an error occurs, log it and exit. Otherwise, notify the user that the download succeeded.
        file, err := os.Create(item)
        numBytes, err := downloader.Download(file,
            &s3.GetObjectInput{
                Bucket: aws.String(bucket),
                Key:    aws.String(item),
        })

        if err != nil {
            log.Fatalf("Unable to download item %q, %v", item, err)
        }

        fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
	// Open our jsonFile
	jsonFile, err := os.Open("test.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Unmarshal using a generic interface
	var f interface{}
	err = json.Unmarshal(byteValue, &f)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
	}

	// JSON object parses into a map with string keys
	itemsMap := f.([]interface{})

	var data1 []map[string]interface{}
	var data2 []map[string]interface{}
	for _, value := range itemsMap {
		// Each value is an interface{} type, that is type asserted as a string
		//fmt.Println(key, value.(map[string]interface {}))
		item := value.(map[string]interface{})
		switch item["Name"] {
		case "EWMInDlvr":
			data1 = append(data1, item)
		case "PART":
			data2 = append(data2, item)
		default:
			fmt.Println("Unknown key for Item found in JSON")
		}
	}
	e, _ := json.Marshal(data1)
	err = ioutil.WriteFile("EWMInDlvr.json", e, 0644)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	p, _ := json.Marshal(data2)
	err = ioutil.WriteFile("PART.json", p, 0644)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
}

	/*

	
	dat, err := ioutil.ReadFile("test.json")
   	 check(err)
   	 type Message struct {
		Name, Text string
	}
	data:=(string(dat))
	dec := json.NewDecoder(strings.NewReader(data))

	// read open bracket
	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)
	// while the array contains values
	for dec.More() {
		var m Message
		// decode an array value (Message)
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v: %v\n", m.Name, m.Text)
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)
}*/

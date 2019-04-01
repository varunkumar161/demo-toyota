package main

    import (
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/s3"
        "github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gorilla/mux"
	"net/http"
        "io/ioutil"
        "fmt"
        "log"
        "os"
        "encoding/json"
        
  )
  func man() {// NOTE: you need to store your AWS credentials in ~/.aws/credentials

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

        if err != nil {log.Fatalf("Unable to download item %q, %v", item, err)
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
type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}
type Address struct {
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}

var people []Person

// Display a single data
func GetPerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range people {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Person{})
}


 func main(){
	
	man()
	router := mux.NewRouter()
         people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
          people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
         router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
          log.Fatal(http.ListenAndServe(":8000", router))


}

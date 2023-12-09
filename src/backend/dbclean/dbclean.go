package main

import (
  "fmt"
  "database/sql"
  "io/ioutil"
  "log"
  "os"
  "os/signal"
  "syscall"
  "time"

  "github.com/go-co-op/gocron"
  _ "github.com/lib/pq"
  "gopkg.in/yaml.v3"
)

// Initialize DB configuration
const (
  dbhost = "postgres"
  dbport = 5432
  dbuser = "postgres"
  dbname = "workflow"
)

var dbpassword = os.Getenv("POSTGRES_PASSWORD")

var delQueries = [2]string{
  `DELETE FROM process WHERE timestamp < $1`,
  `DELETE FROM message WHERE timestamp < $1`,
}

// Struct for application configuration
type config struct {
  Tz string `yaml:"tz"`
  Cron string `yaml:cron`
  Days int `yaml:"days"`
}

// Function for reading an external YAML file
//  and injecting its context to a config struct
func (c *config) getConf() *config {
  yamlFile, err := ioutil.ReadFile("config.yaml")
  if err != nil {
    log.Fatalf("Error reading config gile: %v\n", err)
  }
  err = yaml.Unmarshal(yamlFile, &c)
  if err != nil {
    log.Fatalf("Unmarshal error: %v\n", err)
  }

  return c
}

func main() {
  // Create a signal for monitoring termination commands
  signals := make(chan os.Signal, 1)
  signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

  // Read application configuration
  var conf config
  conf.getConf()

  // Connect to DB
  conninfo := fmt.Sprintf(
    "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    dbhost, dbport, dbuser, dbpassword, dbname)
  db, err := sql.Open("postgres", conninfo)
  if err != nil {
    log.Fatalf("DB failure: %v\n", err)
  }

  // Create a scheduler which runs DELETE queries to DB
  // as scheduled in configuration file
  loc, _ := time.LoadLocation(conf.Tz)
  s := gocron.NewScheduler(loc)
  s.Cron(conf.Cron).Do(func(){
    tx := time.Now().Add(-24 * time.Duration(conf.Days) * time.Hour).UnixMilli()
    fmt.Printf("Deleting rows older than timestamp %d\n", tx)
    for _, query := range delQueries {
      _, err := db.Exec(query, tx)
      if err != nil {
        fmt.Printf("Data deletion failed: %s\n", err)
      }
    }
  })

  // Start a new thread which monitors termination signals
  go func() {
    signal := <-signals
    fmt.Printf("\n%s\n", signal)
    s.StopBlockingChan()
  }()

  // Start scheduler
  s.StartBlocking()
}

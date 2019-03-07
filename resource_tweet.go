package main

import (
	"os"
	"strconv"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// tweet factory method which returns a pointer to the created tweet resource
func tweet() *schema.Resource {
	return &schema.Resource{
		Create: tweetCreate,
		Read:   tweetRead,
		Update: tweetUpdate,
		Delete: tweetDelete,

		Schema: map[string]*schema.Schema{
			"message": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func tweetCreate(d *schema.ResourceData, m interface{}) error {
	create(d, d.Get("message").(string))

	return tweetRead(d, m)
}

func tweetRead(d *schema.ResourceData, m interface{}) error {
	client := createClient()
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	tweet, _, _ := client.Statuses.Show(id, nil)
	d.Set("message", tweet.Text)

	return nil
}

func tweetUpdate(d *schema.ResourceData, m interface{}) error {
	delete(d)
	create(d, d.Get("message").(string))

	return nil
}

func tweetDelete(d *schema.ResourceData, m interface{}) error {
	delete(d)
	d.SetId("")

	return nil
}

func createClient() *twitter.Client {
	config := oauth1.NewConfig(os.Getenv("consumerKey"), os.Getenv("consumerSecret"))
	token := oauth1.NewToken(os.Getenv("token"), os.Getenv("tokenSecret"))
	httpClient := config.Client(oauth1.NoContext, token)
	
	return twitter.NewClient(httpClient)
}

func create(d *schema.ResourceData, message string) {
	client := createClient()
	tweet, _, _ := client.Statuses.Update(message, nil)
	id := strconv.FormatInt(tweet.ID, 10)
	d.SetId(id)
}

func delete(d *schema.ResourceData) {
	client := createClient()
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	client.Statuses.Destroy(id, nil)
}
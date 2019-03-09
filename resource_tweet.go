package main

import (
	"strconv"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/dghubble/go-twitter/twitter"
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
	client := m.(*twitter.Client)
	tweet, _, _ := client.Statuses.Update(d.Get("message").(string), nil)
	id := strconv.FormatInt(tweet.ID, 10)
	d.SetId(id)

	return tweetRead(d, m)
}

func tweetRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*twitter.Client)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	tweet, _, _ := client.Statuses.Show(id, nil)
	d.Set("message", tweet.Text)

	return nil
}

func tweetUpdate(d *schema.ResourceData, m interface{}) error {
	tweetDelete(d, m)
	tweetCreate(d, m)

	return nil
}

func tweetDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*twitter.Client)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	client.Statuses.Destroy(id, nil)

	d.SetId("")

	return nil
}
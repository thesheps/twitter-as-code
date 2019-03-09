package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/dghubble/oauth1"
	"github.com/dghubble/go-twitter/twitter"
)

// Provider function returns instance of Twitter Provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"consumer_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Your Twitter API Consumer Key",
			},

			"consumer_secret": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Your Twitter API Consumer Secret",
			},

			"token": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Your Twitter API Token",
			},

			"token_secret": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Your Twitter API Token Secret",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"twitter_tweet": tweet(),
		},

		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	config := oauth1.NewConfig(d.Get("consumer_key").(string), d.Get("consumer_secret").(string))
	token := oauth1.NewToken(d.Get("token").(string), d.Get("token_secret").(string))
	httpClient := config.Client(oauth1.NoContext, token)
	
	return twitter.NewClient(httpClient), nil
}
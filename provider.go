package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider function returns the ResourcesMap for this simple example
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"twitter_tweet": tweet(),
		},
	}
}

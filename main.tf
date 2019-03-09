provider "twitter" {
  consumer_key    = "XXXXXXXXXXXXXXXXXXXXXXXXX"
  consumer_secret = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  token           = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  token_secret    = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

resource "twitter_tweet" "hello_world" {
  message = "Hello World"
}

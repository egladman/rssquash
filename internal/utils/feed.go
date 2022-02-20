package utils

import (
        "os"
)

const (
	EnvFeedBaseUrl          = "RSSQUASH_BASEURL"
	EnvFeedPrefixUrl        = "RSSQUASH_PREFIX"
	EnvFeedBaseName         = "RSSQUASH_BASENAME"
	EnvFeedTitle            = "RSSQUASH_TITLE"
)

func GetenvFallback(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetFeedBaseUrl() string {
	return GetenvFallback(EnvFeedBaseUrl, "")
}

func GetFeedPrefixUrl() string {
	return GetenvFallback(EnvFeedPrefixUrl, "")

}

func GetFeedBaseName() string {
	return GetenvFallback(EnvFeedBaseName, "feed.atom")
}

func GetFeedTitle() string {
	return GetenvFallback(EnvFeedTitle, "rssquash")
}



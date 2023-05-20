package main

func getAccessTime(url string) float64 {
	responseTime, _ := SiteAvailability.Load(url)
	return responseTime.(float64)
}

func findMinAccessTimeSite() string {
	minAccessTime := float64(1<<63 - 1)
	var siteWithMinAccessTime string
	SiteAvailability.Range(func(k, v interface{}) bool {
		if v.(float64) < minAccessTime {
			minAccessTime = v.(float64)
			siteWithMinAccessTime = k.(string)
		}
		return true
	})
	return siteWithMinAccessTime
}

func findMaxAccessTimeSite() string {
	maxAccessTime := float64(-1 << 63)
	var siteWithMaxAccessTime string
	SiteAvailability.Range(func(k, v interface{}) bool {
		if v.(float64) > maxAccessTime {
			maxAccessTime = v.(float64)
			siteWithMaxAccessTime = k.(string)
		}
		return true
	})
	return siteWithMaxAccessTime
}

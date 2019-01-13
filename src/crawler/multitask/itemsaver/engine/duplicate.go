package engine

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool  {
	if visitedUrls[url] {
		//log.Error("重重的url:%s", url)
		return true
	}
	visitedUrls[url] = true
	return false
}


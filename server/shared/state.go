package shared

import "sync"

// ScrapeMutex is the global mutex for synchronizing scraping operations.
var ScrapeMutex = &sync.Mutex{}

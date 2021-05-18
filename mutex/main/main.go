package main

import "sync"

type UrlCheckerType struct {
	MapUrlsMutex *sync.Mutex
	MapUrls map[string]string
}

func main()  {
	ckecker := UrlCheckerType{
		MapUrls: make(map[string]string, 0),
		MapUrlsMutex: new(sync.Mutex),
	}
	writeToMap(&ckecker);
}

func writeToMap(UrlChecker *UrlCheckerType)  {
	UrlChecker.MapUrlsMutex.Lock();
	defer UrlChecker.MapUrlsMutex.Unlock();

	UrlChecker.MapUrls["simple-key"] = "123";
}
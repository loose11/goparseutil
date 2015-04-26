# goparseutil

## Introduce
This Library is for the Parse.com service. I tried to implement a similiar behaviour like the Android SDK. Please consider, that is an pre-alpha and there is much more to do.
Maybe this library helps someone.

##Example


First of all you have to initialize the library with the AppId and the ApiKey:
```golang
goparseutil.Initialize("i7kQ1yfBexampleAppIdpe8PGLAmFGV",
		"NgFexampleApiKeywIbafKubu")
		
```

After that you can initialize your ParseObject and put some variables in it:
```golang
	element := goparseutil.NewParseObject("gamescore")
	element.Add("score", 1234)
	element.Add("cheatMode", true)
	element.Add("playerName", "Matthias")
	element.AddFile("picture", filename)
```
If you filled the element you can save it to the cloud:
```golang
	element.Save()
```


---
For question open an [issue](https://github.com/loose11/goparseutil/issues/new) under the tag: 'Question' :-)

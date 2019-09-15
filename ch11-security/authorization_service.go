package main

import (
	"fmt"
	"github.com/keets2012/Micro-Go-Pracrise/basic"
	"net/http"
	"strconv"
)

func startAuthorizationHttpListener(host string, port int)  {
	basic.Server = &http.Server{
		Addr: host + ":" +strconv.Itoa(port),
	}
	http.HandleFunc("/health", basic.CheckHealth)
	http.HandleFunc("/discovery", basic.DiscoveryService)
	http.Handle("/oauth/token", clientAuthorizationMiddleware(http.HandlerFunc(getOAuthToken)))
	err := basic.Server.ListenAndServe()
	if err != nil{
		basic.Logger.Println("Service is going to close...")
	}
}

func clientAuthorizationMiddleware(next http.Handler) http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		basic.Logger.Println("Executing client authorization handler")

		// 对客户端信息进行校验
		// clientId:clientSecret Base64加密

		authorization := r.Header.Get("Authorization")

		if authorization == ""{
			http.Error(w, "Please provide the clientId and clientSecret in authorization", http.StatusForbidden)
		}

		//decodeBytes, err := base64.StdEncoding.DecodeString(authorization)
		//if err != nil{
		//	http.Error(w, "Please provide corrent base64 encoding", http.StatusForbidden)
		//}

		fmt.Println("authorization is " + authorization)

		next.ServeHTTP(w, r)

	})
	
}


func getOAuthToken(writer http.ResponseWriter, reader *http.Request)  {
	
}



//clientDetailsService *ClientDetailsService


func main()  {



	//clientDetailsService := NewInMemoryClientDetailService({
	//
	//
	//})
	basic.StartService("Authorization", "127.0.0.1", 10087, startAuthorizationHttpListener)
}

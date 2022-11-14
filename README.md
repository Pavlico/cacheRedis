How to use:
Create config file in the root of your project:

{
"Addr": "cache:6379",
"Password": "",
"DB": 0,
"ExpireTimeSeconds": 10
}


Wrap handle func in middleware function.
routes.HandleFunc("/user/", cacheRedis.HandlerCacheMiddleware(handler.GetUserById()))
thats is!
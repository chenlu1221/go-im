# go-im
前端基于vue3 pinia websocket  后端gin+gorm+mongodb的一个前后端分离的im程序
本程序通过golang 的协程，channel以及io多路复用机制实现多用户发送消息以及消息接收。

#数据库
使用了mysql和mongodb，MySQL用来存储用户信息等，mongodb用来缓存消息

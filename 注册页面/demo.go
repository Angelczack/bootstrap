package main  
  
import (  
 "encoding/json"  
 "fmt"  
 "log"  
 "net/http"  
)  
  
// 定义用户结构体  
type User struct {  
 ID   int    `json:"id"`  
 Name string `json:"name"`  
 Age  int    `json:"age"`  
}  
  
// 用户存储管理器  
var users []User  
  
func main() {  
 // 初始化用户存储管理器  
 users = []User{  
 {ID: 1, Name: "Alice", Age: 25},  
 {ID: 2, Name: "Bob", Age: 30},  
 {ID: 3, Name: "Charlie", Age: 35},  
 }  
  
 // 注册路由处理函数并启动服务器  
 http.HandleFunc("/users", getUsers)  
 http.HandleFunc("/users/", createUser)  
 http.HandleFunc("/users/<id>", getUser)  
 http.HandleFunc("/users/<id>/update", updateUser)  
 http.HandleFunc("/users/<id>/delete", deleteUser)  
 log.Fatal(http.ListenAndServe(":8080", nil))  
}  
  
// 获取所有用户  
func getUsers(w http.ResponseWriter, r *http.Request) {  
 w.Header().Set("Content-Type", "application/json")  
 json.NewEncoder(w).Encode(users)  
}  
  
// 创建用户  
func createUser(w http.ResponseWriter, r *http.Request) {  
 var user User  
 json.NewDecoder(r.Body).Decode(&user)  
 users = append(users, user)  
 fmt.Fprint(w, "User created successfully!")  
}  
  
// 获取指定用户信息  
func getUser(w http.ResponseWriter, r *http.Request) {  
 id := r.URL.Path[len("/users/"):]  
 for _, user := range users {  
 if user.ID == id {  
 w.Header().Set("Content-Type", "application/json")  
 json.NewEncoder(w).Encode(user)  
 return  
 }  
 }  
 http.NotFound(w, r)  
}  
  
// 更新指定用户信息  
func updateUser(w http.ResponseWriter, r *http.Request) {  
 id := r.URL.Path[len("/users/"):]  
 for i, user := range users {  
 if user.ID == id {  
 var updateUser User  
 json.NewDecoder(r.Body).Decode(&updateUser)  
 users[i].Name = updateUser.Name  
 users[i].Age = updateUser.Age  
 fmt.Fprint(w, "User updated successfully!")  
 return  
 }  
 }  
 http.NotFound(w, r)  
}  
  
// 删除指定用户信息  
func deleteUser(w http.ResponseWriter, r *http.Request) {  
 id := r.URL.Path[len("/users/"):]  
 for i, user := range users {  
 if user.ID == id {  
 users = append(users[:i], users[i+1:]...)  
 fmt.Fprint(w, "User deleted successfully!")  
 return  
 }  
 }  
 http.NotFound(w, r)  
}